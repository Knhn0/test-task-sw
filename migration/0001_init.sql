-- +goose Up

-- Extensions
create extension if not exists "pg_trgm";
create extension if not exists "uuid-ossp";

-- Users
create table if not exists "users"
(
    id  int primary key,
    user_name varchar(255),
    surname varchar(255),
    phone varchar(255),
    company_id int,
    passport_id uuid unique default uuid_generate_v4(),
    department_id uuid unique default uuid_generate_v4()
    );

create table if not exists "passports"
(
    id uuid unique default uuid_generate_v4(),
    passport_name varchar(255),
    passport_number_hash varchar(255)
);

create table if not exists "departments"
(
    id uuid unique default uuid_generate_v4()
    department_name varchar(255),
    deaprtment_phone varchar(255)
);

-- +goose Down
drop extension if exists "pg_trgm";
drop extension if exists "uuid-ossp";
drop extension if exists "users";
drop extension if exists "passsports";
drop extension if exists "departments";