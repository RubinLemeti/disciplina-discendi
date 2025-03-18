create table if not exists go_backend.users (
    id serial primary key,
    username varchar(255) not null,
    email varchar(255),
    password varchar(255) not null,
    created_at timestamp default (timezone('utc', now())),
    updated_at timestamp default (timezone('utc', now()))
);

