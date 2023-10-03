SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

DROP SCHEMA IF EXISTS `todo`;
CREATE SCHEMA IF NOT EXISTS `todo` DEFAULT CHARACTER SET utf8mb4;
USE `todo`;

SET CHARSET utf8mb4;

CREATE TABLE IF NOT EXISTS `user` (
    `user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザーID',
    `password` VARCHAR(128) NOT NULL COMMENT 'パスワード',
    `token` VARCHAR(128) NOT NULL COMMENT '認証用トークン',
    `user_name` VARCHAR(128) NOT NULL COMMENT 'ユーザーネーム',
    `mail` VARCHAR(32) NOT NULL COMMENT 'メールアドレス',
    `work` VARCHAR(32) NOT NULL COMMENT '職業',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成時',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新時',
    PRIMARY KEY (`user_id`),
    INDEX `idx_auth_token` (`user_id` ASC)
);

CREATE TABLE if NOT EXISTS `tasks` (
    `id` VARCHAR(64) NOT NULL COMMENT 'ユーザーID',
    `todo_id` VARCHAR(64) NOT NULL COMMENT 'タスクID',
    `title` VARCHAR(64) NOT NULL COMMENT 'タイトル',
    `description` VARCHAR(512) NOT NULL COMMENT 'タスク概要',
    `is_complete` TINYINT(1) NOT NULL COMMENT 'タスク完了判定',
    `created_at`  DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成時',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新時',
    PRIMARY KEY (`todo_id`)
);

INSERT INTO `user` (`user_id`,`password`,`token`,`user_name`,`mail`,`work`) VALUES ("1","password","token","test_name","test_email","test_work");
INSERT INTO `user` (`user_id`,`password`,`token`,`user_name`,`mail`,`work`) VALUES ("2","password","token","test_name","test_email","test_work");

INSERT INTO `tasks` (`id`,`todo_id`,`title`,`description`,`is_complete`) VALUES ("1","1","test_tiele","test_description",false);
INSERT INTO `tasks` (`id`,`todo_id`,`title`,`description`,`is_complete`) VALUES ("1","2","test_tiele","test_description",false);
INSERT INTO `tasks` (`id`,`todo_id`,`title`,`description`,`is_complete`) VALUES ("2","1","test_tiele","test_description",false);