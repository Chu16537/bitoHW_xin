package test

import (
	"bitohw_xin/app/consts"
	"bitohw_xin/app/module/database"
	"bitohw_xin/app/proto"
	"bitohw_xin/app/server"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

var port = 8080
var s *server.Handler
var db database.IDatabase

func init() {

	config := &server.Config{
		Addr: fmt.Sprintf("0.0.0.0:%v", port),
	}

	db = database.New()

	s = server.New(config, db)

	initPerson()

}

func getURL(path string) string {
	return fmt.Sprintf("http://127.0.0.1:%v/%v", port, path)
}

// 產生預設資料
func initPerson() {
	for i := 0; i < 5; i++ {
		p := &proto.AddPerson{
			Name:        strconv.Itoa(i),
			Height:      170,
			Gender:      i % 2,
			WantedDates: 10,
		}
		db.Add(p)
	}
}
func TestQuerySinglePerson(t *testing.T) {

	url := getURL("person/1?count=3")
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		t.Fatal(err)
		return
	}

	rec := httptest.NewRecorder()

	s.GetServer().Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("失敗 狀態號碼:%v", rec.Code)
		return
	}

	fmt.Println("TestQuerySinglePerson success", rec.Body.String())
}

func TestAddSinglePersonAndMatch(t *testing.T) {

	reqData := &proto.AddPerson{
		Name:        "TestAddSinglePersonAndMatch",
		Height:      170,
		Gender:      consts.Boy,
		WantedDates: 10,
	}

	bytes, err := json.Marshal(reqData)
	if err != nil {
		t.Fatal(err)
		return
	}

	url := getURL("addSinglePersonAndMatch")
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(bytes)))
	if err != nil {
		t.Fatal(err)
		return
	}

	rec := httptest.NewRecorder()

	s.GetServer().Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("失敗 狀態號碼:%v", rec.Code)
		return
	}

	fmt.Println("TestAddSinglePersonAndMatch success", rec.Body.String())
}
func TestRemoveSinglePerson(t *testing.T) {
	url := getURL("person/1")
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		t.Fatal(err)
		return
	}

	rec := httptest.NewRecorder()

	s.GetServer().Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("失敗 狀態號碼:%v", rec.Code)
		return
	}

	fmt.Println("TestRemoveSinglePerson success", rec.Body.String())
}
