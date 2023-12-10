package swagger

import (
	"github.com/mycandys/orders/docs"
	"github.com/mycandys/orders/internal/env"
)

func InitInfo() {
	swaggerUri, _ := env.GetEnvVar(env.SWAGGER_URI)

	docs.SwaggerInfo.Title = "MyCandy's Orders Microservice API"
	docs.SwaggerInfo.Description = "This is MyCandy's Orders Microservice API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = swaggerUri

}
