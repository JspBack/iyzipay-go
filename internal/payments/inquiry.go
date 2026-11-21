package payments

import (
	"context"
	"net/http"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

// İşlem sorgulama isteği
func (i *Service) PaymentInquiryRequest(ctx context.Context, req *InquiryRequest) (response InquiryResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, inquiryURI, req, &response)
	return response, err
}
