package main
// go build -buildmode=c-shared -o ./imgproc.so ./src/main/main.go
import "github.com/fogleman/gg"
//import "github.com/disintegration/gift"
import "github.com/nfnt/resize"
import (
	"C"
	"fmt"
	"strings"
	"strconv"
	"math"
)
// argument struct for name, type, image-filename/text/bool
type argu struct{
	name string
	typ string
}
// struct for parsed routestring
type route struct {
	dimensions []int	// canvas size
	arguments []argu	// arguments
	commands []string		// commands
}

func stoi(str string) int{
	x,_ := strconv.Atoi(str)
	return x
}
func stof(str string) float64{
	x, _ := strconv.ParseFloat(str,64)
	return x
}

//generates a empty transparent png
//export generate_blank
func generate_blank(filename string, width int, height int) {
	dc := gg.NewContext(width,height)
	dc.SetRGB(0, 0, 0)
    dc.Fill()
	dc.SavePNG(filename)
}

func draw_text(infile,outfile,posx,posy,text,font,fontsize,color string) {
	// type conversion
	x := stof(posx)
	y := stof(posy)
	fs := stof(fontsize)

	img, _ := gg.LoadPNG(infile)
	cont := gg.NewContextForImage(img)
	cont.SetHexColor(color)
	if err := cont.LoadFontFace("./static/fonts/"+ font +".ttf", fs); err != nil {
		cont.LoadFontFace("./static/fonts/impact.ttf", fs);
	}
	cont.DrawString(text, x, y)
	cont.SavePNG(outfile)
}
func draw_shape(infile,outfile,posx,posy,shape,width,height,color string){
	img, _ := gg.LoadPNG(infile)
	cont := gg.NewContextForImage(img)
	cont.SetHexColor(color)
	// type conversion
	x := stof(posx)
	y := stof(posy)
	w := stof(width)
	h := stof(height)

	switch shape {
	case "rectangle":
		cont.DrawRectangle(x,y,w,h)
	case "rounded_rect":
		cont.DrawRoundedRectangle(x,y,w,h,w*0.18)
	case "reg_triangle":
		cont.DrawRegularPolygon(3,x,y,w,0.0)
	case "reg_pentagon":
		cont.DrawRegularPolygon(5,x,y,w,0.0)
	case "reg_hexagon":
		cont.DrawRegularPolygon(6,x,y,w,0.0)
	case "ellipse":
		cont.DrawEllipse(x,y,w,h)
	default:
		fmt.Println("Bad Shape: " + shape)
	}
	cont.Fill()
	cont.SavePNG(outfile)
}
func draw_image(infile,outfile,posx,posy,addfile,width,height,aspect string){
	
	img, _ := gg.LoadPNG(infile)
	cont := gg.NewContextForImage(img)
	addImg, err := gg.LoadImage(addfile)
	fmt.Println("addFile:" + addfile)
	if err != nil{
		fmt.Println("addFile load error")
		fmt.Println(err)
		addImg, err = gg.LoadImage("placeholder.png")
		if err != nil{
			fmt.Println("placeholder load error")
			fmt.Println(err)
		}
	}
	addCont := gg.NewContextForImage(addImg)
	// type conversion
	x := stoi(posx)
	y := stoi(posy)
	w := stoi(width)
	h := stoi(height)
	aspect_bool := (aspect == "true")

	scaleWidth := uint(w)
	scaleHeight := uint(h)
	if aspect_bool { // use width scaling factor if aspect ratio should be preserved
		scaleHeight = uint(math.Round(float64(addCont.Height()) * float64(w)/float64(addCont.Width())))
	}
	addImg = resize.Resize(scaleWidth,scaleHeight,addImg,resize.Bicubic)
	cont.DrawImage(addImg, x, y)
	cont.SavePNG(outfile)
}	

//converts routestring to route struct
func parse_routestring(routestr string) route {
	break0 := "\\sn\n" 
	break1 := "\\s: "
	// paritioning commands/data is done by spliting at break0
	lines := strings.Split(routestr, break0) // command array
	length := len(lines)
	arguArr := []argu{} // argument array
	commArr := []string{} // command string array
	i := 1
	// name and description does nothing in backend image generation
	for i<length {
		if lines[i] == "CANVAS" {
			i++
			break
		}
		i++
	}
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
		newArgu := argu{splitArgu[0], splitArgu[1]}
		arguArr = append(arguArr, newArgu)
		i++
	}
	for i<length {
		if lines[i] == "END" || lines[i] == "\nEND"{ 
			break
		}
		commArr = append(commArr, lines[i])
		i++
	}
	return route{dimArr, arguArr, commArr}
}

func is_arg(input string) bool {
	break2 := "\\arg: " //argument indiator
	return strings.HasPrefix(input, break2)
}

// split single command line to array
func split_command_str(commandStr string) []string {
	break1 := "\\s: "
	c := strings.Split(commandStr, break1)
	return c
}

//export run_command_str
func run_command_str(infile string, outfile string, commandStr string) {
	commandStr = commandStr[:len(commandStr)-4] // drop newline character
	c := split_command_str(commandStr)
	run_command(infile, outfile, c)
}

func run_command(infile string, outfile string, c[]string) {
	fmt.Println("running: " + c[0] + " at (" + c[1] + "," +  c[2] + ")")
	placeholder := "placeholder.png"
	switch c[0] {
	case "draw_text":
		if is_arg(c[3]) { //argument used (only relevant in creation process)
			c[3] = "[ARG]:" + c[3][5:] //display arg name
			fmt.Println(c[3])
		}
		draw_text(infile,outfile,c[1],c[2],c[3],c[4],c[5],c[6])
	case "draw_shape":
		draw_shape(infile,outfile,c[1],c[2],c[3],c[4],c[5],c[6])
	case "draw_image":
		if is_arg(c[3]) { //argument used (only relevant in creation process)
			c[3] = placeholder	//display placeholder image
		}
		draw_image(infile,outfile,c[1],c[2],c[3],c[4],c[5],c[6])
	default:
		fmt.Println("Improper Command Type")
	}
	return
}

// returns map[argname => value] from user argstr
func parse_argstr(argstr string) map[string][]string {
	if argstr == "" {
		return nil
	}
	m := make(map[string][]string)
	strs := strings.Split(argstr, "\\sn\n")
	for _,line := range strs {
		parts := strings.Split(line, "\\s: ")
		if len(parts) > 1{
			m[parts[0]] = []string{parts[1],parts[2]}
		}
	}
	return m
}

// creates the image to filename according to route string, also does the argument logic
//export perform_route
func perform_route(filename string, routestr string, argDir string, argstr string) {
	rt := parse_routestring(routestr)
	// 1. generate blank canvas
	generate_blank(filename, rt.dimensions[0], rt.dimensions[1])

	// 2. initialize argstr
	argmap := make(map[string][]string)
	if argstr != ""{
		argmap = parse_argstr(argstr) 
	}
	// 3. run commands in sequence
	for _,cstr := range rt.commands {
		// argument replacement
		c := split_command_str(cstr)
		if is_arg(c[3]) { // argument found, replace
			noPrefix := c[3][6:]
			temp, ok := argmap[noPrefix]
			if ok { // found, replace with given argument
				switch temp[1] {
				case "img":
					c[3] = argDir + temp[0]
					fmt.Println("THIS:" + argDir + temp[0])
				case "text":
					c[3] = temp[0]
				default:
					c[3] = temp[0]
				}
			} else { // not found, change to "ERROR"
				c[3] = "ERROR"
				fmt.Println("map Error, perform_route()")
				continue
			}
		}
		run_command(filename, filename, c)
	}
}

func main() {
// main function used only for testing
	test_noarg := `
NAME\sn
Test1\sn
DESCRIPTION\sn
descition\sn
descction\sn
CANVAS\sn
600\s: 600\sn
ARGUMENTS\sn
COMMANDS\sn
draw_text\s: 400\s: 400\s: test\s: impact\s: 80\s: #000000\sn
draw_text\s: 200\s: 200\s: test\s: impact\s: 80\s: #000000\sn
draw_text\s: 100\s: 100\s: test\s: impact\s: 80\s: #000000\sn
draw_shape\s: 100\s: 100\s: rectangle\s: 400\s: 80\s: #000124\sn
draw_image\s: 100\s: 100\s: testing/taka.jpg\s: 400\s: 80\s: true\sn
END\sn
`

	test_yesarg := `NAME\sn
Test2\sn
DESCRIPTION\sn
descition\sn
descction\sn
CANVAS\sn
600\s: 600\sn
ARGUMENTS\sn
taka\s: img\sn
COMMANDS\sn
draw_text\s: 400\s: 400\s: test\s: impact\s: 80\s: #000000\sn
draw_text\s: 200\s: 200\s: \arg: input\s: impact\s: 80\s: #000000\sn
draw_text\s: 100\s: 100\s: test\s: impact\s: 80\s: #000000\sn
draw_shape\s: 100\s: 100\s: rectangle\s: 400\s: 80\s: #000124\sn
draw_image\s: 100\s: 100\s: \arg: taka\s: 400\s: 80\s: true\sn
END\sn
`
	test_argstr := `taka\s: taka.jpg\s: img\sn
input\s: argumenttext \s: text\sn
`

	dir := "testing/"
	test0 := parse_routestring(test_noarg)
	fmt.Println(test0)
	generate_blank("test.png",600,600)
	//run_command_str("test.png","edited.png",`draw_text\s: 400\s: 400\s: test\s: impact\s: 80\s: #000000`)
	//test1 := parse_routestring(test_yesarg)
	//fmt.Println(test1)
	//draw_text("test.png","edited.png","200","200","HOLY BASED","impact","36.0","#ffffff")
	//draw_shape("edited.png","edited2.png","100","100","rectangle","440","364","#012434")
	//draw_image("edited2.png","edited3.png","100","100","edited.png","400","200","true")
	perform_route(dir+"test.png",test_noarg,"anything","")
	perform_route(dir+"test2.png",test_yesarg,"testing", test_argstr)
}
