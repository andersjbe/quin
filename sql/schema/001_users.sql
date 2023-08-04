-- +goose Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` BINARY(16) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `first_name` varchar(127) NOT NULL,
  `last_name` varchar(127) NOT NULL,
  `email` varchar(127) NOT NULL,
  `password` varchar(127) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
);

-- +goose Down
DROP TABLE `users`;
