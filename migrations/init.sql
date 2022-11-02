CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

drop table if exists "user", administrator, country, city, citizenship,
    company, skill, skill_category, specialization, "role", team, achievement,
    employment, eduspeciality, university, profile, profile_skill, lineup, category, project cascade;

create table if not exists "user" (
    id serial primary key,
    email varchar(255) unique not null,
    password varchar(255) not null
);

create table if not exists administrator (
    user_id bigint not null references "user"(id),
    email varchar(255) unique not null,
    firstname varchar(255) not null,
    lastname varchar(255) not null,
    patronymic varchar(255) not null,
    uuid uuid default uuid_generate_v4()
    );

create table if not exists country (
    id serial primary key,
    "name" varchar(255) not null,
    code varchar(2) unique not null
    );

create table if not exists city (
    id serial primary key,
    "name" varchar(255) not null,
    country_code varchar(2) not null references country(code)
    );

create table if not exists citizenship (
    id serial primary key,
    "name" varchar(100) not null,
    country_code varchar(2) not null references country(code)
    );

create table if not exists company (
    id serial primary key,
    "name" varchar(255) not null,
    inn bigint not null
    );

create table if not exists skill_category (
    id serial primary key,
    "name" varchar(255) not null,
    "value" varchar(255) not null
    );

create table if not exists skill (
    id serial primary key,
    "name" varchar(255) not null,
    "value" varchar(255) not null,
    skill_category_id bigint not null references skill_category(id)
    );

create table if not exists specialization (
    id serial primary key,
    "name" varchar(255) not null,
    "value" varchar(255) not null
    );

create table if not exists "role" (
    id serial primary key,
    "name" varchar(255) not null,
    specialization_id bigint references specialization(id)
    );

create table if not exists team (
    id serial primary key,
    "name" varchar(255) not null,
    "value" varchar(255) not null
    );

create table if not exists achievement (
    id serial primary key,
    "text" text not null
);

create table if not exists employment (
    id serial primary key,
    "name" varchar(15) not null,
    "value" varchar(15) not null
    );

create table if not exists eduspeciality (
    id serial primary key,
    "name" varchar(255) not null,
    code varchar(8) not null
    );

create table if not exists university (
    id serial primary key,
    "name" varchar(255) not null,
    city_id bigint not null references city(id)
    );

create table if not exists profile (
    user_id bigint not null references "user"(id),
    firstname varchar(255) not null,
    lastname varchar(255) not null,
    patronymic varchar(255),
    country_code varchar(2) not null references country(code),
    city_id bigint not null references city(id),
    citizenship_id bigint not null references citizenship(id),
    gender varchar(6) check ( gender in ('male', 'female', 'other')) not null,
    phone varchar(255) not null,
    email varchar(255) not null,
    university_id bigint references university(id),
    eduspeciality_id bigint references eduspeciality(id),
    graduation_year int check ( graduation_year > 1900 ),
    employment_id bigint not null references employment(id),
    experience int check ( experience >= 0 ) not null,
    achievement_id bigint unique not null references achievement(id),
    team_id bigint references team(id),
    specialization_id bigint not null references specialization(id),
    company_id bigint references company(id),
    uuid uuid default uuid_generate_v4()
    );

create table if not exists profile_skill (
    id serial primary key,
    profile_id bigint not null references "user"(id),
    skill_id bigint not null references skill(id)
    );

create table if not exists lineup (
    id serial primary key,
    team_id bigint not null references team(id),
    role_id bigint not null references "role"(id),
    profile_id bigint not null references "user"(id)
    );

create table if not exists category (
    id serial primary key,
    "name" varchar(255) not null
    );

create table if not exists project (
    id serial primary key,
    uuid uuid default uuid_generate_v4(),
    "name" varchar(255) not null,
    description text not null,
    category_id bigint not null references category(id),
    project_link text not null,
    presentation_link text not null,
    creator_id bigint not null references "user"(id),
    lineup_id bigint references lineup(id),
    is_visible varchar(9) check ( is_visible in ('visible', 'invisible') ) not null
    );
