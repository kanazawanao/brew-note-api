-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT 'user_id',
  `name` VARCHAR(255) COMMENT 'user_name',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'create_date',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_date',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete_date',
  PRIMARY KEY (`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `users`;