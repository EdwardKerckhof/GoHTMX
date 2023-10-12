-- +goose Up
CREATE TABLE "todos" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "title" VARCHAR(255) NOT NULL,
  "completed" BOOLEAN NOT NULL DEFAULT FALSE,
  "user_id" UUID NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ
);

CREATE UNIQUE INDEX "todos_id_idx" ON "todos" ("id");
CREATE INDEX "todos_user_id_idx" ON "todos" ("user_id");
ALTER TABLE "todos" ADD CONSTRAINT "todos_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;


-- +goose Down
DROP TABLE IF EXISTS "todos";