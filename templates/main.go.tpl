{{file main-doc.tpl}}
{{let local-import %q" github.com/westarver/messenger}}
{{let writer os.Stderr}}
package main

import (
	writer {{var local-import}}
)

func main() {
	writer := messenger.New()
	writer.SetOut({{var writer}})
	
	os.Exit(run(writer))
}


