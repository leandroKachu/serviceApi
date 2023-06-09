package repositories

import (
	"api/src/model"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

func NewRepositoryOfUsers(db *gorm.DB) *Users {
	return &Users{db}
}

func (repository *Users) Create(user model.User) (model.User, error) {
	user.Name = strings.ToLower(user.Name)
	user.Name = strings.ToLower(user.Nick)

	result := repository.db.Create(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (repository *Users) FindUsersByNameOrNick(nameOrNick string) ([]model.User, error) {

	var users []model.User

	newParam := "Ron"

	result := repository.db.Table("users").
		Select("id, name, nick, email, created_at").
		Where("name LIKE ?", newParam).
		Or("nick LIKE ?", newParam).
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, user := range users {
		// Faça o processamento desejado com cada usuário
		fmt.Println(user.ID, user.Name, user.Nick, user.Email, user.Created_at)
	}

	return users, nil

}

func (repository *Users) FindUserByID(ID uint64) (model.User, error) {

	var user model.User

	result := repository.db.First(&user, ID)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	fmt.Println(*result)
	return user, nil
}

func (repository *Users) UpdateUser(user model.User, ID uint64) (model.User, error) {

	query := "UPDATE users SET name = $1, nick = $2 WHERE id = $3"
	result := repository.db.Exec(query, user.Name, user.Nick, ID).Find(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return model.User{Name: user.Name, Nick: user.Nick}, nil
}

func (repository *Users) DeletedUser(ID uint64) string {

	query := "DELETE FROM users WHERE id = $1"

	result := repository.db.Exec(query, ID)

	if result.Error != nil {
		log.Fatal(result.Error)
		fmt.Println(result.Error)
	}

	return "sucessfully deleted"
}

func (repository *Users) CheckEmail(email string) (model.User, error) {

	var user model.User
	result := repository.db.Table("users").
		Select("id, password").
		Where("email = ?", email).
		Find(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
		fmt.Println(result.Error)
	}

	fmt.Println(user)
	return user, nil

}

func (repository *Users) FollowUser(current_userID, follow_userID uint64) error {

	query := "insert into followers (user_id, follower_id) values ($1, $2)"

	result := repository.db.Exec(query, current_userID, follow_userID)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil

}

func (repository *Users) UnFollowUser(current_userID, follow_userID uint64) error {
	query := "delete from followers where user_id = $1 and follower_id = $2"

	result := repository.db.Exec(query, current_userID, follow_userID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *Users) FindFollowers(follower_userID uint64) ([]model.User, error) {
	var user []model.User
	// essa queria de baixo existe um detalhe que impossibilitou de fazer a querie de forma desejada. usei left join e nao inner join
	// result := repository.db.Table("users").Select("users.id, users.name, users.nick, users.email").Joins("left join followers on  users.id = followers.follower_id").Where("follower_id = ?", follower_userID).Scan(&user)
	result := repository.db.Table("users u").Select("u.id, u.name, u.nick, u.email, u.created_at").Joins("inner join followers s on u.id = s.follower_id").Where("s.user_id = ?", follower_userID).Scan(&user)

	if result.Error != nil {
		return []model.User{}, result.Error
	}
	return user, nil
}

func (repository *Users) Following(userID uint64) ([]model.User, error) {
	var user []model.User

	result := repository.db.Table("users u").Select("u.id, u.name, u.nick, u.email, u.created_at").Joins("inner join followers s on u.id = s.user_id").Where("s.follower_id = ?", userID).Scan(&user)

	if result.Error != nil {
		return []model.User{}, result.Error
	}
	return user, nil

}

func (repository *Users) FindPassHash(userID uint64) (string, error) {
	var user model.User
	result := repository.db.Table("users").Select("password").Where("id = ?", userID).Find(&user)

	if result.Error != nil {
		return "", result.Error
	}
	return user.Password, nil
}

func (repository *Users) UpdatePass(pass string, userID uint64) (string, error) {
	var user model.User

	result := repository.db.Table("users").Where("id = ?", userID).Update("password", pass)
	if result.Error != nil {
		return user.Name, result.Error
	}
	selectuser := repository.db.Table("users").Select("name").Where("id = ?", userID).First(&user)

	if selectuser.Error != nil {
		return user.Name, result.Error
	}

	return user.Name, nil
}
