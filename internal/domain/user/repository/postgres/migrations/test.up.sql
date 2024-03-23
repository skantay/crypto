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

CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    chat_id BIGINT NOT NULL,
    coin VARCHAR(255) NOT NULL,
    notification_interval INTERVAL NOT NULL,
    last_notification_time TIMESTAMP NOT NULL,
    FOREIGN KEY (coin) REFERENCES coins(name)
);
INSERT INTO users (id, chat_id, coin, notification_interval, last_notification_time)
VALUES (1,  123456, 'BTC', '1 mins', '2024-02-21  00:00:00');
