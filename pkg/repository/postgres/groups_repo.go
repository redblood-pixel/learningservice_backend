package postgres_repo

import (
	"fmt"

	"github.com/redblood-pixel/learning-service-go/pkg/domain"
	"github.com/sirupsen/logrus"
)

type GroupsRepository struct {
	db *domain.Database
}

func NewGropsRepository(db *domain.Database) *GroupsRepository {
	return &GroupsRepository{db: db}
}

func (gr *GroupsRepository) Create(name string) (id int, err error) {
	q := fmt.Sprintf("INSERT INTO groups (name) VALUES('%s') RETURNING id", name)

	row := gr.db.QueryRow(q)
	if err = row.Scan(); err != nil {
		return 0, err
	}

	return
}

func (gr *GroupsRepository) Update(id int, name string) (err error) {
	q := fmt.Sprintf("UPDATE groups SET name = '%s' WHERE id = %d", name, id)

	_, err = gr.db.Exec(q)
	return
}

func (gr *GroupsRepository) Delete(id int) (err error) {
	q := fmt.Sprintf("DELETE FROM groups WHERE id = %d", id)

	_, err = gr.db.Exec(q)
	return
}

func (gr *GroupsRepository) GetAll() (res []domain.Group) {
	q := "SELECT * FROM groups"

	if err := gr.db.Select(&res, q); err != nil {
		logrus.Errorf("Error occured while getting all groups: %v", err)
		return nil
	}
	return
}

func (gr *GroupsRepository) Get(id int) (res domain.Group, err error) {
	q := fmt.Sprintf("SELECT * FROM groups WHERE id = %d", id)

	if err = gr.db.Get(&res, q); err != nil {
		logrus.Errorf("Error occured while getting group: %v", err)
		return domain.Group{}, err
	}
	return
}

// Join queries

func (gr *GroupsRepository) GetWordsInGroup(id int) (res []domain.Word) {
	return
}
