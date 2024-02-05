package main

import (
	"path/filepath"
	router "yuanaturego/routers"

	"github.com/gin-contrib/multitemplate"
)

func main() {
	router := router.InitRouter()
	//載入CSS、JS
	router.Static("/static", "static")
	//樣板
	//LoadHTMLGlob只能載入同一層的檔案
	//router.LoadHTMLGlob("templates/*.tmpl")
	router.HTMLRender = loadTemplates("./templates")

	router.Run(":8099")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	//前台
	frontbase, err := filepath.Glob(templatesDir + "/common/frontBase.tmpl")
	if err != nil {
		panic(err.Error())
	}

	frontends, err := filepath.Glob(templatesDir + "/frontend/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	for _, frontend := range frontends {
		layoutCopy := make([]string, len(frontbase))
		copy(layoutCopy, frontbase)
		files := append(layoutCopy, frontend)
		r.AddFromFiles(filepath.Base(frontend), files...)
	}

	//後台
	backbase, err := filepath.Glob(templatesDir + "/common/backBase.tmpl")
	if err != nil {
		panic(err.Error())
	}

	backends, err := filepath.Glob(templatesDir + "/backend/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	for _, backend := range backends {
		layoutCopy := make([]string, len(backbase))
		copy(layoutCopy, backbase)
		files := append(layoutCopy, backend)
		r.AddFromFiles(filepath.Base(backend), files...)
	}

	return r
}
