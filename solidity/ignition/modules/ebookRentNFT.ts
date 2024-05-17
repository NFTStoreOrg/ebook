import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const EBookModule = buildModule("EBookModule", (m) => {
    const deployer = m.getAccount(0);

    const ebook = m.contract("YiSinEBook", [deployer], {
        from: deployer,
    });
    return { ebook };
})

export default EBookModule;