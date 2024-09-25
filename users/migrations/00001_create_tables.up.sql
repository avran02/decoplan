CREATE TABLE IF NOT EXISTS "groups"(
    "id" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "avatar_url" VARCHAR(255),
    PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "users"(
    "id" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "avatar_url" VARCHAR(255),
    "birth_date" TIMESTAMP(0) WITHOUT TIME ZONE,
    PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "user_groups"(
    "user_id" VARCHAR(255) NOT NULL,
    "group_id" VARCHAR(255) NOT NULL,
    PRIMARY KEY ("user_id", "group_id"),
    FOREIGN KEY("group_id") REFERENCES "groups"("id") ON DELETE CASCADE,
    FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE
);
