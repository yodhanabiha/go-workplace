CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) unique NOT NULL,
    password VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    at_created timestamp,
    at_updated timestamp
);