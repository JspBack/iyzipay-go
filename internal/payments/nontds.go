package payments

import (
	"context"
	"net/http"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

// 3D Secure olmayan ödeme işlemi için istek oluşturur
func (i *Service) CreatePayment(ctx context.Context, req *PaymentRequest) (response Non3DSPaymentResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, nonTDSURI, req, &response)
	return response, err
}
