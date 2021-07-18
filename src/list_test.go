package semv

import (
	"reflect"
	"testing"

	"github.com/blang/semver"
)

func TestGetList(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *List
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := GetList(tt.args.url)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. GetList() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GetList() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestList_FindSimilar(t *testing.T) {
	type fields struct {
		data semver.Versions
	}
	type args struct {
		v semver.Version
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
		l := &List{
			data: tt.fields.data,
		}
		if got := l.FindSimilar(tt.args.v); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. List.FindSimilar() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestList_String(t *testing.T) {
	type fields struct {
		data semver.Versions
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &List{
			data: tt.fields.data,
		}
		if got := l.String(); got != tt.want {
			t.Errorf("%q. List.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestList_Latest(t *testing.T) {
	type fields struct {
		data semver.Versions
	}
	tests := []struct {
		name   string
		fields fields
		want   *Semv
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &List{
			data: tt.fields.data,
		}
		if got := l.Latest(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. List.Latest() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestList_WithoutPreRelease(t *testing.T) {
	type fields struct {
		data semver.Versions
	}
	tests := []struct {
		name   string
		fields fields
		want   *List
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &List{
			data: tt.fields.data,
		}
		if got := l.WithoutPreRelease(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. List.WithoutPreRelease() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestList_OnlyPreRelease(t *testing.T) {
	type fields struct {
		data semver.Versions
	}
	tests := []struct {
		name   string
		fields fields
		want   *List
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		l := &List{
			data: tt.fields.data,
		}
		if got := l.OnlyPreRelease(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. List.OnlyPreRelease() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_getContent(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := getContent(tt.args.url)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. getContent() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. getContent() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_getVersions(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    semver.Versions
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := getVersions(tt.args.url)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. getVersions() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. getVersions() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
