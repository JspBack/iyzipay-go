package iyzipay

import (
	"github.com/JspBack/iyzipay-go/internal/app2app"
	"github.com/JspBack/iyzipay-go/internal/cards"
	"github.com/JspBack/iyzipay-go/internal/iyzilink"
	"github.com/JspBack/iyzipay-go/internal/marketplace"
	"github.com/JspBack/iyzipay-go/internal/payments"
	"github.com/JspBack/iyzipay-go/internal/subscriptions"
)

type (
	PaymentRequest              = payments.PaymentRequest
	Buyer                       = payments.Buyer
	PaymentCard                 = payments.PaymentCard
	BillingAddress              = payments.BillingAddress
	ShippingAddress             = payments.ShippingAddress
	BasketItem                  = payments.BasketItem
	Non3DSPaymentResponse       = payments.Non3DSPaymentResponse
	ItemTransaction             = payments.ItemTransaction
	ConvertedPayout             = payments.ConvertedPayout
	InitTDSRequest              = payments.InitTDSRequest
	InitTDSResponse             = payments.InitTDSResponse
	TDSPaymentRequest           = payments.TDSPaymentRequest
	TDSPaymentResponse          = payments.TDSPaymentResponse
	InquiryRequest              = payments.InquiryRequest
	InquiryResponse             = payments.InquiryResponse
	RefundPaymentRequest        = payments.RefundPaymentRequest
	RefundPaymentV2Request      = payments.RefundPaymentV2Request
	CancelPaymentRequest        = payments.CancelPaymentRequest
	RefundPaymentResponse       = payments.RefundPaymentResponse
	RefundPaymentIyzicoRequest  = payments.RefundPaymentIyzicoRequest
	CancelPaymentResponse       = payments.CancelPaymentResponse
	RefundPaymentIyzicoResponse = payments.RefundPaymentIyzicoResponse
	InitPWIRequest              = payments.InitPWIRequest
	InitPWIResponse             = payments.InitPWIResponse
	PWIInquiryRequest           = payments.PWIInquiryRequest
	PWIInquiryResponse          = payments.PWIInquiryResponse
	CFRequest                   = payments.CFRequest
	CFResponse                  = payments.CFResponse
	CFInquiryRequest            = payments.CFInquiryRequest
	CFInquiryResponse           = payments.CFInquiryResponse
)

type (
	BinRequest                  = cards.BinRequest
	InstallmentRequest          = cards.InstallmentRequest
	BinResponse                 = cards.BinResponse
	InstallmentResponse         = cards.InstallmentResponse
	InstallmentDetail           = cards.InstallmentDetail
	InstallmentPrice            = cards.InstallmentPrice
	CardStorageCard             = cards.CardStorageCard
	StoredCard                  = cards.StoredCard
	CardStorageNewUserRequest   = cards.CardStorageNewUserRequest
	CardStorageExUserRequest    = cards.CardStorageExUserRequest
	CardStorageRetrieveRequest  = cards.CardStorageRetrieveRequest
	CardStorageDeleteRequest    = cards.CardStorageDeleteRequest
	CardStorageResponse         = cards.CardStorageResponse
	CardStorageNewUserResponse  = cards.CardStorageNewUserResponse
	CardStorageExUserResponse   = cards.CardStorageExUserResponse
	CardStorageDeleteResponse   = cards.CardStorageDeleteResponse
	CardStorageRetrieveResponse = cards.CardStorageRetrieveResponse
)

type (
	MarketplaceProductRequest              = marketplace.MarketplaceProductRequest
	MarketplaceProductResponse             = marketplace.MarketplaceProductResponse
	IndividualSubMerchantUpdateRequest     = marketplace.IndividualSubMerchantUpdateRequest
	PrivateSubMerchantUpdateRequest        = marketplace.PrivateSubMerchantUpdateRequest
	LimitedCompanySubMerchantUpdateRequest = marketplace.LimitedCompanySubMerchantUpdateRequest
	IndividualSubMerchantRequest           = marketplace.IndividualSubMerchantRequest
	PrivateSubMerchantRequest              = marketplace.PrivateSubMerchantRequest
	LimitedCompanySubMerchantRequest       = marketplace.LimitedCompanySubMerchantRequest
	SubMerchantInquiryRequest              = marketplace.SubMerchantInquiryRequest
	SubMerchantProductUpdateRequest        = marketplace.SubMerchantProductUpdateRequest
	SubMerchantResponse                    = marketplace.SubMerchantResponse
	UpdateSubMerchantResponse              = marketplace.UpdateSubMerchantResponse
	SubMerchantInquiryResponse             = marketplace.SubMerchantInquiryResponse
	SubMerchantProductUpdateResponse       = marketplace.SubMerchantProductUpdateResponse
)

type (
	CreateSubscriptionProductRequest           = subscriptions.CreateSubscriptionProductRequest
	UpdateSubscriptionProductRequest           = subscriptions.UpdateSubscriptionProductRequest
	DeleteSubscriptionProductRequest           = subscriptions.DeleteSubscriptionProductRequest
	GetSubscriptionProductDetailRequest        = subscriptions.GetSubscriptionProductDetailRequest
	GetSubscriptionProductListRequest          = subscriptions.GetSubscriptionProductListRequest
	CreateSubscriptionPlanRequest              = subscriptions.CreateSubscriptionPlanRequest
	UpdateSubscriptionPlanRequest              = subscriptions.UpdateSubscriptionPlanRequest
	DeleteSubscriptionPlanRequest              = subscriptions.DeleteSubscriptionPlanRequest
	GetSubscriptionPlanDetailRequest           = subscriptions.GetSubscriptionPlanDetailRequest
	GetSubscriptionPlanListRequest             = subscriptions.GetSubscriptionPlanListRequest
	SubscriptionPlanResponse                   = subscriptions.SubscriptionPlanResponse
	DeleteSubscriptionPlanResponse             = subscriptions.DeleteSubscriptionPlanResponse
	GetSubscriptionPlanListResponse            = subscriptions.GetSubscriptionPlanListResponse
	GetSubscriptionProductListResponse         = subscriptions.GetSubscriptionProductListResponse
	SubscriptionProductResponse                = subscriptions.SubscriptionProductResponse
	Subscription                               = subscriptions.Subscription
	PricingPlan                                = subscriptions.PricingPlan
	SubscriptionPlan                           = subscriptions.SubscriptionPlan
	PaginatedSubscription                      = subscriptions.PaginatedSubscription
	PaginatedSubscriptionPlan                  = subscriptions.PaginatedSubscriptionPlan
	InitSubscriptionWithCheckoutFormRequest    = subscriptions.InitSubscriptionWithCheckoutFormRequest
	InquirySubscriptionWithCheckoutFormRequest = subscriptions.InquirySubscriptionWithCheckoutFormRequest
	InitSubscriptionWithNonTDSRequest          = subscriptions.InitSubscriptionWithNonTDSRequest
	InitSubscriptionWithCheckoutFormResponse   = subscriptions.InitSubscriptionWithCheckoutFormResponse
	SubscriptionPaymentResponse                = subscriptions.SubscriptionPaymentResponse
	CfSubscriptionData                         = subscriptions.CfSubscriptionData
	UpdateSubscriberRequest                    = subscriptions.UpdateSubscriberRequest
	GetSubscriberDetailRequest                 = subscriptions.GetSubscriberDetailRequest
	GetSubscriberListRequest                   = subscriptions.GetSubscriberListRequest
	SubscriberResponse                         = subscriptions.SubscriberResponse
	GetSubscriberListRequestResponse           = subscriptions.GetSubscriberListRequestResponse
	SubscriberData                             = subscriptions.SubscriberData
	PaginatedSubscriberData                    = subscriptions.PaginatedSubscriberData
)

type (
	IyzilinkCreateRequest        = iyzilink.IyzilinkCreateRequest
	ProductBuyerInfo             = iyzilink.ProductBuyerInfo
	IyzilinkUpdateRequest        = iyzilink.IyzilinkUpdateRequest
	IyzilinkDeleteRequest        = iyzilink.IyzilinkDeleteRequest
	IyzilinkGetDetailRequest     = iyzilink.IyzilinkGetDetailRequest
	IyzilinkGetListRequest       = iyzilink.IyzilinkGetListRequest
	IyzilinkStatusUpdateRequest  = iyzilink.IyzilinkStatusUpdateRequest
	IyzilinkResponse             = iyzilink.IyzilinkResponse
	ResponseData                 = iyzilink.ResponseData
	IyzilinkDeleteResponse       = iyzilink.IyzilinkDeleteResponse
	IyzilinkGetDetailResponse    = iyzilink.IyzilinkGetDetailResponse
	DetailResponseData           = iyzilink.DetailResponseData
	IyzilinkGetListResponse      = iyzilink.IyzilinkGetListResponse
	ListResponseData             = iyzilink.ListResponseData
	IyzilinkStatusUpdateResponse = iyzilink.IyzilinkStatusUpdateResponse
)

type (
	A2ACreatePaymentRequest   = app2app.A2ACreatePaymentRequest
	A2ARefundPaymentRequest   = app2app.A2ARefundPaymentRequest
	A2AInquiryPaymentRequest  = app2app.A2AInquiryPaymentRequest
	A2AInquiryPaymentResponse = app2app.A2AInquiryPaymentResponse
	InStoreCompleteOperation  = app2app.InStoreCompleteOperation
	Transaction               = app2app.Transaction
	Receipt                   = app2app.Receipt
	DetailItem                = app2app.DetailItem
	A2AListResponse           = app2app.A2AListResponse
	UserInfo                  = app2app.UserInfo
	A2ARefundPaymentResponse  = app2app.A2ARefundPaymentResponse
	A2ACreatePaymentResponse  = app2app.A2ACreatePaymentResponse
)
