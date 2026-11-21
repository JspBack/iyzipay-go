package marketplace

import (
	"context"
	"net/http"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

// Bireysel Alt Üye Oluşturma İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
//
// SubMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i *Service) CreateIndividualSubMerchant(ctx context.Context, req *IndividualSubMerchantRequest) (response SubMerchantResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return
	}

	err = i.Transport.Do(ctx, http.MethodPost, subMerchantURI, req, &response)
	return response, err
}

// Şahıs Şirketi Alt Üye Oluşturma (Pazaryeri müşterisi olmanız gerekmektedir)
//
// SubMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i *Service) CreatePrivateSubMerchant(ctx context.Context, req *PrivateSubMerchantRequest) (response SubMerchantResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return
	}

	err = i.Transport.Do(ctx, http.MethodPost, subMerchantURI, req, &response)
	return response, err
}

// Limited/Anonim Şirket Alt Üye Oluşturma (Pazaryeri müşterisi olmanız gerekmektedir)
//
// SubMerchantKey alanı için dönen sonuç saklanmalıdır.
func (i *Service) CreateLimitedCompanySubMerchant(ctx context.Context, req *LimitedCompanySubMerchantRequest) (response SubMerchantResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return
	}

	err = i.Transport.Do(ctx, http.MethodPost, subMerchantURI, req, &response)
	return response, err
}

// Bireysel Alt Üye Güncelleme İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i *Service) UpdateIndividualSubMerchant(ctx context.Context, req *IndividualSubMerchantUpdateRequest) (response UpdateSubMerchantResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPut, subMerchantURI, req, &response)
	return response, err
}

// Şahıs Şirketi Alt Üye Güncelleme İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i *Service) UpdatePrivateSubMerchant(ctx context.Context, req *PrivateSubMerchantUpdateRequest) (response UpdateSubMerchantResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPut, subMerchantURI, req, &response)
	return response, err
}

// Limited/Anonim Şirket Alt Üye Güncelleme İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i *Service) UpdateLimitedCompanySubMerchant(ctx context.Context, req *LimitedCompanySubMerchantUpdateRequest) (response UpdateSubMerchantResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPut, subMerchantURI, req, &response)
	return response, err
}

// Alt Üye Sorgulama İsteği (Pazaryeri müşterisi olmanız gerekmektedir)
func (i *Service) SubMerchantInquiry(ctx context.Context, req *SubMerchantInquiryRequest) (response SubMerchantInquiryResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, subMerchantInquiryURI, req, &response)
	return response, err
}

func (i *Service) UpdateSubMerchantProduct(ctx context.Context, req *SubMerchantProductUpdateRequest) (response SubMerchantProductUpdateResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPut, subMerchantProductURI, req, &response)
	return response, err
}
