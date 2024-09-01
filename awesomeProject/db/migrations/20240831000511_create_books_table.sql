-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       author VARCHAR(255) NOT NULL,
                       user_id INT NOT NULL,
                       FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
