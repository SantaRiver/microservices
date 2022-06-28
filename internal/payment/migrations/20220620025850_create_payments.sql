-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table public.payments
(
    id          serial primary key,
    uuid        uuid           not null default uuid_generate_v4(),
    order_id    int            not null,
    status      int            not null,
    total_price numeric(10, 2) not null,
    currency    int            not null,
    created_at  timestamp               default now(),
    updated_at  timestamp               default now(),
    constraint payments_order_id_fkey foreign key (order_id) references public.orders (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table public.payments;
-- +goose StatementEnd
