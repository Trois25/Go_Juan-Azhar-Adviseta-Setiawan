package route

import(
	"github.com/labstack/echo/v4"
	"ai-task/controller"
	"ai-task/usecase"
)

func SetLaptopRecomendedRoutes(e *echo.Echo) {
    RecomendationLaptopUseCase := usecase.NewUseCase()
    RecomendationLaptopController := controller.NewRecomendationLaptopController(RecomendationLaptopUseCase)

    e.POST("/laptop", RecomendationLaptopController.GetLaptopRecomended)
}