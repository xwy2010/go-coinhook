package version

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xwy2010/go-admin-coinhook/common/global"
)

var (
	StartCmd = &cobra.Command{
		Use:     "ver",
		Short:   "Get version info",
		Example: "version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println(global.Version)
	return nil
}
