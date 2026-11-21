package iyzilink

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

// Iyzilink oluşturma isteği
func (i *Service) IyzilinkCreate(ctx context.Context, req *IyzilinkCreateRequest) (response IyzilinkResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, iyzilinkURI, req, &response)
	return response, err
}

// Iyzilink güncelleme isteği
func (i *Service) IyzilinkUpdate(ctx context.Context, req *IyzilinkUpdateRequest) (response IyzilinkResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPut, iyzilinkURI, req, &response)
	return response, err
}

// Iyzilink durum güncelleme isteği
func (i *Service) IyzilinkStatusUpdate(ctx context.Context, req *IyzilinkStatusUpdateRequest) (response IyzilinkStatusUpdateResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	updateStatusURI := iyzilinkURI + "/" + req.Token + "/status/" + req.ProductStatus

	err = i.Transport.Do(ctx, http.MethodPatch, updateStatusURI, req, &response)
	return response, err
}

// Iyzilink silme isteği
func (i *Service) IyzilinkDelete(ctx context.Context, req *IyzilinkDeleteRequest) (response IyzilinkDeleteResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	path := iyzilinkURI + "/" + url.PathEscape(req.Token)
	query := url.Values{}
	if req.Locale != "" {
		query.Set("locale", req.Locale)
	}
	if req.ConversationID != "" {
		query.Set("conversationId", req.ConversationID)
	}
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	err = i.Transport.Do(ctx, http.MethodDelete, path, nil, &response)
	return response, err
}

// Iyzilink detay sorgulama isteği
// Token bazlı çalışmıyor nedense o yüzden clientta filtreleme ile çözdüm(şimdilik).[liste içerisinden token ile arama yapılıyor]
// Veri sayısı çok olursa bu yöntem performansı düşürebilir.
func (i *Service) IyzilinkGetDetail(ctx context.Context, req *IyzilinkGetDetailRequest) (response IyzilinkGetDetailResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	resp, err := i.IyzilinkGetList(ctx, &IyzilinkGetListRequest{
		ConversationID: req.ConversationID,
		Locale:         req.Locale,
		Page:           req.Page,
		Count:          req.Count,
	})
	if err != nil {
		return response, err
	}

	for _, item := range resp.Data.Items {
		if item.Token == req.Token {
			response = IyzilinkGetDetailResponse{
				Status: "success",
				Locale: req.Locale,
				Data:   item,
			}
			return response, nil
		}
	}
	return response, errors.New("token not found")
}

// Iyzilink listeleme isteği
func (i *Service) IyzilinkGetList(ctx context.Context, req *IyzilinkGetListRequest) (response IyzilinkGetListResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodGet, iyzilinkURI, req, &response)
	return response, err
}
