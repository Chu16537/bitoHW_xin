package server

import (
	"bitohw_xin/app/module/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Addr string
}

type Handler struct {
	c  *Config
	s  *http.Server
	db database.IDatabase
}

func New(c *Config, db database.IDatabase) *Handler {
	h := new(Handler)
	h.c = c
	h.db = db
	r := gin.Default()

	r.GET("/person/", h.getAll)
	r.GET("/person/:id", h.querySinglePerson)

	r.POST("/addSinglePersonAndMatch", h.addSinglePersonAndMatch)

	r.DELETE("/person/:id", h.removeSinglePerson)

	h.s = &http.Server{
		Addr:    c.Addr,
		Handler: r,
	}

	return h
}

func (h *Handler) GetServer() *http.Server {
	return h.s
}

func (h *Handler) Run() {
	go func() {
		if err := h.s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error: %v\n", err)
		}
	}()
}

func (h *Handler) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.s.Shutdown(ctx); err != nil {
		fmt.Println("server Shutdown err:", err)
		return
	}

	fmt.Println("server Shutdown")
}
