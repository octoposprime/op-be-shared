package tconfig

import (
	"os"
	"testing"
)

func Test_GetRedisConfigInstance(t *testing.T) {

	type args struct {
		local bool
		test  bool
	}

	tests := []struct {
		name string
		args args
		want *RedisConfig
	}{
		{
			name: "Test GetRedisConfigInstance()",
			args: args{
				local: false,
				test:  true,
			},
			want: &RedisConfig{},
		},
		{
			name: "Local GetRedisConfigInstance()",
			args: args{
				local: true,
				test:  false,
			},
			want: &RedisConfig{},
		},
		{
			name: "Prod GetRedisConfigInstance()",
			args: args{
				local: false,
				test:  false,
			},
			want: &RedisConfig{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.local {
				os.Setenv("LOCAL", "true")
				RedisConfigTestPath = "../../config/redis_local.yml"
			} else if tt.args.test {
				os.Setenv("TEST", "true")
				RedisConfigTestPath = "../../config/redis_test.yml"
			} else {
				os.Setenv("LOCAL", "")
				os.Setenv("TEST", "")
				RedisConfigTestPath = "../../config/redis.yml"
			}
			if got := GetRedisConfigInstance(); got == nil || got == tt.want {
				t.Errorf("GetRedisConfigInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
