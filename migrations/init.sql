create table project (
    id serial primary key,
    name varchar(255) not null,
    description text not null,
    link varchar(255),
    presentation varchar(255)
);
