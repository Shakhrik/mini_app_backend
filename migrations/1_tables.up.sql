CREATE DATABASE IF NOT EXISTS mini_app_backend;

CREATE TABLE IF NOT EXISTS telegram_bot (
    id uuid PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    token VARCHAR NOT NULL,
    description TEXT,
    image_url VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS shop (
    id uuid PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    domain VARCHAR(150) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    telegram_bot_id uuid NOT NULL,
    FOREIGN KEY (telegram_bot_id) REFERENCES telegram_bot(id)
    ON DELETE CASCADE
);