package iyzilink

import (
	"context"
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

	path := iyzilinkURI + "/" + url.PathEscape(req.Token)
	err = i.Transport.Do(ctx, http.MethodPut, path, req, &response)
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

	path := linkPath(req.Token, req.Locale, req.ConversationID)
	err = i.Transport.Do(ctx, http.MethodDelete, path, nil, &response)
	return response, err
}

// IyzilinkGetDetail, bağlantı bilgilerini token ile sorgular.
func (i *Service) IyzilinkGetDetail(ctx context.Context, req *IyzilinkGetDetailRequest) (response IyzilinkGetDetailResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	path := linkPath(req.Token, req.Locale, req.ConversationID)
	err = i.Transport.Do(ctx, http.MethodGet, path, nil, &response)
	return response, err
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

func linkPath(token, locale, conversationID string) string {
	path := iyzilinkURI + "/" + url.PathEscape(token)
	query := url.Values{}
	if locale != "" {
		query.Set("locale", locale)
	}
	if conversationID != "" {
		query.Set("conversationId", conversationID)
	}
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	return path
}
