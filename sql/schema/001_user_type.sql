-- +goose Up
create type user_type as enum ('user', 'admin');

-- +goose Down
drop type user_type;