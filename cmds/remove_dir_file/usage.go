package remove_dir_file

const Usage = `
rf [commands|flags]
`

/*
The commands & flags are:
  -w      关键字按逗号分隔: 广告,游戏
  -d      文件夹

Examples:
  # 删除指定文件夹文件
  go-third-cli-tool -w hello,world -d (project dir)
*/