package cmd

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/mhristof/go-alfred"
	"github.com/spf13/cobra"
)

var (
	alfredCmd = &cobra.Command{
		Use:   "alfred",
		Short: "List alfred options",
		Run: func(cmd *cobra.Command, args []string) {
			dir := "memes"
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				panic(err)
			}

			var opts alfred.ScriptFilter
			for _, file := range files {
				path := filepath.Join(dir, file.Name())

				item := opts.Add(path, path)
				item.SetMatch(strings.ReplaceAll(file.Name(), "-", " "))

			}
			opts.Print()
		},
	}
)

func match(haystack, needle string) string {
	for _, char := range needle {
		haystack = strings.ReplaceAll(haystack, string(char), " ")
	}
	return haystack
}

func init() {
	rootCmd.AddCommand(alfredCmd)
}
