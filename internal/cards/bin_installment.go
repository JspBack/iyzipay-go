package cards

import (
	"context"
	"net/http"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

// Kart bilgilerini öğrenmek için kullanılır.
func (i *Service) BinControlRequest(ctx context.Context, req *BinRequest) (response BinResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, binCheckURI, req, &response)
	return response, err
}

// Kartın taksit bilgilerini sorgulamak için kullanılır.
//
// Talepte BIN numarası belirtilmemişse, mevcut tüm taksit seçenekleri görüntülenecektir. Ancak, BIN numarası verilirse, yalnızca ilgili ve o karta özel taksit seçeneklerini içerecektir.
//
// force3ds değeri 1 olarak dönerse, işlemin 3DS ile işlenmesi gerektiği anlamına gelir. 0 ise tercihlere göre işlem yapılabilir. İşyeri hesabında 3DS zorunlu olarak ayarlanmışsa, bu değer sürekli olarak 1 döndürür.
func (i *Service) InstallmentControlRequest(ctx context.Context, req *InstallmentRequest) (response InstallmentResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, installmentCheckURI, req, &response)
	return response, err
}
