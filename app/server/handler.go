package server

import (
	"bitohw_xin/app/consts"
	"bitohw_xin/app/proto"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 增加
// 增加新用戶時 回傳n(預設3個)個匹配的人員
func (h *Handler) addSinglePersonAndMatch(c *gin.Context) {
	var addPerson proto.AddPerson

	if err := c.ShouldBindJSON(&addPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	p, err := h.db.Add(&addPerson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	f := proto.NewMatchFilter(p.ID, p.Name, p.Height, p.Gender, consts.BaseCount, p.NotMatchPersonIDs)

	persons, err := h.db.Match(f)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}

// 移除
func (h *Handler) removeSinglePerson(c *gin.Context) {
	fmt.Println("removeSinglePerson")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is error",
		})

		return
	}

	if err := h.db.Remove(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) getAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"persons": h.db.GetAll(),
	})
}

// 查詢
func (h *Handler) querySinglePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is error",
		})

		return
	}

	p, err := h.db.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "count is error",
		})

		return
	}

	f := &proto.MatchFilter{
		ID:                id,
		Height:            p.Height,
		Gender:            p.Gender,
		Count:             count,
		NotMatchPersonIDs: p.NotMatchPersonIDs,
	}

	persons, err := h.db.Match(f)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}
