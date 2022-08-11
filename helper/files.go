package helper

import (
	"os"

	"github.com/ymzuiku/gowebdev/pkg/fsx"
)

var baseHtml = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
    <link rel="stylesheet" href="/public/tailwind.css" />
    <script src="/main.go.js"></script>
  </head>
  <body></body>
</html>
`

var baseTailwindConfig = `
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["**/*.go", "./index.html"],
  theme: {
    extend: {},
  },
  plugins: [],
};

`

var baseTailwindCss = `
@tailwind base;
@tailwind components;
@tailwind utilities;

`

func InitBaseFiles() {
	if !fsx.Exists("index.html") {
		os.WriteFile("index.html", []byte(baseHtml), 0o777)
	}
	if !fsx.Exists("tailwind.config.js") {
		os.WriteFile("tailwind.config.js", []byte(baseTailwindConfig), 0o777)
	}
	if !fsx.Exists("tailwind.css") {
		os.WriteFile("tailwind.css", []byte(baseTailwindCss), 0o777)
	}
}
