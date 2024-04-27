-- +migrate Up
CREATE TABLE IF NOT EXISTS `roast_levels` (
  `id` SMALLINT PRIMARY KEY COMMENT 'roast id',
  `level` SMALLINT NOT NULL COMMENT '焙煎度合い',
  `name` VARCHAR(255) NOT NULL COMMENT '名前'
) ENGINE=InnoDB COMMENT='roast level table';

INSERT INTO `roast_levels` (`id`, `level`, `name`) VALUES
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
  `roast_level_id` SMALLINT NOT NULL COMMENT 'roast id',
  `production_area` VARCHAR(255) NOT NULL COMMENT '産地',
  `kind` VARCHAR(255) NOT NULL COMMENT '品種',
  `price` SMALLINT NOT NULL COMMENT '購入価格',
  `gram` SMALLINT NOT NULL COMMENT '購入グラム',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'create_date',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_date',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete_date',
  FOREIGN KEY (`roast_level_id`) REFERENCES `roast_levels` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB COMMENT='bean table';

CREATE TABLE IF NOT EXISTS `recipes` (
  `id` SERIAL PRIMARY KEY NOT NULL COMMENT 'recipes id',
  `bean_id` BIGINT UNSIGNED NOT NULL COMMENT 'bean id',
  `user_id` VARCHAR(255) NOT NULL COMMENT 'user id',
  `step_number` VARCHAR(255) NOT NULL COMMENT 'ステップ数',
  `memo` VARCHAR(255) NOT NULL COMMENT 'メモ',
  `temperature` VARCHAR(255) NOT NULL COMMENT '温度',
  `seconds` VARCHAR(255) NOT NULL COMMENT '秒数',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'create_date',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_date',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete_date',
  FOREIGN KEY (`bean_id`) REFERENCES `beans` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB COMMENT='recipe table';

-- +migrate Down
DROP TABLE IF EXISTS `recipes`;
DROP TABLE IF EXISTS `beans`;
DROP TABLE IF EXISTS `roast_levels`;