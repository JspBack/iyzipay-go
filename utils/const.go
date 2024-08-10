package utils

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Random string oluşturmak için kullanılan karakterler
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits

	NonTDSURI       = "/payment/auth"                                  // 3DS olmayan ödeme isteği için URI yolu
	InguiryURI      = "/payment/detail"                                // Ödeme sorgulama isteği için URI yolu
	BinControlURI   = "/payment/bin/check"                             // Bin kontrol isteği için URI yolu
	TDSInitilizeURI = "/payment/3dsecure/initialize"                   // 3DS ödeme başlatma isteği için URI yolu
	TDSFinalizeURI  = "/payment/3dsecure/auth"                         // 3DS ödeme isteği için URI yolu
	InitPWIURI      = "/payment/pay-with-iyzico/initialize"            // PWI başlatma isteği için URI yolu
	CheckPWIURI     = "/payment/iyzipos/checkoutform/auth/ecom/detail" // PWI ödeme isteği için URI yolu
)

var (
	TDSAcceptedCardAssociations = map[string]bool{ // 3DS işlemi için desteklenen kart tipleri
		"TROY":        true,
		"MASTER_CARD": true,
		"MASTERCARD":  true,
		"VISA":        true,
		"AMEX":        true,
	}
)
