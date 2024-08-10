package utils

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Random string oluşturmak için kullanılan karakterler
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits

	NonTDSURI     = "/payment/auth"      // 3DS olmayan ödeme isteği için URI yolu
	TDSURI        = "/payment/3d"        // 3DS ödeme isteği için URI yolu
	InguiryURI    = "/payment/detail"    // Ödeme sorgulama isteği için URI yolu
	BinControlURI = "/payment/bin/check" // Bin kontrol isteği için URI yolu
)
