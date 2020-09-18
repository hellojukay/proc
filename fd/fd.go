package fd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	Fd   string
	Link string
}

func (f File) IsSocket() bool {
	return false
}

func (f File) isNomalFile() bool {
	return false
}

func ReadFd(pid int) ([]File, error) {
	dirPath := fmt.Sprintf("/proc/%d/fd", pid)
	dirInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	var files []File
	for _, info := range dirInfos {
		var file File
		file.Fd = info.Name()
		l, _ := os.Readlink(filepath.Join(dirPath, file.Fd))
		file.Link = l
		files = append(files, file)
	}
	return files, nil
}
