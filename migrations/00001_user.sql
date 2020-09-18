-- +goose Up
CREATE TABLE `user` (
  `id`             bigint(20)    NOT NULL AUTO_INCREMENT,
  `username`       varchar(50)   NOT NULL,
  `first_name`     varchar(100)  NOT NULL,
  `last_name`      varchar(100)  NOT NULL,
  `phone`          varchar(32)   NOT NULL,
  `email`          varchar(100)  NOT NULL,
  `password`       varchar(128)  NOT NULL,
  `last_login_at`  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`     timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at`     timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`) USING HASH
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `user_token` (
  `id`             bigint(20)   NOT NULL AUTO_INCREMENT,
  `token`          varchar(64)  NOT NULL,
  `ip`             varchar(40)  NOT NULL,
  `user_id`        bigint(20)   NOT NULL,
  `expire_at`      timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`     timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- +goose Down
DROP TABLE `user`;
DROP TABLE `user_token`;
