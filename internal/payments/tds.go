package payments

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/JspBack/iyzipay-go/internal/cards"
	"github.com/JspBack/iyzipay-go/internal/transport"
)

type TDSOption func(*tdsRequestOptions)

type tdsRequestOptions struct {
	binRequest        bool
	htmlDecodeRequest bool
}

// Otomatik olarak bin kontrolü yapar. Eğer bin kontrol yapmak istemiyorsanız binRequest parametresini false yapabilirsiniz.
func WithBinRequest(binRequest bool) TDSOption {
	return func(o *tdsRequestOptions) {
		o.binRequest = binRequest
	}
}

// Otomatik olarak html içeriğini decode eder. Eğer decode etmek istemiyorsanız htmlDecodeRequest parametresini false yapabilirsiniz.
func WithHtmlDecodeRequest(htmlDecodeRequest bool) TDSOption {
	return func(o *tdsRequestOptions) {
		o.htmlDecodeRequest = htmlDecodeRequest
	}
}

// TROY, MASTERCARD, VISA ve AMEX markalı kartlarla yapılan işlemleri destekler.
// BONUS, WORLD, MAXIMUM, AXESS, CARDFINANS, PARAF, ADVANTAGE gibi taksit programlarına katılmış olan kartlara 2, 3, 6, 9 ve 12 taksit seçenekleri sunulmaktadır.
// Response da dönen paymentId saklanmalıdır.
//
// Bincontrol işlemi otomatik olarak yapılyor eğer bin kontrol yapmak istemiyorsanız binRequest parametresini false yapabilirsiniz.(Eğer kayıtlı kart ile ödeme yapılıyorsa bin kontrolü yapılmıyor.)
//
// Html içeriği otomatik olarak decode ediliyor eğer decode etmek istemiyorsanız htmlDecodeRequest parametresini false yapabilirsiniz.
func (i *Service) InitializeTDSPayment(ctx context.Context, req *InitTDSRequest, opts ...TDSOption) (response InitTDSResponse, decodedHtmlContent string, err error) {
	if req == nil {
		return response, decodedHtmlContent, transport.ErrNilRequest
	}

	options := &tdsRequestOptions{
		binRequest:        true, // default
		htmlDecodeRequest: true, // default
	}

	for _, opt := range opts {
		opt(options)
	}

	if req.PaymentCard.CardUserKey != "" && req.PaymentCard.CardToken != "" {
		options.binRequest = false
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, decodedHtmlContent, err
	}

	if options.binRequest && len(req.PaymentCard.CardNumber) < 6 {
		return response, decodedHtmlContent, errors.New("iyzipay: card number too short for BIN lookup")
	}
	if options.binRequest {
		binReq := cards.BinRequest{
			Locale:         req.Locale,
			BinNumber:      req.PaymentCard.CardNumber[:6],
			ConversationId: req.ConversationID,
		}
		binResp, err := (&cards.Service{Transport: i.Transport, Validator: i.Validator}).BinControlRequest(ctx, &binReq)
		if err != nil {
			return response, decodedHtmlContent, err
		}

		if binResp.Status == "failure" {
			return response, decodedHtmlContent, errors.New("bin control failed")
		}

		if !acceptedCardAssociations[binResp.CardAssociation] {
			return response, decodedHtmlContent, errors.New("card association not accepted")
		}
	}

	err = i.Transport.Do(ctx, http.MethodPost, tDSInitilizeURI, req, &response)
	if err != nil {
		return response, decodedHtmlContent, err
	}

	if options.htmlDecodeRequest {
		decodedHtmlContent, err = decodeHTML(response.ThreeDSHtmlContent)
		if err != nil {
			return response, decodedHtmlContent, errors.New("failed to decode html content")
		}
	}

	return response, decodedHtmlContent, nil
}

// 3DS işlemlerinin tamamlanması için kullanılır.
//
// ÖNEMLİ: Size geri dönen htmlde 3D Secure işlemini tamamlamanız gerekmektedir.
func (i *Service) FinalizeTDSPayment(ctx context.Context, req *TDSPaymentRequest) (response TDSPaymentResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, tDSFinalizeURI, req, &response)
	return response, err
}

func decodeHTML(value string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(value)
	return string(data), err
}
