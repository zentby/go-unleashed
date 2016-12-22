package unleashed

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	//"time"
)

func TestProductService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", jsontype)

		fmt.Fprint(w, `{"Pagination": {"NumberOfItems": 9, "PageSize": 200, "PageNumber": 1, "NumberOfPages": 1 }, "Items": [{"ProductCode": "1"} ] }`)
	})

	issues, _, err := client.Products.List(nil, nil)
	if err != nil {
		t.Errorf("Issues.List returned error: %v", err)
	}

	want := &ProductList{Pagination: Pagination{NumberOfItems: 9, PageSize: 200, PageNumber: 1, NumberOfPages: 1},
		Items: []Product{{ProductCode: String("1")}}}
	if !reflect.DeepEqual(issues, want) {
		t.Errorf("Issues.List returned %+v, want %+v", issues, want)
	}
}
