-- +migrate Up

-- 企業

create table if not exists companies (
	"id" bigserial primary key,
	"name" varchar(255) not null,
	"created_at" timestamp not null,
	"updated_at" timestamp not null
);

-- ユーザー

create table if not exists users (
	"id" bigserial primary key,
	"email" varchar(255) not null unique,
	"name" varchar(255) not null,
	"company_id" bigint references companies (id) on delete cascade on update cascade,
	"created_at" timestamp not null,
	"updated_at" timestamp not null
);

-- 記事

create table if not exists articles (
	"id" bigserial primary key,
	"title" varchar(255) not null,
	"content" text not null,
	"user_id" bigint references users (id) on delete
	set
		null on update
	set
		null,
		"created_at" timestamp not null,
		"updated_at" timestamp not null
);

-- +migrate Down

drop table if exists articles;

drop table if exists users;

drop table if exists companies;
