package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"medical-record-api/internal/auth"
	"medical-record-api/internal/doctor"
	mdcHandler "medical-record-api/internal/medicine/handler"
	mdcRepo "medical-record-api/internal/medicine/repository"
	mdcService "medical-record-api/internal/medicine/service"
	"medical-record-api/internal/menu"
	"medical-record-api/internal/nurse"
	ptHandler "medical-record-api/internal/patient/handler"
	ptRepo "medical-record-api/internal/patient/repository"
	ptService "medical-record-api/internal/patient/service"
	"medical-record-api/internal/specialization"
	"medical-record-api/pkg/errors"
	"time"
)

func SetupRouter(db *sqlx.DB) *gin.Engine {
	// init validator
	errors.InitValidator()

	r := gin.Default()

	// middleware cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// error handler
	r.Use(errors.ErrorHandler())

	// auth routes
	userRepo := auth.NewRepository(db)
	authService := auth.NewService(userRepo)
	authHandler := auth.NewHandler(authService)
	authGroup := r.Group("/api/v1/auth")
	RegisterAuthRoutes(authGroup, authHandler)

	// menu routes
	menuRepo := menu.NewRepository(db)
	menuService := menu.NewService(menuRepo)
	menuHandler := menu.NewHandler(menuService)
	menuGroup := r.Group("api/v1/menus")
	RegisterMenuRoutes(menuGroup, menuHandler)

	// specialization routes
	specializationRepo := specialization.NewRepository(db)
	specializationService := specialization.NewService(specializationRepo)
	specializationHandler := specialization.NewHandler(specializationService)
	specializationGroup := r.Group("/api/v1/specializations")
	RegisterSpecializationRoutes(specializationGroup, specializationHandler)

	// doctor routes
	doctorRepo := doctor.NewRepository(db)
	doctorService := doctor.NewService(doctorRepo)
	doctorHandler := doctor.NewHandler(doctorService)
	doctorGroup := r.Group("/api/v1/doctors")
	RegisterDoctorRoutes(doctorGroup, doctorHandler)

	// nurse routes
	nurseRepo := nurse.NewRepository(db)
	nurseService := nurse.NewService(nurseRepo)
	nurseHandler := nurse.NewHandler(nurseService)
	nurseGroup := r.Group("/api/v1/nurses")
	RegisterNurseRoutes(nurseGroup, nurseHandler)

	// medicine routes
	medicineGroup := r.Group("api/v1/medicines")

	unitRepo := mdcRepo.NewUnitRepository(db)
	unitService := mdcService.NewUnitService(unitRepo)
	unitHandler := mdcHandler.NewUnitHandler(unitService)
	RegisterUnitRoutes(medicineGroup, unitHandler)

	categoryRepo := mdcRepo.NewCategoryRepository(db)
	categoryService := mdcService.NewCategoryService(categoryRepo)
	categoryHandler := mdcHandler.NewCategoryHandler(categoryService)
	RegisterCategoryRoutes(medicineGroup, categoryHandler)

	medicineRepo := mdcRepo.NewRepository(db)
	medicineService := mdcService.NewService(medicineRepo)
	medicineHandler := mdcHandler.NewHandler(medicineService)
	RegisterMedicineRoutes(medicineGroup, medicineHandler)

	// patient routes
	patientGroup := r.Group("/api/v1/patients")

	patientRepo := ptRepo.NewRepository(db)
	patientService := ptService.NewService(patientRepo)
	patientHandler := ptHandler.NewHandler(patientService)
	RegisterPatientRoutes(patientGroup, patientHandler)

	emergencyContactRepo := ptRepo.NewEmergencyContactRepository(db)
	emergencyContactService := ptService.NewEmergencyContactService(emergencyContactRepo)
	emergencyContactHandler := ptHandler.NewEmergencyContactHandler(emergencyContactService)
	RegisterEmergencyContactRoutes(patientGroup, emergencyContactHandler)

	insurancePatientRepo := ptRepo.NewInsurancePatientRepository(db)
	insurancePatientService := ptService.NewInsurancePatientService(insurancePatientRepo)
	insurancePatientHandler := ptHandler.NewInsurancePatientHandler(insurancePatientService)
	RegisterInsurancePatientRoutes(patientGroup, insurancePatientHandler)

	return r
}
