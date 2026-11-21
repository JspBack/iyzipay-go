package payments

const (
	nonTDSURI                = "/payment/auth"
	tDSFinalizeURI           = "/payment/3dsecure/auth"
	tDSInitilizeURI          = "/payment/3dsecure/initialize"
	inquiryURI               = "/payment/detail"
	cancelURI                = "/payment/cancel"
	refundToIyzicoAccountURI = "/payment/iyzipos/item/approve"
	refundURI                = "/payment/refund"
	refundV2URI              = "/v2/payment/refund"
	checkPWIURI              = "/payment/iyzipos/checkoutform/auth/ecom/detail"
	initPWIURI               = "/payment/pay-with-iyzico/initialize"
	checkoutFormInquiryURI   = "/payment/iyzipos/checkoutform/auth/ecom/detail"
	checkoutFormURI          = "/payment/iyzipos/checkoutform/initialize/auth/ecom"
)
