package database

import (
	"bitohw_xin/app/consts"
	"bitohw_xin/app/proto"
)

// 配對機制
// 女 配 高男
func (h *Handler) getBoy(height int, count int, notMatchPersonIDs map[int]struct{}) []*proto.Person {
	persons := make([]*proto.Person, 0, count)

	for _, v := range h.personMap {
		// 數量達標
		if len(persons) == count {
			break
		}
		// 不是女生
		if v.Gender != consts.Boy {
			continue
		}

		// 是否可以配對
		if !v.IsMatch() {
			continue
		}

		// 檢查id是否存在
		if _, ok := notMatchPersonIDs[v.ID]; ok {
			continue
		}

		// 條件
		// 身高太矮
		if v.Height < height {
			continue
		}

		// 配對成功
		persons = append(persons, v)
		v.Match(1)
	}

	return persons
}

// 配對機制
// 男 配 低女
func (h *Handler) getGril(height int, count int, notMatchPersonIDs map[int]struct{}) []*proto.Person {
	persons := make([]*proto.Person, 0, count)

	for _, v := range h.personMap {
		// 數量達標
		if len(persons) == count {
			break
		}
		// 不是女生
		if v.Gender != consts.Girl {
			continue
		}

		// 是否可以配對
		if !v.IsMatch() {
			continue
		}

		// 檢查id是否存在
		if _, ok := notMatchPersonIDs[v.ID]; ok {
			continue
		}

		// 條件
		// 身高太高
		if v.Height > height {
			continue
		}

		// 配對成功
		persons = append(persons, v)
		v.Match(1)
	}

	return persons
}
