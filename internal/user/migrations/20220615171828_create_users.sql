-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table public.users
(
    id         serial primary key,
    name       varchar(255) not null,
    surname    varchar(255) not null,
    patronymic varchar(255) not null,
    birthday   date         not null,
    type       smallint        not null,
    created_at timestamp    not null default current_timestamp,
    updated_at timestamp    not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table public.users;
-- +goose StatementEnd
