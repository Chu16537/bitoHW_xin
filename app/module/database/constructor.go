package database

import (
	"bitohw_xin/app/proto"
	"sync"
)

type Handler struct {
	mu sync.RWMutex

	personMap map[int]*proto.Person

	lastId int //最後的id
}

type IDatabase interface {
	// 取得使用者
	Get(id int) (*proto.Person, error)
	// 取得所有使用者
	GetAll() map[int]*proto.Person
	// 增加使用者
	Add(data *proto.AddPerson) (*proto.Person, error)
	// 移除
	Remove(id int) error
	// 配對
	Match(filter *proto.MatchFilter) ([]*proto.Person, error)
}

func New() IDatabase {
	h := new(Handler)

	h.personMap = make(map[int]*proto.Person)
	h.lastId = 0

	return h
}
