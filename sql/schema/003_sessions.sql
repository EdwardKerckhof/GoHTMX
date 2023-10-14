-- +goose Up
CREATE TABLE "sessions" (
  "id" UUID PRIMARY KEY,
  "user_id" UUID NOT NULL,
  "refresh_token" VARCHAR NOT NULL,
  "user_agent" VARCHAR(255) NOT NULL,
  "client_ip" VARCHAR(255) NOT NULL,
  "is_blocked" BOOLEAN NOT NULL DEFAULT FALSE,
  "expires_at" TIMESTAMPTZ NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE "sessions" ADD CONSTRAINT "sessions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;


-- +goose Down
DROP TABLE IF EXISTS "sessions";