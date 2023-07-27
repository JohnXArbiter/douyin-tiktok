CREATE TABLE IF NOT EXISTS `file_video`
(
    `id`         BIGINT       NOT NULL,
    `user_id`    BIGINT       NOT NULL,
    `video_id`   BIGINT       NOT NULL COMMENT '作品视频id',
    `objectname` VARCHAR(255) NOT NULL,
    `url`        VARCHAR(255) NOT NULL,
    `upload_at`  DATETIME     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `file_user`
(
    `id`         BIGINT       NOT NULL,
    `user_id`    BIGINT       NOT NULL,
    `type`       tinyint      NOT NULL COMMENT '头像or背景图',
    `objectname` VARCHAR(255) NOT NULL,
    `url`        VARCHAR(255) NOT NULL,
    `upload_at`  VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `file_cover`
(
    `id`            BIGINT       NOT NULL,
    `user_id`       BIGINT       NOT NULL,
    `file_video_id` BIGINT       NOT NULL COMMENT '作品id',
    `objectname`    VARCHAR(255) NOT NULL,
    `url`           VARCHAR(255) NOT NULL,
    `upload_at`     DATETIME     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;