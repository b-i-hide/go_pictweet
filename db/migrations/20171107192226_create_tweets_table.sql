-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `tweets` (
		`tweet_id` INT NOT NULL AUTO_INCREMENT,
		`image` VARCHAR (255) NOT NULL,
		`description` VARCHAR (255),
		`created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
		`updated_at` TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
		`user_id` INT NOT NULL,
		PRIMARY KEY (`tweet_id`),
		FOREIGN KEY (`user_id`) REFERENCES Users(`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `tweets`;
