begin;

CREATE TABLE accounts (
    id              uuid primary key,
    name            varchar(255) not null,
    cpf             varchar(255) not null,
    secret          varchar(255) not null,
    balance 	    decimal not null,
    created_at		timestamptz not null default (now()),
    updated_at 		timestamptz not null default (now())
);

CREATE TABLE  transfers (
    id                      uuid primary key,
    account_origin_id       uuid references accounts,
    account_destination_id  uuid references accounts,
    amount 	                decimal not null,
    created_at 		        timestamptz   not null default (now()),
    updated_at 		        imestamptz    not null default (now())
);

commit;