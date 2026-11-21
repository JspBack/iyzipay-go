package iyzipay

import "github.com/JspBack/iyzipay-go/internal/transport"

// APIError, HTTP hatalarında ve ödeme sağlayıcısının bildirdiği başarısız işlemlerde döndürülür.
type APIError = transport.APIError

// ErrNilRequest, zorunlu istek parametresinin nil olduğunu belirtir.
var ErrNilRequest = transport.ErrNilRequest
