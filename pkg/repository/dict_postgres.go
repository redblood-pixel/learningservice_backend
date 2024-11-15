package repository

import (
	"fmt"

	"github.com/redblood-pixel/learning-service-go/pkg/domain"
	"github.com/sirupsen/logrus"
)

type DictRepository struct {
	db *domain.Database
}

func NewDictRepository(db *domain.Database) *DictRepository {
	return &DictRepository{
		db: db,
	}
}

func (dr *DictRepository) GetAll() []domain.Word {
	var res []domain.Word
	q := "SELECT * FROM words"

	if err := dr.db.Select(&res, q); err != nil {
		logrus.Error(err.Error())
	}
	fmt.Println(res)
	return res
}

func (dr *DictRepository) Get(id int) (domain.Word, error) {
	var res domain.Word
	q := fmt.Sprintf("SELECT * FROM words WHERE id=%d", id)

	if err := dr.db.Get(&res, q); err != nil {
		logrus.Error(err)
		return domain.Word{}, err
	}

	fmt.Println(res)
	return res, nil
}

func (dr *DictRepository) Create(word domain.CreateWordRequest) error {
	q := fmt.Sprintf("INSERT INTO words (rus_word, translation) VALUES ('%s', '%s')", word.RusWord, word.Translation)

	_, err := dr.db.Exec(q)
	logrus.Error(err)
	return err
}

func (dr *DictRepository) Update(word domain.Word) error {
	q := fmt.Sprintf("UPDATE words SET rus_word = '%s', translation = '%s' WHERE id = '%d'", word.RusWord, word.Translation, word.Id)

	_, err := dr.db.Exec(q)
	logrus.Error(err)
	return err
}

func (dr *DictRepository) Delete(wordId int) error {
	q := fmt.Sprintf("DELETE FROM words WHERE id = '%d'", wordId)

	_, err := dr.db.Exec(q)
	logrus.Error(err)
	return err
}
