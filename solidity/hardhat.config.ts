import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox-viem";
import * as dotenv from "dotenv";

dotenv.config();

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.24",
    settings: {
      evmVersion: "london",
    },
  },
  networks: {
    'thunderCoreTestnet': {
      url: 'https://testnet-rpc.thundercore.com',
      chainId: 18,
      accounts: process.env.KEY ? [process.env.KEY] : [],
    },
    'thumderCoreMainnet': {
      url: 'https://mainnet-rpc.thundercore.com',
      chainId: 108,
      accounts: process.env.KEY ? [process.env.KEY] : [],
    },
    'amoyTestnet': {
      url: 'https://polygon-amoy-bor-rpc.publicnode.com',
      chainId: 80002,
      accounts: process.env.KEY ? [process.env.KEY] : [],
    }
  },
  etherscan: {
    apiKey: {
      "thunderCoreTestnet": "unused",
    },
    //  Verify contract on thunderCore
    customChains: [
      {
        network: "thunderCoreTestnet",
        chainId: 18,
        urls: {
          apiURL: "https://explorer-testnet.thundercore.com/api",
          browserURL: "https://explorer-testnet.thundercore.com",
        },
      },
    ],
  },
};

export default config;
