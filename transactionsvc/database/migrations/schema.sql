CREATE DATABASE wp_transaction;
USE wp_transaction;

CREATE TABLE `inbound` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `po_number` varchar(255) NOT NULL,
                           `created_at` datetime(6) NULL DEFAULT CURRENT_TIMESTAMP(6),
                           PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `inbound_detail` (
                                  `id` int NOT NULL AUTO_INCREMENT,
                                  `inbound_id` int NOT NULL,
                                  `item_id` int NOT NULL,
                                  `quantity` int NOT NULL,
                                  PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;