package command

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/hexium310/srcurl/internal/browser"
	"github.com/hexium310/srcurl/internal/config"
	"github.com/hexium310/srcurl/internal/source"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "srcurl [flags] <path>",
	Short: "Generate a URL from the filename",
	Long: "Extract an id from <path> with the pattern in your config file, and generate a URL built with it.",
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		copy, err := cmd.Flags().GetBool("copy")
		if err != nil {
			panic(err)
		}
		open, err := cmd.Flags().GetBool("open")
		if err != nil {
			panic(err)
		}

		_, filename := filepath.Split(args[0])

		url := source.GetUrl(filename)
		if url == "" {
			return errors.New("No matches")
		}

		if copy {
			err := clipboard.WriteAll(url)
			if err != nil {
				return err
			}

		} else if open {
			browser.Open(url)
		} else {
			fmt.Println(url)
		}

		return nil
	},
}

func init() {
	cobra.MousetrapHelpText = ""

	RootCmd.PersistentFlags().StringVar(&config.ConfigFile, "config", "", "config file (default: $XDG_CONFIG_HOME/srcurl/srcurl.toml, $HOME/.config/srcurl/srcurl.toml (in Unix), or %LOCALAPPDATA%\\srcurl\\srcurl.toml (in Windows))")
	RootCmd.Flags().BoolP("copy", "c", false, "Help message for toggle")
	RootCmd.Flags().BoolP("open", "o", false, "Help message for toggle")
	RootCmd.MarkFlagsMutuallyExclusive("copy", "open")
}
