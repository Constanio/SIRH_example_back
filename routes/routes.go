package routes

import (
	"sirh/handlers"
	"sirh/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	api := r.Group("/api")
	{
		// AUTH
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
			auth.GET("/me", middleware.AuthRequired, handlers.GetMe)
		}

		// ROUTES PROTÉGÉES
		protected := api.Group("/")
		protected.Use(middleware.AuthRequired)
		{
			// DASHBOARD
			protected.GET("/dashboard/stats", handlers.GetDashboardStats)
			protected.GET("/dashboard/revenue", handlers.GetMonthlyRevenue)

			// UTILISATEURS
			utilisateurs := protected.Group("/utilisateurs")
			{
				utilisateurs.GET("/", handlers.GetUtilisateurs)
				utilisateurs.GET("/:id", handlers.GetUtilisateur)
				utilisateurs.POST("/", handlers.CreateUtilisateur)
				utilisateurs.PUT("/:id", handlers.UpdateUtilisateur)
				utilisateurs.DELETE("/:id", handlers.DeleteUtilisateur)
			}

			// ORGANISATION
			protected.GET("/departements", handlers.GetDepartements)
			protected.POST("/departements", handlers.CreateDepartement)
			protected.PUT("/departements/:id", handlers.UpdateDepartement)
			protected.DELETE("/departements/:id", handlers.DeleteDepartement)
			protected.GET("/postes", handlers.GetPostes)
			protected.POST("/postes", handlers.CreatePoste)
			protected.PUT("/postes/:id", handlers.UpdatePoste)
			protected.DELETE("/postes/:id", handlers.DeletePoste)

			// CONGÉS
			conges := protected.Group("/conges")
			{
				conges.GET("/types", handlers.GetTypesConges)
				conges.POST("/types", middleware.RequireRoles("admin", "rh"), handlers.CreateTypeConge)
				conges.PUT("/types/:id", middleware.RequireRoles("admin", "rh"), handlers.UpdateTypeConge)
				conges.DELETE("/types/:id", middleware.RequireRoles("admin", "rh"), handlers.DeleteTypeConge)
				conges.GET("/mes-demandes", handlers.GetMesDemandes)
				conges.POST("/demande", handlers.CreateDemandeConge)
				conges.GET("/mes-soldes", handlers.GetMesSoldes)
				
				// Routes RH/Manager
				conges.GET("/toutes-les-demandes", middleware.RequireRoles("admin", "rh", "manager"), handlers.GetAllDemandes)
				conges.PATCH("/approuver/:id", middleware.RequireRoles("admin", "rh", "manager"), handlers.ApprouverDemande)
				conges.PATCH("/refuser/:id", middleware.RequireRoles("admin", "rh", "manager"), handlers.RefuserDemande)
			}

			// PAIE (RH/Admin)
			paie := protected.Group("/paie")
			paie.Use(middleware.RequireRoles("admin", "rh"))
			{
				paie.GET("/fiches", handlers.GetFichesPaie)
				paie.GET("/fiches/:id", handlers.GetFichePaie)
				paie.POST("/fiches", handlers.CreateFichePaie)
				paie.PUT("/fiches/:id", handlers.UpdateFichePaie)
				paie.DELETE("/fiches/:id", handlers.DeleteFichePaie)
			}

			// SALAIRES (RH/Admin)
			salaires := protected.Group("/salaires")
			salaires.Use(middleware.RequireRoles("admin", "rh"))
			{
				salaires.GET("/", handlers.GetSalaires)
				salaires.POST("/", handlers.CreateSalaire)
				salaires.PUT("/:id", handlers.UpdateSalaire)
				salaires.DELETE("/:id", handlers.DeleteSalaire)
			}

			// EVALUATIONS (RH/Admin/Manager)
			evals := protected.Group("/evaluations")
			evals.Use(middleware.RequireRoles("admin", "rh", "manager"))
			{
				evals.GET("/", handlers.GetEvaluations)
				evals.POST("/", handlers.CreateEvaluation)
				evals.PUT("/:id", handlers.UpdateEvaluation)
				evals.DELETE("/:id", handlers.DeleteEvaluation)
			}
		}

		// PING
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
	}
}
