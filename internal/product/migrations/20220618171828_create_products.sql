-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table public.products
(
    id          serial primary key,
    name        varchar(255) not null,
    description varchar(255) not null,
    price       float        not null,
    image       varchar(255) not null,
    category    varchar(255) not null,
    subcategory varchar(255) not null,
    brand       varchar(255) not null,
    model       varchar(255) not null,
    color       varchar(255) not null,
    size        varchar(255) not null,
    material    varchar(255) not null,
    country     varchar(255) not null,
    created_at  timestamp    not null default now(),
    updated_at  timestamp    not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table public.products;
-- +goose StatementEnd
