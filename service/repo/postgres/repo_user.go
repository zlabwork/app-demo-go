package postgres

import (
	"app/entity"
	"context"
	"gorm.io/gorm"
	"time"
)

func NewUserRepository() (*UserRepository, error) {

	conn, err := getHandle()
	if err != nil {
		return nil, err
	}
	return &UserRepository{
		Conn: conn,
	}, nil
}

type UserRepository struct {
	Conn *gorm.DB
}

func (ur *UserRepository) GetOne(ctx context.Context, id int64) (*entity.User, error) {

	data := &entity.User{}
	err := ur.Conn.First(data, id).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ur *UserRepository) GetMany(ctx context.Context, id []int64) ([]*entity.User, error) {
	// TODO:
	return nil, nil
}

func (ur *UserRepository) Create(ctx context.Context, user *entity.User) error {

	user.CreatedAt = time.Now()
	return ur.Conn.Create(&user).Error
}

func (ur *UserRepository) Modify(ctx context.Context, user *entity.User) error {
	// TODO:
	return nil
}

func (ur *UserRepository) Delete(ctx context.Context, id int64) error {

	return ur.Conn.Delete(&entity.User{}, id).Error
}
