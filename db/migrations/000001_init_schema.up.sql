CREATE TABLE IF NOT EXISTS "blog_posts" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "body" text NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp
);

COMMENT ON COLUMN "blog_posts"."body" IS 'Content of the blog post';