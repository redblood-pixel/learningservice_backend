package service

import (
	"github.com/redblood-pixel/learning-service-go/pkg/domain"
	"github.com/redblood-pixel/learning-service-go/pkg/repository"
)

type GroupService struct {
	GroupsRepository repository.Group
}

func NewGroupService(db repository.Group) *GroupService {
	return &GroupService{GroupsRepository: db}
}

func (gs *GroupService) GetAllGroups() []domain.Group {
	return gs.GroupsRepository.GetAll()
}

func (gs *GroupService) GetGroup(groupID int) (domain.Group, error) {
	return gs.GroupsRepository.Get(groupID)
}

func (gs *GroupService) CreateGroup(input domain.CreateGroupRequest) (int, error) {
	return gs.GroupsRepository.Create(input)
}

func (gs *GroupService) UpdateGroup(group domain.Group) error {
	return gs.GroupsRepository.Update(group)
}

func (gs *GroupService) DeleteGroup(groupID int) error {
	return gs.GroupsRepository.Delete(groupID)
}

func (gs *GroupService) GetWordsInGroup(groupID int) ([]domain.Word, error) {
	return gs.GroupsRepository.WordsInGroup(groupID)
}

func (gs *GroupService) GetGroupsOfUser(userID int) ([]domain.Group, error) {
	return gs.GroupsRepository.GroupsOfUser(userID)
}
