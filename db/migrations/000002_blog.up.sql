CREATE TABLE IF NOT EXISTS "blog" (
    id serial not null
        constraint blog_pk
            primary key,
    created_by int not null
        constraint created_by_user_fk
            references "users"
                on delete RESTRICT on update cascade,
    title varchar not null,
    body text not null,
    image varchar not null,
    status smallint not null default 0,
    created_at timestamp not null,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE INDEX "blog_created_index" ON "blog" (created_by);

CREATE INDEX "blog_status_index" ON "blog" (status DESC);