package payments

import (
	"context"
	"net/http"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

// Yapılan ödemenin iadesi için kullanılır.
func (i *Service) RefundPayment(ctx context.Context, req *RefundPaymentRequest) (response RefundPaymentResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return
	}

	err = i.Transport.Do(ctx, http.MethodPost, refundURI, req, &response)
	return response, err
}

// Yapılan ödemenin iadesi için kullanılır.
//
// v1 ile benzerdir sadece paymentTransactionId yerine paymentId kullanır.
func (i *Service) RefundPaymentV2(ctx context.Context, req *RefundPaymentV2Request) (response RefundPaymentResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return
	}

	err = i.Transport.Do(ctx, http.MethodPost, refundV2URI, req, &response)
	return response, err
}

// Yapılan ödemenin iptali için kullanılır.
func (i *Service) CancelPayment(ctx context.Context, req *CancelPaymentRequest) (response CancelPaymentResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return
	}

	err = i.Transport.Do(ctx, http.MethodPost, cancelURI, req, &response)
	return response, err
}

// Yapılan ödemenin Iyzico hesabına iadesi için kullanılır.
func (i *Service) RefundPaymentToIyzicoAccount(ctx context.Context, req *RefundPaymentIyzicoRequest) (response RefundPaymentIyzicoResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return
	}

	err = i.Transport.Do(ctx, http.MethodPost, refundToIyzicoAccountURI, req, &response)
	return response, err
}
