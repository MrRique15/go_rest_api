package repositories

import (
	"errors"

	"main_api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersRepository struct {
	db DBHandlerUsers
}

type DBHandlerUsers interface {
	InitiateCollection()
	findOneEmail(email string) (models.User, error)
	insertOne(user models.User) (models.User, error)
	findOneID(id primitive.ObjectID) (models.User, error)
	updateOne(id primitive.ObjectID, user models.User) (models.User, error)
}

func NewUsersRepository(dbh DBHandlerUsers) *UsersRepository {
	usersRepository := UsersRepository{
		db: dbh,
	}

	usersRepository.db.InitiateCollection()

	return &usersRepository
}

// ------------------------------------------------- Functions ----------------------------
func (us UsersRepository) GetUserByEmail(email string) (models.User, error) {
	user, err := us.db.findOneEmail(email)

	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (us UsersRepository) GetUserByID(id primitive.ObjectID) (models.User, error) {
	user, err := us.db.findOneID(id)

	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (us UsersRepository) RegisterUser(newUser models.User) (models.User, error) {
	user, err := us.db.insertOne(newUser)

	if err != nil {
		return user, errors.New("error during user registration")
	}

	return user, nil
}

func (us UsersRepository) UpdateUser(user models.User) (models.User, error) {
	result, err := us.db.updateOne(user.ID, user)

	if err != nil {
		return result, errors.New("error during user update")
	}

	return result, nil
}
