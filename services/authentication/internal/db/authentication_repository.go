package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dsantaguida/idle-clicker/pkg/config"
	"github.com/dsantaguida/idle-clicker/pkg/idle_errors"
	"github.com/dsantaguida/idle-clicker/services/authentication/internal/models"
	"github.com/dsantaguida/idle-clicker/services/authentication/internal/password"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
)

type AuthenticationRepository struct {
	db *sql.DB;
}

func (authenticationRepository *AuthenticationRepository) Close(){
	authenticationRepository.db.Close()
}

func CreateAuthenticationRepository(dbConfig config.DBConfig) (*AuthenticationRepository, error){
	url := fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.SslMode)

	db, err := sql.Open(dbConfig.Driver, url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil{
		return nil, err
	}

	return &AuthenticationRepository{db: db}, nil
}

func (b *AuthenticationRepository) RegisterUser(ctx context.Context, user *models.User) (*models.User, error){
	err := password.CheckPasswordCompliance(user.Password)
	if err != nil {
		return nil, err
	}

	existingUser, err := b.FindUser(ctx, user.Username)
	if err != nil && err != idle_errors.ErrUserNotFound {
		return nil, err
	} else if existingUser != nil {
		return nil, idle_errors.ErrUsernameTaken
	}

	createdUser := &models.User{}
	query := fmt.Sprintf(createUserQuery, user.Username, user.Password)
	err = b.db.QueryRowContext(ctx, query).Scan(&createdUser.Id, &createdUser.Username, &createdUser.Password)
	if err != nil {
		return nil, errors.Wrap(err, "RegisterUser")
	}

	return createdUser, nil
}

func (b *AuthenticationRepository) FindUser(ctx context.Context, username string) (*models.User, error){
	query := fmt.Sprintf(findUserQuery, username)

	existingUser := &models.User{}
	err := b.db.QueryRowContext(ctx, query).Scan(&existingUser.Id, &existingUser.Username, &existingUser.Password)
	if err != nil {
		return nil, idle_errors.ErrUserNotFound
	}

	return existingUser, nil
}

func (b *AuthenticationRepository) UpdateUserPassword(ctx context.Context, user *models.User, newPassword string) (*models.User, error){
	_, err := b.FindUser(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{}
	query := fmt.Sprintf(updateUserQuery, newPassword, user.Username)
	err = b.db.QueryRowContext(ctx, query).Scan(&newUser.Id, &newUser.Username, &newUser.Password)
	if err != nil {
		return nil, errors.Wrap(err, "UpdateUserPassword")
	}

	return newUser, nil
}