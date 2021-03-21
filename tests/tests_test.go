package tests

import (
	"github.com/ksenia-portfolio/crud-in-go"
	"net/http"
	"reflect"
	"testing"
	"time"
)

// test methods are in progress

func TestColumn_Order(t *testing.T) {
	type fields struct {
		Name string
		Asc  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &main.Column{
				Name: tt.fields.Name,
				Asc:  tt.fields.Asc,
			}
			if got := c.Order(); got != tt.want {
				t.Errorf("Order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColumn_changeAsc(t *testing.T) {
	type fields struct {
		Name string
		Asc  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &main.Column{
				Name: tt.fields.Name,
				Asc:  tt.fields.Asc,
			}
			if got := c.ChangeAsc(); got != tt.want {
				t.Errorf("changeAsc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerForm_IsValidAge(t *testing.T) {
	type fields struct {
		Id        uint16
		FirstName string
		LastName  string
		Gender    string
		DOB       string
		Email     string
		Address   string
		Errors    map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cf := &main.CustomerForm{
				Id:        tt.fields.Id,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Gender:    tt.fields.Gender,
				DOB:       tt.fields.DOB,
				Email:     tt.fields.Email,
				Address:   tt.fields.Address,
				Errors:    tt.fields.Errors,
			}
			if got := cf.IsValidAge(); got != tt.want {
				t.Errorf("IsValidAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerForm_ValidateCreateForm(t *testing.T) {
	type fields struct {
		Id        uint16
		FirstName string
		LastName  string
		Gender    string
		DOB       string
		Email     string
		Address   string
		Errors    map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cf := &main.CustomerForm{
				Id:        tt.fields.Id,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Gender:    tt.fields.Gender,
				DOB:       tt.fields.DOB,
				Email:     tt.fields.Email,
				Address:   tt.fields.Address,
				Errors:    tt.fields.Errors,
			}
			if got := cf.ValidateCreateForm(); got != tt.want {
				t.Errorf("ValidateCreateForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerForm_ValidateUpdateForm(t *testing.T) {
	type fields struct {
		Id        uint16
		FirstName string
		LastName  string
		Gender    string
		DOB       string
		Email     string
		Address   string
		Errors    map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cf := &main.CustomerForm{
				Id:        tt.fields.Id,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Gender:    tt.fields.Gender,
				DOB:       tt.fields.DOB,
				Email:     tt.fields.Email,
				Address:   tt.fields.Address,
				Errors:    tt.fields.Errors,
			}
			if got := cf.ValidateUpdateForm(); got != tt.want {
				t.Errorf("ValidateUpdateForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomer_FormatDOB(t *testing.T) {
	type fields struct {
		Id        uint16
		DOB       time.Time
		FirstName string
		LastName  string
		Gender    string
		Email     string
		Address   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := main.Customer{
				Id:        tt.fields.Id,
				DOB:       tt.fields.DOB,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Gender:    tt.fields.Gender,
				Email:     tt.fields.Email,
				Address:   tt.fields.Address,
			}
			if got := c.FormatDOB(); got != tt.want {
				t.Errorf("FormatDOB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDOB(t *testing.T) {
	type args struct {
		dob string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main.GetDOB(tt.args.dob); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDOB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchForm_ValidateSearchForm(t *testing.T) {
	type fields struct {
		FirstName string
		LastName  string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sf := &main.SearchForm{
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
			}
			if got := sf.ValidateSearchForm(); got != tt.want {
				t.Errorf("ValidateSearchForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addRandomCustomersToDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_ascValue(t *testing.T) {
	type args struct {
		col main.Column
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main.AscValue(tt.args.col); got != tt.want {
				t.Errorf("ascValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_confirmationPageHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_createCustomer(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_createCustomerPageHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_createTable(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_customerInfoPageHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_deleteCustomer(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_handleFunctions(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_homePageHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_loadOptions(t *testing.T) {
	tests := []struct {
		name string
		want []main.Column
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main.LoadOptions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchCustomer(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sortByAddress(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sortByDOB(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sortByEmail(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sortByFirstName(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sortByGender(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sortById(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sortByLastName(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_sortByOrder(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_updateCustomer(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_updateCustomerPageHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_viewCustomersPageHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
