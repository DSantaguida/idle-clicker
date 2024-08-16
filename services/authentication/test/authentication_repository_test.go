package test

import (
	"testing"

	"github.com/dsantaguida/idle-clicker/pkg/config"
	"github.com/dsantaguida/idle-clicker/pkg/idle_errors"
	"github.com/dsantaguida/idle-clicker/services/authentication/internal/db"
	"github.com/dsantaguida/idle-clicker/services/authentication/internal/models"
	"golang.org/x/net/context"
)

func Setup(t *testing.T) (*db.AuthenticationRepository, context.Context) {
	c, err := config.GetConfig("../config/", "config_test")
	if err != nil {
		t.Fatal("Failed to find test config: ", err)
	}

	db, err := db.CreateAuthenticationRepository(c.Db)
	if err != nil {
		t.Fatal("Failed to create database connection: ", err)
	}

	return db, context.TODO()
}

func TestCreateUser(t *testing.T) {
	db, ctx := Setup(t)
	defer db.Close()

	users := map[string]string {
		"dan": "san",
		"user1": "pw1",
		"user2": "pw2",
		"user3": "pw3",
		"user4": "pw4",
	}

	//Create users
	for k, v := range users {
		user := &models.User{Username: k, Password: v}
		_ = CreateBaseUser(user, ctx, db, t)
	}

	//Try to create user that already exists
	user := &models.User{Username: "dan", Password: ""}
	_, err := db.RegisterUser(ctx, user)
	if err != idle_errors.ErrUsernameTaken {
		t.Fatal("Failed to detect user that already exists.")
	}
}

func TestFindUser(t *testing.T) {
	db, ctx := Setup(t)

	//Create a base user
	user := &models.User{Username: "user123", Password: "password123"}
	_ = CreateBaseUser(user, ctx, db, t)

	//Find user that exists
	foundUser, err := db.FindUser(ctx, "user123")
	if err != nil {
		t.Fatal(err)
	}
	if foundUser == nil {
		t.Fatal("Failed to find user")
	}

	//Find user that doesn't exist
	foundUser, err = db.FindUser(ctx, "user124")
	if err != nil && err != idle_errors.ErrUserNotFound {
		t.Fatal(err)
	}
	if foundUser != nil && err == idle_errors.ErrUsernameTaken {
		t.Fatal("Found user that does not exist")
	}
}

 func TestUpdatePassword(t *testing.T) {
 	db, ctx := Setup(t)

	//Create a base user
	user := &models.User{Username: "username456", Password: "password456"}
	_ = CreateBaseUser(user, ctx, db, t)

	//Update base user value
	_, err := db.UpdateUserPassword(ctx, user, "password678")
	if err != nil {
		t.Fatal(err)
	}

	//Check base user new value
	newUser, err := db.FindUser(ctx, "username456")
	if err != nil {
		t.Fatal(err)
	}
	if newUser.Password != "password678" {
		t.Fatal("Failed to update user value")
	}
 }

 func CreateBaseUser(user *models.User, ctx context.Context, db *db.AuthenticationRepository, t *testing.T) (*models.User) {
	createdUser, err := db.RegisterUser(ctx, user)
	if err != nil {
		t.Fatal(err)
	}
	if createdUser == nil {
		t.Fatal("Failed to create users")
	}
	return createdUser
}