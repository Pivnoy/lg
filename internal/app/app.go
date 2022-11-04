package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lg/config"
	v1 "lg/internal/controller/http/v1"
	"lg/internal/usecase"
	"lg/internal/usecase/repo"
	"lg/pkg/httpserver"
	"lg/pkg/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	pg, err := postgres.New(cfg)

	if err != nil {
		log.Fatal("Error in creating postgres instance")
	}
	lineUpUseCase := usecase.NewLineupUseCase(repo.NewLineupRepo(pg))
	projectUseCase := usecase.NewProjectUseCase(repo.NewProjectRepo(pg), lineUpUseCase)
	userUseCase := usecase.NewUserUseCase(repo.NewUserRepo(pg))
	signInUseCase := usecase.NewSignInUseCase(userUseCase)
	jwtUseCase := usecase.NewJwtUseCase(userUseCase, cfg.SecretKey)
	profileUseCase := usecase.NewProfileUseCase(repo.NewProfileRepo(pg))
	countryUseCase := usecase.NewCountryUseCase(repo.NewCountryRepo(pg))
	citizenshipUseCase := usecase.NewCitizenshipUseCase(repo.NewCitizenshipRepo(pg))
	eduspecialityUseCase := usecase.NewEduspecialityUseCase(repo.NewEduspecialityRepo(pg))
	employmentUseCase := usecase.NewEmploymentUseCase(repo.NewEmploymentRepo(pg))
	specializationUseCase := usecase.NewSpecializationUseCase(repo.NewSpecializationRepo(pg))
	universityUseCase := usecase.NewUniversityUseCase(repo.NewUniversityRepo(pg))
	cityUseCase := usecase.NewCityUseCase(repo.NewCityRepo(pg))

	handler := gin.New()

	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1.NewRouter(handler,
		projectUseCase,
		signInUseCase,
		jwtUseCase,
		userUseCase,
		profileUseCase,
		countryUseCase,
		citizenshipUseCase,
		eduspecialityUseCase,
		employmentUseCase,
		specializationUseCase,
		universityUseCase,
		cityUseCase)

	serv := httpserver.New(handler, httpserver.Port(cfg.AppPort))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err = <-serv.Notify():
		log.Printf("Notify from http server")
	}

	err = serv.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}
}
