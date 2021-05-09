package service

import (
	"code.lhc.org/we-blog/dal"
	"code.lhc.org/we-blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type UserService struct {
}

var (
	userDao = &dal.UserDao{}
)

func (u *UserService) LoginByPassword(c *gin.Context) (*model.User, error) {
	username := c.Query("username")
	password := c.Query("password")
	user, err := userDao.GetUserByName(c, username)
	if err != nil {
		fmt.Println("err is &v", err)
		return nil, err
	}
	if password != user.Password {
		return nil, fmt.Errorf("pass err")
	}
	return user, nil
}

func (u *UserService) RegisterByPassword(c *gin.Context) (int64, error) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Printf("username: %v  password%v", username, password)
	user := model.User{
		Name:       username,
		Password:   password,
		Data:       "{}",
		CreateTime: time.Now().UTC(),
		ModifyTime: time.Now().UTC(),
	}
	id, err := userDao.Insert(c, &user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
