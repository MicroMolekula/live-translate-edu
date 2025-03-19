package repository

import (
	"fmt"
	"github.com/live-translate-edu/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindAll() ([]*models.User, error) {
	var users []*models.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) FindOneById(id uint) (models.IModel, error) {
	var user models.IModel
	if err := ur.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) FindByIds(ids []uint) (map[uint]*models.User, error) {
	var users []*models.User
	if err := ur.db.Where("id in (?)", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	result := make(map[uint]*models.User)
	for _, user := range users {
		result[user.ID] = user
	}
	return result, nil
}

func (ur *UserRepository) FindOneByEmail(email string) (*models.User, error) {
	var user *models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) AddUserInGroup(group *models.Group, user *models.User) error {
	return ur.db.Model(user).Update("group_id", group.ID).Error
}

func (ur *UserRepository) AddUsersInGroupByIds(groupId int, usersIds []int) error {
	sql := ur.db.ToSQL(func(db *gorm.DB) *gorm.DB {
		return db.Model(&models.User{}).
			Where("id IN (?)", usersIds).
			Update("group_id", groupId)
	})
	fmt.Println(sql)
	return ur.db.Model(&models.User{}).
		Where("id IN ?", usersIds).
		Update("group_id", groupId).
		Error
}

func (ur *UserRepository) AddUsersInGroup(group *models.Group, users []*models.User) error {
	ids := make([]int, len(users))
	for i, user := range users {
		ids[i] = int(user.ID)
	}
	return ur.AddUsersInGroupByIds(int(group.ID), ids)
}
