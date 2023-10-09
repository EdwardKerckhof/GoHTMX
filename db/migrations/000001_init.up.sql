CREATE TABLE "todos" (
  "id" UUID PRIMARY KEY,
  "title" VARCHAR(255) NOT NULL,
  "completed" BOOLEAN NOT NULL DEFAULT FALSE,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ
);

CREATE INDEX "todos_deleted_at_idx" ON "todos" ("deleted_at");