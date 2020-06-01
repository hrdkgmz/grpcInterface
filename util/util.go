package util

import (
	"io/ioutil"
	"strings"
)

//PackArgs define
func PackArgs(paras []string) [][]byte {
	var args [][]byte
	for _, k := range paras {
		args = append(args, []byte(k))
	}
	return args
}

//FindFiles define
func FindFiles(path string) [20]string {
	files, _ := ioutil.ReadDir(path)
	var fNames [20]string
	for i := 0; i < len(files); i++ {
		fNames[i] = files[i].Name()
	}
	return fNames
}

//FineOrgStr define
func FineOrgStr(str string) string {
	start := strings.Index(str, "Org")
	end := strings.Index(str, ".yaml")
	if start < 0 || end < 0 {
		return ""
	} else {
		str = str[start:end]
	}
	return str
}

