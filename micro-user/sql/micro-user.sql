CREATE TABLE `users`  (
  `id` VARCHAR(36) UNIQUE NOT NULL COMMENT 'UUID',
  `username` VARCHAR(64) UNIQUE NOT NULL COMMENT 'username',
  `password` VARCHAR(255) NOT NULL COMMENT 'argon2hash',
  `pk` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'auto incrementing primary key',
  PRIMARY KEY (`pk`),
  INDEX `idx_users_id`(`id`) USING BTREE,
  INDEX `idx_users_username`(`username`) USING BTREE,
);

CREATE TABLE `user_settings`  (
  `user_pk` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'pk of user',
  `pk` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'auto incrementing primary key',
  `summary_length` SMALLINT UNSIGNED NOT NULL DEFAULT 50,
  `summary_prompt` TEXT NOT NULL COMMENT 'summary prompt when call chatgpt',
  `emphasis_prompt` TEXT NOT NULL COMMENT 'emphasis prompt when call chatgpt',

  PRIMARY KEY (`pk`),
  INDEX `idx_user_profiles_user_pk`(`user_pk`) USING BTREE,
);

CREATE TABLE `summary`  (
  `user_pk` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'pk of user',
  `pk` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'auto incrementing primary key',
  `content` TEXT NOT NULL COMMENT 'summary content',

  PRIMARY KEY (`pk`),
  INDEX `idx_summary_user_pk`(`user_pk`) USING BTREE,
);

CREATE TABLE `emphasis`  (
  `user_pk` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'pk of user',
  `pk` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'auto incrementing primary key',
  `content` TEXT NOT NULL COMMENT 'emphasis content',

  PRIMARY KEY (`pk`),
  INDEX `idx_emphasis_user_pk`(`user_pk`) USING BTREE,
);

ALTER TABLE `user_settings` ADD CONSTRAINT `fk_user_settings` FOREIGN KEY (`user_pk`) REFERENCES `users` (`pk`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `summary` ADD CONSTRAINT `fk_summary` FOREIGN KEY (`user_pk`) REFERENCES `users` (`pk`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `emphasis` ADD CONSTRAINT `fk_emphasis` FOREIGN KEY (`user_pk`) REFERENCES `users` (`pk`) ON DELETE RESTRICT ON UPDATE RESTRICT;
