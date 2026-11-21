package subscriptions

import (
	"github.com/JspBack/iyzipay-go/internal/transport"
	"github.com/JspBack/iyzipay-go/internal/validation"
)

type Service struct {
	Transport transport.Executor
	Validator *validation.Validator
}
