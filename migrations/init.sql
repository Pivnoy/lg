CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

drop table if exists "user", administrator, country, city, citizenship,
    company, skill, skill_category, specialization, "role", team, achievement,
    employment, eduspeciality, university, profile, profile_skill, lineup, category, project cascade;

create table if not exists "user" (
                                      id serial primary key,
                                      uuid uuid default uuid_generate_v4(),
                                      email varchar(255) not null,
                                      "password" text not null
);

create table if not exists administrator (
                                             id serial primary key,
                                             user_uuid uuid not null references "user"(uuid),
                                             email varchar(255) unique not null,
                                             firstname varchar(255) not null,
                                             lastname varchar(255) not null,
                                             patronymic varchar(255) not null
);

create table if not exists country (
                                       id serial primary key,
                                       uuid uuid default uuid_generate_v4(),
                                       "name" varchar(255) not null,
                                       code varchar(2) unique not null
);

create table if not exists city (
                                    id serial primary key,
                                    uuid uuid default uuid_generate_v4(),
                                    "name" varchar(255) not null,
                                    country_uuid uuid not null references country(uuid)
);

create table if not exists citizenship (
                                           id serial primary key,
                                           uuid uuid default uuid_generate_v4(),
                                           "name" varchar(100) not null,
                                           country_uuid uuid not null references country(uuid)
);

create table if not exists company (
                                       id serial primary key,
                                       uuid uuid default uuid_generate_v4(),
                                       "name" varchar(255) not null,
                                       inn varchar(10) not null
);

create table if not exists skill_category (
                                              id serial primary key,
                                              uuid uuid default uuid_generate_v4(),
                                              "name" varchar(255) not null,
                                              "value" varchar(255) not null
);

create table if not exists skill (
                                     id serial primary key,
                                     uuid uuid default uuid_generate_v4(),
                                     "name" varchar(255) not null,
                                     "value" varchar(255) not null,
                                     skill_category_uuid uuid not null references skill_category(uuid)
);

create table if not exists specialization (
                                              id serial primary key,
                                              uuid uuid default uuid_generate_v4(),
                                              "name" varchar(255) not null,
                                              "value" varchar(255) not null
);

create table if not exists "role" (
                                      id serial primary key,
                                      uuid uuid default uuid_generate_v4(),
                                      "name" varchar(255) not null,
                                      specialization_uuid uuid references specialization(uuid)
);

create table if not exists team (
                                    id serial primary key,
                                    uuid uuid default uuid_generate_v4(),
                                    "name" varchar(255) not null,
                                    "value" varchar(255) not null
);

create table if not exists achievement (
                                           id serial primary key,
                                           uuid uuid default uuid_generate_v4(),
                                           "text" text not null
);

create table if not exists employment (
                                          id serial primary key,
                                          uuid uuid default uuid_generate_v4(),
                                          "name" varchar(15) not null,
                                          "value" varchar(15) not null
);

create table if not exists eduspeciality (
                                             id serial primary key,
                                             uuid uuid default uuid_generate_v4(),
                                             "name" varchar(255) not null,
                                             code varchar(8) not null
);

create table if not exists university (
                                          id serial primary key,
                                          uuid uuid default uuid_generate_v4(),
                                          "name" varchar(255) not null,
                                          city_uuid uuid not null references city(uuid)
);

create table if not exists profile (
                                       user_uuid uuid not null references "user"(uuid),
                                       firstname varchar(255) not null,
                                       lastname varchar(255) not null,
                                       patronymic varchar(255),
                                       country_uuid uuid not null references country(uuid),
                                       city_uuid uuid not null references city(uuid),
                                       citizenship_uuid uuid not null references citizenship(uuid),
                                       gender varchar(6) check ( gender in ('male', 'female', 'other')) not null,
                                       phone varchar(255) not null,
                                       email varchar(255) not null,
                                       university_uuid uuid references university(uuid),
                                       eduspeciality_uuid uuid references eduspeciality(uuid),
                                       graduation_year int check ( graduation_year > 1900 ),
                                       employment_uuid uuid not null references employment(uuid),
                                       experience int check ( experience >= 0 ) not null,
                                       achievement_uuid uuid unique not null references achievement(uuid),
                                       team_uuid uuid references team(uuid),
                                       specialization_uuid uuid not null references specialization(uuid),
                                       company_uuid uuid references company(uuid),
                                       uuid uuid default uuid_generate_v4()
);

create table if not exists profile_skill (
                                             id serial primary key,
                                             uuid uuid default uuid_generate_v4(),
                                             profile_uuid uuid not null references "user"(uuid),
                                             skill_uuid uuid not null references skill(uuid)
);

create table if not exists lineup (
                                      id serial primary key,
                                      uuid uuid default uuid_generate_v4(),
                                      team_uuid uuid not null references team(uuid),
                                      role_uuid uuid not null references "role"(uuid),
                                      profile_uuid uuid references "user"(uuid),
                                      project_uuid uuid not null references project(uuid)
);

create table if not exists category (
                                        id serial primary key,
                                        uuid uuid default uuid_generate_v4(),
                                        "name" varchar(255) not null
);

create table if not exists project (
                                       id serial primary key,
                                       uuid uuid default uuid_generate_v4(),
                                       "name" varchar(255) not null,
                                       description text not null,
                                       category_uuid uuid not null references category(uuid),
                                       project_link text not null,
                                       presentation_link text not null,
                                       creator_uuid uuid not null references "user"(uuid),
                                       lineup_uuid uuid references lineup(uuid),
                                       is_visible varchar(9) check ( is_visible in ('visible', 'invisible') ) not null
);