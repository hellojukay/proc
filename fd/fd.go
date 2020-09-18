package fd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type File struct {
	Fd   string
	Link string
}

func (f File) IsSocket() bool {
	r := regexp.MustCompile("socket:\\[[0-9]+\\]")
	return r.MatchString(f.Link)
}

func (f File) isNomalFile() bool {
	return strings.HasPrefix(f.Link, "/")
}

func (f File) Inode() (string, error) {
	if !f.IsSocket() {
		return "", errors.New("not socket description")
	}
	r := regexp.MustCompile("socket:\\[([0-9]+)\\]")
	return r.FindAllStringSubmatch(f.Link, -1)[0][0], nil
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
