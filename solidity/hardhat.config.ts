import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox-viem";
import * as dotenv from "dotenv";
import { vars } from "hardhat/config";

dotenv.config();
var ETHERSCAN_API_KEY = vars.get("ETHERSCAN_API_KEY")

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
    },
    'sepolia': {
      url: 'http://211.75.24.92:8545/',
      chainId: 11155111,
      accounts: process.env.KEY ? [process.env.KEY] : [],
    },
    'holesky': {
      url: 'https://rpc.holesky.ethpandaops.io',
      chainId: 17000,
      accounts: process.env.KEY ? [process.env.KEY] : [],
    }
  },
  etherscan: {
    apiKey: ETHERSCAN_API_KEY,
  },
};

export default config;
