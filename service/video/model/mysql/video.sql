USE douyin_video;

CREATE TABLE IF NOT EXISTS `video_info`
(
    `id`                BIGINT                  NOT NULL,
    `title`             VARCHAR(255)            NOT NULL COMMENT '视频标题',
    `play_url`          VARCHAR(255) DEFAULT '' NOT NULL COMMENT '视频播放地址',
    `cover_url`         VARCHAR(255) DEFAULT '' NOT NULL COMMENT '视频封面地址',
    `favorite_count`    BIGINT       DEFAULT 0  NOT NULL COMMENT '视频的点赞总数',
    `comment_count`     BIGINT       DEFAULT 0  NOT NULL COMMENT '视频的点赞总数',
    `user_id`           BIGINT                  NOT NULL COMMENT '外键',
    `publish_at`        BIGINT                  NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `video_comment`
(
    `id`        BIGINT       NOT NULL AUTO_INCREMENT,
    `user_id`   BIGINT       NOT NULL,
    `video_id`  BIGINT       NOT NULL,
    `content`   VARCHAR(255) NOT NULL,
    `create_at` BIGINT       NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;