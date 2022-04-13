package cmd

import (
	"errors"
	"fmt"

	"github.com/xwy2010/go-core/sdk/pkg"

	// "go-admin/cmd/app"
	"os"

	"github.com/xwy2010/go-admin-coinhook/common/global"

	"github.com/spf13/cobra"

	// "go-admin/cmd/api"
	// "go-admin/cmd/config"
	// "go-admin/cmd/migrate"
	"github.com/xwy2010/go-admin-coinhook/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:          "coinhook",
	Short:        "coinhook",
	SilenceUsage: true,
	Long:         `coinhook`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(pkg.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`go-admin-coinhook v`+global.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	// rootCmd.AddCommand(api.StartCmd)
	// rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	// rootCmd.AddCommand(config.StartCmd)
	// rootCmd.AddCommand(app.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
