CREATE TABLE "file_cover" (
  "id" int8 NOT NULL,
  "user_id" int8 NOT NULL,
  "file_video_id" int8 NOT NULL,
  "objectname" varchar(255) NOT NULL,
  "url" varchar(255) NOT NULL,
  "upload_at" timestamp NOT NULL,
  PRIMARY KEY ("id")
);
COMMENT ON COLUMN "file_cover"."file_video_id" IS '作品id';

CREATE TABLE "file_user" (
  "id" int8 NOT NULL,
  "user_id" int8 NOT NULL,
  "type" int2 NOT NULL,
  "objectname" varchar(255) NOT NULL,
  "url" varchar(255) NOT NULL,
  "upload_at" varchar(255) NOT NULL,
  CONSTRAINT "_copy_2" PRIMARY KEY ("id")
);
COMMENT ON COLUMN "file_user"."type" IS '头像or背景图';

CREATE TABLE "file_video" (
  "id" int8 NOT NULL,
  "user_id" int8 NOT NULL,
  "video_id" int8 NOT NULL,
  "object_name" varchar(255) NOT NULL,
  "url" varchar(255) NOT NULL,
  "upload_at" timestamp NOT NULL,
  CONSTRAINT "_copy_1" PRIMARY KEY ("id")
);
COMMENT ON COLUMN "file_video"."video_id" IS '作品视频id';

