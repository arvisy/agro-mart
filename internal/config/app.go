package config

import (
	"agro-mart/internal/atlas"
	"agro-mart/internal/constant"
	"log"

	"github.com/gin-gonic/gin"
)

func InitServer(gin *gin.Engine) {
	db, err := InitPostgresDB()
	if err != nil {
		log.Fatal("failed connecting to database: ", err.Error())
	}

	err = atlas.Migration(
		atlas.MigrationConfig{
			DSN:                NewPostgresConfig().MigrationDSN(),
			DirUrl:             "../../migrations",
			RevisionNameSchema: constant.AGRO_MART_MIGRATION_LOG,
		},
	)
	if err != nil {
		log.Fatal("failed to apply migration: ", err.Error())
	}

	container := NewContainer(db)
	InitRoute(gin, container)
}
