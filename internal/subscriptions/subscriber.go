package subscriptions

import (
	"context"
	"net/http"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

// Abone bilgilerini güncellemek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) UpdateSubscriber(ctx context.Context, req *UpdateSubscriberRequest) (response SubscriberResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	updateSbscrbrURI := subscriptionSubscribersURI + "/" + req.CustomerReferenceCode
	err = i.Transport.Do(ctx, http.MethodPost, updateSbscrbrURI, req, &response)
	return response, err
}

// Abone bilgilerini sorgulamak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) GetSubscriberDetail(ctx context.Context, req *GetSubscriberDetailRequest) (response SubscriberResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	getSbscrbrURI := subscriptionSubscribersURI + "/" + req.CustomerReferenceCode
	err = i.Transport.Do(ctx, http.MethodGet, getSbscrbrURI, req, &response)
	return response, err
}

// Abone listesini sorgulamak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) GetSubscriberList(ctx context.Context, req *GetSubscriberListRequest) (response GetSubscriberListRequestResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodGet, subscriptionSubscribersURI, req, &response)
	return response, err
}
