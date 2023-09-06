USE douyin_user;

CREATE TABLE IF NOT EXISTS `user_info`
(
    `id`               BIGINT                  NOT NULL,
    `username`         VARCHAR(20)             NOT NULL UNIQUE,
    `password`         VARCHAR(255)            NOT NULL,
    `name`             VARCHAR(20)             NOT NULL COMMENT '用户名称',
    `avatar`           VARCHAR(255) DEFAULT '' NOT NULL COMMENT '用户头像',
    `background_image` VARCHAR(255) DEFAULT '' NOT NULL COMMENT '用户个人页顶部大图',
    `signature`        VARCHAR(200) DEFAULT '' NOT NULL COMMENT '个人简介',
    `work_count`       BIGINT       DEFAULT 0  NOT NULL COMMENT '作品数量',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE INDEX login_idx ON user_info(username, password, name);

CREATE TABLE IF NOT EXISTS `user_message`
(
    `id`          BIGINT       NOT NULL AUTO_INCREMENT,
    `user_id`     BIGINT       NOT NULL COMMENT '发送者',
    `to_user_id`  BIGINT       NOT NULL COMMENT '对方',
    `content`     VARCHAR(100) NOT NULL,
    `create_time` BIGINT       NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

ALTER TABLE user_message
    ADD INDEX uid_2uid_time (user_id, to_user_id, create_time);