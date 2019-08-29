package pkg

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/codeskyblue/go-sh"
)

var data map[string]interface{}

// GenerateCLOC generates the CLOC of a project using https://github.com/AlDanial/cloc binary
func GenerateCLOC(path string) ([]byte, error) {
	log.Println("Generating CLOC json for the project: ", path)

	tempDir, err := ioutil.TempDir(os.TempDir(), "")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir)

	session := sh.NewSession()
	session.SetDir(path)
	session.Stdout = nil
	session.Stderr = nil

	err = session.Command("cloc", ".", "--json", "--exclude-dir=vendor,tmp", "--by-file", "--report-file="+filepath.Join(tempDir, "data.json")).Run()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(filepath.Join(tempDir, "data.json"))
}
