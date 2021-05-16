package service

import (
	"code.lhc.org/we-blog/common"
	"code.lhc.org/we-blog/dal"
	"code.lhc.org/we-blog/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type UserService struct {
}

var (
	userDao = &dal.UserDao{}
)

func (u *UserService) LoginByPassword(c *gin.Context) (*model.User, string, error) {
	username := c.Query("username")
	password := c.Query("password")
	password = common.Md5(password)
	user, err := userDao.GetUserByName(c, username)
	if err != nil {
		fmt.Println("err is &v", err)
		return nil, "", err
	}
	if password != user.Password {
		return nil, "", fmt.Errorf("pass err")
	}
	claims := common.JWTClaims{
		UserID: user.Id,
		Role:   user.Role,
	}
	token, err := common.JwtToToken(&claims)
	if err != nil {
		fmt.Println("[login by password:token generate] err is &v", err)
		return nil, "", err
	}
	return user, token, nil
}

func (u *UserService) GetUserInfo(c *gin.Context) (*model.User, error) {
	id := c.Query("id")
	claims, err := common.VerifyAction(c.Request.Header.Get("we-token"))
	if err != nil {
		fmt.Printf("err is %v \n", err)
		return nil, err
	}
	if id != claims.Id && claims.Role != 1 {
		fmt.Println("err is no permission")
		return nil, errors.New("no permission")
	}
	user_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Printf("id is not int, %s", id)
		return nil, err
	}
	user, err := userDao.GetUserById(c, user_id)
	if err != nil {
		fmt.Println("err is &v", err)
		return nil, err
	}
	return user, nil
}

func (u *UserService) RegisterByPassword(c *gin.Context) (int64, error) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	password = common.Md5(password)
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
