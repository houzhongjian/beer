package beer

import "os"

//SetTemplateDir 设置视图文件夹路径.
func (srv *Handler) SetTemplateDir(dir string)   {
	_, err := os.Stat(dir)
	if err != nil {
		panic(err)
	}
	srv.templateDir = dir
}
