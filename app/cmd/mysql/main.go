package main

import (
	"fmt"
	"log"
	"time"
	//"database/sql"

	userDomain "github.com/maitakedayo/go-dbmanager/app/domain/user"
	userAppli "github.com/maitakedayo/go-dbmanager/app/application/user"
	userRepository "github.com/maitakedayo/go-dbmanager/app/infrastructure/mysql/repository"
	_ "github.com/lib/pq" // PostgreSQL用サードパーティードライバ
)


func main() {
	// PostgreSQLを別途立ち上げてね(GUIのpgAdminアプリなど)
	// Initialize repository with the PostgreSQL database
	postgreSQLUserRepository, err := userRepository.NewPostgreSQLUserRepository()
	if err != nil {
		log.Fatal("Error initializing PostgreSQL repository:", err)
	}
	defer postgreSQLUserRepository.DB.Close()

	// Initialize service with the repository
	userService := userDomain.NewUserService(postgreSQLUserRepository)

	// Initialize use cases with the service
	saveUserUseCase := userAppli.NewSaveUserUseCase(userService)
	findUserUseCase := userAppli.NewFindUserUseCase(userService)
	findAllUsersUseCase := userAppli.NewFindAllUsersUseCase(userService)

	// Save user
	saveDto := userAppli.SaveUseCaseDto{
		LastName:  "Smith",
		FirstName: "John",
		Email:     "john.smith@example.com",
		Posts:     []string{"Post1", "Post2"}, // Change: Add posts
		Idlimit:   time.Now().AddDate(50, 0, 0),
	}
	err = saveUserUseCase.Run(saveDto)
	if err != nil {
		log.Fatal("Error saving user:", err)
	}

	// Save another user
	saveDto2 := userAppli.SaveUseCaseDto{
		LastName:  "Doe",
		FirstName: "Jane",
		Email:     "jane.doe@example.com",
		Posts:     []string{"Post1", "Post2"}, // Change: Add posts
		Idlimit:   time.Now().AddDate(50, 0, 0),
	}
	err = saveUserUseCase.Run(saveDto2)
	if err != nil {
		log.Fatal("Error saving user:", err)
	}

	// Find user by full name
	findDto := userAppli.FindUseCaseDto{
		LastName:  "Smith",
		FirstName: "John",
	}
	foundUserDto, err := findUserUseCase.Run(findDto)
	if err != nil {
		log.Fatal("Error finding user:", err)
	}
	fmt.Println("===Found User:===")
	fmt.Printf("ID: %s\nLastName: %s\nFirstName: %s\nEmail: %s\nPosts: %v\nIdlimit: %s\n",
		foundUserDto.ID, foundUserDto.LastName, foundUserDto.FirstName, foundUserDto.Email, foundUserDto.Posts, foundUserDto.Idlimit)

	// Fetch and display all users
	allUsersDto, err := findAllUsersUseCase.Run()
	if err != nil {
		log.Fatal("Error fetching all users:", err)
	}
	fmt.Println("\n===All Users:===")
	for _, userDto := range allUsersDto.Users {
		fmt.Printf("ID: %s\nLastName: %s\nFirstName: %s\nEmail: %s\nPosts: %v\nIdlimit: %s\n",
			userDto.ID, userDto.LastName, userDto.FirstName, userDto.Email, userDto.Posts, userDto.Idlimit)
	}
}