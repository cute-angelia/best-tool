package remove_dir_file

import (
	"errors"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var includefiles = []string{}

func Run(cli *cli.Context) (err error) {
	filterWords := []string{}

	if option.filterWords == "" {
		err = errors.New("Please input filter words")
		return
	} else {
		filterWords = strings.Split(option.filterWords, ",")
	}
	if option.dir == "" {
		err = errors.New("Please input dir")
		return
	}

	getFiles(option.dir)
	for _, v := range includefiles {
		for _, bad := range filterWords {
			if strings.Contains(v, bad) {
				log.Println("删除 --> ", v)
				os.Remove(v)
			}
		}
	}
	return
}

func getFiles(folder string) {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		pathson := folder + "/" + file.Name()
		includefiles = append(includefiles, pathson)

		if file.IsDir() {
			getFiles(pathson)
		}
	}

}
