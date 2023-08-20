package main

import (
	"database/sql"
	"encoding/json"

	"github.com/LucBecker/imersao12-go-esquenta/internal/infra/akafka"
	"github.com/LucBecker/imersao12-go-esquenta/internal/infra/repository"
	"github.com/LucBecker/imersao12-go-esquenta/internal/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	repository := repository.NewProductRepositoryMysql(db)
	createProductUsecase := usecase.NewCreateProductUseCase(repository)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			// logar o erro
		}
		_, err = createProductUsecase.Execute(dto)
	}
}
