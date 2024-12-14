pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Counters.sol";
import "./NFT.sol";
import "./Token.sol";

contract Controller {
    using Counters for Counters.Counter;

    //
    // STATE VARIABLES
    //
    Counters.Counter private _sessionIdCounter;
    GeneNFT public geneNFT;
    PostCovidStrokePrevention public pcspToken;

    struct UploadSession {
        uint256 id;
        address user;
        string proof;
        bool confirmed;
    }

    struct DataDoc {
        string id;
        string hashContent;
    }

    mapping(uint256 => UploadSession) sessions;
    mapping(string => DataDoc) docs;
    mapping(string => bool) docSubmits;
    mapping(uint256 => string) nftDocs;

    //
    // EVENTS
    //
    event UploadData(string docId, uint256 sessionId);
    event GeneDataSubmitted(string indexed docId, address indexed user, string hashContent);
    event GeneNFTMinted(address indexed owner, uint256 indexed tokenId, string docId);
    event PCSPRewarded(address indexed user, uint256 amount, uint256 riskScore);

    constructor(address nftAddress, address pcspAddress) {
        geneNFT = GeneNFT(nftAddress);
        pcspToken = PostCovidStrokePrevention(pcspAddress);
        _sessionIdCounter = Counters.Counter(0);
    }

    function uploadData(string memory docId) public returns (uint256) {
        // Check if doc has been submitted before
        require(!docSubmits[docId], "Doc already been submitted");
        
        uint256 sessionId = _sessionIdCounter.current();
        // Increment session counter
        _sessionIdCounter.increment();
        
        // Create new upload session
        sessions[sessionId] = UploadSession({
            id: sessionId,
            user: msg.sender,
            proof: "",
            confirmed: false
        });
        
        // Mark doc as initiated
        docSubmits[docId] = true;
        
        // Emit upload event
        emit UploadData(docId, sessionId);
        
        return sessionId;
    }

    function confirm(
        string memory docId,
        string memory contentHash,
        string memory proof,
        uint256 sessionId,
        uint256 riskScore
    ) public {
        // Verify session exists and is not confirmed
        require(bytes(docs[docId].id).length == 0, "Doc already been submitted");
        require(sessions[sessionId].id == sessionId, "Invalid session ID");
        require(!sessions[sessionId].confirmed, "Session is ended");
        require(sessions[sessionId].user == msg.sender, "Invalid session owner");

        // Update doc content
        docs[docId] = DataDoc({
            id: docId,
            hashContent: contentHash
        });

        // Emit gene data submission event
        emit GeneDataSubmitted(docId, msg.sender, contentHash);

        // Mint NFT
        uint256 tokenId = geneNFT.safeMint(msg.sender);
        nftDocs[tokenId] = docId;

        // Emit NFT minting event
        emit GeneNFTMinted(msg.sender, tokenId, docId);

        // Reward PCSP tokens based on risk score
        uint256 rewardAmount = pcspToken.reward(msg.sender, riskScore);

        // Emit PCSP reward event
        emit PCSPRewarded(msg.sender, rewardAmount, riskScore);

        // Close session
        sessions[sessionId].proof = proof;
        sessions[sessionId].confirmed = true;
    }

    function getSession(uint256 sessionId) public view returns(UploadSession memory) {
        return sessions[sessionId];
    }

    function getDoc(string memory docId) public view returns(DataDoc memory) {
        return docs[docId];
    }
}