package utilstests

import (
	"fmt"
	"os"
	"regexp"

	"github.com/freddiemo/healthcare-api/config"
)

// this should match the project root folder name.
const rootFolder = "healthcare-api"

// LoadEnv load the env files relative to the root folder from any nested test directory.
func LoadEnv() map[string]string {
	root := os.Getenv("APP_ROOT")
	if len(root) == 0 {
		root = rootFolder
	}
	re := regexp.MustCompile(fmt.Sprintf("^(.*%s)", root))
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	errMsg := "update the `testutils.rootFolder` variable to match the name of the project root folder"

	if rootPath == nil {
		panic(errMsg)
	}

	if err := os.Chdir(string(rootPath)); err != nil {
		panic(errMsg)
	}

	os.Setenv("GO_ENVIRONMENT", "TEST")

	return config.Init()
}
