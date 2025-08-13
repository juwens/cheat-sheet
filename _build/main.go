package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

func getCmpYmlFiles(root string, suffix string) []string {
	var cmp_yml_files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		path_segments := strings.Split(path, string(os.PathSeparator))
		if slices.Contains(path_segments, ".git") {
			return nil
		}

		if slices.Contains(path_segments, ".build") {
			return nil
		}

		if strings.HasSuffix(path, suffix) {
			cmp_yml_files = append(cmp_yml_files, path)
			return nil
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cmp_yml_files)

	return cmp_yml_files
}

func main() {
	var root string

	var rootCmd = &cobra.Command{
		Use:   "listfiles [root path]",
		Short: "Recursively list files starting from a root path",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			root = args[0]
			files := getCmpYmlFiles(root, ".yml")

			for _, inputFile := range files {
				data := ReadAllBytes(inputFile)

				if strings.HasSuffix(inputFile, ".cmp_nuget.yml") {
					re := regexp.MustCompile(regexp.QuoteMeta(".cmp_nuget.yml") + "$")

					outFilename := re.ReplaceAllString(inputFile, ".adoc")
					fmt.Println("outfile: " + outFilename)

					outFile, outFileErr := os.Create(outFilename)
					if outFileErr != nil {
						panic(outFileErr)
					}

					main_nugets(data, outFile)
					err := outFile.Close()
					if err != nil {
						panic(err)
					}
				}

			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ReadAllBytes(inputFile string) []byte {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
