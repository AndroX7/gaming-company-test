package module

import (
	"go.uber.org/fx"

	responseCacheUsecase "gaming-company-test/service/response_cache/usecase"
)

var Module = fx.Options(
	fx.Provide(
		responseCacheUsecase.New,
	),
)
