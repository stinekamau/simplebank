CREATE TABLE "users" (
    "username" varchar PRIMARY KEY,
    "hashed_password" varchar NOT NULL, 
    "full_name" varchar  NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password_changed_at" timestamptz DEFAULT '0001-01-01 0:00:00.00+10',                                            
    "created_at" timestamptz not null DEFAULT (now())

);

CREATE UNIQUE INDEX ON "accounts"  ("owner","currency");
ALTER TABLE "accounts" ADD FOREIGN KEY  ("owner") REFERENCES "users" ("username");





