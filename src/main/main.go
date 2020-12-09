package main
// CGO_ENABLED=0 go build -buildmode=c-shared -o ./imgproc.so ./src/main/main.go
import "github.com/fogleman/gg"
//import "github.com/disintegration/gift"
import "github.com/nfnt/resize"
import (
	"C"
	"fmt"
	"strings"
	"strconv"

)
// argument struct for name, type, image-filename/text/bool
type argu struct{
	name string
	typ string
	text string
}
// 
type comm struct {
	typ string // type of command (which function)
	arguments []string // rest of string
}
// struct for parsed routestring
type route struct {
	dimensions []int	// canvas size
	arguments []argu	// arguments
	commands []comm		// commands
}

//generates a empty transparent png
//export generate_blank
func generate_blank(filename string, width int, height int) {
	dc := gg.NewContext(width,height)
	dc.SetRGB(0, 0, 0)
    dc.Fill()
	dc.SavePNG(filename)
}

func draw_text(infile string,outfile string,posx float64,posy float64,text string,font string,fontsize float64,color string) {
	img, _ := gg.LoadPNG(infile)
	cont := gg.NewContextForImage(img)
	cont.SetHexColor(color)
	if err := cont.LoadFontFace("./static/fonts/"+ font +".ttf", fontsize); err != nil {
		panic(err)
	}
	cont.DrawString(text, float64(posx), float64(posy))
	cont.SavePNG(outfile)
}
func draw_shape(infile string,outfile string,posx float64,posy float64,shape string,width float64,height float64,color string){
	img, _ := gg.LoadPNG(infile)
	cont := gg.NewContextForImage(img)
	cont.SetHexColor(color)

	switch shape {
	case "rectangle":
		fmt.Println("hello")
		cont.DrawRectangle(posx,posy,width,height)
	case "rounded_rect":
		cont.DrawRoundedRectangle(posx,posy,width,height,float64(width)*0.06)
	case "reg_triangle":
		cont.DrawRegularPolygon(3,posx,posy,width,float64(0))
	case "reg_pentagon":
		cont.DrawRegularPolygon(5,posx,posy,width,float64(0))
	case "reg_hexagon":
		cont.DrawRegularPolygon(6,posx,posy,width,float64(0))
	case "ellipse":
		cont.DrawEllipse(posx,posy,width,height)
	default:
		fmt.Println("Bad Shape: " + shape)
	}
	cont.Fill()
	cont.SavePNG(outfile)
}
func draw_image(infile string,outfile string,posx int,posy int,addfile string,width int,height int,scale bool,aspect bool){
	img, _ := gg.LoadPNG(infile)
	cont := gg.NewContextForImage(img)
	addImg, _ := gg.LoadImage(addfile)
	if scale {
		scaleWidth := uint(width)
		scaleHeight := uint(height)
		if aspect { // use width scaling factor if aspect ratio should be preserved

		}
		addImg = resize.Resize(scaleWidth,scaleHeight,addImg,resize.NearestNeighbor)
	}
	cont.DrawImage(addImg, posx, posy)
	cont.SavePNG(outfile)
}	

// change command line to command struct
func parse_command(line string) comm {
	break1 := "\\s: "
	parts := strings.Split(line, break1)
	return comm{parts[0], parts[1:len(parts)]}
}

// run command 
func run_command(line string) {
	// command := parse_command(line)
	return
}

//converts routestring to route struct
func parse_routestring(routestr string) route {
	break0 := "\\sn\n" 
	break1 := "\\s: "
	// paritioning commands/data is done by spliting at break0
	lines := strings.Split(routestr, break0) // command array
	length := len(lines)
	arguArr := []argu{} // argument array
	commArr := []comm{} // argument array
	i := 1

	// canvas size
	tempArr := strings.Split(lines[i], break1)
	canvasWidth, _:= strconv.Atoi(tempArr[0])
	canvasHeight, _ := strconv.Atoi(tempArr[1])
	dimArr := []int{canvasWidth, canvasHeight} // dimension array
	i+=2

	// build argument array
	for i<length {
		if lines[i] == "COMMANDS" {
			i++
			break
		}
		// argument types split at break1
		splitArgu := strings.Split(lines[i], break1)
		newArgu := argu{splitArgu[0], splitArgu[1], splitArgu[2]}
		arguArr = append(arguArr, newArgu)
		i++
	}
	for i<length {
		commArr = append(commArr, parse_command(lines[i]))
		i++
	}
	return route{dimArr, arguArr, commArr}
}

// creates the image to filename according to route
// export perform route
func perform_route(routestr string, filename string) *C.char {
	return C.CString("not done")
}

func main() {
// main function used only for testing
	test_noarg := `
CANVAS\sn
600\s: 600\sn
ARGUMENTS\sn
COMMANDS\sn
scale TRUE 200 200\sn
draw_2d TRUE 200 200\sn 
	`

	test_yesarg := `
CANVAS\sn
600\s: 600\sn
ARGUMENTS\sn
imgname\s: img\s: test.png\sn
imgname2\s: img\s: test.png\sn
COMMANDS\sn
scale TRUE 200 200\sn
draw_2d TRUE 200 200\sn 
	`
	fmt.Println(parse_routestring(test_noarg))
	test1 := parse_routestring(test_yesarg)
	fmt.Println(test1)
	generate_blank("test.png",600,600)
	draw_text("test.png","edited.png",200,200,"HOLY BASED","impact",36.0,"#ffffff")
	//draw_shape("edited.png","edited.png",100,100,"reg_triangle",440.0,364.0,"#012434")
	draw_shape("edited.png","edited2.png",100,100,"reg_triangle",440.0,364.0,"#012434")
	draw_image("edited2.png","edited3.png",100,100,"edited.png",60,60,true,false)
}
