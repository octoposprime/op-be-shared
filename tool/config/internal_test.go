package tconfig

import (
	"os"
	"testing"
)

func Test_GetInternalConfigInstance(t *testing.T) {

	type args struct {
		local bool
		test  bool
	}

	tests := []struct {
		name string
		args args
		want *InternalConfig
	}{
		{
			name: "Test GetInternalConfigInstance()",
			args: args{
				local: false,
				test:  true,
			},
			want: &InternalConfig{},
		},
		{
			name: "Local GetInternalConfigInstance()",
			args: args{
				local: true,
				test:  false,
			},
			want: &InternalConfig{},
		},
		{
			name: "Prod GetInternalConfigInstance()",
			args: args{
				local: false,
				test:  false,
			},
			want: &InternalConfig{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.local {
				os.Setenv("LOCAL", "true")
				InternalConfigTestPath = "../../config/internal_local.yml"
			} else if tt.args.test {
				os.Setenv("TEST", "true")
				InternalConfigTestPath = "../../config/internal_test.yml"
			} else {
				os.Setenv("LOCAL", "")
				os.Setenv("TEST", "")
				InternalConfigTestPath = "../../config/internal.yml"
			}
			if got := GetInternalConfigInstance(); got == nil || got == tt.want {
				t.Errorf("GetInternalConfigInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
