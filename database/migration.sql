CREATE TABLE `price_data`
(
    id bigint(20) unsigned  auto_increment,
    unix bigint unsigned NOT NULL,
    symbol varchar(255) NOT NULL,
    open_price DECIMAL(27,8) NOT NULL,
    high_price DECIMAL(27,8) NOT NULL,
    low_price DECIMAL(27,8) NOT NULL,
    close_price DECIMAL(27,8) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE INDEX price_symbol ON price_data(symbol);
CREATE INDEX close_price ON price_data(close_price);
CREATE INDEX low_price ON price_data(low_price);
CREATE INDEX open_close_price ON price_data(open_price,close_price);
CREATE INDEX high_low_price ON price_data(high_price,low_price);
