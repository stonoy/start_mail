-- +goose Up
create table email(
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    subject text not null,
    body text not null,
    sender uuid not null
    references users(id),
    reciever uuid not null
    references users(id)
);

-- +goose Down
drop table email;