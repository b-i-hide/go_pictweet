-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `users` (
		`user_id` INT NOT NULL AUTO_INCREMENT,
		`name` VARCHAR(255) NOT NULL COMMENT 'user name',
		`email` VARCHAR (255) NOT NULL COMMENT 'email',
		`is_session_alive` bool NOT NULL DEFAULT FALSE,
		`password_digest` VARCHAR(255) NOT NULL,
		`created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
		`updated_at` TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (`user_id`),
		UNIQUE KEY (`email`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `users`;

