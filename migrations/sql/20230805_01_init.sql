-- +migrate Up
CREATE TABLE IF NOT EXISTS `roast_levels` (
  `id` SERIAL PRIMARY KEY COMMENT 'roast id',
  `level` SMALLINT NOT NULL COMMENT '焙煎度合い',
  `name` VARCHAR(255) NOT NULL COMMENT '名前'
) ENGINE=InnoDB COMMENT='roast level table';

INSERT INTO `roast_levels`(`level`, `name`) VALUES
('1','Light roast'),
('2', 'Cinnamon roast'),
('3', 'Medium roast'),
('4', 'High roast'),
('5', 'City roast'),
('6', 'FUll City roast'),
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
  `roast_level_id` INT NOT NULL COMMENT '焙煎度',
  `processing_id` INT NOT NULL COMMENT '精製方法',
  `production_region` VARCHAR(255) NOT NULL COMMENT '生産地域',
  `farm` VARCHAR(255) NOT NULL COMMENT '農園',
  `altitude` INT NOT NULL COMMENT '標高',
  `kind` VARCHAR(255) NOT NULL COMMENT '品種',
  `flavor` VARCHAR(255) NOT NULL COMMENT '香り',
  `memo` VARCHAR(255) NOT NULL COMMENT 'メモ',
  `roasted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '焙煎日',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'create_date',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_date',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete_date'
) ENGINE=InnoDB COMMENT='bean table';

CREATE TABLE IF NOT EXISTS `recipes` (
  `id` SERIAL PRIMARY KEY NOT NULL COMMENT 'recipe id',
  `user_id` VARCHAR(255) NOT NULL COMMENT 'user id',
  `extraction_equipment` VARCHAR(255) NOT NULL COMMENT '抽出器具',
  `coffee_type` VARCHAR(255) NOT NULL COMMENT '珈琲の種類',
  `water_temperature` INT NOT NULL COMMENT '抽出温度',
  `bean_dose` INT NOT NULL COMMENT '豆の量（グラム）',
  `grind` VARCHAR(255) NOT NULL COMMENT '挽き具合',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'create_date',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_date',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete_date'
) ENGINE=InnoDB COMMENT='recipe table';

CREATE TABLE IF NOT EXISTS `recipe_steps` (
  `id` SERIAL PRIMARY KEY NOT NULL COMMENT 'recipe step id',
  `recipe_id` BIGINT UNSIGNED NOT NULL COMMENT 'recipe id',
  `step_number` SMALLINT NOT NULL COMMENT '順番',
  `seconds` INT NOT NULL COMMENT '秒数',
  `amount_water` INT NOT NULL COMMENT '注ぐ量',
  `memo` VARCHAR(255) NOT NULL COMMENT 'メモ',
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
