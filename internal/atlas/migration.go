package atlas

import (
	"context"
	"errors"
	"fmt"
	"os"

	"ariga.io/atlas/atlasexec"
)

type MigrationConfig struct {
	DSN                string
	DirUrl             string
	RevisionNameSchema string
}

func Migration(config MigrationConfig) error {
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(
			os.DirFS(config.DirUrl),
		),
	)
	if err != nil {
		return errors.New("failed to load working directory" + err.Error())
	}
	defer workdir.Close()

	client, err := atlasexec.NewClient(
		workdir.Path(), "atlas",
	)
	if err != nil {
		return errors.New("failed to initialize client: " + err.Error())
	}

	res, err := client.MigrateApply(
		context.Background(),
		&atlasexec.MigrateApplyParams{
			URL:             config.DSN,
			RevisionsSchema: config.RevisionNameSchema,
			AllowDirty:      true,
		},
	)
	if err != nil {
		return errors.New("failed to apply migrations: " + err.Error())
	}
	if res.Error != "" {
		return fmt.Errorf("error apply migrations on response: %v", res.Error)
	}

	fmt.Printf("applied %d migrations", len(res.Applied))

	return nil
}
