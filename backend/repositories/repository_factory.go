package repositories

import "log"

type RepositoryFactoryStruct struct {
	repositories []IRepository
}

var RepositoryFactory *RepositoryFactoryStruct = &RepositoryFactoryStruct{
	repositories: make([]IRepository, 0),
}

func (factory *RepositoryFactoryStruct) Register(repository IRepository) {
	log.Printf("Repository registered: %s", repository.GetName())
	factory.repositories = append(factory.repositories, repository)
}
func (factory *RepositoryFactoryStruct) GetRepository(connectionString string) IRepository {
	errors := make([]error, 0)
	for _, repository := range factory.repositories {
		if repository.IsValidConnectionString(connectionString) {
			newRepo, err := repository.CreateRepository(connectionString)
			if err == nil {
				log.Printf("Repository created: %s - %s", newRepo.GetName(), connectionString)
				return newRepo
			}
			errors = append(errors, err)
		}
	}
	for _, err := range errors {
		log.Printf("Failed to create repository: %v", err)
	}
	return nil
}
