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
	q := fmt.Sprintf("INSERT INTO %s (name) VALUES('%s') RETURNING id", GroupTable, name)

	row := gr.db.QueryRow(q)
	if err = row.Scan(); err != nil {
		return 0, err
	}

	return
}

func (gr *GroupsRepository) Update(groupID int, name string) (err error) {
	q := fmt.Sprintf("UPDATE %s SET name = '%s' WHERE id = %d", GroupTable, name, groupID)

	_, err = gr.db.Exec(q)
	return
}

func (gr *GroupsRepository) Delete(groupID int) (err error) {
	q := fmt.Sprintf("DELETE FROM %s WHERE id = %d", GroupTable, groupID)

	_, err = gr.db.Exec(q)
	return
}

func (gr *GroupsRepository) GetAll() (res []domain.Group) {
	q := fmt.Sprintf("SELECT * FROM %s", GroupTable)

	if err := gr.db.Select(&res, q); err != nil {
		logrus.Errorf("Error occured while getting all groups: %v", err)
		return nil
	}
	return
}

func (gr *GroupsRepository) Get(groupID int) (res domain.Group, err error) {
	q := fmt.Sprintf("SELECT * FROM %s WHERE id = %d", GroupTable, groupID)

	if err = gr.db.Get(&res, q); err != nil {
		logrus.Errorf("Error occured while getting group: %v", err)
		return domain.Group{}, err
	}
	return
}

// Join queries
// ? Как лучше всего хранить sql запросы
// ? Как быть менее зависимым от имен таблиц
// ? Правильно ли хранить sql запросы на гитхабе

func (gr *GroupsRepository) GroupsOfUser(userID int) (res []domain.Group, err error) {
	q := fmt.Sprintf(`SELECT group_id, name_group FROM %s
		INNER JOIN %s ON group_user.group_id = group.group_id WHERE user_id = %d`, GroupUserTable, GroupTable, userID)
	fmt.Println(q)
	if err = gr.db.Select(&res, q); err != nil {
		logrus.Errorf("Error occured while getting groups of user %d: %v", userID, err)
		return nil, err
	}
	return
}

func (gr *GroupsRepository) WordsInGroup(groupID int) (res []domain.Word, err error) {
	q := fmt.Sprintf(`SELECT word_id, rus_word, translation FROM %s
		INNER JOIN %s ON group_word.word_id = word.word_id WHERE group_id = %d`, GroupWordTable, WordTable, groupID)
	if err = gr.db.Select(&res, q); err != nil {
		logrus.Errorf("Error occured while getting word in groups %d: %v", groupID, err)
		return nil, err
	}
	return
}
