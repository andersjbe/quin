-- +goose Up
CREATE TABLE IF NOT EXISTS `products` (
  `id` BINARY(16) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `name` varchar(127) NOT NULL,
  `description` text,
  `sourceUrl` varchar(255) DEFAULT NULL,
  `available_to_buy` tinyint(1) NOT NULL,
  `gender` enum('male','female','unisex') DEFAULT NULL,
  `categories` enum('headwear','eyewear','tops','bottoms','dresses','outerwear','accessories','shoes','swimwear','piercings','necklaces','rings') DEFAULT NULL,
  `materials` json NOT NULL DEFAULT (json_array()),
  `image_url` varchar(255) NOT NULL,
  `profile_id` binary(16) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `profile_id` (`profile_id`),
  CONSTRAINT `products_ibfk_1` FOREIGN KEY (`profile_id`) REFERENCES `profiles` (`id`)
);

-- +goose Down
DROP TABLE `products`;
