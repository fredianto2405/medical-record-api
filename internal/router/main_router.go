package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"medical-record-api/internal/auth"
	"medical-record-api/internal/clinic"
	"medical-record-api/internal/doctor"
	emrHandler "medical-record-api/internal/medical_record/handler"
	emrRepo "medical-record-api/internal/medical_record/repository"
	emrService "medical-record-api/internal/medical_record/service"
	medHandler "medical-record-api/internal/medicine/handler"
	medRepo "medical-record-api/internal/medicine/repository"
	medService "medical-record-api/internal/medicine/service"
	"medical-record-api/internal/menu"
	"medical-record-api/internal/nurse"
	patHandler "medical-record-api/internal/patient/handler"
	patRepo "medical-record-api/internal/patient/repository"
	patService "medical-record-api/internal/patient/service"
	payHandler "medical-record-api/internal/payment/handler"
	payRepo "medical-record-api/internal/payment/repository"
	payService "medical-record-api/internal/payment/service"
	"medical-record-api/internal/specialization"
	"medical-record-api/internal/treatment"
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

	unitRepo := medRepo.NewUnitRepository(db)
	unitService := medService.NewUnitService(unitRepo)
	unitHandler := medHandler.NewUnitHandler(unitService)
	RegisterUnitRoutes(medicineGroup, unitHandler)

	categoryRepo := medRepo.NewCategoryRepository(db)
	categoryService := medService.NewCategoryService(categoryRepo)
	categoryHandler := medHandler.NewCategoryHandler(categoryService)
	RegisterCategoryRoutes(medicineGroup, categoryHandler)

	medicineRepo := medRepo.NewRepository(db)
	medicineService := medService.NewService(medicineRepo)
	medicineHandler := medHandler.NewHandler(medicineService)
	RegisterMedicineRoutes(medicineGroup, medicineHandler)

	// patient routes
	patientGroup := r.Group("/api/v1/patients")

	patientRepo := patRepo.NewRepository(db)
	patientService := patService.NewService(patientRepo)
	patientHandler := patHandler.NewHandler(patientService)
	RegisterPatientRoutes(patientGroup, patientHandler)

	emergencyContactRepo := patRepo.NewEmergencyContactRepository(db)
	emergencyContactService := patService.NewEmergencyContactService(emergencyContactRepo)
	emergencyContactHandler := patHandler.NewEmergencyContactHandler(emergencyContactService)
	RegisterEmergencyContactRoutes(patientGroup, emergencyContactHandler)

	insurancePatientRepo := patRepo.NewInsurancePatientRepository(db)
	insurancePatientService := patService.NewInsurancePatientService(insurancePatientRepo)
	insurancePatientHandler := patHandler.NewInsurancePatientHandler(insurancePatientService)
	RegisterInsurancePatientRoutes(patientGroup, insurancePatientHandler)

	// treatment routes
	treatmentGroup := r.Group("/api/v1/treatments")
	treatmentRepo := treatment.NewRepository(db)
	treatmentService := treatment.NewService(treatmentRepo)
	treatmentHandler := treatment.NewHandler(treatmentService)
	RegisterTreatmentRoutes(treatmentGroup, treatmentHandler)

	// insurance routes
	insuranceGroup := r.Group("/api/v1/insurances")
	insuranceRepo := payRepo.NewInsuranceRepository(db)
	insuranceService := payService.NewInsuranceService(insuranceRepo)
	insuranceHandler := payHandler.NewInsuranceHandler(insuranceService)
	RegisterInsuranceRoutes(insuranceGroup, insuranceHandler)

	// payment routes
	paymentGroup := r.Group("/api/v1/payment")
	paymentRepo := payRepo.NewRepository(db)
	paymentService := payService.NewService(paymentRepo)
	paymentHandler := payHandler.NewHandler(paymentService)
	RegisterPaymentRoutes(paymentGroup, paymentHandler)

	// clinic routes
	clinicGroup := r.Group("/api/v1/clinics")
	clinicRepo := clinic.NewRepository(db)
	clinicService := clinic.NewService(clinicRepo)
	clinicHandler := clinic.NewHandler(clinicService)
	RegisterClinicRoutes(clinicGroup, clinicHandler)

	// medical record routes
	medicalRecordRepo := emrRepo.NewRepository(db)
	nurseAssignmentRepo := emrRepo.NewNurseAssignmentRepository(db)
	treatmentDetailRepo := emrRepo.NewTreatmentDetailRepository(db)
	recipeRepo := emrRepo.NewRecipeRepository(db)
	historyRepo := emrRepo.NewHistoryRepository(db)
	statusRepo := emrRepo.NewStatusRepository(db)

	medicalRecordService := emrService.NewService(db, medicalRecordRepo, nurseAssignmentRepo, treatmentDetailRepo, recipeRepo, historyRepo)
	nurseAssignmentService := emrService.NewNurseAssignmentService(db, nurseAssignmentRepo)
	treatmentDetailService := emrService.NewTreatmentDetailService(db, treatmentDetailRepo)
	recipeService := emrService.NewRecipeService(db, recipeRepo)
	statusService := emrService.NewStatusService(statusRepo)

	medicalRecordGroup := r.Group("/api/v1/medical-records")
	medicalRecordHandler := emrHandler.NewHandler(medicalRecordService, nurseAssignmentService, treatmentDetailService, recipeService, statusService)
	RegisterMedicalRecordRoutes(medicalRecordGroup, medicalRecordHandler)

	return r
}
