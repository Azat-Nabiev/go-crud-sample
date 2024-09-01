-- +goose Up
-- +goose StatementBegin
INSERT INTO users (name, surname, age, since) VALUES
('Azat', 'Nabiev', 23, '2023-01-01 12:00:00'),
('Jane', 'Jane', 25, '2022-06-15 08:30:00'),
('Alice', 'Alice', 35, '2021-03-22 17:45:00');

INSERT INTO books (name, author, user_id) VALUES
('some book', 'some author', 1),
('some book2', 'some author2', 1),
('some book3', 'some author3', 2),
('some book4', 'some author4', 3);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM books WHERE user_id IN (1, 2, 3);
DELETE FROM users WHERE id IN (1, 2, 3);
-- +goose StatementEnd
