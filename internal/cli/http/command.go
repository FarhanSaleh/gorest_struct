package http

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: "http server",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				return err
			}
			fmt.Printf("Server Run at %v:%d\n", "http://localhost",port)
			return nil
		},
	}
	cmd.Flags().Int("port", 8080, "http port")
	return cmd
}