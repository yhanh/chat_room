-- up.sql help us create a new database or run commands that we have in here

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "username" varchar NOT NULL, 
    "email" varchar NOT NULL, 
    "password" varchar NOT NULL
)