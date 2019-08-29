package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/realphant0m/SOLFLOWER/template"

	"github.com/realphant0m/SOLFLOWER/pkg"
	"github.com/spf13/cobra"
)

var (
	projectPath       string
	projectURL        string
	outputPath        string
	open              bool
	defaultOutputPath = filepath.Join(os.TempDir(), "SOLFLOWER")
)

var rootCmd = &cobra.Command{
	Use:   "SOLFLOWER",
	Short: "Visualize codebase repositories using an interactive tree",
	Long:  `SOLFLOWER is a tool for visualizing source repositories using an interactive tree`,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		if projectURL != "" {
			setOutputPath()
			projectPath, err = pkg.DownloadRepository(projectURL, outputPath)
			if err != nil {
				log.Println("error downloading repository. Reason: ", err)
				return
			}
		}

		if projectPath != "" {
			setOutputPath()
			jsonData, err := pkg.GenerateCLOC(projectPath)
			if err != nil {
				log.Println("error generating cloc. Reason: ", err)
				return
			}

			jsonByte, err := pkg.GenerateSolFlowerJSON(jsonData, outputPath)
			if err != nil {
				log.Println("error generating SOLFLOWER JSON. Reason: ", err)
				return
			}

			indexByte := append([]byte(template.INDEX_HTML1), jsonByte...)
			indexByte = append(indexByte, []byte(template.INDEX_HTML2)...)
			err = ioutil.WriteFile(filepath.Join(outputPath, "index.html"), indexByte, 0777)
			if err != nil {
				log.Println("error writing index.html. Reason: ", err)
				return
			}

			log.Println("Written files in the directory: ", outputPath)

			if open {
				fs := http.FileServer(http.Dir(outputPath))
				http.Handle("/", fs)

				f := pkg.IsPortOpen(3500)
				if !f {
					log.Println("Error opening server. Reason: port 3500 is in use")
					return
				}
				go func() {
					err := http.ListenAndServe(":3500", nil)
					if err != nil {
						log.Println("Error opening server. Reason: ", err)
						return
					}
				}()

				log.Println("Server is running at: http://localhost:3500")

				pkg.OpenBrowser("http://localhost:3500")
				log.Println("Press ctrl+c to close the server")

				ch := make(chan os.Signal, 1)
				signal.Notify(ch, os.Interrupt)
				<-ch
			}
		} else {
			help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&projectPath, "path", "", "local path of the project")
	rootCmd.PersistentFlags().StringVar(&projectURL, "url", "", "github url of the project")
	rootCmd.PersistentFlags().StringVar(&outputPath, "output", "", "outputPath path of the index.html and json file")
	rootCmd.PersistentFlags().BoolVar(&open, "open-in-browser", false, "opens the shape in browser")
}

func help() {
	helpText := `SOLFLOWER is a tool for codebase visualization using an interactive tree.

Available Flags:
  --path=<project-path>             Path of the project of which interactive tree need to be generated.

  --url=<project-github-url>        URL of the project of which interactive tree need to be generated.
                                    eg. --url=github.com/janforys/30-in-30-solidity-challenge

  --output=<output-path>            Path of the output directory where the index.html will be generated.
                                    Please make sure that you give a valid path.
  
  --open-in-browser                 Opens the output shape in browser.
`
	fmt.Println(helpText)
}

func setOutputPath() {
	if outputPath == "" {
		log.Println("No output path is set. Setting default output path to: ", defaultOutputPath)
		outputPath = defaultOutputPath
	}
}
