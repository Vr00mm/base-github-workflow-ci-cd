package semv

import (
	"reflect"
	"testing"

	"github.com/blang/semver"
)

func TestMustNew(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Semv
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := MustNew(tt.args.s); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. MustNew() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestLatest(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *Semv
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Latest(tt.args.url)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Latest() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Latest() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSemv_String(t *testing.T) {
	type fields struct {
		data semver.Version
		list *List
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		v := &Semv{
			data: tt.fields.data,
			list: tt.fields.list,
		}
		if got := v.String(); got != tt.want {
			t.Errorf("%q. Semv.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSemv_IsEmpty(t *testing.T) {
	type fields struct {
		data semver.Version
		list *List
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		v := &Semv{
			data: tt.fields.data,
			list: tt.fields.list,
		}
		if got := v.IsEmpty(); got != tt.want {
			t.Errorf("%q. Semv.IsEmpty() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSemv_Next(t *testing.T) {
	type fields struct {
		data semver.Version
		list *List
	}
	type args struct {
		target string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Semv
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		v := &Semv{
			data: tt.fields.data,
			list: tt.fields.list,
		}
		if got := v.Next(tt.args.target); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Semv.Next() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSemv_PreRelease(t *testing.T) {
	type fields struct {
		data semver.Version
		list *List
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Semv
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		v := &Semv{
			data: tt.fields.data,
			list: tt.fields.list,
		}
		got, err := v.PreRelease(tt.args.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Semv.PreRelease() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Semv.PreRelease() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSemv_Build(t *testing.T) {
	type fields struct {
		data semver.Version
		list *List
	}
	type args struct {
		name string
		url  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Semv
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		v := &Semv{
			data: tt.fields.data,
			list: tt.fields.list,
		}
		got, err := v.Build(tt.args.name, tt.args.url)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Semv.Build() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Semv.Build() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSemv_cloneForPreRelease(t *testing.T) {
	type fields struct {
		data semver.Version
		list *List
	}
	type args struct {
		vv *Semv
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Semv
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		v := &Semv{
			data: tt.fields.data,
			list: tt.fields.list,
		}
		if got := v.cloneForPreRelease(tt.args.vv); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Semv.cloneForPreRelease() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSemv_incrementMajor(t *testing.T) {
	type fields struct {
		data semver.Version
		list *List
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		v := &Semv{
			data: tt.fields.data,
			list: tt.fields.list,
		}
		v.incrementMajor()
	}
}

func TestSemv_incrementMinor(t *testing.T) {
	type fields struct {
		data semver.Version
		list *List
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		v := &Semv{
			data: tt.fields.data,
			list: tt.fields.list,
		}
		v.incrementMinor()
	}
}

func TestSemv_incrementPatch(t *testing.T) {
	type fields struct {
		data semver.Version
		list *List
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		v := &Semv{
			data: tt.fields.data,
			list: tt.fields.list,
		}
		v.incrementPatch()
	}
}
