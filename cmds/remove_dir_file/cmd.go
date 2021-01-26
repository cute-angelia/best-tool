package remove_dir_file

import "github.com/urfave/cli/v2"

var Cmd = &cli.Command{
	Name:            "remove_dir_file",
	Aliases:         []string{"rf"},
	Usage:           "删除文件夹制定文件",
	SkipFlagParsing: false,
	Action:          Run,
	UsageText:       Usage,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "word",
			Value:       "",
			Aliases:     []string{"w"},
			Usage:       "关键字按逗号分隔",
			Destination: &option.filterWords,
		},
		&cli.StringFlag{
			Name:        "dir",
			Value:       ".",
			Aliases:     []string{"d"},
			Usage:       "指定文件夹",
			Destination: &option.dir,
		},
	},
}
