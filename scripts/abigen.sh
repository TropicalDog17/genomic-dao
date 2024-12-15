#!/bin/bash

set -e  # Exit on error

# Create directories
mkdir -p internal/onchain/bindings
mkdir -p build

echo "Compiling contracts..."
# Compile contracts
cd genomicdao
npx hardhat compile
cd ..

# Print available artifacts
echo "Available artifacts:"
ls genomicdao/artifacts/contracts/

echo "Extracting ABIs..."
# Extract ABIs with error checking
if [ -f "genomicdao/artifacts/contracts/NFT.sol/GeneNFT.json" ]; then
    jq .abi "genomicdao/artifacts/contracts/NFT.sol/GeneNFT.json" > build/GeneNFT.abi
    echo "Extracted GeneNFT ABI"
else
    echo "Error: GeneNFT artifact not found"
    exit 1
fi

if [ -f "genomicdao/artifacts/contracts/Token.sol/PostCovidStrokePrevention.json" ]; then
    jq .abi "genomicdao/artifacts/contracts/Token.sol/PostCovidStrokePrevention.json" > build/PCSP.abi
    echo "Extracted PCSP ABI"
else
    echo "Error: PCSP artifact not found"
    exit 1
fi

if [ -f "genomicdao/artifacts/contracts/Controller.sol/Controller.json" ]; then
    jq .abi "genomicdao/artifacts/contracts/Controller.sol/Controller.json" > build/Controller.abi
    echo "Extracted Controller ABI"
else
    echo "Error: Controller artifact not found"
    exit 1
fi

echo "Generating Go bindings..."
# Generate bindings
abigen --abi build/GeneNFT.abi --pkg contracts --type GeneNFT --out internal/onchain/bindings/gene_nft.go
abigen --abi build/PCSP.abi --pkg contracts --type PCSPToken --out internal/onchain/bindings/pcsp_token.go
abigen --abi build/Controller.abi --pkg contracts --type Controller --out internal/onchain/bindings/controller.go

# Clean up
rm -rf build

echo "Done!"