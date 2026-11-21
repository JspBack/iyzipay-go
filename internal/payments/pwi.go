package payments

import (
	"context"
	"net/http"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

// PWI ile ödeme başlatma isteği yapılır.
//
// Response içerisinde gelen token ve payWithIyzicoPageUrl saklanmalıdır.
func (i *Service) InitializePWIPaymentRequest(ctx context.Context, req *InitPWIRequest) (response InitPWIResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, initPWIURI, req, &response)
	return response, err
}

// PWI ile gerçekleştirilen ödeme sorgulama isteği yapılır.
//
// Response içerisinde gelen paymentId saklanmalıdır.
func (i *Service) CheckPWIPaymentRequest(ctx context.Context, req *PWIInquiryRequest) (response PWIInquiryResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, checkPWIURI, req, &response)
	return response, err
}
