package lifecycle

import "context"

type Cycle interface {
	Start(context.Context) error
	Stop(context.Context) error
}
