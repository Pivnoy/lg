CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table project (
    id serial primary key,
    uuid uuid DEFAULT uuid_generate_v4(),
    name varchar(255) not null,
    description text not null,
    project_link varchar(255),
    presentation_link varchar(255),
    creator_id int not null             -- позже сделать references
);
