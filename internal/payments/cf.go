package payments

import (
	"context"
	"errors"
	"net/http"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

// Checkout Fom ile ödeme isteği yapmak için kullanılır.
//
// Bu istekte dönecek olan html elementi için 3 farklı CFFormType belirleyebilirsiniz. "responsive", "popup", "iframe".
func (i *Service) CheckoutFormPaymentRequest(ctx context.Context, req *CFRequest, CFFormType string) (response *CFResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if !acceptedFormTypes[CFFormType] {
		return nil, errors.New("CFFormType is not valid")
	}
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, checkoutFormURI, req, &response)
	if err != nil {
		return response, err
	}

	if CFFormType == "iframe" {
		response.PaymentPageUrl += "&iframe=true"
	}

	return response, nil
}

// Checkout Form ile yapılan ödeme isteği için sorgulama yapmak için kullanılır.
//
// Ödemenin tamamlanması gerekli.
func (i *Service) CheckoutFormPaymentInquiryRequest(ctx context.Context, req *CFInquiryRequest) (response *CFInquiryResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, checkoutFormInquiryURI, req, &response)
	return response, err
}
