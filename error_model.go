package iyzipay

type errorModel struct {

	// İşlem sonucu başarılı ise success, başarısız ise failure döner.
	Status string `json:"status"`

	// İşlem başarısız ise hata kodunu döner.
	ErrorCode string `json:"errorCode"`

	// İşlem başarısız ise hata mesajını döner.
	ErrorMessage string `json:"errorMessage"`

	// Hata grubunu döner.
	Errorgroup string `json:"errorGroup"`

	// Requeste göre dönen response'un dili.
	Locale string `json:"locale"`

	// Sistem zamanı ?
	SystemTime int64 `json:"systemTime"`

	// İstek ve yanıtı eşleştirmek için kullanılan Konuşma ID'si.
	ConversationID string `json:"conversationId"`
}
