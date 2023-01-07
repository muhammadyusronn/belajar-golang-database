package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	commet := entity.Comment{
		Email:   "yusron@gmail.com",
		Comment: "Tes Komentar",
	}
	result, err := commentRepository.Insert(ctx, commet)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindByID(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	comment, err := commentRepository.FindByID(context.Background(), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
