-- +goose Up
CREATE TABLE IF NOT EXISTS `product_variants` (
  `id` BINARY(16) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `name` varchar(127) NOT NULL,
  `colors` json NOT NULL DEFAULT (json_array()),
  `pattern` varchar(127) NOT NULL,
  `image_urls` json DEFAULT (json_array()),
  `product_id` binary(16) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `product_id` (`product_id`),
  CONSTRAINT `product_variants_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);

-- +goose Down
DROP TABLE `product_variants`;
