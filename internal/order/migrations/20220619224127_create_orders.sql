-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

create table public.orders
(
    id          serial primary key,
    uuid uuid not null default uuid_generate_v4(),
    user_id     int            not null,
    basket_id   int            not null,
    status      int            not null,
    total_price numeric(10, 2) not null,
    currency    int            not null,
    error       varchar(255),
    created_at  timestamp default now(),
    updated_at  timestamp default now(),
    constraint orders_user_id_fkey foreign key (user_id) references public.users (id),
    constraint orders_basket_id_fkey foreign key (basket_id) references public.baskets (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table public.orders;
-- +goose StatementEnd
