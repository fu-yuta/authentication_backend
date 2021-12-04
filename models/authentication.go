package models

import (
	"authentication_backend/errorValues"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type User struct {
	Username          string `gorm:"primary_key"`
	Password          string
	AccessToken       string `gorm:"unique"`
	RefreshToken      string `gorm:"unique"`
	AccessTokenExpire int64
}

func init() {
	db = setupDB()
	db.AutoMigrate(User{})
}

func setupDB() *gorm.DB {
	DBDriverName := beego.AppConfig.String("db_driver_name")
	DBName := beego.AppConfig.String("db_name")
	DBUserName := beego.AppConfig.String("db_user_name")
	DBUserPassword := beego.AppConfig.String("db_user_password")
	DBHost := beego.AppConfig.String("db_host")

	connectionInfo := fmt.Sprintf("%s:%s@%s/%s", DBUserName, DBUserPassword, DBHost, DBName)
	db, err := gorm.Open(DBDriverName, connectionInfo)

	if err != nil {
		log.Println(err.Error())
	}

	return db

}

func SignUp(name string, password string) (*User, error) {
	var user User
	db.Find(&user, "username = ?", name)
	if user.Username != "" {
		log.Println("すでにユーザー登録されています。")
		return nil, errorValues.AlreadyExistUserError
	}

	user = User{
		Username:          name,
		Password:          password,
		AccessToken:       uuid.New().String(),
		RefreshToken:      uuid.New().String(),
		AccessTokenExpire: time.Now().Add(30 * time.Minute).Unix(),
	}

	err := db.Create(&user).Error
	if err != nil {
		log.Println("CreateUserError error")
		return nil, errorValues.CreateUserError
	}
	return &user, nil
}

func Login(name, password string) (*User, error) {
	var user User
	db.Find(&user, "username = ?", name)
	if user.Username == "" {
		log.Println("ユーザーが登録されていません")
		return nil, errorValues.NotRegisterUserError
	}

	if user.Password != password {
		log.Println("パスワードが違います")
		return nil, errorValues.MissmatchPasswordError
	}

	user.AccessToken = uuid.New().String()
	user.RefreshToken = uuid.New().String()
	user.AccessTokenExpire = time.Now().Add(30 * time.Minute).Unix()

	err := db.Save(&user).Error
	if err != nil {
		log.Println("ユーザーのトークン保存に失敗しました")
		return nil, errorValues.LoginError
	}

	return &user, nil
}

func Logout(accessToken string) {
	var user User
	db.Where("access_token = ?", accessToken).First(&user)
	if user.Username == "" {
		return
	}

	user.AccessTokenExpire = 0

	db.Save(&user)
}

func GetUser(accessToken string) string {
	var user User
	db.Where("access_token = ?", accessToken).First(&user)
	if user.Username == "" {
		return ""
	}

	return user.Username
}

func AuthenticationAccessToken(accessToken string) error {
	var user User
	db.Where("access_token = ?", accessToken).First(&user)
	if user.Username == "" {
		log.Println("ユーザーが見つかりません")
		return errorValues.NotRegisterUserError
	}

	if time.Now().Unix() > user.AccessTokenExpire {
		log.Println("トークンの期限が切れています")
		return errorValues.AccessTokenExpireError
	}

	return nil
}

func RefreshToken(refreshToken string) (*User, error) {
	var user User
	db.Where("refresh_token = ?", refreshToken).First(&user)
	if user.Username == "" {
		log.Println("ユーザーが見つかりません")
		return nil, errorValues.NotRegisterUserError
	}

	user.AccessToken = uuid.New().String()
	user.AccessTokenExpire = time.Now().Add(30 * time.Minute).Unix()

	err := db.Save(&user).Error
	if err != nil {
		log.Println("ユーザーのトークン保存に失敗しました")
		return nil, errorValues.LoginError
	}

	return &user, nil
}
