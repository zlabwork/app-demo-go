-- Adminer 4.8.1 PostgreSQL 15.2 dump

DROP TABLE IF EXISTS "zl_articles";
DROP SEQUENCE IF EXISTS articles_id_seq;
CREATE SEQUENCE articles_id_seq INCREMENT 1 MINVALUE 1000 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."zl_articles" (
    "id" integer DEFAULT nextval('articles_id_seq') NOT NULL,
    "state" smallint DEFAULT '0' NOT NULL,
    "category" integer DEFAULT '0' NOT NULL,
    "title" character varying(256) DEFAULT '' NOT NULL,
    "cover" character varying(128) DEFAULT '' NOT NULL,
    "content" text DEFAULT '' NOT NULL,
    "created_at" timestamp(0) DEFAULT '0001-01-01 00:00:00' NOT NULL,
    "updated_at" timestamp(0) DEFAULT '0001-01-01 00:00:00' NOT NULL,
    CONSTRAINT "articles_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "zl_categories";
DROP SEQUENCE IF EXISTS categories_id_seq;
CREATE SEQUENCE categories_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."zl_categories" (
    "id" integer DEFAULT nextval('categories_id_seq') NOT NULL,
    "parent" integer DEFAULT '0' NOT NULL,
    "sort" integer DEFAULT '0' NOT NULL,
    "state" smallint DEFAULT '0' NOT NULL,
    "name" character varying(32) DEFAULT '' NOT NULL,
    "cover" character varying(128) DEFAULT '' NOT NULL,
    "style" character varying(128) DEFAULT '' NOT NULL,
    "created_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    "updated_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    CONSTRAINT "categories_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "zl_users";
DROP SEQUENCE IF EXISTS users_id_seq;
CREATE SEQUENCE users_id_seq INCREMENT 1 MINVALUE 1000 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."zl_users" (
    "id" integer DEFAULT nextval('users_id_seq') NOT NULL,
    "alias" character varying(32) DEFAULT '' NOT NULL,
    "name" character varying(32) DEFAULT '' NOT NULL,
    "gender" smallint DEFAULT '0' NOT NULL,
    "phone" character varying(32) DEFAULT '' NOT NULL,
    "birth" date DEFAULT '0001-01-01' NOT NULL,
    "avatar" character varying(128) DEFAULT '' NOT NULL,
    "state" smallint DEFAULT '0' NOT NULL,
    "passwd" character varying(40) DEFAULT '' NOT NULL,
    "points" bigint DEFAULT '0' NOT NULL,
    "credit" bigint DEFAULT '0' NOT NULL,
    "created_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    "updated_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "zl_users_map";
DROP SEQUENCE IF EXISTS users_map_id_seq;
CREATE SEQUENCE users_map_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."zl_users_map" (
    "id" integer DEFAULT nextval('users_map_id_seq') NOT NULL,
    "user_id" integer DEFAULT '0' NOT NULL,
    "map_type" smallint DEFAULT '0' NOT NULL,
    "map_id" character varying(32) DEFAULT '' NOT NULL,
    "data" character varying(128) DEFAULT '' NOT NULL,
    "created_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    "updated_at" timestamp DEFAULT '0001-01-01 00:00:00' NOT NULL,
    CONSTRAINT "users_map_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "users_map_uniq" UNIQUE ("map_id", "map_type")
) WITH (oids = false);


-- 2023-05-06 09:48:20.648809+08
