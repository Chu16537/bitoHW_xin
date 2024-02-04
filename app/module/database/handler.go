package database

import (
	"bitohw_xin/app/consts"
	"bitohw_xin/app/proto"
	"fmt"
)

// 取得使用者
func (h *Handler) Get(id int) (*proto.Person, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	v, ok := h.personMap[id]
	if !ok {
		return nil, fmt.Errorf("get person %v is not exist", id)
	}

	return v, nil
}

// 取得所有使用者
func (h *Handler) GetAll() map[int]*proto.Person {
	return h.personMap
}

// 增加使用者
func (h *Handler) Add(data *proto.AddPerson) (*proto.Person, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.lastId++

	// 添加
	h.personMap[h.lastId] = proto.NewPerson(h.lastId, data.Name, data.Height, data.Gender, data.WantedDates)

	return h.personMap[h.lastId], nil
}

// 移除
func (h *Handler) Remove(id int) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	v, ok := h.personMap[id]
	if !ok {
		// 使用者不存在
		return fmt.Errorf("remove Person %v not exist", id)
	}

	// 軟刪除
	v.SetMatchStatus(consts.MatchStatus_NotMatch)

	return nil
}

// 配對
func (h *Handler) Match(filter *proto.MatchFilter) ([]*proto.Person, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	v, ok := h.personMap[filter.ID]
	if !ok {
		return nil, fmt.Errorf("match Person %v not exist", filter.ID)
	}

	if !v.IsMatch() {
		return nil, fmt.Errorf("match Person %v not match", filter.ID)
	}

	// 次數不夠，改為剩餘次數
	if v.WantedDates < filter.Count {
		fmt.Printf("match Person %v Count less \n", filter.ID)
		filter.Count = v.WantedDates
	}

	var persons []*proto.Person

	if filter.Gender == consts.Boy {
		persons = h.getGril(filter.Height, filter.Count, filter.NotMatchPersonIDs)
	} else {
		persons = h.getBoy(filter.Height, filter.Count, filter.NotMatchPersonIDs)
	}

	v.Match(len(persons))

	// notMatchIds := make([]int, len(persons))
	// for i, v := range persons {
	// 	notMatchIds[i] = v.ID
	// }
	// v.AddnotMatchPersonIDs(notMatchIds)

	return persons, nil
}
