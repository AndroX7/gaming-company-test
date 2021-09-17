package module

import (
	artistHTTP "gaming-company-test/service/artist/delivery/http"
	artistRepository "gaming-company-test/service/artist/repository/mysql"
	artistUsecase "gaming-company-test/service/artist/usecase"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		artistRepository.New,
		artistUsecase.New,
		artistHTTP.New,
	),
)
