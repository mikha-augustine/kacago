CREATE DATABASE `kacago`;
USE `kacago`;

CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `user` varchar(255) NOT NULL,
  `pass` varchar(255) NOT NULL
);

INSERT INTO `user` (`user`, `pass`)
VALUES ('admin', 'Q94Qqtqp_g8uzA3aqfeTlQNQpek=');

ALTER TABLE `user`
ADD `secret_message` tinytext COLLATE 'latin1_swedish_ci' NOT NULL;