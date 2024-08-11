package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// Checkout Fom ile ödeme isteği yapmak için kullanılır.
//
// Bu istekte dönecek olan html elementi için 3 farklı CFFormType belirleyebilirsiniz. "responsive", "popup", "iframe".
//
// Eğer iframe kullanılcaksa, dönen html elementini önemsemeyiniz.
func (i iyzipayClient) CheckoutFormPaymentRequest(req *CFRequest, CFFormType string) (response *CFResponse, HtmlElement string, err error) {
	if err = req.validate(); err != nil {
		return response, "", err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, "", errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, i.baseURI, i.apiKey, i.apiSecret, utils.CheckoutFormURI)
	if err != nil {
		return response, "", err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil {
		return response, "", errors.New("failed to unmarshal response")
	}

	if response.Status == "failure" {
		var respError errorModel
		err = json.Unmarshal(httpresp, &respError)
		if err != nil {
			return response, "", errors.New("failed to unmarshal response")
		}
		return response, "", errors.New(
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

	if utils.AcceptedCFTypes[CFFormType] {
		if CFFormType == "iframe" {
			response.PaymentPageUrl = response.PaymentPageUrl + "&iframe=true"
			return response, "", nil
		} else {
			htmlelement := "<div id='iyzipay-checkout-form' class='" + CFFormType + "'></div>"
			return response, htmlelement, nil
		}
	}

	return response, "", errors.New("CFFormType is not valid")
}

// Checkout Form ile yapılan ödeme isteği için sorgulama yapmak için kullanılır.
//
// Ödemenin tamamlanması gerekli.
func (i iyzipayClient) CheckoutFormPaymentInquiryRequest(req *CFInquiryRequest) (response *CFInquiryResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, i.baseURI, i.apiKey, i.apiSecret, utils.CheckoutFormInquiryURI)
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

	return response, err
}
