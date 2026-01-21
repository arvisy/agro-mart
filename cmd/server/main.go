package main

import (
	"agro-mart/internal/atlas"
)

func main() {
	atlas.Migration(atlas.MigrationConfig{
		DSN: "postgres://postgres:qwe123@localhost:5432/agro_mart?sslmode=disable",
		// path ke folder migrations
		DirUrl: "../../migrations",
		// schema tempat atlas simpan history migration
		RevisionNameSchema: "migration_log",
	})
}
