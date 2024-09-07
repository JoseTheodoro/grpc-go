package repository

import (
	"github.com/JoseTheodoro/grpc-go/domain"
	"github.com/google/uuid"
)

func CreatePerson(person *domain.Person) *domain.Person {
	// implements logic persist person

	return &domain.Person{
		ID:   uuid.NewString(),
		Name: person.Name,
		Age:  person.Age,
		City: person.City,
	}
}
