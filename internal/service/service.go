package service

import (
	"entryTask/internal/dao"
)

type Service struct {
	dao *dao.Dao
}

func New() *Service {
	return &Service{
		dao: dao.New(),
	}
}

func (s *Service) Close() {
	s.dao.Close()
}
