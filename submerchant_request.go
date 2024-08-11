package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// Bireysel Alt Üye Oluşturma İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
//
// SubMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i iyzipayClient) CreateIndividualSubMerchant(req IndividualSubMerchantRequest) (response SubMerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantURI)
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
// SubMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i iyzipayClient) CreatePrivateSubMerchant(req PrivateSubMerchantRequest) (response SubMerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantURI)
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
// SubMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i iyzipayClient) CreateLimitedCompanySubMerchant(req LimitedCompanySubMerchantRequest) (response SubMerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantURI)
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

// Bireysel Alt Üye Güncelleme İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i iyzipayClient) UpdateIndividualSubMerchant(req IndividualSubMerchantUpdateRequest) (response UpdateSubMerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "PUT", i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantURI)
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

// Şahıs Şirketi Alt Üye Güncelleme İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i iyzipayClient) UpdatePrivateSubMerchant(req PrivateSubMerchantUpdateRequest) (response UpdateSubMerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "PUT", i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantURI)
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

// Limited/Anonim Şirket Alt Üye Güncelleme İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i iyzipayClient) UpdateLimitedCompanySubMerchant(req LimitedCompanySubMerchantUpdateRequest) (response UpdateSubMerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "PUT", i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantURI)
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

// Alt Üye Sorgulama İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i iyzipayClient) SubMerchantInquiry(req SubMerchantInquiryRequest) (response SubMerchantInquiryResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantInquiryURI)
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

func (i iyzipayClient) UpdateSubMerchantProduct(req SubMerchantProductUpdateRequest) (response SubMerchantProductUpdateResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "PUT", i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantProductURI)
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
