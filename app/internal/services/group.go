package services

import (
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/repository"
)

type GroupService struct {
	userRepository  *repository.UserRepository
	groupRepository *repository.GroupRepository
}

func NewGroupService(userRepository *repository.UserRepository, groupRepository *repository.GroupRepository) *GroupService {
	return &GroupService{userRepository, groupRepository}
}

func (g *GroupService) AddNewGroup(group *dto.Group) error {
	groupModel := dto.GroupToModel(group)
	if err := g.groupRepository.Create(groupModel); err != nil {
		return err
	}
	return nil
}

func (g *GroupService) DeleteGroup() error {
	return nil
}

func (g *GroupService) GetGroups() ([]dto.Group, error) {
	groupsModel, err := g.groupRepository.GetAll()
	if err != nil {
		return nil, err
	}
	groups := make([]dto.Group, len(groupsModel))
	for i, group := range groupsModel {
		groups[i] = dto.Group{
			Id:    int(group.ID),
			Title: group.Title,
			Code:  group.Code,
		}
	}
	return groups, nil
}

func (g *GroupService) AddUsersInGroup(groupId int, userIds []int) error {
	return g.userRepository.AddUsersInGroupByIds(groupId, userIds)
}

func (g *GroupService) GetUsersByGroupId(groupId int) ([]*dto.UserDTO, error) {
	group, err := g.groupRepository.GetById(groupId)
	if err != nil {
		return nil, err
	}
	users := dto.UsersArrayToDTO(group.Users)
	return users, nil
}

func (g *GroupService) DeleteUsersFromGroup() error {
	return nil
}
