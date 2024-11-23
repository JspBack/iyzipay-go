package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

type tdsOptions func(*tdsRequestOptions)

type tdsRequestOptions struct {
	binRequest        bool
	htmlDecodeRequest bool
}

// Otomatik olarak bin kontrolü yapar. Eğer bin kontrol yapmak istemiyorsanız binRequest parametresini false yapabilirsiniz.
func WithBinRequest(binRequest bool) tdsOptions {
	return func(o *tdsRequestOptions) {
		o.binRequest = binRequest
	}
}

// Otomatik olarak html içeriğini decode eder. Eğer decode etmek istemiyorsanız htmlDecodeRequest parametresini false yapabilirsiniz.
func WithHtmlDecodeRequest(htmlDecodeRequest bool) tdsOptions {
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
func (i *IyzipayClient) InitializeTDSPayment(req *InitTDSRequest, opts ...tdsOptions) (response InitTDSResponse, decodedHtmlContent string, err error) {
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

	if err = req.validate(); err != nil {
		return response, decodedHtmlContent, err
	}

	if options.binRequest {
		binReq := BinRequest{
			Locale:         req.Locale,
			BinNumber:      req.PaymentCard.CardNumber[:6],
			ConversationId: req.ConversationID,
		}
		binResp, err := i.BinControlRequest(&binReq)
		if err != nil {
			return response, decodedHtmlContent, err
		}

		if binResp.Status == "failure" {
			return response, decodedHtmlContent, errors.New("bin control failed")
		}

		if !utils.TDSAcceptedCardAssociations[binResp.CardAssociation] {
			return response, decodedHtmlContent, errors.New("card association not accepted")
		}

	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, decodedHtmlContent, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.TDSInitilizeURI)
	if err != nil {
		return response, decodedHtmlContent, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, "", err
		} else {
			return response, "", errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	if options.htmlDecodeRequest {
		decodedHtmlContent, err = utils.Base64Decode(response.ThreeDSHtmlContent)
		if err != nil {
			return response, decodedHtmlContent, errors.New("failed to decode html content")
		}
	}

	return response, decodedHtmlContent, nil
}

// 3DS işlemlerinin tamamlanması için kullanılır.
//
// ÖNEMLİ: Size geri dönen htmlde 3D Secure işlemini tamamlamanız gerekmektedir.
func (i *IyzipayClient) FinalizeTDSPayment(req *TDSPaymentRequest) (response TDSPaymentResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.TDSFinalizeURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}
