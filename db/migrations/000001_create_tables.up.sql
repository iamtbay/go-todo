CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name" VARCHAR(200) NOT NULL,
    "email" VARCHAR(200) NOT NULL UNIQUE,
    "password" VARCHAR NOT NULL,
    "createdAt" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS "todos" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "userId" uuid NOT NULL,
    "title" VARCHAR(100) NOT NULL,
    "body" TEXT NOT NULL,
    "isCompleted" BOOLEAN NOT NULL DEFAULT FALSE,
    "questTime" TIMESTAMP NOT NULL,
    "createdAt" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

ALTER TABLE "todos" ADD FOREIGN KEY ("userId") REFERENCES "users" ("id");
