package wunderlist

import (
	"testing"
	"net/http"
	"fmt"
	"reflect"
)

func TestListService_All(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists", func(w http.ResponseWriter, r *http.Request){
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[
  {
    "id": 102721804,
    "title": "inbox",
    "type": "list"
  }
 ]
		`)
	})

	lists, err := client.Lists.All()
	if err != nil {
		t.Errorf("Lists.All returned error: %+v", err)
	}

	want := []*List{{ID: Int(102721804), Title: String("inbox"), Type: String("list")}}

	if !reflect.DeepEqual(want, lists){
		t.Errorf("want: %+v, got: %+v", want, lists)
	}

}

func TestListService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists/1", func(w http.ResponseWriter, r *http.Request){
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
		{
			"id": 1
		}
		`)
	})

	list, err := client.Lists.Get(1)
	if err != nil {
		t.Errorf("Lists.Get returned error: %+v", err)
	}

	want := &List{ID: Int(1)}

	if !reflect.DeepEqual(want, list){
		t.Errorf("want: %+v, got: %+v", want, list)
	}
}
