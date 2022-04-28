CREATE TABLE IF NOT EXISTS "users" (
    id serial not null
        constraint users_pk
            primary key,
    name varchar not null,
    email varchar(100) not null,
    password varchar(100) not null,
    created_at timestamp not null,
    updated_at timestamp
);

CREATE INDEX "users_email_index" ON "users" (email);