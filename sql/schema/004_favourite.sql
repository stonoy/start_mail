-- +goose Up
create table favourite(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    emailid uuid not null
    references email(id)
    on delete cascade,
    userid uuid not null
    references users(id)
    on delete cascade,
    unique(emailid, userid)
);

-- +goose Down
drop table favourite;