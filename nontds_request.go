package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// 3D Secure olmayan ödeme işlemi için istek oluşturur
func (i Iyzipay) NTdsPaymentRequest(req *Non3DSPaymentRequest) (response Non3DSPaymentResponse, err error) {
	if err = req.validate(); err != nil {
		return Non3DSPaymentResponse{}, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return Non3DSPaymentResponse{}, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, i.baseURI, i.apiKey, i.apiSecret, utils.NonTDSURI)
	if err != nil {
		return Non3DSPaymentResponse{}, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil {
		return Non3DSPaymentResponse{}, errors.New("failed to unmarshal response")
	}

	// Mükemmel hata yönetimi XD
	if response.Status == "failure" {
		var respError errorModel
		err = json.Unmarshal(httpresp, &respError)
		if err != nil {
			return Non3DSPaymentResponse{}, errors.New("failed to unmarshal response")
		}
		return Non3DSPaymentResponse{}, errors.New(
			"{" +
				"Status: " + respError.Status + "," +
				"ErrorCode: " + respError.ErrorCode + "," +
				"ErrorMessage: " + respError.ErrorMessage + "," +
				"Locale: " + respError.Locale + "," +
				"SystemTime: " + string(rune(respError.SystemTime)) + "," +
				"ConversationID: " + respError.ConversationID +
				"}",
		)
	}

	return response, nil
}
