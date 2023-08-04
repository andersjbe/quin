-- +goose Up
CREATE TABLE IF NOT EXISTS `outfits` (
  `id` binary(16) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `description` text,
  `image_urls` json NOT NULL DEFAULT (json_array()),
  `profile_id` binary(16) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `profile_id` (`profile_id`),
  CONSTRAINT `outfits_ibfk_1` FOREIGN KEY (`profile_id`) REFERENCES `profiles` (`id`)
);

-- +goose Down
DROP TABLE `outfits`;
