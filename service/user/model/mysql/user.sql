USE douyin_user;

CREATE TABLE IF NOT EXISTS `user_info`
(
    `id`               BIGINT                  NOT NULL,
    `username`         VARCHAR(20)             NOT NULL,
    `password`         VARCHAR(255)            NOT NULL,
    `name`             VARCHAR(20)             NOT NULL COMMENT '用户名称',
    `follow_count`     BIGINT       DEFAULT 0 COMMENT '关注总数',
    `follower_count`   BIGINT       DEFAULT 0 COMMENT '粉丝总数',
    `avatar`           VARCHAR(255) DEFAULT '' NOT NULL COMMENT '用户头像',
    `background_image` VARCHAR(255) DEFAULT '' NOT NULL COMMENT '用户个人页顶部大图',
    `signature`        VARCHAR(200) DEFAULT '' NOT NULL COMMENT '个人简介',
    `total_favorited`  BIGINT       DEFAULT 0  NOT NULL COMMENT '获赞数量',
    `work_count`       BIGINT       DEFAULT 0  NOT NULL COMMENT '作品数量',
    `favorite_count`   BIGINT       DEFAULT 0  NOT NULL COMMENT '点赞数量',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;


CREATE TABLE IF NOT EXISTS `user_message`
(
    `id`          BIGINT       NOT NULL AUTO_INCREMENT,
    `user_id`     BIGINT       NOT NULL COMMENT '发送者',
    `to_user_id`  BIGINT       NOT NULL COMMENT '对方',
    `content`     VARCHAR(100) NOT NULL,
    `create_time` DATETIME     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;