SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

DROP SCHEMA IF EXISTS `todo`;
CREATE SCHEMA IF NOT EXISTS `todo` DEFAULT CHARACTER SET utf8mb4;
USE `todo`;

SET CHARSET utf8mb4;

CREATE TABLE IF NOT EXISTS `user` (
    `id` VARCHAR(64) NOT NULL COMMENT 'ユーザーID',
    `name` VARCHAR(128) NOT NULL COMMENT 'ユーザーネーム',
    `password` VARCHAR(128) NOT NULL COMMENT 'パスワード',
    `mail` VARCHAR(32) NOT NULL COMMENT 'メールアドレス',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成時',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新時',
    PRIMARY KEY (`id`)
);

CREATE TABLE if NOT EXISTS `tasks` (
    `id` VARCHAR(64) NOT NULL COMMENT 'タスクID',
    `user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザーID',
    `title` VARCHAR(64) NOT NULL COMMENT 'タイトル',
    `description` VARCHAR(512) NOT NULL COMMENT 'タスク概要',
    `is_complete` TINYINT(1) NOT NULL COMMENT 'タスク完了判定',
    `created_at`  DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成時',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新時',
    PRIMARY KEY (`id`)
);

INSERT INTO `user` (`id`,`name`,`password`,`mail`) VALUES ("1","test1","testpass1","testmail1");
INSERT INTO `user` (`id`,`name`,`password`,`mail`) VALUES ("2","test2","testpass2","testmail2");

INSERT INTO `tasks` (`id`,`user_id`,`title`,`description`,`is_complete`) VALUES ("1","1","test_tiele1","test_description1",false);
INSERT INTO `tasks` (`id`,`user_id`,`title`,`description`,`is_complete`) VALUES ("2","2","test_tiele2","test_description2",false);
INSERT INTO `tasks` (`id`,`user_id`,`title`,`description`,`is_complete`) VALUES ("3","1","test_tiele3","test_description3",false);