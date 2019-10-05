package domain

import (
	"reflect"
	"strconv"
	"testing"
)

var p *Page

func TestCreatePage(t *testing.T) {
	p = &Page{
		Title:   "test",
		Content: "test",
	}
	ID, err := CreatePage(p)
	if err != nil {
		t.Errorf("Failed to create page: %s\n", err.Error())
	}
	p.Id = int64(ID)
}
func TestGetPage(t *testing.T) {
	page, err := GetPage(strconv.Itoa(int(p.Id)))
	if err != nil {
		t.Errorf("Failed to get page: %s\n", err.Error())
	}
	if page.Id != p.Id {
		t.Errorf("Page Ids do not match: %d\n vs %d\n", page.Id, p.Id)
	}
	if reflect.DeepEqual(page, p) != true {
		t.Errorf("Pages do not match: %+v\n vs %+v\nn", page, p)
	}
}
