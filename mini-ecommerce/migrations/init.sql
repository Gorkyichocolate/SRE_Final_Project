CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,

    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,

    name VARCHAR(255) NOT NULL,
    description TEXT,

    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),

    stock INT NOT NULL DEFAULT 0 CHECK (stock >= 0),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,

    user_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,

    quantity INT NOT NULL CHECK (quantity > 0),

    total_price DECIMAL(10,2) NOT NULL CHECK (total_price >= 0),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
            ON DELETE CASCADE,

    CONSTRAINT fk_product
        FOREIGN KEY(product_id)
            REFERENCES products(id)
            ON DELETE CASCADE
);

INSERT INTO products (name, description, price, stock)
VALUES
(
    'Mechanical Keyboard',
    'RGB mechanical keyboard',
    89.99,
    15
),
(
    'Gaming Mouse',
    'Wireless gaming mouse',
    49.99,
    20
),
(
    'Monitor 27',
    '27 inch IPS monitor',
    299.99,
    8
);