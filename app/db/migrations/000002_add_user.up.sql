CREATE TABLE "users"
(
    "username"            varchar PRIMARY KEY,
    "password"            varchar UNIQUE NOT NULL,
    "full_name"           varchar        NOT NULL,
    "email"               varchar UNIQUE NOT NULL,
    "created_at"          timestamptz    NOT NULL DEFAULT (now()),
    "password_changed_at" timestamptz    NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE UNIQUE INDEX "accounts_owner_currency_idx" ON "accounts" ("owner", "currency");
-- the same solution
-- ALTER Table "accounts" add constraint "unique_owner_currency" unique ("owner", "currency");

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");