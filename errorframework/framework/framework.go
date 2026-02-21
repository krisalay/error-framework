package framework

import (
	"sync"

	pgxadapter "github.com/krisalay/error-framework/adapters/pgx"
	validatoradapter "github.com/krisalay/error-framework/adapters/validator"
	"github.com/krisalay/error-framework/core"
)

type Framework struct {
	manager          *core.Manager
	dbAdapter        *pgxadapter.Adapter
	validatorAdapter *validatoradapter.Adapter
}

var instance *Framework
var once sync.Once

func Init(
	manager *core.Manager,
	dbAdapter *pgxadapter.Adapter,
	validatorAdapter *validatoradapter.Adapter,
) {

	once.Do(func() {
		instance = &Framework{
			manager:          manager,
			dbAdapter:        dbAdapter,
			validatorAdapter: validatorAdapter,
		}
	})
}

func get() *Framework {

	if instance == nil {
		panic("framework not initialized")
	}

	return instance
}
