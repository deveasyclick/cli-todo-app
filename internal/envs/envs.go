package envs

import (
	"log"
	"os"
	"path"
)

type ENV_NAMES struct {
	AuthFileName string
	HomeDir      string
	AuthFilepath string
}

var homeDir = getHomeDir()
var authFileName = os.Getenv("AUTH_FILENAME")
var ENV_VARIABLES = ENV_NAMES{
	AuthFileName: authFileName,
	HomeDir:      homeDir,
	AuthFilepath: path.Join(homeDir, authFileName),
}

func getHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Error getting home dir", err)
	}

	return homeDir
}
