package response

import "api-exemple/app/model"

type Token struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
} //@name Token

type State struct {
	State string `json:"state"`
} //@name State
