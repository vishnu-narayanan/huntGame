package db

import (
	"gopkg.in/mgo.v2/bson"
)

const UsersColl = "users"

// User schema for users collection
type User struct {
	FirstName   string `bson:"firstName" json:"firstName"`
	LastName    string `bson:"lastName" json:"lastName"`
	Email       string `bson:"email" json:"email"`
	Password    string `bson:"password" json:"password"`
	Level       int    `bson:"level" json:"level"`
	AccessLevel string `bson:"accessLevel" json:"accessLevel"`
	AccessToken string `bson:"accessToken" json:"accessToken"`
}

// Model for new user insert query
type InsertUserQuery struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Level       int    `json:"level"`
	AccessLevel string `json:"accessLevel"`
	AccessToken string `json:"accessToken"`
}

func GetUserByEmail(emailId string) (User, error) {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)

	var user User
	err := c.Find(bson.M{"email": emailId}).One(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func GetUserByAccessToken(t string) (User, error) {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)

	var user User
	err := c.Find(bson.M{"accessToken": t}).One(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func InsertNewUser(u *InsertUserQuery) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)
	err := c.Insert(u)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAccessTokenByEmailId(e string, t string) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(UsersColl)
	q := bson.M{"email": e}
	u := bson.M{"$set": bson.M{"accessToken": t}}
	err := c.Update(&q, &u)
	if err != nil {
		return err
	}
	return nil
}
