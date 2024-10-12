package repositories

import (
	"awesomeProject/internal/models"
	"context"
	"database/sql"
	"go.uber.org/zap"
)

type UserRepositoryImpl struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewUserRepository(db *sql.DB, logger *zap.SugaredLogger) UserRepository {
	return &UserRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (repo *UserRepositoryImpl) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := "INSERT INTO users (name, surname, age, since) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRowContext(ctx, query, user.Name, user.Surname, user.Age, user.Since).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepositoryImpl) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, name, surname, age, since FROM users WHERE id = $1"
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Surname, &user.Age, &user.Since)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepositoryImpl) GetAllUserBooks(ctx context.Context) ([]models.User, error) {
	query := "SELECT id, name, surname, age, since FROM users"

	rows, err := repo.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Age, &user.Since); err != nil {
			return nil, err
		}

		bookQuery := "SELECT id, name, author, user_id FROM books where user_id = $1"
		bookRows, err := repo.db.QueryContext(ctx, bookQuery, user.ID)
		if err != nil {
			return nil, err
		}

		var books []models.Book
		for bookRows.Next() {
			var book models.Book
			if err := bookRows.Scan(&book.ID, &book.Name, &book.Author, &book.UserID); err != nil {
				return nil, err
			}
			books = append(books, book)
		}
		err = bookRows.Close()

		user.Books = books
		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepositoryImpl) GetAllUsers(ctx context.Context) ([]models.User, error) {
	query := "SELECT id, name, surname, age, since FROM users"

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Age, &user.Since); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepositoryImpl) UpdateUser(ctx context.Context, id int, user *models.User) (*models.User, error) {
	userPm := &models.User{}
	query := "SELECT id, name, surname, age, since FROM users where id = $1"
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Surname, &user.Age, &user.Since)
	if err != nil {
		return nil, err
	}

	userPm.Surname = user.Surname
	userPm.Name = user.Name
	userPm.Age = user.Age

	updateQuery := "UPDATE users SET name = $1, surname = $2, age = $3, since = $4 WHERE id = $5"
	_, err2 := repo.db.ExecContext(ctx, updateQuery, user.Name, user.Surname, user.Age, user.Since, user.ID)

	if err2 != nil {
		return nil, err2
	}
	return userPm, nil
}

func (repo *UserRepositoryImpl) DeleteUser(ctx context.Context, id int) (*models.User, error) {
	userPm := &models.User{}
	fetchQuery := "SELECT name FROM users where id = $1"
	deleteQuery := "DELETE FROM users where id = $1"

	err := repo.db.QueryRowContext(ctx, fetchQuery, id).Scan(&userPm.Name)
	if err != nil {
		return nil, err
	}

	_, err2 := repo.db.ExecContext(ctx, deleteQuery, id)

	if err2 != nil {
		return nil, err2
	}

	userPm.ID = id

	return userPm, nil
}
