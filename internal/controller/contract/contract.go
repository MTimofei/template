package contract

import "context"

type Service interface {
	Get(context.Context)
}
