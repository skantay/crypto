CREATE TABLE IF NOT EXISTS domain.coins (
    name VARCHAR(255) PRIMARY KEY NOT NULL,
    price INT NOT NULL,
    min_price INT NOT NULL,
    max_price INT NOT NULL,
    hour_change_price DOUBLE PRECISION NOT NULL
);