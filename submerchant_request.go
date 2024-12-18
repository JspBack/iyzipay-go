package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// Bireysel Alt Üye Oluşturma İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
//
// SubMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i *IyzipayClient) CreateIndividualSubMerchant(req *IndividualSubMerchantRequest) (response SubMerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil || response.Status != "success" {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}

// Şahıs Şirketi Alt Üye Oluşturma (Pazaryeri müşterisi olmanız gerekmektedir)
//
// SubMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i *IyzipayClient) CreatePrivateSubMerchant(req *PrivateSubMerchantRequest) (response SubMerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil || response.Status != "success" {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}

// Limited/Anonim Şirket Alt Üye Oluşturma (Pazaryeri müşterisi olmanız gerekmektedir)
//
// SubMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i *IyzipayClient) CreateLimitedCompanySubMerchant(req *LimitedCompanySubMerchantRequest) (response SubMerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil || response.Status != "success" {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}

// Bireysel Alt Üye Güncelleme İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i *IyzipayClient) UpdateIndividualSubMerchant(req *IndividualSubMerchantUpdateRequest) (response UpdateSubMerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPut, i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil || response.Status != "success" {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}

// Şahıs Şirketi Alt Üye Güncelleme İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i *IyzipayClient) UpdatePrivateSubMerchant(req *PrivateSubMerchantUpdateRequest) (response UpdateSubMerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPut, i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil || response.Status != "success" {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}

// Limited/Anonim Şirket Alt Üye Güncelleme İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i *IyzipayClient) UpdateLimitedCompanySubMerchant(req *LimitedCompanySubMerchantUpdateRequest) (response UpdateSubMerchantResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPut, i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil || response.Status != "success" {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}

// Alt Üye Sorgulama İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i *IyzipayClient) SubMerchantInquiry(req *SubMerchantInquiryRequest) (response SubMerchantInquiryResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPst, i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantInquiryURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil || response.Status != "success" {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}

func (i *IyzipayClient) UpdateSubMerchantProduct(req *SubMerchantProductUpdateRequest) (response SubMerchantProductUpdateResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, utils.RPut, i.baseURI, i.apiKey, i.apiSecret, utils.SubMerchantProductURI)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(httpresp, &response)
	if err != nil || response.Status != "success" {
		errorResp, err := utils.HandleError(httpresp)
		if err != nil {
			return response, err
		} else {
			return response, errors.Join(errors.New(errorResp.ErrorCode), errors.New(errorResp.ErrorMessage), errors.New(errorResp.Errorgroup), errors.New(errorResp.Locale), errors.New(errorResp.Status), errors.New(errorResp.ConversationID))
		}
	}

	return response, nil
}
