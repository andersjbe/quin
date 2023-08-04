-- +goose Up
CREATE TABLE IF NOT EXISTS `outfit_likes` (
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `outfit_id` binary(16) NOT NULL,
  `profile_id` binary(16) NOT NULL,
  PRIMARY KEY (`outfit_id`,`profile_id`),
  KEY `profile_id` (`profile_id`),
  CONSTRAINT `outfit_likes_ibfk_1` FOREIGN KEY (`outfit_id`) REFERENCES `outfits` (`id`),
  CONSTRAINT `outfit_likes_ibfk_2` FOREIGN KEY (`profile_id`) REFERENCES `profiles` (`id`)
);

-- +goose Down
DROP TABLE `outfit_likes`;
