package tjwt

import (
	"testing"

	"github.com/google/uuid"
)

func Test_GenerateJWT(t *testing.T) {
	type user struct {
		userId   uuid.UUID
		username string
		email    string
	}
	testUser1 := user{
		userId:   uuid.UUID{},
		username: "test",
		email:    "test@test.com",
	}
	type args struct {
		claim *Claims
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test GenerateJWT()",
			args: args{
				claim: NewClaims(testUser1.userId.String(), testUser1),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := tt.args.claim.GenerateJWT()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
