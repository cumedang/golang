package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/cumedang/go/explorer"
	"github.com/cumedang/go/rest"
)

func useage() {
	fmt.Printf("태양 코인에 어서오세욤\n\n")
	fmt.Printf("여기있는 커맨드들을 사용해주세요!:\n\n")
	fmt.Printf("-port:		서버의 포트를 설정합니다\n")
	fmt.Printf("-mode: 		'html' 과 'rest'중 하나를 선택해주세요\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		useage()
	}
	port := flag.Int("port", 4000, "서버의 포트를 설정하세요")
	mode := flag.String("mode", "rest", "'html' 과 'rest'중 하나를 선택해주세요")
	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		useage()
	}
}
