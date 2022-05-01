package entity

import "github.com/o-t-k-t/sql_boiler_exercise/models"

type User models.User

func (u User) LocalName() string {
	return u.Name + "さん"
}
