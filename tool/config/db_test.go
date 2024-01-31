package tconfig

import (
	"os"
	"testing"
)

func Test_GetDbConfigInstance(t *testing.T) {

	type args struct {
		local bool
		test  bool
	}

	tests := []struct {
		name string
		args args
		want *DbConfig
	}{
		{
			name: "Test GetDbConfigInstance()",
			args: args{
				local: false,
				test:  true,
			},
			want: &DbConfig{},
		},
		{
			name: "Local GetDbConfigInstance()",
			args: args{
				local: true,
				test:  false,
			},
			want: &DbConfig{},
		},
		{
			name: "Prod GetDbConfigInstance()",
			args: args{
				local: false,
				test:  false,
			},
			want: &DbConfig{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.local {
				os.Setenv("LOCAL", "true")
				DbConfigTestPath = "../../config/postgres_local.yml"
			} else if tt.args.test {
				os.Setenv("TEST", "true")
				DbConfigTestPath = "../../config/postgres_test.yml"
			} else {
				os.Setenv("LOCAL", "")
				os.Setenv("TEST", "")
				DbConfigTestPath = "../../config/postgres.yml"
			}
			if got := GetDbConfigInstance(); got == nil || got == tt.want {
				t.Errorf("GetDbConfigInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
