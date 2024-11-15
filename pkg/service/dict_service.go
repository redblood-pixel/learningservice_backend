package service

import (
	"github.com/redblood-pixel/learning-service-go/pkg/domain"
	"github.com/redblood-pixel/learning-service-go/pkg/repository"
)

type DictService struct {
	DictRepository repository.Dictionary
}

func NewDictService(dr repository.Dictionary) *DictService {
	return &DictService{
		DictRepository: dr,
	}
}

func (ds *DictService) GetAllWords() []domain.Word {
	return ds.DictRepository.GetAll()
}

func (ds *DictService) GetWord(id int) (domain.Word, error) {
	return ds.DictRepository.Get(id)
}

func (ds *DictService) CreateWord(word domain.CreateWordRequest) error {
	return ds.DictRepository.Create(word)
}

func (ds *DictService) UpdateWord(word domain.Word) error {
	return ds.DictRepository.Update(word)
}

func (ds *DictService) DeleteWord(id int) error {
	return ds.DictRepository.Delete(id)
}
