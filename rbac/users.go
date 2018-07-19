package rbac

import (
	"encoding/json"

	"github.com/crusttech/crust/rbac/types"
)

type (
	Users struct {
		*Client
	}

	UsersInterface interface {
		Create(username, password string) error
		Get(username string) (*types.User, error)
		Delete(username string) error
	}
)

func (u *Users) Create(username, password string) error {
	body := struct {
		Password string `json:"password"`
	}{password}

	resp, err := u.Client.Post("/users/"+username, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (u *Users) Get(username string) (*types.User, error) {
	resp, err := u.Client.Get("/users/" + username)
	if err != nil {
		return nil, err
	}
	user := &types.User{}
	defer resp.Body.Close()
	return user, json.NewDecoder(resp.Body).Decode(user)
}

func (u *Users) Delete(username string) error {
	resp, err := u.Client.Delete("/users/" + username)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

var _ UsersInterface = &Users{}