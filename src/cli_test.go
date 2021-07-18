package semv

import (
	"reflect"
	"testing"
)

func TestRunCLI(t *testing.T) {
	type args struct {
		env Env
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := RunCLI(tt.args.env); got != tt.want {
			t.Errorf("%q. RunCLI() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_cli_buildHelp(t *testing.T) {
	type fields struct {
		env       Env
		command   string
		RepoURL   string
		Pre       bool
		PreName   string
		Build     bool
		BuildName string
		All       bool
		Prefix    string
		Help      bool
		Version   bool
	}
	type args struct {
		names []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &cli{
			env:       tt.fields.env,
			command:   tt.fields.command,
			RepoURL:   tt.fields.RepoURL,
			Pre:       tt.fields.Pre,
			PreName:   tt.fields.PreName,
			Build:     tt.fields.Build,
			BuildName: tt.fields.BuildName,
			All:       tt.fields.All,
			Prefix:    tt.fields.Prefix,
			Help:      tt.fields.Help,
			Version:   tt.fields.Version,
		}
		if got := c.buildHelp(tt.args.names); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. cli.buildHelp() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_cli_showHelp(t *testing.T) {
	type fields struct {
		env       Env
		command   string
		RepoURL   string
		Pre       bool
		PreName   string
		Build     bool
		BuildName string
		All       bool
		Prefix    string
		Help      bool
		Version   bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &cli{
			env:       tt.fields.env,
			command:   tt.fields.command,
			RepoURL:   tt.fields.RepoURL,
			Pre:       tt.fields.Pre,
			PreName:   tt.fields.PreName,
			Build:     tt.fields.Build,
			BuildName: tt.fields.BuildName,
			All:       tt.fields.All,
			Prefix:    tt.fields.Prefix,
			Help:      tt.fields.Help,
			Version:   tt.fields.Version,
		}
		c.showHelp()
	}
}

func Test_cli_run(t *testing.T) {
	type fields struct {
		env       Env
		command   string
		RepoURL   string
		Pre       bool
		PreName   string
		Build     bool
		BuildName string
		All       bool
		Prefix    string
		Help      bool
		Version   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &cli{
			env:       tt.fields.env,
			command:   tt.fields.command,
			RepoURL:   tt.fields.RepoURL,
			Pre:       tt.fields.Pre,
			PreName:   tt.fields.PreName,
			Build:     tt.fields.Build,
			BuildName: tt.fields.BuildName,
			All:       tt.fields.All,
			Prefix:    tt.fields.Prefix,
			Help:      tt.fields.Help,
			Version:   tt.fields.Version,
		}
		if got := c.run(); got != tt.want {
			t.Errorf("%q. cli.run() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
