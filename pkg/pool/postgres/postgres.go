package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/mhthrh/common-lib/config/model"
	customModelError "github.com/mhthrh/common-lib/errors"
	customError "github.com/mhthrh/common-lib/errors/pool"
	"github.com/mhthrh/common-lib/pkg/pool"
	"github.com/samber/lo"
	"strings"
	"time"
)

const (
	psql                   = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	checkCnnStatusDuration = time.Millisecond * 10
	totalWaitForReleaseAll = time.Second * 10
)

var (
	connections map[string]pool.Connection
	stop        chan struct{}
)

type Config struct {
	db model.DB
}

func init() {
	connections = make(map[string]pool.Connection)

}

func New(db model.DB) (pool.IConnection, *customModelError.XError) {
	if strings.Trim(db.Host, " ") == "" {
		return nil, customError.InputParamsMismatch(nil)
	}
	return Config{db: db}, nil
}

func (c Config) Maker(request chan pool.Request) (response chan pool.Response) {
	for {
		select {
		case r := <-request:
			switch {
			case r.Count == 0:
				response <- pool.Response{
					Total: uint(len(connections)),
					InUse: uint(len(lo.PickBy(connections, func(key string, value pool.Connection) bool {
						return value.InUse == true
					}))),
					Error: nil,
				}
				break
			case r.Count > 0:
				if len(connections) >= int(r.Count) {
					response <- pool.Response{
						Total: r.Count,
						InUse: uint(len(lo.PickBy(connections, func(key string, value pool.Connection) bool {
							return value.InUse == true
						}))),
						Error: customError.MaximumConnection(nil),
					}
					break
				}
				s := true
				for range int(r.Count) - len(connections) {
					cnn, err := sql.Open(c.db.Driver, fmt.Sprintf(psql, c.db.Host, c.db.Port, c.db.UserName, c.db.Password, c.db.DbName, c.db.SSLMode))
					if err != nil {
						response <- pool.Response{
							Total: uint(len(connections)),
							InUse: uint(0),
							Error: customError.DbConnectionFailed(customModelError.RunTimeError(err)),
						}
						s = false
						break
					}
					key := uuid.New()
					connections[key.String()] = pool.Connection{
						Id:    key,
						Cnn:   cnn,
						InUse: false,
					}
				}
				if s {
					response <- pool.Response{
						Total: uint(len(connections)),
						InUse: uint(len(lo.PickBy(connections, func(key string, value pool.Connection) bool {
							return value.InUse == true
						}))),
						Error: nil,
					}
					break
				}
			}
		}
	}
}

func (c Config) Manager(connections chan pool.Connection) {
	//TODO implement me
	panic("implement me")
}

func (c Config) Refresh(c2 chan struct{}) chan pool.Response {
	//TODO implement me
	panic("implement me")
}

func (c Config) Release(uuids chan uuid.UUID) chan pool.Response {
	//TODO implement me
	panic("implement me")
}

func (c Config) ReleaseAll() *customModelError.XError {
	//TODO implement me
	panic("implement me")
}
