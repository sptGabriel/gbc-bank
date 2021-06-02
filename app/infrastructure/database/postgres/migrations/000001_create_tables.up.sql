begin;

CREATE TABLE accounts (
    id              UUID PRIMARY KEY,
    name            VARCHAR (255) NOT NULL,
    cpf             VARCHAR (255) NOT NULL,
    secret          VARCHAR (255) NOT NULL,
    balance 	    DECIMAL NOT NULL,
    created_at		timestamptz NOT NULL DEFAULT (now()),
    updated_at 		timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE  transfers (
    id                      UUID PRIMARY KEY,
    account_origin_id       UUID REFERENCES accounts (id),
    account_destination_id  UUID REFERENCES accounts (id),
    amount 	                DECIMAL NOT NULL,
    created_at 		        timestamptz NOT NULL DEFAULT (now()),
    updated_at 		        timestamptz NOT NULL DEFAULT (now())
);

commit;