## GenomicDAO Implementation

### Overview
```mermaid
sequenceDiagram
    participant User
    participant Genomic as GenomicService
    participant TEE
    participant Storage
    participant Chain

    %% Register Flow
    User->>Genomic: Register(pubkey)
    Genomic->>Storage: Store user details
    Storage-->>User: Return userID

    %% Upload Flow
    User->>Genomic: ProcessAndUploadGenomicData
    Note over Genomic: Authenticate user
    Genomic->>TEE: Process & encrypt data
    Note over TEE: Calculate risk score
    Genomic->>Storage: Store encrypted data
    Storage-->>Genomic: Return fileID
    Genomic->>Chain: Upload & confirm
    Chain-->>User: Return sessionID & fileID

    %% Retrieve Flow
    User->>Genomic: RetrieveGenomicData(fileID)
    Genomic->>Storage: Fetch encrypted data
    Genomic->>TEE: Decrypt data
    TEE-->>User: Return decrypted data
```

### Details

#### Techstack

+ API: Gin framework
+ Database: SQLite
+ Blockchain: Avalanche L1(Subnet-EVM)

#### Setting up the subnet

First, run the `create` command and enter the neccessary information(chainID, token symbol, subnet type)
```
avalanche blockchain create lifeNetwork
```

Then, deploy the subnet. In this case I deploy to localnet. More information [here](https://docs.avax.network/avalanche-l1s/build-first-avalanche-l1)

```
avalanche blockchain deploy lifeNetwork
```

#### Swagger docs
```bash
./scripts/gen-swagger.sh
```

The swagger will be available at localhost:8080/swagger/index.html
#### Tests

+ Complete user flow: [server_test.go](./server/server_test.go)

##### Result

```
Confirmed upload with tx hash: 0x25481776b12f5b3cb054d3a65d1f5b4ce139c2c5682623e46e225f23b18d6b88
Minted GeneNFT with token ID: 14
Rewarded PCSP with amount 225000000000000000000 to 0x62f563A2e09c7987dECBFF61fdcC89cd74717721
[GIN] 2024/12/15 - 16:38:43 | 200 |  2.085183958s |             ::1 | POST     "/upload"
Raw response body: {"sessionId":"14","message":"Genomic data uploaded successfully","fileId":"c29814719d660d3f"}
[GIN] 2024/12/15 - 16:38:43 | 200 |     635.083µs |             ::1 | GET      "/retrieve?fileID=c29814719d660d3f"
Genomic data: LifeNetwork, Decentralized Science, Blockchain, Genomic Data
[GIN] 2024/12/15 - 16:38:43 | 200 |    2.436708ms |             ::1 | GET      "/pcsp/balance?address=0x62f563A2e09c7987dECBFF61fdcC89cd74717721"
PCSP balance: 225000000000000000000
```

