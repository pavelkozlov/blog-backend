CREATE TABLE IF NOT EXISTS post
(
    id          uuid PRIMARY KEY   DEFAULT gen_random_uuid(),
    title       VARCHAR   NOT NULL,
    slug        VARCHAR,
    body        VARCHAR   NOT NULL,
    description VARCHAR,
    is_draft    BOOLEAN   NOT NULL DEFAULT false,
    author_id   uuid      NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP,
    CHECK (title <> ''),
    CHECK (body <> '')
);

CREATE TABLE IF NOT EXISTS "user"
(
    id         uuid PRIMARY KEY   DEFAULT gen_random_uuid(),
    email      VARCHAR   NOT NULL,
    name       VARCHAR,
    password   VARCHAR   NOT NULL,
    salt       VARCHAR   NOT NULL,
    version    NUMERIC   NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    is_admin   BOOLEAN   NOT NULL DEFAULT false,
    CHECK (email <> ''),
    CHECK (password <> ''),
    CHECK (salt <> '')
);

CREATE TABLE IF NOT EXISTS tag
(
    id         uuid PRIMARY KEY   DEFAULT gen_random_uuid(),
    title      VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    CHECK (title <> '')
);

CREATE TABLE IF NOT EXISTS category
(
    id          uuid PRIMARY KEY   DEFAULT gen_random_uuid(),
    title       VARCHAR   NOT NULL,
    description VARCHAR,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP,
    CHECK (title <> '')
);

CREATE TABLE IF NOT EXISTS post_tags
(
    id         uuid PRIMARY KEY   DEFAULT gen_random_uuid(),
    post_id    uuid      NOT NULL,
    tag_id     uuid      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS post_category
(
    id          uuid PRIMARY KEY   DEFAULT gen_random_uuid(),
    post_id     uuid      NOT NULL,
    category_id uuid      NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP
);

ALTER TABLE post_tags
    ADD CONSTRAINT fk_post_tags_post FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE;

ALTER TABLE post_tags
    ADD CONSTRAINT fk_post_tags_tag FOREIGN KEY (tag_id) REFERENCES tag (id) ON DELETE CASCADE;

ALTER TABLE post
    ADD CONSTRAINT fk_post_user FOREIGN KEY (author_id) REFERENCES "user" (id) ON DELETE CASCADE;

ALTER TABLE post_category
    ADD CONSTRAINT fk_post_category_post FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE;

ALTER TABLE post_category
    ADD CONSTRAINT fk_post_category_category FOREIGN KEY (category_id) REFERENCES category (id) ON DELETE CASCADE;

