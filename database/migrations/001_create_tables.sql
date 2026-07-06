create table users
(
    id         uuid not null
        primary key,
    Name       varchar,
    email      varchar
        unique,
    cpf        varchar
        unique,
    phone      varchar,
    address    varchar,
    birth_date date,
    created_at timestamp,
    updated_at timestamp
);

alter table users
    owner to payment_user;

create table accounts
(
    id         uuid not null
        primary key,
    user_id    uuid
        references users
            deferrable,
    balance    numeric(15, 2),
    status     varchar,
    created_at timestamp,
    updated_at timestamp
);

alter table accounts
    owner to payment_user;

create table transactions
(
    id              uuid not null
        primary key,
    from_account_id uuid
        references accounts
            deferrable,
    to_account_id   uuid
        references accounts
            deferrable,
    amount          numeric(15, 2),
    transfer_type   varchar,
    status          varchar,
    metadata        jsonb,
    created_at      timestamp,
    completed_at    timestamp
);

alter table transactions
    owner to payment_user;