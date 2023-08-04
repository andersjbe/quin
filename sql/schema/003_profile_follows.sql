-- +goose Up
CREATE TABLE IF NOT EXISTS `profile_follows` (
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `follower_id` binary(16) NOT NULL,
  `following_id` binary(16) NOT NULL,
  PRIMARY KEY (`follower_id`,`following_id`),
  KEY `following_id` (`following_id`),
  CONSTRAINT `profile_follows_ibfk_1` FOREIGN KEY (`follower_id`) REFERENCES `profiles` (`id`),
  CONSTRAINT `profile_follows_ibfk_2` FOREIGN KEY (`following_id`) REFERENCES `profiles` (`id`)
);

-- +goose Down
DROP TABLE `profile_follows`;
