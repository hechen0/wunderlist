package wunderlist

import (
	"testing"
	"net/http"
	"fmt"
	"reflect"
	"context"
)

func TestListService_All(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists", func(w http.ResponseWriter, r *http.Request) {
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

	lists, err := client.Lists.All(context.Background())
	if err != nil {
		t.Errorf("Lists.All returned error: %+v", err)
	}

	want := []*List{{Id: Int(102721804), Title: String("inbox"), Type: String("list")}}

	if !reflect.DeepEqual(want, lists) {
		t.Errorf("want: %+v, got: %+v", want, lists)
	}
}

func TestListService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
		{
			"id": 1
		}
		`)
	})

	list, err := client.Lists.Get(context.Background(), 1)
	if err != nil {
		t.Errorf("Lists.Get returned error: %+v", err)
	}

	want := &List{Id: Int(1)}

	if !reflect.DeepEqual(want, list) {
		t.Errorf("want: %+v, got: %+v", want, list)
	}
}

func TestListService_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		w.WriteHeader(http.StatusCreated)

		w.Write([]byte(`{"title": "test title"}`))
	})

	want := &List{Title: String("test title")}

	got, err := client.Lists.Create(context.Background(), want)
	if err != nil {
		t.Errorf("Lists.Create returned error: %+v", err)
	}

	if *got.Title != *want.Title {
		t.Errorf("want: %+v, got: %+v", want, got)
	}
}

