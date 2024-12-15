-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(20) PRIMARY KEY,
    pubkey TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    UNIQUE(pubkey)
);

-- Create GeneData table
CREATE TABLE IF NOT EXISTS gene_data (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    file_id TEXT NOT NULL UNIQUE,
    user_id INTEGER NOT NULL,
    encrypted_data BLOB NOT NULL,
    data_hash BLOB NOT NULL,
    signature BLOB NOT NULL
);