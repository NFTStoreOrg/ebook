// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;
import {ERC1155Pausable} from "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Pausable.sol";
import {ERC1155} from "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {AutomationCompatibleInterface} from "@chainlink/contracts/src/v0.8/automation/AutomationCompatible.sol";

/// @title The Yi Sin EBook store ERC1155 contract
/// @author Yi-Sin NFT
/**
 * @notice This contract provides functions for book borrowing, returning, and shelving,
 *         using NFTs held under ERC1155 as authentication for data access rights.
 */
contract YiSinEBook is ERC1155, ERC1155Pausable, Ownable, ReentrancyGuard {
    uint256 constant DECIMALS = 3; //  10^3
    uint256 constant DECIMAL_FACTOR = 10 ** DECIMALS;
    uint256 public fee = 30; //  Marketplace fee（fee/1000）
    struct BookInfo {
        address writer;
        uint256 supplyAmount;
        uint256 rentPrice;
        uint256 maxRentTime;
    }
    struct RentInfo {
        address renter;
        uint256 tokenId;
        uint256 endTime;
    }
    //  All type of books.
    uint256 public totalSupplyBook = 0;
    //  tokenId => bookInfo
    mapping(uint256 => BookInfo) public bookInfos;
    //  tokenId => rentInfo
    mapping(uint256 => RentInfo[]) public rentInfos;
    //  Book renter now, use for rentInfos index. tokenId => renter amount
    mapping(uint256 => uint256) public booksOnRent;
    //  renter => bookId => index
    mapping(address => mapping(uint256 => uint256)) public renterRentInfoIndex;
    mapping(uint256 => bool) private _isBookBeBurned;

    constructor(address initialOwner) ERC1155("") Ownable(initialOwner) {}

    // Ebook store function support for depolyer.
    function uploadEBook(
        uint256 bookAmount,
        address uploader,
        uint256 price,
        uint256 time
    ) external onlyOwner {
        _mint(msg.sender, totalSupplyBook, bookAmount, "");
        bookInfos[totalSupplyBook] = BookInfo({
            writer: uploader,
            supplyAmount: bookAmount,
            rentPrice: price,
            maxRentTime: time
        });
        totalSupplyBook++;
    }

    function burnEBook(uint256 bookId) external onlyOwner {
        _burn(owner(), bookId, bookInfos[bookId].supplyAmount);
        _isBookBeBurned[bookId] = true;
        delete bookInfos[bookId];
        delete rentInfos[bookId];
        //  No need totalSupplyBook--
    }

    /// @dev The following function are provided for borrower.
    /**
     * @dev Record booksOnRent when user borrow the book,
     *      and record the index to read rentInfos when I need to check the user's rent information.
     * @param bookId Rent what type of book.
     * @param rentTime Rent duration(in seconds).
     */
    function rentBook(
        uint256 bookId,
        uint256 rentTime
    ) external payable nonReentrant {
        require(!_isBookBeBurned[bookId], "This book is not exist.");
        require(msg.sender != address(0), "Invalid call by address 0");
        require(
            balanceOf(owner(), bookId) > 0,
            "Invalid book id or insufficient balance"
        );
        require(balanceOf(msg.sender, bookId) == 0, "Already rented");
        require(isApprovedForAll(msg.sender, owner()), "Not Approved");
        require(msg.value == bookInfos[bookId].rentPrice, "Invalid price");
        require(bookId < totalSupplyBook, "Invalid bookId");
        require(
            rentTime <= bookInfos[bookId].maxRentTime,
            "Exceed max rent time."
        );

        rentInfos[bookId].push(
            RentInfo({
                renter: msg.sender,
                tokenId: bookId,
                endTime: rentTime + block.timestamp
            })
        );
        renterRentInfoIndex[msg.sender][bookId] = booksOnRent[bookId];
        booksOnRent[bookId]++;
        transferValueToWriter(bookInfos[bookId].writer, bookId);
        safeTransferFrom(owner(), msg.sender, bookId, 1, "");
    }

    function returnBook(uint256 bookId) external nonReentrant {
        require(balanceOf(msg.sender, bookId) > 0, "Isn't rented this book");
        require(msg.sender != address(0), "Cannot transfer to address 0");
        require(
            rentInfos[bookId][renterRentInfoIndex[msg.sender][bookId]].renter ==
                msg.sender,
            "Should be renter"
        );
        require(bookId < totalSupplyBook, "Invalid bookId");

        transferBook(bookId, msg.sender);
    }

    //  Chainlink automation return book.
    function checkUpkeep(
        bytes calldata checkData
    ) external view returns (bool upkeepNeeded, bytes memory performData) {
        upkeepNeeded = false;
        for (uint256 i = 0; i < checkData.length; i += 32) {
            uint256 tokenId = abi.decode(checkData[i:i + 32], (uint256));
            /*  condition 1: Reach end time.
             *  condition 2: The tokenId is not on rent.
             */
            for (uint256 index = 0; index < booksOnRent[tokenId]; index++) {
                if (
                    block.timestamp >= rentInfos[tokenId][index].endTime &&
                    rentInfos[tokenId][index].endTime != 0
                ) {
                    upkeepNeeded = true;
                    performData = abi.encode(tokenId, index);
                    break;
                }
                if (upkeepNeeded) {
                    break;
                }
            }
        }
    }

    function performUpkeep(bytes calldata performData) external {
        (uint256 tokenId, uint256 index) = abi.decode(
            performData,
            (uint256, uint256)
        );
        if (
            block.timestamp > rentInfos[tokenId][index].endTime &&
            rentInfos[tokenId][index].endTime != 0
        ) {
            transferBook(tokenId, rentInfos[tokenId][index].renter);
        }
    }

    function transferBook(uint256 bookId, address renterAddress) private {
        rentInfos[bookId][
            renterRentInfoIndex[renterAddress][bookId]
        ] = rentInfos[bookId][rentInfos[bookId].length - 1]; //  Move the last rent information to user's index.
        renterRentInfoIndex[
            rentInfos[bookId][renterRentInfoIndex[renterAddress][bookId]].renter
        ][bookId] = renterRentInfoIndex[renterAddress][bookId]; //  The last user rent information index = user's index
        rentInfos[bookId].pop(); //  pop the last rent information.
        booksOnRent[bookId]--;
        safeTransferFrom(renterAddress, owner(), bookId, 1, "");
    }

    /**
     * @notice Check have authentication for data access rights.
     * @param signer the user address.
     * @param id the tokenId.
     */
    function isAddressHaveTokenId(
        address signer,
        uint256 id
    ) external view returns (bool) {
        if (balanceOf(signer, id) == 0) {
            return false;
        }
        return true;
    }

    //  Function only owner.
    function setURI(string memory newuri) public onlyOwner {
        _setURI(newuri);
    }

    /**
     * @dev pause or unpause ebook store by owner.
     */
    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function transferValueToWriter(address writer, uint256 bookId) private {
        payable(writer).transfer(devFee(bookInfos[bookId].rentPrice));
    }

    function devFee(uint256 price) private view returns(uint256) {
        return price - (price * fee) / DECIMAL_FACTOR;
    }

    function _update(
        address from,
        address to,
        uint256[] memory ids,
        uint256[] memory values
    ) internal override(ERC1155, ERC1155Pausable) {
        super._update(from, to, ids, values);
    }
}
