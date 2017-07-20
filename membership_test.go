package wunderlist

import (
	"net/http"
	"testing"
)

func TestMembershipService_All(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/memberships", func(w http.ResponseWriter, r *http.Request) {
	})
}

//func TestMembershipService_All2(t *testing.T) {
//
//	token, clientId := os.Getenv("X-Access-Token"), os.Getenv("X-Client-ID")
//	if token == "" || clientId == "" {
//		return
//	}
//
//	client := NewClient()
//	client.SetAuth(&Auth{Token: token, ClientId: clientId})
//
//	_, err := client.Folders.All(context.Background())
//	if err != nil {
//		t.Errorf("TestMembershipService_All2 got error: %v", err)
//	}
//
//	list := &List{Title: String("hello world form golang api")}
//	list, err = client.Lists.Create(context.Background(), list)
//	if err != nil {
//		panic(err)
//	}
//}
