-- Adminer 4.8.1 PostgreSQL 15.2 dump

DROP TABLE IF EXISTS "bi_apps";
DROP SEQUENCE IF EXISTS bi_apps_id_seq;
CREATE SEQUENCE bi_apps_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."bi_apps" (
    "id" integer DEFAULT nextval('bi_apps_id_seq') NOT NULL,
    "app_id" integer DEFAULT '0' NOT NULL,
    "app_name" character varying DEFAULT '' NOT NULL,
    "sort" integer DEFAULT '0' NOT NULL,
    "state" integer DEFAULT '0' NOT NULL,
    "created_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    "updated_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    CONSTRAINT "bi_apps_app_id" UNIQUE ("app_id"),
    CONSTRAINT "bi_apps_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "bi_categories";
DROP SEQUENCE IF EXISTS bi_categories_id_seq;
CREATE SEQUENCE bi_categories_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."bi_categories" (
    "id" integer DEFAULT nextval('bi_categories_id_seq') NOT NULL,
    "app_id" integer DEFAULT '0' NOT NULL,
    "parent" integer DEFAULT '0' NOT NULL,
    "type" integer DEFAULT '0' NOT NULL,
    "sort" integer DEFAULT '0' NOT NULL,
    "name" character varying(32) DEFAULT '' NOT NULL,
    "state" smallint DEFAULT '0' NOT NULL,
    "created_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    CONSTRAINT "bi_categories_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "bi_categories_app_id" ON "public"."bi_categories" USING btree ("app_id");


DROP TABLE IF EXISTS "bi_channels";
DROP SEQUENCE IF EXISTS bi_channels_id_seq;
CREATE SEQUENCE bi_channels_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."bi_channels" (
    "id" integer DEFAULT nextval('bi_channels_id_seq') NOT NULL,
    "app_id" integer DEFAULT '0' NOT NULL,
    "channel_id" integer DEFAULT '0' NOT NULL,
    "title" character varying(32) DEFAULT '' NOT NULL,
    "created_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    CONSTRAINT "bi_channels_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "bi_channels_app_id" ON "public"."bi_channels" USING btree ("app_id");


DROP TABLE IF EXISTS "bi_events";
DROP SEQUENCE IF EXISTS bi_events_id_seq;
CREATE SEQUENCE bi_events_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."bi_events" (
    "id" integer DEFAULT nextval('bi_events_id_seq') NOT NULL,
    "app_id" integer DEFAULT '0' NOT NULL,
    "category" integer DEFAULT '0' NOT NULL,
    "state" smallint DEFAULT '0' NOT NULL,
    "event_id" character varying(64) DEFAULT '' NOT NULL,
    "event_name" character varying(64) DEFAULT '' NOT NULL,
    "remark" character varying(64) DEFAULT '' NOT NULL,
    "policy" character varying(2048) DEFAULT '' NOT NULL,
    "created_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    "updated_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    CONSTRAINT "bi_events_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

COMMENT ON TABLE "public"."bi_events" IS '事件管理';

COMMENT ON COLUMN "public"."bi_events"."app_id" IS 'APP';

COMMENT ON COLUMN "public"."bi_events"."category" IS '分类';

COMMENT ON COLUMN "public"."bi_events"."state" IS '状态';

COMMENT ON COLUMN "public"."bi_events"."event_id" IS '事件ID';

COMMENT ON COLUMN "public"."bi_events"."event_name" IS '事件标题';

COMMENT ON COLUMN "public"."bi_events"."remark" IS '备注';

COMMENT ON COLUMN "public"."bi_events"."policy" IS '事件规则';

COMMENT ON COLUMN "public"."bi_events"."created_at" IS '创建时间';

COMMENT ON COLUMN "public"."bi_events"."updated_at" IS '修改时间';


DROP TABLE IF EXISTS "bi_keys";
DROP SEQUENCE IF EXISTS bi_keys_id_seq1;
CREATE SEQUENCE bi_keys_id_seq1 INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."bi_keys" (
    "id" integer DEFAULT nextval('bi_keys_id_seq1') NOT NULL,
    "app_id" integer DEFAULT '0' NOT NULL,
    "key_id" character varying(32) DEFAULT '' NOT NULL,
    "key_name" character varying(32) DEFAULT '' NOT NULL,
    "key_type" smallint DEFAULT '0' NOT NULL,
    "created_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    CONSTRAINT "bi_keys_app_id_key_id" UNIQUE ("app_id", "key_id"),
    CONSTRAINT "bi_keys_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

COMMENT ON TABLE "public"."bi_keys" IS 'Key值管理';

COMMENT ON COLUMN "public"."bi_keys"."app_id" IS 'APP';

COMMENT ON COLUMN "public"."bi_keys"."key_id" IS 'KeyID';

COMMENT ON COLUMN "public"."bi_keys"."key_name" IS 'Key名称';

COMMENT ON COLUMN "public"."bi_keys"."key_type" IS 'Key类型';

COMMENT ON COLUMN "public"."bi_keys"."created_at" IS '创建时间';


DROP TABLE IF EXISTS "bi_label_name";
DROP SEQUENCE IF EXISTS bi_label_name_id_seq;
CREATE SEQUENCE bi_label_name_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."bi_label_name" (
    "id" integer DEFAULT nextval('bi_label_name_id_seq') NOT NULL,
    "app_id" integer DEFAULT '0' NOT NULL,
    "category" integer DEFAULT '0' NOT NULL,
    "label_id" character varying(32) DEFAULT '' NOT NULL,
    "label_title" character varying(32) DEFAULT '' NOT NULL,
    "label_data" character varying(256) DEFAULT '' NOT NULL,
    "remark" character varying(64) DEFAULT '' NOT NULL,
    "created_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    "updated_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    CONSTRAINT "bi_label_name_app_id_label_id" UNIQUE ("app_id", "label_id"),
    CONSTRAINT "bi_label_name_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

COMMENT ON COLUMN "public"."bi_label_name"."app_id" IS 'APP';

COMMENT ON COLUMN "public"."bi_label_name"."category" IS '分类';

COMMENT ON COLUMN "public"."bi_label_name"."label_id" IS '标签ID';

COMMENT ON COLUMN "public"."bi_label_name"."label_title" IS '标签Title';

COMMENT ON COLUMN "public"."bi_label_name"."label_data" IS '标签数据';

COMMENT ON COLUMN "public"."bi_label_name"."remark" IS '备注';

COMMENT ON COLUMN "public"."bi_label_name"."created_at" IS '创建时间';

COMMENT ON COLUMN "public"."bi_label_name"."updated_at" IS '更新时间';


DROP TABLE IF EXISTS "bi_label_tasks";
DROP SEQUENCE IF EXISTS bi_label_tasks_id_seq;
CREATE SEQUENCE bi_label_tasks_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."bi_label_tasks" (
    "id" integer DEFAULT nextval('bi_label_tasks_id_seq') NOT NULL,
    "app_id" integer DEFAULT '0' NOT NULL,
    "state" smallint DEFAULT '0' NOT NULL,
    "type" smallint DEFAULT '0' NOT NULL,
    "title" character varying(128) DEFAULT '' NOT NULL,
    "remark" character varying(128) DEFAULT '' NOT NULL,
    "scheduled_at" character varying(64) DEFAULT '' NOT NULL,
    "policy" character varying(4096) DEFAULT '' NOT NULL,
    "created_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    "updated_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    CONSTRAINT "bi_label_tasks_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

COMMENT ON COLUMN "public"."bi_label_tasks"."app_id" IS 'APP';

COMMENT ON COLUMN "public"."bi_label_tasks"."state" IS '状态';

COMMENT ON COLUMN "public"."bi_label_tasks"."type" IS '分类';

COMMENT ON COLUMN "public"."bi_label_tasks"."title" IS '主题';

COMMENT ON COLUMN "public"."bi_label_tasks"."remark" IS '备注';

COMMENT ON COLUMN "public"."bi_label_tasks"."scheduled_at" IS '定时';

COMMENT ON COLUMN "public"."bi_label_tasks"."policy" IS '规则';

COMMENT ON COLUMN "public"."bi_label_tasks"."created_at" IS '创建时间';

COMMENT ON COLUMN "public"."bi_label_tasks"."updated_at" IS '更新时间';


DROP TABLE IF EXISTS "bi_st_dnu";
DROP SEQUENCE IF EXISTS bi_st_dnu_id_seq;
CREATE SEQUENCE bi_st_dnu_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."bi_st_dnu" (
    "id" integer DEFAULT nextval('bi_st_dnu_id_seq') NOT NULL,
    "app_id" integer DEFAULT '0' NOT NULL,
    "channel_id" integer DEFAULT '0' NOT NULL,
    "sdk_ver" character varying(32) DEFAULT '' NOT NULL,
    "os" smallint DEFAULT '0' NOT NULL,
    "key" date DEFAULT '0001-01-01' NOT NULL,
    "count" integer DEFAULT '0' NOT NULL,
    "created_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    CONSTRAINT "bi_st_day_new_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "bi_st_dnu_app_id_key" ON "public"."bi_st_dnu" USING btree ("app_id", "key");

COMMENT ON COLUMN "public"."bi_st_dnu"."app_id" IS 'APP ID';

COMMENT ON COLUMN "public"."bi_st_dnu"."channel_id" IS '渠道';

COMMENT ON COLUMN "public"."bi_st_dnu"."sdk_ver" IS 'SDK版本';

COMMENT ON COLUMN "public"."bi_st_dnu"."os" IS '系统类型';

COMMENT ON COLUMN "public"."bi_st_dnu"."key" IS '日期';

COMMENT ON COLUMN "public"."bi_st_dnu"."count" IS '计数';

COMMENT ON COLUMN "public"."bi_st_dnu"."created_at" IS '创建时间';


-- 2023-07-19 16:57:47.200957+08
