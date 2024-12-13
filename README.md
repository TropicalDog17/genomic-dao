# Backend engineer assignment

## Table of contents
* [Description](#description)
* [Components](#components)
* [Requirements](#requirements)
* [Instruction](#instruction)
* [How to run it](#how-to-run-it)
* [File structure](#file-structure)
 
## Description
GenomicDAO is a cutting-edge DeSci (Decentralized Science) platform harnessing the power of AI to revolutionize precision medicine. The platform is leading the way in Post-Covid Stroke Prevention (PCSP), offering users not only access to advanced genetic testing services but also a unique opportunity to participate in the governance of drug discovery and development processes.

The flagship product, G-Stroke, empowers users by providing personalized insights into their risk of stroke based on genetic data, delivered securely through our blockchain-powered infrastructure.

Your homework assignment is to design and implement an end-to-end user flow within the GenomicDAO, one of the application that built on top of the LIFE AI.

### User flow
![E2E flow](./images/end-user-flow.jpg?raw=true "e2e flow")

1. Purchase and Sample Collection:
- After registering, the user purchases the G-Stroke service. A saliva sample is then collected from the user and sent to a laboratory for genetic decoding. 
2. Data Generation and Storage:
- The laboratory processes the saliva sample and generates the user’s raw genetic data in a text file format. This genetic data is encrypted and securely stored within the system, ensuring privacy and security.
3. Secure AI Computation:
- All AI-driven computations on the user’s genetic data are conducted within a Trusted Execution Environment (TEE). This ensures that the computations are tamper-proof and that no intermediate data can be accessed by unauthorized parties.
4. Report Generation:
- The output of the AI computation provides genetic insights, which are then compiled into a comprehensive report for the user, detailing their risk of stroke.
5. Token Allocation Based on Risk Score:
- Depending on the user’s stroke risk score, they receive a corresponding amount of PCSP tokens:
	- Extremely High Risk: 15,000 PCSP
 	- High Risk: 3,000 PCSP
  - Slightly High Risk: 225 PCSP
  - Normal or Low Risk: 30 PCSP
6. NFT Minting:
- Simultaneously, a G-NFT is minted on the blockchain. This NFT represents the ownership of the user's genetic profile, securely linking their unique genetic data to a digital asset.

### Data format
**Gene Data Storage**

The genetic information for each user is stored as a text file. Importantly, these files contain no identifiable information that could trace back to the user's real-world identity. Each file is unique to ensure no duplicates exist within the system.

**File Characteristics**

Typically, the size of each text file is around 40MB. To streamline this assignment, the gene text file will contain one of the following predefined risk indicators:
- "extremely high risk"
- "high risk"
- "slightly high risk"
- "low risk"

**Processing and Risk Scoring**

When processed within the Trusted Execution Environment (TEE), the program will assign a risk score corresponding to the gene data content.

## Components

![Components](./images/components.jpg?raw=true "Components")

### Gateway
The client makes requests to the Gateway using conventional REST endpoints.

### Data Storage
Where the encrypted gene data is stored.

### Trusted Execution Environment
To protect algorithms and data from unauthorized observers, ensuring that inputs and outputs are end-to-end encrypted, we use the trusted execution environment (TEE).

### Blockchain
The following main functions need to be implemented:
- Submit gene data
- Mint G-NFT
- Reward PCSP

Those features should be implemented with contracts on Avax Subnet. For more information:
- [Avalanche Docs](https://docs.avax.network/build/subnet/hello-subnet)

## Requirements
- To complete this assignment you should implement the main features in the #Components section to demonstrate the user’s end-to-end flow.
- Propose solutions to protect the user’s gene data so that no one could access it directly without the end user’s permission.
- Implement the solution on the Avalanche Subnet.

### Avalanche Subnet

**Build the the Subnet**

[Avalanche Docs](https://docs.avax.network/build/subnet/hello-subnet)

Follow the above instruction to build the subnet with the following configuration:
- Network name: LIFENetwork
- Chain ID: 9999
- Currency Symbol: LIFE

## Instructions
To complete this assignment you should implement the `/// TODO:` with solution code. Most of the changes will be in `/contracts/`

- Contracts:
    - NFT.sol: for GeneNFT
    - Token.sol: for PCSP Token
    - Storage.sol: for managing the document info on-chain

- Blockchain event when:
   - Gene data is submitted
   - GNFT is minted
   - PCSP is rewarded

### Run the test
Requirements
- [Node](https://nodejs.org/en)

Go to folder `/genomicdao/` and run the following command to install package requirements
```
npm install
```

To run the test
```
npx hardhat test
```

## File structure
```
genomicdao/                 # Directory contains main
│   
├── contracts/                          
│   ├── NFT.sol
│   ├── Token.sol
│   ├── Storage.sol
│   
├── scripts
│   ├── deploy.js
└── ...
```

