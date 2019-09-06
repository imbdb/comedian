-- +goose Up
-- +goose StatementBegin
CREATE TABLE `notifications_thread` (
    `id` INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `channel_id` VARCHAR(255) NOT NULL,
    `real_name` VARCHAR(255) NOT NULL,
    `notification_time` DATETIME NOT NULL,
    `reminder_counter` INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `notifications_thread`;
-- +goose StatementEnd