-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table public.baskets
(
    id          serial primary key,
    user_id     integer not null,
    products    jsonb,
    created_at  timestamp    not null default now(),
    updated_at  timestamp    not null default now(),
    constraint basket_user_id_fkey
        foreign key (user_id)
        references public.users (id)
        on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table public.baskets;
-- +goose StatementEnd
