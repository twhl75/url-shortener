CREATE DATABASE IF NOT EXISTS `urlDB`;

USE `urlDB`;

CREATE TABLE `url` (
    `id` int NOT NULL AUTO_INCREMENT,
    `original` varchar(255),
    `shortened` varchar(255),
    PRIMARY KEY (`id`)
);

INSERT INTO `url` (`original`, `shortened`)
VALUES ('test.com/thisIsForInit', '123');