package gorm_repo

import "github.com/redblood-pixel/learning-service-go/pkg/domain"

type DictRepository struct {
	db *domain.DatabaseORM
}

func NewDictRepository(db *domain.DatabaseORM) *DictRepository {
	return &DictRepository{db: db}
}

func (dr *DictRepository) GetAll() []domain.Word {
	return nil
}

func (dr *DictRepository) Get(id int) (domain.Word, error) {
	return domain.Word{}, nil
}

func (dr *DictRepository) Create(word domain.CreateWordRequest) error {
	return nil
}

func (dr *DictRepository) Update(word domain.Word) error {
	return nil
}

func (dr *DictRepository) Delete(wordId int) error {
	return nil
}
