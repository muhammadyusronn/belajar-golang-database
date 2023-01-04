package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExcSQL(t *testing.T) {
	t.Skip()
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	qwr := "INSERT INTO customer  (id, name, email, balance, rating, birth_date, married) VALUES ('09013', 'Nadia Lutviana', NULL, 100000, 4.9, NULL, false)"
	_, err := db.ExecContext(ctx, qwr)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

}

func TestSelectSQL(t *testing.T) {
	t.Skip()
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	qwr := "SELECT id, name, email, balance, rating, birth_date, married, created_at  FROM customer"
	rows, err := db.QueryContext(ctx, qwr)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("-----------------------------------------")
		fmt.Println("ID :", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("Email :", email)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		if birth_date.Valid {
			fmt.Println("Birth Date :", birth_date)
		}
		fmt.Println("Married :", married)
		fmt.Println("Created At :", created_at)
		fmt.Println("-----------------------------------------")

	}
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	pass := "admin"
	qwr := "SELECT username  FROM users WHERE username = '" + username + "' AND password ='" + pass + "' LIMIT 1"
	fmt.Println(qwr)
	rows, err := db.QueryContext(ctx, qwr)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success Login! Selamat datang", username)
	} else {
		fmt.Println("Gagal login!")
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	pass := "admin"
	qwr := "SELECT username  FROM users WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(qwr)
	rows, err := db.QueryContext(ctx, qwr, username, pass)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success Login! Selamat datang", username)
	} else {
		fmt.Println("Gagal login!")
	}
}

func TestExcSafeSQL(t *testing.T) {
	t.Skip()
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "Pimpinan'; DELETE FROM users where username='Pimpinan'; #"
	pass := "pimpinan"

	qwr := "INSERT INTO users  (username, password) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, qwr, username, pass)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new users")

}
