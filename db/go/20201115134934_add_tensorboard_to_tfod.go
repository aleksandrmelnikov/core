package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func initialize20201115134934() {
	if _, ok := initializedMigrations[20201115134934]; !ok {
		goose.AddMigration(Up20201115134934, Down20201115134934)
		initializedMigrations[20201115134934] = true
	}
}

//Up20201115134934 add TensorBoard sidecar to TFODs
func Up20201115134934(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return updateWorkflowTemplateManifest(
		"20201115134934_tfod.yaml",
		tensorflowObjectDetectionWorkflowTemplateName,
		map[string]string{
			"used-by": "cvat",
		},
	)
}

//Down20201115134934 do nothing
func Down20201115134934(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
