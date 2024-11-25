package postgres_repo

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
	q := fmt.Sprintf("SELECT * FROM %s", WordTable)

	if err := dr.db.Select(&res, q); err != nil {
		logrus.Error(err.Error())
		return nil
	}
	fmt.Println(res)
	return res
}

func (dr *DictRepository) Get(wordID int) (domain.Word, error) {
	var res domain.Word
	q := fmt.Sprintf("SELECT * FROM %s WHERE word_id=%d", WordTable, wordID)

	if err := dr.db.Get(&res, q); err != nil {
		logrus.Error(err)
		return domain.Word{}, err
	}

	fmt.Println(res)
	return res, nil
}

func (dr *DictRepository) Create(word domain.CreateWordRequest) error {
	q := fmt.Sprintf("INSERT INTO %s (rus_word, translation) VALUES ('%s', '%s')", WordTable, word.RusWord, word.Translation)

	_, err := dr.db.Exec(q)
	logrus.Error(err)
	return err
}

func (dr *DictRepository) Update(word domain.Word) error {
	q := fmt.Sprintf("UPDATE %s SET rus_word = '%s', translation = '%s' WHERE word_id = '%d'", WordTable, word.RusWord, word.Translation, word.ID)

	_, err := dr.db.Exec(q)
	logrus.Error(err)
	return err
}

func (dr *DictRepository) Delete(wordID int) error {
	q := fmt.Sprintf("DELETE FROM %s WHERE word_id = '%d'", WordTable, wordID)

	_, err := dr.db.Exec(q)
	logrus.Error(err)
	return err
}
