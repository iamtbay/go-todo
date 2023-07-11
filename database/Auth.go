package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var dbTimeout = time.Second * 3

// LOGIN
func (op *User) Login(user *User) (*User, error) {
	
	//check session user exist or not


	//context part
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	//
	query := `SELECT id,name,email,password FROM users WHERE email=$1`

	//try to find email on db
	dbRes := dbConn.DB.QueryRowContext(ctx, query, user.Email)

	//create a var to get values
	var passControl User
	err := dbRes.Scan(
		&passControl.Id,
		&passControl.Name,
		&passControl.Email,
		&passControl.Password,
	)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("invalid email or password")
	}

	//compare password true or not
	err = bcrypt.CompareHashAndPassword([]byte(passControl.Password), []byte(user.Password))
	if err != nil {
		return nil, errors.New("invalid email or password comp")
	}

	//return user infos and error
	return &passControl, nil
}

// REGISTER
func (op *User) Register(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		panic("something went wrong while hashing the password")
	}

	query := `INSERT INTO users(name,email,password) VALUES($1,$2,$3)`
	dbRes, err := dbConn.DB.ExecContext(ctx, query, user.Name, user.Email, hashedPassword)
	if err != nil {
		return err
	}
	fmt.Println(dbRes.RowsAffected())
	return nil

}
