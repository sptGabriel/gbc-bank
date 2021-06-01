begin;

CREATE TABLE "accounts" (
                           "id"                 varchar(255) primary KEY,
                           "name"               varchar(255) NOT NULL,
                           "cpf"                varchar(255) NOT NULL,
                           "secret"             varchar(255) NOT NULL,
                           "balance" 	        decimal NOT NULL,
                           "created_at" 		timestamptz NOT NULL DEFAULT (now()),
                           "updated_at" 		timestamptz NOT NULL DEFAULT (now())
);

commit;