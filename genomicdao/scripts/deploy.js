const { ethers } = require("hardhat");
async function main() {
    const [deployer] = await ethers.getSigners();
    
    console.log("deployer account:", deployer.address);
    const balance = (await deployer.provider.getBalance(deployer.address)).toString();
    console.log("Account balance:", balance);

    // Deploy GeneNFT contract
    const GeneNFTToken = await ethers.getContractFactory("GeneNFT");
    const geneNftToken = await GeneNFTToken.deploy();
    await geneNftToken.waitForDeployment();
    console.log("GeneNFT Token deployed at address:", geneNftToken.target);

    // Deploy PostCovidStrokePrevention contract
    const PCSPToken = await ethers.getContractFactory("PostCovidStrokePrevention");
    const pcspToken = await PCSPToken.deploy();
    await pcspToken.waitForDeployment();
    console.log("PCSP Token deployed at address:", pcspToken.target);

    // Deploy Controller
    const Controller = await ethers.getContractFactory("Controller");
    const controller = await Controller.deploy(geneNftToken.target, pcspToken.target);
    await controller.waitForDeployment();
    console.log("Controller deployed at address:", controller.target);

    // Transfer ownership of tokens to Controller
    const transferNFTTx = await geneNftToken.transferOwnership(controller.target);
    await transferNFTTx.wait();
    console.log("Ownership of GeneNFTToken transferred to Controller");

    const transferPCSPTx = await pcspToken.transferOwnership(controller.target);
    await transferPCSPTx.wait();
    console.log("Ownership of PCSPToken transferred to Controller");

    // Verify ownership
    const nftOwner = await geneNftToken.owner();
    const pcspOwner = await pcspToken.owner();
    
    console.log("Current NFT owner:", nftOwner);
    console.log("Current PCSP owner:", pcspOwner);
    console.log("Controller address:", controller.target);
    
    if (nftOwner !== controller.target || pcspOwner !== controller.target) {
        console.error("Ownership transfer failed!");
    }
}

main().catch((error) => {
    console.error("Error during deployment:", error);
    process.exit(1);
});