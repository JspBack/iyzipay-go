package cards

import (
	"context"
	"net/http"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

// Kart saklama işlemi için yeni kullanıcı ile kart saklama isteği
//
// Response içerisinde dönen cardUserKey ve cardToken saklanmalıdır.
func (i *Service) CreateCardStorageNewUser(ctx context.Context, req *CardStorageNewUserRequest) (response CardStorageNewUserResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, cardStorageURI, req, &response)
	return response, err
}

// Kart saklama işlemi için mevcut kullanıcı ile kart saklama isteği
//
// Response içerisinde dönen cardToken saklanmalıdır.
func (i *Service) CreateCardStorageExUser(ctx context.Context, req *CardStorageExUserRequest) (response CardStorageExUserResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, cardStorageURI, req, &response)
	return response, err
}

// Kart saklama işlemi için kart silme isteği
func (i *Service) DeleteCardStorageCard(ctx context.Context, req *CardStorageDeleteRequest) (response CardStorageDeleteResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodDelete, cardStorageURI, req, &response)
	return response, err
}

func (i *Service) GetCardStorageCards(ctx context.Context, req *CardStorageRetrieveRequest) (response CardStorageRetrieveResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, cardsStorageURI, req, &response)
	return response, err
}
