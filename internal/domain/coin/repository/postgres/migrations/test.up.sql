CREATE TABLE IF NOT EXISTS coins (
    name VARCHAR(255) PRIMARY KEY NOT NULL,
    price INT NOT NULL,
    min_price INT NOT NULL,
    max_price INT NOT NULL,
    hour_change_price DOUBLE PRECISION NOT NULL
);

INSERT INTO coins(name, price, min_price, max_price, hour_change_price)
VALUES
    ('BTC',10,10,10,10.5),
    ('ETH',10,10,10,10.5);