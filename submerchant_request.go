package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// Bireysel Alt Üye Oluşturma İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
//
// subMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i iyzipayClient) CreateIndividualSubmerchant(req IndividualSubmerchantRequest) (response SubmerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, i.baseURI, i.apiKey, i.apiSecret, utils.SubmerchantCreateURI)
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

// Şahıs Şirketi Alt Üye Oluşturma (Pazaryeri müşterisi olmanız gerekmektedir)
//
// subMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i iyzipayClient) CreatePrivateSubmerchant(req PrivateSubmerchantRequest) (response SubmerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, i.baseURI, i.apiKey, i.apiSecret, utils.SubmerchantCreateURI)
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

// Limited/Anonim Şirket Alt Üye Oluşturma (Pazaryeri müşterisi olmanız gerekmektedir)
//
// subMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i iyzipayClient) CreateLimitedCompanySubmerchant(req LimitedCompanySubmerchantRequest) (response SubmerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, i.baseURI, i.apiKey, i.apiSecret, utils.SubmerchantCreateURI)
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
