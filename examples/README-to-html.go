// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// build +ignore

package main

import (
	"bytes"
	"os"
	"strings"

	"github.com/cpmech/gosl/io"
	"github.com/russross/blackfriday"
)

func main() {

	// read README.md file
	md := io.ReadFile("README.md")

	// process markdown
	//html := string(blackfriday.MarkdownCommon(md))

	flags := 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	extensions := 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS

	renderer := blackfriday.HtmlRenderer(flags, "", "")
	html := string(blackfriday.MarkdownOptions(md, renderer, blackfriday.Options{Extensions: extensions}))

	// environment variable => figures path
	path := os.ExpandEnv("${GOPATH}/src/github.com/cpmech/goga/")
	io.Pforan("path = %v\n", path)

	// set path of figures
	html = strings.Replace(html, "img src=\"", io.Sf("img src=\"%s/examples/", path), -1)

	// set header and footer
	html = `<!DOCTYPE HTML>
<html>
<head>
<title>Goga &ndash; Examples</title>
<meta charset="utf-8" />

<style>
h1 {color:#0064cb; font-family:verdana; font-size:200%;}
h2 {color:#0064cb}
h3 {color:#0064cb}
a:hover {background-color:#5397dc;}
#container {
	width:500px;
	text-align:center;
}
#container img {
	max-width:100%;
	height:auto;
}
</style>

</head>
<body>
` + html + `
</body>
</html>`

	// write file
	io.WriteFileVD("/tmp", "goga-README.html", bytes.NewBuffer([]byte(html)))
}
