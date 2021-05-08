CREATE DATABASE wp_inventory;
USE wp_inventory;

CREATE TABLE `inventory` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `inbound_detail_id` int NOT NULL,
                             `item_id` int NOT NULL,
                             `quantity` int NOT NULL,
                             `created_at` datetime(6) NULL DEFAULT CURRENT_TIMESTAMP(6),
                             PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;