package message

import "echo-server/database"

type Repository interface {
	Create(*database.Message) error
	FindById(int64) (*database.Message, error)
}
