# Optimail

## How to install the project

### Step 1

Deploy your own email server, then change the email configuration in **mailagent.py**:

```py
IMAP_SERVER = "YOUR.IMAP.SREVER"
EMAIL = "YOUR_EMAIL_ADDR@TO.RECEIVE.EMAIL"
PASSWORD = getpass()
OPTIMAILAPI = "http://OPTIMAIL_HOST_ADDR/api"

#Example:
IMAP_SERVER = "imap.mailserver.com"
EMAIL = "admin@mailserver.com"
PASSWORD = getpass()
OPTIMAILAPI = "http://localhost:80/api"
```

### Step 2

Create database for micro-user:

```sql
create database micro_user;
use micro_user;

-- The following code is copied from Optimail/micro-user/sql/micro-user.sql

CREATE TABLE `users`  (
  `id` VARCHAR(36) UNIQUE NOT NULL COMMENT 'UUID',
  `username` VARCHAR(64) UNIQUE NOT NULL COMMENT 'username',
  `password` VARCHAR(512) NOT NULL COMMENT 'argon2hash',
  `pk` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'auto incrementing primary key',
  PRIMARY KEY (`pk`),
  INDEX `idx_users_id`(`id`) USING BTREE,
  INDEX `idx_users_username`(`username`) USING BTREE
);

CREATE TABLE `user_settings`  (
  `user_pk` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'pk of user',
  `pk` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'auto incrementing primary key',
  `email` VARCHAR(255) NOT NULL,
  `summary_length` SMALLINT UNSIGNED NOT NULL DEFAULT 50,
  `summary_prompt` TEXT NOT NULL COMMENT 'summary prompt when call chatgpt',
  `emphasis_prompt` TEXT NOT NULL COMMENT 'emphasis prompt when call chatgpt',

  PRIMARY KEY (`pk`),
  INDEX `idx_user_profiles_user_pk`(`user_pk`) USING BTREE
);

CREATE TABLE `summary`  (
  `user_pk` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'pk of user',
  `pk` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'auto incrementing primary key',
  `content` TEXT NOT NULL COMMENT 'summary content',

  PRIMARY KEY (`pk`),
  INDEX `idx_summary_user_pk`(`user_pk`) USING BTREE
);

CREATE TABLE `emphasis`  (
  `user_pk` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'pk of user',
  `pk` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'auto incrementing primary key',
  `content` TEXT NOT NULL COMMENT 'emphasis content',

  PRIMARY KEY (`pk`),
  INDEX `idx_emphasis_user_pk`(`user_pk`) USING BTREE
);

ALTER TABLE `user_settings` ADD CONSTRAINT `fk_user_settings` FOREIGN KEY (`user_pk`) REFERENCES `users` (`pk`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `summary` ADD CONSTRAINT `fk_summary` FOREIGN KEY (`user_pk`) REFERENCES `users` (`pk`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `emphasis` ADD CONSTRAINT `fk_emphasis` FOREIGN KEY (`user_pk`) REFERENCES `users` (`pk`) ON DELETE RESTRICT ON UPDATE RESTRICT;
```

### Step 3

Configure **server** and **micro-user**:

In `Optimail/server/config.toml`:

```toml
[server]
jwtkey = "Generated by `openssl rand -hex 32`"
[server.webpage]
base_path = "/go/src/server/webpage" # DO NOT MODIFY UNLESS YOU KNOW WHAT YOU ARE DOING
[server.gpt]
apiurl = "" # Leave blank to use default
model = "" # Leave blank to use default
apikey = "YOUR OPENAI_API_KEY"
[server.microuser]
grpc_addr = "localhost:50051" # grpc address for micro-user
```

In `Optimail/micro-user/config.toml`:

```toml
[server]
dsn = "username:password@tcp(your.mysql.endpoint)/micro_user"
grpc_addr = "localhost:50051" # grpc address for micro-user. Keep consistent with the server.microuser.grpc_addr
```

### Step 4

Launch everything by `./runall.sh`!

## How to clean docker?

After running, clean docker by `./cleandocker.sh`.

