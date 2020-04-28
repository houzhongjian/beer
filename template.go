package beer

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var templateData = make(map[string]string)

//SetTemplateDir 设置视图文件夹路径.
func (srv *Handler) SetTemplateDir(templateDir string) {
	_, err := os.Stat(templateDir)
	if err != nil {
		panic(err)
	}
	templateDir = strings.Replace(templateDir, "./", "", -1)
	err = filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			//读取视图.
			b, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("err:%+v\n", err)
				return err
			}
			//替换名称.
			path = strings.Replace(path, templateDir, "", -1)
			templateData[path] = string(b)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	srv.templateDir = templateDir
}
