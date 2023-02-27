CREATE TABLE "user" (
  "userrname" varchar PRIMARY KEY,
  "hash_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varcar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT (0001-01-01 00:00:00Z),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

--CREATE UNIQUE INDEX ON "account" ("owner", "currency");
ALTER TABLE "account" ADD FOREIGN KEY ("owner") REFERENCES "user" ("userrname");
ALTER TABLE "account" ADD CONSTRAINT ("owner_currency_key") UNIQUE ("userrname");