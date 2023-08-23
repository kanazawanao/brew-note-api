-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` VARCHAR(255) NOT NULL COMMENT 'user_id',
  `name` VARCHAR(255) COMMENT 'user_name',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'create_date',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_date',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete_date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='user table';
CREATE TABLE IF NOT EXISTS `place_types` (
  `id` VARCHAR(255) NOT NULL COMMENT 'place_type_id',
  `key` VARCHAR(255) NOT NULL COMMENT 'キー',
  `name` VARCHAR(255) NOT NULL COMMENT '名称',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'create_date',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_date',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete_date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='place type table';
CREATE TABLE IF NOT EXISTS `stores` (
  `id` VARCHAR(255) NOT NULL COMMENT 'store id',
  `store_type` VARCHAR(255) NOT NULL COMMENT 'store type(store or ec_store)',
  `name` VARCHAR(255) NOT NULL COMMENT '名称',
  `address` VARCHAR(255) NOT NULL COMMENT '住所',
  `url` VARCHAR(255) NOT NULL COMMENT 'url',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'create_date',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_date',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete_date',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='store table';

CREATE TABLE IF NOT EXISTS `beans` (
  `id` VARCHAR(255) NOT NULL COMMENT 'beans id',
  `store_id` VARCHAR(255) NOT NULL COMMENT 'store id(store or ec_store)',
  `production_area` VARCHAR(255) NOT NULL COMMENT '産地',
  `plantation_name` VARCHAR(255) NOT NULL COMMENT '農園',
  `kind` VARCHAR(255) NOT NULL COMMENT '品種',
  `roast_level` VARCHAR(255) NOT NULL COMMENT '焙煎度合い',
  `price` VARCHAR(255) NOT NULL COMMENT '100g当たりの単価',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'create_date',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update_date',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete_date',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB COMMENT='store table';

-- +migrate Down
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `stores`;
DROP TABLE IF EXISTS `place_types`;
DROP TABLE IF EXISTS `beans`;