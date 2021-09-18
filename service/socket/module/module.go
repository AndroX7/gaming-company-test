package module

import (
	"go.uber.org/fx"

	socketHTTP "gaming-company-test/service/socket/delivery/http"
	socketUsecase "gaming-company-test/service/socket/usecase"
)

var Module = fx.Options(
	fx.Provide(
		socketHTTP.New,
		socketUsecase.New,
	),
)
