CREATE TABLE beers (
    id int not null auto increment,
    name    varchar(150),
    price   int(10),
    country varchar(150),
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME ON UPDATE CURRENT_TIMESTAMP;
    PRIMARY KEY (`id`)
) engine = InnoDB
  DEFAULT charset = utf8;