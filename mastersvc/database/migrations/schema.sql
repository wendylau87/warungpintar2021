CREATE DATABASE wp_master;

CREATE TABLE IF NOT EXISTS `item` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `total` int,
    PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO item(name) VALUES('Chitato');
INSERT INTO item(name) VALUES('Baygon');
