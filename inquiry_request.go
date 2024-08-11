package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// İşlem sorgulama isteği
func (i iyzipayClient) PaymentInquiryRequest(req *InquiryRequest) (response InquiryResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, utils.InguiryURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil {
		return response, errors.New("failed to unmarshal response")
	}

	if response.Status == "failure" {
		var respError errorModel
		err = json.Unmarshal(httpresp, &respError)
		if err != nil {
			return response, errors.New("failed to unmarshal response")
		}
		return response, errors.New(
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
