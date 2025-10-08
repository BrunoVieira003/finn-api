CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE accounts(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    amount NUMERIC(1000, 2)
);