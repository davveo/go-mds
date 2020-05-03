package service

import "github.com/zbrechave/go-mds/models"

type MessageService interface {
	Get(id int) *models.Message
	Delete(id int) error
	Update(user *models.Message, columns []string) error
	Create(user *models.Message) error
}

type messageService struct {
	dao *dao.messageDao
}

func (ms *messageService) Get(id int) *models.Message {
	return nil
}
func (ms *messageService) Delete(id int) error {
	return nil
}
func (ms *messageService) Update(message *models.Message, columns []string) error {
	return nil
}

func (ms *messageService) Create(message *models.Message) error {
	return nil
}
