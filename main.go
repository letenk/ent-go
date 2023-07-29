package main

import (
	"context"
	"fmt"
	"github.com/letenk/ent-go/ent"
	"github.com/letenk/ent-go/ent/user"
	"log"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=True", "root", "localhost", "3306", "ent_go")
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// create new user
	newUser, err := CreateUser(ctx, client)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("user was created: ", newUser)

	// query user
	user, err := QueryUser(ctx, client, newUser.ID)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Query user: ", user)

	// update email by id
	email := fmt.Sprintf("%d@mail.test", rand.Int())
	update, err := UpdateUserEmail(ctx, client, user.ID, email)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("updated email user request %s successfully, now in database value %s", email, update.Email)

}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	newUser, err := client.User.
		Create().
		SetAge(29).
		SetName("Rizky").
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return newUser, nil
}

func QueryUser(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {
	u, err := client.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func UpdateUserEmail(ctx context.Context, client *ent.Client, id int, email string) (*ent.User, error) {
	u, err := client.User.UpdateOneID(id).SetEmail(email).Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}

	return u, nil

}
