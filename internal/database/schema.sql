-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
SET TIMEZONE="America/Fortaleza";

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true
);

-- Create users table
CREATE TABLE IF NOT EXISTS products (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    description VARCHAR(45) NOT NULL,
    price NUMERIC(15, 2),
    active BOOLEAN NOT NULL DEFAULT true,
    type INT NOT NULL DEFAULT 1
);