package test

import (
	"testing"

	"github.com/dsantaguida/idle-clicker/pkg/config"
	"github.com/dsantaguida/idle-clicker/pkg/idle_errors"
	"github.com/dsantaguida/idle-clicker/services/authentication/internal/db"
	"github.com/dsantaguida/idle-clicker/services/authentication/internal/models"
	"github.com/dsantaguida/idle-clicker/services/authentication/internal/password"
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
		"dan": "sandan585",
		"username1": "password1",
		"username2": "password2",
		"username3": "password3",
		"username4": "password4",
	}

	//Create users
	for k, v := range users {
		user := &models.User{Username: k, Password: v}
		_ = CreateBaseUser(user, ctx, db, t)
	}

	//Try to create user that already exists
	user := &models.User{Username: "dan", Password: "password1234"}
	_, err := db.RegisterUser(ctx, user)
	if err != idle_errors.ErrUsernameTaken {
		t.Fatal("Failed to detect user that already exists", err)
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

	newPassword := "password678"
	_, err := db.UpdateUserPassword(ctx, user, newPassword)
	if err != nil {
		t.Fatal(err)
	}

	newUser, err := db.FindUser(ctx, "username456")
	if err != nil {
		t.Fatal(err)
	}
	if !password.VerifyPassword(newPassword, newUser.Password) {
		t.Fatal("Failed to update user value")
	}
 }

 func TestLogin(t *testing.T) {
	db, ctx := Setup(t)

   //Create a base user
   user := &models.User{Username: "username456", Password: "password456"}
   _ = CreateBaseUser(user, ctx, db, t)

   _, err := db.UserLogin(ctx, user)
   if err != nil {
	   t.Fatal(err)
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