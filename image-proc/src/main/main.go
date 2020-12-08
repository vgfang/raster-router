package main
// go build -buildmode=c-shared -o imgproc.so main.go
import "github.com/fogleman/gg"
import (
	"C"
	"fmt"
	"strings"
)
// argument struct for name, type, image-filename/text/bool
type argu struct{
	name string
	typ string
	text string
}

//export test_circle
func test_circle(filename string) *C.char {
	dc := gg.NewContext(1000, 1000)
    dc.DrawCircle(500, 500, 400)
    dc.SetRGB(0, 0, 0)
    dc.Fill()
    dc.SavePNG(filename)
    return C.CString(filename)
}

func generate_blank(filename string, width int, height int){
	dc := gg.NewContext(width,height)
	dc.SetRGB(0, 0, 0)
    dc.Fill()
	dc.SavePNG(filename)
}

func filter() {

}
func scale() {

}

func parse_routestring(route string) ([]argu) {
	break0 := "\\sn\n" 
	break1 := "\\s: "
	// paritioning commands/data is done by spliting at break0
	comms := strings.Split(route, break0) // command array
	length := len(comms)
	arguArr := []argu{} // argument array

	i := 0
	// name, description
	for i<length {
		if comms[i] == "ARGUMENTS" {
			i++
			break
		}		
		i++
	}
	// build argument array
	for i<length {
		if comms[i] == "COMMANDS" {
			i++
			break
		}
		// argument types split at break1
		splitArgu := strings.Split(comms[i], break1)
		fmt.Println(splitArgu[0], splitArgu[1], splitArgu[2])
		newArgu := argu{splitArgu[0], splitArgu[1], splitArgu[2]}
		arguArr = append(arguArr, newArgu)
		i++
	}
	return arguArr
}

// parses routestring and creates the image
func perform_route(route string, filename string) *C.char {
	

	return C.CString("not done")
}



func main() {
	test_noarg := `
NAME\sn
First test\sn
DESCRIPTION\sn
First Description\sn
ARGUMENTS\sn
COMMANDS\sn
scale TRUE 200 200\sn
draw_2d TRUE 200 200\sn 
	`
	fmt.Println(test_noarg)
	parse_routestring(test_noarg)

	test_yesarg := `
NAME\sn
First test\sn
DESCRIPTION\sn
First Description\sn
ARGUMENTS\sn
imgname\s: img\s: test.png\sn
imgname2\s: img\s: test.png\sn
COMMANDS\sn
scale TRUE 200 200\sn
draw_2d TRUE 200 200\sn 
	`
	test := parse_routestring(test_yesarg)
	fmt.Println(test)
	generate_blank("test.png",600,600)
}
