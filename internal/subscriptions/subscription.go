package subscriptions

import (
	"context"
)

// Abonelik ürünü oluşturmak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) CreateSubscriptionProduct(ctx context.Context, req CreateSubscriptionProductRequest) (response SubscriptionProductResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, "POST", subscriptionProductURI, req, &response)
	return response, err
}

// Abonelik ürünü güncellemek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) UpdateSubscriptionProduct(ctx context.Context, req UpdateSubscriptionProductRequest) (response SubscriptionProductResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	UpdatedPath := subscriptionProductURI + "/" + req.ProductReferenceCode

	err = i.Transport.Do(ctx, "POST", UpdatedPath, req, &response)
	return response, err
}

// Abonelik ürünü silmek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) DeleteSubscriptionProduct(ctx context.Context, req DeleteSubscriptionProductRequest) (response SubscriptionProductResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	deleteSubscriptionCustomURI := subscriptionProductURI + "/" + req.ProductReferenceCode
	err = i.Transport.Do(ctx, "DELETE", deleteSubscriptionCustomURI, nil, &response)
	return response, err
}

// Abonelik ürünü detaylarını getirmek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) GetSubscriptionProductDetail(ctx context.Context, req GetSubscriptionProductDetailRequest) (response SubscriptionProductResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	getSubscriptionDetailCustomURI := subscriptionProductURI + "/" + req.ProductReferenceCode
	err = i.Transport.Do(ctx, "GET", getSubscriptionDetailCustomURI, nil, &response)
	return response, err
}

// Abonelik ürünlerini listelemek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) GetSubscriptionProductList(ctx context.Context, req GetSubscriptionProductListRequest) (response GetSubscriptionProductListResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, "GET", subscriptionProductURI, req, &response)
	return response, err
}

// Abonelik ürünü için fiyatlandırma planı oluşturmak için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) CreateSubscriptionPlan(ctx context.Context, req CreateSubscriptionPlanRequest) (response SubscriptionPlanResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	subscriptionPlanURI := subscriptionProductURI + "/" + req.ProductReferenceCode + "/" + "pricing-plans"
	err = i.Transport.Do(ctx, "POST", subscriptionPlanURI, req, &response)
	return response, err
}

// Abonelik ürünü için fiyatlandırma planı güncellemek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) UpdateSubscriptionPlan(ctx context.Context, req UpdateSubscriptionPlanRequest) (response SubscriptionPlanResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, "POST", subscriptionPlanURI, req, &response)
	return response, err
}

// Abonelik ürünü için fiyatlandırma planı silmek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) DeleteSubscriptionPlan(ctx context.Context, req DeleteSubscriptionPlanRequest) (response DeleteSubscriptionPlanResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	deletePlanURI := subscriptionPlanURI + "/" + req.PricingPlanReferenceCode
	err = i.Transport.Do(ctx, "DELETE", deletePlanURI, req, &response)
	return response, err
}

// Abonelik ürünü için fiyatlandırma planı detaylarını getirmek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) GetSubscriptionPlanDetail(ctx context.Context, req GetSubscriptionPlanDetailRequest) (response SubscriptionPlanResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	getPlanURI := subscriptionPlanURI + "/" + req.PricingPlanReferenceCode
	err = i.Transport.Do(ctx, "GET", getPlanURI, req, &response)
	return response, err
}

// Abonelik ürünleri için fiyatlandırma planlarını listelemek için kullanılır.
//
// Not: Sandbox ortamında çalışmamaktadır.(Iyzico panelden eklentiler kısmından abonelik eklentisini aktif etmeniz gerekmektedir.)
func (i *Service) GetSubscriptionPlanList(ctx context.Context, req GetSubscriptionPlanListRequest) (response GetSubscriptionPlanListResponse, err error) {
	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, "GET", subscriptionPlanURI, req, &response)
	return response, err
}
