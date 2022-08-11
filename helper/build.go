package helper

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/ymzuiku/gowebdev/pkg/execx"
	"github.com/ymzuiku/gowebdev/pkg/fsx"
)

func Md5String(d []byte) string {
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func Build(clienPath, outDir string) {
	InitBaseFiles()
	outPublcDir := outDir + "/public"
	fsx.CreateIfNotExists(outDir, 0o777)
	fsx.CreateIfNotExists(outPublcDir, 0o777)
	log.Printf("%v", "go build js...")
	execx.Run(context.Background(), nil, "gopherjs", "build", clienPath)
	jsName := fmt.Sprintf("main_%s.js", Md5String(fsx.LoadFileByte("main.js"))[:8])
	outFile := outDir + "/" + jsName
	log.Printf("%v", "mini file to "+outFile)
	result := api.Build(api.BuildOptions{
		EntryPoints:       []string{"main.js"},
		Outfile:           outFile,
		Bundle:            true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Engines: []api.Engine{
			{Name: api.EngineChrome, Version: "58"},
			{Name: api.EngineFirefox, Version: "57"},
			{Name: api.EngineSafari, Version: "11"},
			{Name: api.EngineEdge, Version: "16"},
		},
		Write: true,
	})

	if len(result.Errors) > 0 {
		log.Printf("esbuild: %v", result.Errors)
		os.Exit(1)
		return
	}

	log.Printf("%v", "remove temp files...")
	if err := os.Remove("main.js"); err != nil {
		panic(err)
	}
	if err := os.Remove("main.js.map"); err != nil {
		panic(err)
	}

	log.Printf("%v", "copy public dir...")
	fsx.CopyDirectory("public", outPublcDir)
	html := fsx.LoadFile("index.html")
	html = strings.Replace(html, "main.go.js", jsName, 1)
	if fsx.Exists(outPublcDir + "/tailwind.css") {
		cssByte := fsx.LoadFileByte(outPublcDir + "/tailwind.css")
		cssName := fmt.Sprintf("tailwind_%s.css", Md5String(cssByte)[:8])
		os.WriteFile(outPublcDir+"/"+cssName, cssByte, 0o777)
		html = strings.Replace(html, "tailwind.css", cssName, 1)
		os.Remove(outPublcDir + "/tailwind.css")
	}

	os.WriteFile(outDir+"/index.html", []byte(html), 0o777)

	log.Printf("%v", "Done")
}
