// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;
import "hardhat/console.sol";
import {ERC1155Pausable} from "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Pausable.sol";
import {ERC1155} from "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/// @title The Yi Sin EBook store ERC1155 contract
/// @author Yi-Sin NFT
/**
 * @notice This contract provides functions for book borrowing, returning, and shelving,
 *         using NFTs held under ERC1155 as authentication for data access rights.
 */
contract YiSinEBook is ERC1155, ERC1155Pausable, Ownable, ReentrancyGuard {
    struct BookInfo {
        address writer;
        uint256 supplyAmount;
        uint256 rentPrice;
    }
    struct RentInfo {
        address renter;
        uint256 tokenId;
        uint256 endTime;
    }
    //  All type of books.
    uint256 public totalSupplyBook;
    //  tokenId => bookInfo
    mapping(uint256 => BookInfo) public bookInfos;
    //  tokenId => rentInfo
    mapping(uint256 => RentInfo[]) public rentInfos;
    //  Book renter now, use for rentInfos index. tokenId => renter amount
    mapping(uint256 => uint256) public booksOnRent;
    mapping(address => uint256) public renterRentInfoIndex;

    constructor(address initialOwner) ERC1155("") Ownable(initialOwner) {}

    // Ebook store function support for depolyer.

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
    ) public payable nonReentrant {
        require(msg.sender != address(0), "Invalid call by address 0");
        require(
            balanceOf(address(this), bookId) > 0,
            "Invalid book id or insufficient balance"
        );
        require(balanceOf(msg.sender, bookId) == 0, "Already rented");
        require(isApprovedForAll(msg.sender, address(this)), "Not Approved");
        require(msg.value == bookInfos[bookId].rentPrice, "Invalid price");

        rentInfos[bookId].push(
            RentInfo({
                renter: msg.sender,
                tokenId: bookId,
                endTime: rentTime + block.timestamp
            })
        );
        renterRentInfoIndex[msg.sender] = booksOnRent[bookId];
        booksOnRent[bookId]++;
        safeTransferFrom(address(this), msg.sender, bookId, 1, "");
    }

    function returnBook(uint256 bookId) public nonReentrant {
        require(balanceOf(msg.sender, bookId) > 0, "Isn't rented this book");
        require(msg.sender != address(0), "Cannot transfer to address 0");
        require(
            rentInfos[bookId][renterRentInfoIndex[msg.sender]].renter ==
                msg.sender,
            "Should be renter"
        );

        rentInfos[bookId][renterRentInfoIndex[msg.sender]] = rentInfos[bookId][
            rentInfos[bookId].length - 1
        ]; //  Move the last rent information to user's index.
        rentInfos[bookId].pop(); //  pop the last rent information.
        renterRentInfoIndex[
            rentInfos[bookId][renterRentInfoIndex[msg.sender]].renter
        ] = renterRentInfoIndex[msg.sender]; //  The last user rent information index = user's index
        booksOnRent[bookId]--;
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

    function _update(
        address from,
        address to,
        uint256[] memory ids,
        uint256[] memory values
    ) internal override(ERC1155, ERC1155Pausable) {
        super._update(from, to, ids, values);
    }
}
