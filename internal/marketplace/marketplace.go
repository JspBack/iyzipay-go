package marketplace

import (
	"context"
	"net/http"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

// Pazaryeri çözümünde, ödeme iyzico’dan geçtikten sonra, üye işyeri ödeme içinde yer alan kırılıma / ürüne onay verene dek para korumalı havuz hesapta bekletilir.
// Üye işyeri bu sürede ödemeyi iptal edebilir, ödemenin kırılımını iade edebilir ya da ürün alıcıya ulaştı ve işlem sorunsuz tamamlandıysa para transferi için ürüne onay verebilir veya verdiği ürün onayını geri çekebilir.
//
// Ürün onayı vermek için kullanılır.
func (i *Service) ApproveProduct(ctx context.Context, req *MarketplaceProductRequest) (response MarketplaceProductResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, approveProductURI, req, &response)
	return response, err
}

// Ürün onayını geri çekmek için kullanılır.
func (i *Service) DisapproveProduct(ctx context.Context, req *MarketplaceProductRequest) (response MarketplaceProductResponse, err error) {
	if req == nil {
		return response, transport.ErrNilRequest
	}

	if err = i.Validator.Struct(req); err != nil {
		return response, err
	}

	err = i.Transport.Do(ctx, http.MethodPost, disapproveProductURI, req, &response)
	return response, err
}
