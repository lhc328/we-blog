package dal

import (
	"code.lhc.org/we-blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserDao struct {
}

func (a *UserDao) GetUserByName(c *gin.Context, username string) (*model.User, error) {
	var user model.User
	err := model.Db.Model(&model.User{}).Select("*").Where("name = (?)", username).First(&user).Error
	if err != nil {
		fmt.Printf("userdao get by name error name is %s", username)
		return nil, err
	}
	return &user, nil
}

func (a *UserDao) GetUserById(c *gin.Context, id int64) (*model.User, error) {
	var user model.User
	err := model.Db.Model(&model.User{}).Select("*").Where("id = (?)", id).First(&user).Error
	if err != nil {
		fmt.Printf("userdao get by id error id is %s", id)
		return nil, err
	}
	return &user, nil
}

func (a *UserDao) Insert(c *gin.Context, user *model.User) (int64, error) {
	err := model.Db.Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}
