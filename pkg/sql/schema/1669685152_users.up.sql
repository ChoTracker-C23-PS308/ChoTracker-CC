CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(150) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    birth_date VARCHAR(255) NOT NULL,
    gender VARCHAR(100) NOT NULL,
    image_url VARCHAR(255) NOT NULL,


    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

