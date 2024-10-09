CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Auto-generated UUID for category
    name VARCHAR(255) UNIQUE NOT NULL             -- Category name must be unique
);

CREATE TABLE category_links (
    international_branch_id VARCHAR(255) PRIMARY KEY,  -- InternationalBranchID field, primary key
    category_id UUID REFERENCES categories(id)         -- Foreign key referencing categories by id
);

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Auto-generated UUID
    name VARCHAR(255) NOT NULL,                   -- Name field, cannot be null
    amount NUMERIC(10, 2) NOT NULL,               -- Amount field, numeric with 2 decimal precision
    date TIMESTAMP NOT NULL,                      -- Date field, cannot be null
    address VARCHAR(255),                         -- Address field, can be null
    card_unique_id VARCHAR(255),                  -- CardUniqueId field, can be null
    category_id UUID REFERENCES categories(id),   -- Foreign key referencing categories
    merchant_phone_no VARCHAR(20),                -- MerchantPhoneNo, can be null
    international_branch_id VARCHAR(255) NOT NULL -- InternationalBranchID field, cannot be null
);
