package cmd

import (
	"context"
	"github.com/JoeyLearnsToCode/staticweb-spider/global"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "staticweb-spider",
		Short: "A static website spider",
		Long:  `A static website spider.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			global.PrintOptions()
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	c, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	go func() {
		select {
		case <-c.Done():
			global.Logger.Printf("user interrupted, exiting...")
			os.Exit(0)
		}
	}()

	err := rootCmd.ExecuteContext(c)
	if err != nil {
		global.Logger.Fatalf("quit with error: %v", err)
	}
}

func init() {
	parseArgs()
}

func parseArgs() {
	rootCmd.PersistentFlags().StringVarP(&global.Proxy, "proxy", "p", "", "http proxy url")
	rootCmd.PersistentFlags().StringVar(&global.ProxyMode, "proxy-mode", "none", "proxy mode, values: none, primary, all")
}
