package view

import "github.com/pufferpanel/pufferpanel/models"

type UserViewModel struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	//ONLY SHOW WHEN COPYING
	Password string `json:"password,omitempty"`
}

func FromUser(model *models.User) *UserViewModel {
	return &UserViewModel{
		Username: model.Username,
		Email:    model.Email,
	}
}

func (model *UserViewModel) CopyToModel(newModel *models.User) {
	if model.Username != "" {
		newModel.Username = model.Username
	}

	if model.Email != "" {
		newModel.Email = model.Email
	}

	if model.Password != "" {
		newModel.SetPassword(model.Password)
	}
}
