CREATE TABLE "user_info" (
  "id" int8 NOT NULL,
  "username" varchar(20) NOT NULL,
  "password" varchar(255) NOT NULL,
  "name" varchar(20) NOT NULL,
  "follow_count" int8,
  "follower_count" int8,
  "avatar" varchar(255) NOT NULL,
  "background_image" varchar(255) NOT NULL,
  "signature" varchar(200) NOT NULL,
  "total_favorited" int8 NOT NULL,
  "work_count" int8 NOT NULL,
  "favorite_count" int8 NOT NULL,
  PRIMARY KEY ("id")
);
COMMENT ON COLUMN "user_info"."name" IS '用户名称';
COMMENT ON COLUMN "user_info"."follow_count" IS '关注总数';
COMMENT ON COLUMN "user_info"."follower_count" IS '粉丝总数';
COMMENT ON COLUMN "user_info"."avatar" IS '用户头像';
COMMENT ON COLUMN "user_info"."background_image" IS '用户个人页顶部大图';
COMMENT ON COLUMN "user_info"."signature" IS '个人简介';
COMMENT ON COLUMN "user_info"."total_favorited" IS '获赞数量';
COMMENT ON COLUMN "user_info"."work_count" IS '作品数量';
COMMENT ON COLUMN "user_info"."favorite_count" IS '点赞数量';

CREATE TABLE "user_message" (
  "id" int8 NOT NULL,
  "user_id" int8 NOT NULL,
  "to_user_id" int8 NOT NULL,
  "content" varchar(255) NOT NULL,
  "create_at" timestamp NOT NULL,
  CONSTRAINT "_copy_2" PRIMARY KEY ("id")
);
COMMENT ON COLUMN "user_message"."user_id" IS '发送者';
COMMENT ON COLUMN "user_message"."to_user_id" IS '对方';

CREATE TABLE "user_relation" (
  "id" int8 NOT NULL,
  "user_id" int8 NOT NULL,
  "to_user_id" int8 NOT NULL,
  "create_at" timestamp NOT NULL,
  CONSTRAINT "_copy_1" PRIMARY KEY ("id")
);
COMMENT ON COLUMN "user_relation"."to_user_id" IS '关注的用户';

