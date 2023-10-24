package data

import "fmt"

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

func (u *User) GetID() int {
	return u.ID
}

func (u *User) Convert(dataMap map[string]interface{}) (Entity, error) {
    var user User

    if email, ok := dataMap["email"].(string); ok {
        user.Email = email
    } else {
        return &user, fmt.Errorf("email not found or not a string")
    }

    if id, ok := dataMap["id"].(float64); ok {
        user.ID = int(id)
    } else {
        return &user, fmt.Errorf("id not found or not a float64")
    }

    if name, ok := dataMap["name"].(string); ok {
        user.Name = name
    } else {
        return &user, fmt.Errorf("name not found or not a string")
    }

    if password, ok := dataMap["password"].(string); ok {
        user.Password = password
    } else {
        return &user, fmt.Errorf("password not found or not a string")
    }

    return &user, nil
}