ALTER TABLE post DROP CONSTRAINT fk_post_user;
ALTER TABLE post_tags DROP CONSTRAINT fk_post_tags_tag;
ALTER TABLE post_tags DROP CONSTRAINT fk_post_tags_post;
ALTER TABLE post_category DROP CONSTRAINT fk_post_category_post;
ALTER TABLE post_category DROP CONSTRAINT fk_post_category_category;

DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS tag;
DROP TABLE IF EXISTS category;
DROP TABLE IF EXISTS post_tags;
DROP TABLE IF EXISTS post_category;