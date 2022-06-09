package transactor

import "context"

type Transactor interface {
	WithinTransaction(func(ctx context.Context) error) error
}
