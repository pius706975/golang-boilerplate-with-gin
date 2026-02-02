package custom

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var migrationName string

var CreateMigrationCMD = &cobra.Command{
	Use:   "create-migration",
	Short: "Generate SQL migration files using nanosecond timestamp versioning",
	RunE:  createMigration,
}

func init() {
	CreateMigrationCMD.Flags().StringVar(
		&migrationName,
		"name",
		"",
		"Migration name (can be multiple words)",
	)
	_ = CreateMigrationCMD.MarkFlagRequired("name")
}

func createMigration(cmd *cobra.Command, args []string) error {
	migrationsDir := "package/database/migrations"

	// ensure migrations directory exists
	if err := os.MkdirAll(migrationsDir, os.ModePerm); err != nil {
		return err
	}

	// normalize migration name
	cleanName := normalizeMigrationName(migrationName)

	// nanosecond-precision timestamp (super safe for teams & CI)
	version := fmt.Sprintf("%d", time.Now().UnixNano())

	upFile := filepath.Join(
		migrationsDir,
		fmt.Sprintf("%s_%s.up.sql", version, cleanName),
	)

	downFile := filepath.Join(
		migrationsDir,
		fmt.Sprintf("%s_%s.down.sql", version, cleanName),
	)

	upContent := "-- +++ UP migration +++\n\n"
	downContent := "-- +++ DOWN migration +++\n\n"

	if err := os.WriteFile(upFile, []byte(upContent), 0644); err != nil {
		return err
	}

	if err := os.WriteFile(downFile, []byte(downContent), 0644); err != nil {
		return err
	}

	fmt.Println("Migration files created:")
	fmt.Println(upFile)
	fmt.Println(downFile)

	return nil
}

func normalizeMigrationName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))

	name = strings.ReplaceAll(name, " ", "_")

	re := regexp.MustCompile(`[^a-z0-9_]+`)
	name = re.ReplaceAllString(name, "")

	reUnderscore := regexp.MustCompile(`_+`)
	name = reUnderscore.ReplaceAllString(name, "_")

	return name
}
