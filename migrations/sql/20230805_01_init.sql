-- +migrate Up
CREATE TABLE IF NOT EXISTS `roast_levels` (
  `id` SERIAL PRIMARY KEY COMMENT 'roast id',
  `level` SMALLINT NOT NULL COMMENT '焙煎度合い',
  `name` VARCHAR(255) NOT NULL COMMENT '名前'
) ENGINE=InnoDB COMMENT='roast level table';

INSERT IGNORE INTO `roast_levels` (`id`, `level`, `name`) VALUES
(1, 1,'Light'),
(2, 2,'Cinnamon'),
(3, 3,'Medium'),
(4, 4,'Hight'),
(5, 5,'City'),
(6, 6,'Full City'),
(7, 7,'French'),
(8, 8,'Italian');

CREATE TABLE IF NOT EXISTS `beans` (
  `id` SERIAL PRIMARY KEY COMMENT 'beans id',
  `user_id` VARCHAR(255) NOT NULL COMMENT 'user id',
  `roast_level_id` BIGINT UNSIGNED NOT NULL COMMENT 'roast id',
  `production_area` VARCHAR(255) NOT NULL COMMENT '産地',
  `kind` VARCHAR(255) NOT NULL COMMENT '品種',
  `price` SMALLINT NOT NULL COMMENT '購入価格',
  `gram` SMALLINT NOT NULL COMMENT '購入グラム',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'create_date',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_date',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete_date'
) ENGINE=InnoDB COMMENT='bean table';

CREATE TABLE IF NOT EXISTS `recipes` (
  `id` SERIAL PRIMARY KEY NOT NULL COMMENT 'recipe id',
  `user_id` VARCHAR(255) NOT NULL COMMENT 'user id',
  `extraction_equipment` VARCHAR(255) NOT NULL COMMENT '抽出器具',
  `coffee_type` VARCHAR(255) NOT NULL COMMENT '種類',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'create_date',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_date',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete_date'
) ENGINE=InnoDB COMMENT='recipe table';

CREATE TABLE IF NOT EXISTS `recipe_steps` (
  `id` SERIAL PRIMARY KEY NOT NULL COMMENT 'recipe step id',
  `user_id` VARCHAR(255) NOT NULL COMMENT 'user id',
  `recipe_id` BIGINT UNSIGNED NOT NULL COMMENT 'recipe id',
  `step_number` SMALLINT NOT NULL COMMENT '順番',
  `memo` VARCHAR(255) NOT NULL COMMENT 'メモ',
  `extraction_temperature` VARCHAR(255) NOT NULL COMMENT '抽出温度',
  `seconds` VARCHAR(255) NOT NULL COMMENT '秒数',
  FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB COMMENT='recipe table';

CREATE TABLE IF NOT EXISTS `reviews` (
  `id` SERIAL PRIMARY KEY NOT NULL COMMENT 'review id',
  `recipe_id`  BIGINT UNSIGNED NOT NULL COMMENT 'recipe id',
  `bean_id` BIGINT UNSIGNED NOT NULL COMMENT 'bean id',
  `user_id` VARCHAR(255) NOT NULL COMMENT 'user id',
  `review` VARCHAR(255) COMMENT 'レビュー',
  FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  FOREIGN KEY (`bean_id`) REFERENCES `beans` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB COMMENT='review table';

-- +migrate Down
DROP TABLE IF EXISTS `reviews`;
DROP TABLE IF EXISTS `recipe_steps`;
DROP TABLE IF EXISTS `recipes`;
DROP TABLE IF EXISTS `beans`;
DROP TABLE IF EXISTS `roast_levels`;