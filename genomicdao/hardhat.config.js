require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.19",
  networks: {
    lifeNetwork: {
      url:"http://127.0.0.1:9650/ext/bc/2DRnyQGGPuypaPCvC3FkpZqjZyHQBskd2nmGrba2Jv4CVGx24X/rpc",
      chainId: 8386,
      accounts: [
        "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
      ]
    }
  }
};