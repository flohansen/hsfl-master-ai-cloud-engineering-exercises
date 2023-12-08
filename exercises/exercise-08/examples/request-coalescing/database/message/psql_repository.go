package message

import (
	"database/sql"
	"echo-server/database"

	_ "github.com/lib/pq"
)

type psqlRepository struct {
	db *sql.DB
}

const queryCreateTable = `
create table if not exists messages (
	id   serial primary key,
	username varchar(100) not null,
	text text not null
)
`

func NewPsqlRepository(config database.Config) (*psqlRepository, error) {
	db, err := sql.Open("postgres", config.Dsn())
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(queryCreateTable); err != nil {
		return nil, err
	}

	return &psqlRepository{db}, nil
}

const queryCreateMessage = `
insert into messages (id, username, text) values ($1, $2, $3)
`

func (repo *psqlRepository) Create(message *database.Message) error {
	_, err := repo.db.Exec(queryCreateMessage, message.ID, message.User, message.Text)
	return err
}

const queryFindMessageById = `
select id, username, text from messages where id = $1
`

func (repo *psqlRepository) FindById(id int64) (*database.Message, error) {
	row := repo.db.QueryRow(queryFindMessageById, id)

	var message database.Message
	if err := row.Scan(&message.ID, &message.User, &message.Text); err != nil {
		return nil, err
	}

	return &message, nil
}
