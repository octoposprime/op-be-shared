package pc_tgorm

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormClientType string

const (
	_                  GormClientType = ""
	PostgresGormClient GormClientType = "postgres"
)

type GormClient struct {
	DbClient       *gorm.DB
	ConnectHandler ConnectHandlerFunc
}

func NewGormClient(clientType GormClientType) *GormClient {
	return &GormClient{
		ConnectHandler: GenerateConnectHandler(clientType),
	}
}

func (c *GormClient) Connect(host string, port string, userName string, password string, database string) (*GormClient, error) {
	var err error
	c.DbClient, err = c.ConnectHandler(host, port, userName, password, database)
	if err != nil {
		return c, err
	}
	return c, nil
}

func GenerateConnectHandler(clientType GormClientType) ConnectHandlerFunc {
	var handler ConnectHandlerFunc
	switch clientType {
	case PostgresGormClient:
		handler = ConnectPostgres
	default:
		handler = ConnectPostgres
	}
	return handler
}

type ConnectHandlerFunc func(host string, port string, userName string, password string, database string) (*gorm.DB, error)

func ConnectPostgres(host string, port string, userName string, password string, database string) (*gorm.DB, error) {
	dbURL := "postgres://" + userName + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=disable"
	dbClient, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, err
	}
	return dbClient, nil
}

type Model struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CreatedBy uuid.UUID `json:"created_by" gorm:"type:uuid;default:uuid_nil()"` // CreatedBy is the creator user's id.
	UpdatedBy uuid.UUID `json:"updated_by" gorm:"type:uuid;default:uuid_nil()"` // UpdatedBy is the updater user's id.
	DeletedBy uuid.UUID `json:"deleted_by" gorm:"type:uuid;default:uuid_nil()"` // DeletedBy is the deleter user's id.
}
