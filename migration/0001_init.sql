-- +goose Up

-- Extensions
create extension if not exists "pg_trgm";
create extension if not exists "uuid-ossp";

-- Users

create table if not exists "passports"
(
    id serial primary key NOT NULL,
    passport_type varchar(255) not null,
    passport_number_hash varchar(255) not null
);

create table if not exists "departments"
(
    id serial primary key NOT NULL,
    department_name varchar(255) not null,
    department_phone varchar(255) not null
);

create table if not exists "users"
(
    id serial primary key NOT NULL,
    user_name varchar(255) not null,
    surname varchar(255) not null,
    phone varchar(255) not null,
    company_id int not null,
    passport_id int not null references passports(id),
    department_id int not null references departments(id)
);


-- +goose Down
drop extension if exists "pg_trgm";
drop extension if exists "uuid-ossp";
drop extension if exists "users";
drop extension if exists "passsports";
drop extension if exists "departments";