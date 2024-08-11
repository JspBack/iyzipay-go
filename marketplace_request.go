package iyzipay

import (
	"encoding/json"
	"errors"

	"github.com/JspBack/iyzipay-go/utils"
)

// Pazaryeri çözümünde, ödeme iyzico’dan geçtikten sonra, üye işyeri ödeme içinde yer alan kırılıma / ürüne onay verene dek para korumalı havuz hesapta bekletilir.
// Üye işyeri bu sürede ödemeyi iptal edebilir, ödemenin kırılımını iade edebilir ya da ürün alıcıya ulaştı ve işlem sorunsuz tamamlandıysa para transferi için ürüne onay verebilir veya verdiği ürün onayını geri çekebilir.
//
// Ürün onayı vermek için kullanılır.
func (i iyzipayClient) ApproveProduct(req MarketplaceProductRequest) (response MarketplaceProductResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, utils.ApproveProductURI)
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

// Ürün onayını geri çekmek için kullanılır.
func (i iyzipayClient) DisapproveProduct(req MarketplaceProductRequest) (response MarketplaceProductResponse, err error) {
	if err = req.validate(); err != nil {
		return response, err
	}

	requestData, err := json.Marshal(req)
	if err != nil {
		return response, errors.New("failed to marshal request")
	}

	httpresp, err := utils.DoRequest(requestData, i.client, "POST", i.baseURI, i.apiKey, i.apiSecret, utils.DisapproveProductURI)
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
