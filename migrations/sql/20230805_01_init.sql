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

INSERT INTO `roast_levels`(`level`, `name`) VALUES
('1','Light roast'),
('2', 'Chinamon roast'),
('3', 'Medium roast'),
('4', 'High roast'),
('5', 'City roast'),
('6', 'FUllcity roast'),
('7', 'French roast'),
('8', 'Italian roast');

CREATE TABLE IF NOT EXISTS `processings` (
  `id` SERIAL PRIMARY KEY COMMENT 'processings id',
  `method` VARCHAR(255) COMMENT '精製方法',
  `description` VARCHAR(255) COMMENT '説明'
) ENGINE=InnoDB COMMENT='processings table';

INSERT INTO `processings`(`method`, `description`) VALUES
('Washed','スッキリとしたクリアなフレーバー。産地特有の味をダイレクトに。'),
('Honey', 'washedとnaturalの中間。クリアな質感もありながら、香りや甘みも。'),
('Natural', '濃厚なフレーバー。香りや甘みが濃厚。フルーティー。'),
('Anaerobic', '空気がない状態(嫌気)で発酵させる特殊な方法。複雑な味・お酒のようなフレーバー。'),
('other', 'その他');

CREATE TABLE IF NOT EXISTS `beans` (
  `id` SERIAL PRIMARY KEY COMMENT 'beans id',
  `user_id` VARCHAR(255) COMMENT 'user id',
  `roast_id` INT NOT NULL COMMENT '焙煎度',
  `processing_id` INT NOT NULL COMMENT '精製方法',
  `producing_region` VARCHAR(255) NOT NULL COMMENT '生産地域',
  `kind` VARCHAR(255) NOT NULL COMMENT '品種',
  `process` VARCHAR(255),
  `price` SMALLINT NOT NULL COMMENT '購入価格',
  `gram` SMALLINT NOT NULL COMMENT '購入グラム',
  `memo` VARCHAR(255) NOT NULL COMMENT 'メモ',
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
DROP TABLE IF EXISTS `processings`;
DROP TABLE IF EXISTS `roast_levels`;
DROP TABLE IF EXISTS `users`;
