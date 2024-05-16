-- +goose Up

-- Extensions
create extension if not exists "pg_trgm";
create extension if not exists "uuid-ossp";

-- Users
create table if not exists "users"
(
    id  int primary key,
    login   varchar unique not null,
    password_hash    varchar(64) not null,
    role_id int
    );

-- +goose Down
drop extension if exists "pg_trgm";
drop extension if exists "uuid-ossp";
drop extension if exists "users";