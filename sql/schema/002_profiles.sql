-- +goose Up
CREATE TABLE IF NOT EXISTS `profiles` (
  `id` BINARY(16) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `username` varchar(127) NOT NULL,
  `image_url` varchar(255) NOT NULL,
  `gender` enum('man','woman','nonbinary') DEFAULT NULL,
  `height_inches` int DEFAULT NULL,
  `weight_lbs` int DEFAULT NULL,
  `skin_pigmentation` int DEFAULT NULL,
  `hair_color` varchar(63) DEFAULT NULL,
  `eye_color` varchar(63) DEFAULT NULL,
  `size_preferences` json DEFAULT NULL,
  `user_id` binary(16) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `profiles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

-- +goose Down
DROP TABLE `profiles`;
