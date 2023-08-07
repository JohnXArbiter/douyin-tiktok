CREATE TABLE "video_comment" (
  "id" int8 NOT NULL,
  "user_id" int8 NOT NULL,
  "video_id" int8 NOT NULL,
  "content" varchar(255) NOT NULL,
  "create_date" timestamp NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "video_favorite" (
  "id" int8 NOT NULL,
  "user_id" int8 NOT NULL,
  "video_id" int8 NOT NULL,
  "create_at" timestamp NOT NULL,
  CONSTRAINT "_copy_2" PRIMARY KEY ("id")
);
COMMENT ON COLUMN "video_favorite"."video_id" IS '喜欢的视频';

CREATE TABLE "video_info" (
  "id" int8 NOT NULL,
  "title" varchar(255) NOT NULL,
  "play_url" varchar(255) NOT NULL,
  "cover_url" varchar(255) NOT NULL,
  "favorite_count" int8 NOT NULL,
  "comment_count" int8 NOT NULL,
  "user_id" int8 NOT NULL,
  "publish_at" timestamp NOT NULL,
  CONSTRAINT "_copy_1" PRIMARY KEY ("id")
);
COMMENT ON COLUMN "video_info"."title" IS '视频标题';
COMMENT ON COLUMN "video_info"."play_url" IS '视频播放地址';
COMMENT ON COLUMN "video_info"."cover_url" IS '视频封面地址';
COMMENT ON COLUMN "video_info"."favorite_count" IS '视频的点赞总数';
COMMENT ON COLUMN "video_info"."comment_count" IS '视频的点赞总数';
COMMENT ON COLUMN "video_info"."user_id" IS '外键';

