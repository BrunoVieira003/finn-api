CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE accounts(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE transactions(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    amount NUMERIC(12,2) NOT NULL,
    type TEXT NOT NULL,
    account_id UUID NOT NULL REFERENCES accounts(id),
    date DATE NOT NULL,
    description TEXT NULL
);