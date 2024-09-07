package services

import (
	"github.com/JoseTheodoro/grpc-go/domain"
	pb "github.com/JoseTheodoro/grpc-go/person"
	"github.com/JoseTheodoro/grpc-go/repository"
)

func CreatePerson(p *pb.PersonRequest) *domain.Person {
	// implements logic create person

	person := &domain.Person{
		Name: p.GetName(),
		Age:  p.GetAge(),
		City: p.GetCity(),
	}

	PersonCreated := repository.CreatePerson(person)

	return PersonCreated
}
