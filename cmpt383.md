# Raster Router
Raster Router is a web application for producing templates for algorithmic image editing. The design goal for these templates, or _routes_ is to have a black-box, functional design for creating images where the user inputs arguments (text or other images), to receive a modified image as a result. There are two main aspects to the webapp:
1. Producing and serving images made of shapes and text according to the user's commands sequentially.
2. Producing and using human-readable templates that can be used to create images using other images and inputed text as arguments.

Single Image generation on the web can be done via HTML and Javascript, but generation of many image with complex algorithms with many steps, or creating images in a single action using a template necessitates a more performant way to produce images. Raster Router can produce images quickly using shapes, text, and other images, and can produce images with or without arguments.

## Programming Languages
It is built on the Flask web-framework, using JQuery for front-end commands, and Golang for image production. Python uses Ctypes for calling a Golang shared binary to produce raster image functions. Python uses Flask as a web-server to communicate the client-side Javascript web app, and the webapp makes many AJAX calls to call python functions to create dynamic web pages.

### External Libraries
- github.com/fogleman/gg
   + A golang graphics library was used to provide simple image manipulation functions
- github.com/nfnt/resize
   + A golang library for resizing images
- github.com/eligrey/FileSaver.js
   + A Javascript file was used to handle client-side downloading of a JS string for the route creation

### Deployment
- Vagrant VM and Chef are used to deploy a Virtual Machine.
1. Create and Start the Vagrant VM `vagrant up`
2. SSH into the Vagrant VM `vagrant ssh`
3. Start the Flask server: `cd project; python3 app.py;`
4. The local IP will be: `192.168.33.10:5000`

## How to Use
### Creating Routes
1. From the main page, select "Create an Image Route"
2. The user generates a blank canvas of the size that they want.
3. The user then can add arguments and commands. This is rendered in a preview image which only shows the difference of a single command step. To add more commands, the user must sequentially write one command, continue to next, then write another command.
4. Arguments can be created by first adding an argument of type text or image. Then, the user must reference the chosen argument name prefixed with "\\arg: " it in either the text of an `ADD_TEXT` command or the image-file of an `ADD_IMAGE` command based on type. Arguments must be the proper syntax, otherwise the code will fall back on a placeholder.
5. Downloading the route serves a `.route` extension RouteString. This can be used to _route_ an image.

### Using a Route
1. To use a route, the user must upload a valid RouteString file and press "Use an Image Route".
2. Users are then redirected to a page where they can choose their arguments and read the route name and description. A dynamic form is generated using the RouteString's argument fields.
3. To give an image argument, the user must upload a valid image file, then upload them to the server, using the upload button. This image uploading must be done before image generation.
4. To give a text argument, the user can just type into a text area element in the same way.
5. To generate the final image, pressing the "Generate Image" button, and after waiting,  will render the image below on the same page. 


## How it works
### RouteStrings
RouteStrings (`.route`) are a custom plaintext markup format to represent all of the data needed to create a __route__. They are designed to be human readable/editable and contain the information for a route's name, description, arguments, and commands. To make these processable with user input, three break characters had to be made. These are the newline(`\sn\n`), space(\s: ), and argument(\arg: ).

```
NAME\sn
4panelart\sn
DESCRIPTION\sn
Demonstrating Image argument.\sn
CANVAS\sn
400\s: 400\sn
ARGUMENTS\sn
test\s: img\sn
COMMANDS\sn
draw_image\s: 0\s: 0\s: \arg: test\s: 200\s: 200\s: false\sn
draw_image\s: 0\s: 200\s: \arg: test\s: 200\s: 200\s: false\sn
draw_image\s: 200\s: 200\s: \arg: test\s: 200\s: 200\s: false\sn
draw_image\s: 200\s: 0\s: \arg: test\s: 200\s: 200\s: false\sn
draw_shape\s: 0\s: 0\s: rectangle\s: 200\s: 200\s: 99994444\sn
draw_shape\s: 200\s: 200\s: rectangle\s: 200\s: 200\s: 22992244\sn
draw_shape\s: 0\s: 200\s: rectangle\s: 200\s: 200\s: 00559944\sn
draw_shape\s: 200\s: 0\s: rectangle\s: 200\s: 200\s: 33224444\sn
END\sn
```
The syntax for RouteStrings is simple and can be easily modified in a text editor. This can be useful if a mistake was made during the template creation process, such as the wrong type for an argument or the wrong argument name. The name and description can also be modified via text editor.

### CommandStrings
A CommandString is a single line of plaintext which represents a single command. These can be ran sequentially and running a RouteString executes a list of commands to produce the final image. These strings can be parsed by the Golang code, to call different functions and are passed from Javascript to Python to the Golang shared library. These are generated during the _Route Creation_ process.
```
draw_shape\s: 200\s: 0\s: rectangle\s: 200\s: 200\s: 3322444
```

### ArgumentStrings
ArgumentStrings are a plaintext string, similar to a RouteString in syntax, which defines the user's arguments during _Image routing_. These are used provide a linking mechanism between arguments and the RouteString commands and are only generated during the Routing process.
```
taka\s: taka.jpg\s: img\sn
```

### Golang Role
The Golang code was designed to be agnostic to file structure, requiring  input and output directory strings to produce images. Golang receives information in string format, and provides many decoding functions to parse RouteStrings, CommandStrings, and ArgumentStrings into different data structures and using them to run commands while detecting argument calls. Golang was chosen for this role because it is performant. The Golang code is able to run single CommandStrings or take a RouteString and ArgumentString to produce an image.

### Python Role
The Python code was used as a web-server and to provide functions that received Javascript AJAX calls, and dealt with most of the file operations, including the renaming and copying of the creation images, using cookies. It was safer and more sensible for the web-server to determine the directories and edit files, rather than having the client send file information directly. Python was chosen for this role because Flask is a simple web framework.

### Javascript Role
Javascript functions are made to provide AJAX calls and certain UI features, such as image reloading. There are many functions to generate the various RouteStrings, CommandStrings, ArgumentsStrings from information on the webpage into the proper syntax and provides the information to produce it and send it to Python. Javascript was chosen because it is required for front-end web programming.

### Creating Routes Inner Workings
During the process of creating routes, the browser client is creating a RouteString in the background using Javascript and sending single-use command strings for image generation for previews. The image generation functionality is achieved though AJAX calls from Javascript to Flask, then from Flask, where the file manipulation is done. Python then calls the Golang binary to generate a new image. Lastly, Javascript updates the browser client image.

To move on to the next image/step, Python renames the files and Javascript refreshes the images as well, as well as adding to the RouteString variable. The RouteString can downloaded as a file as soon as at least one command, name, and description was inputed.

### Using Routes Inner workings
During the process of using routes, the webserver saves the uploaded RouteString file and sends the string via POST. Then the browser extracts the needed arguments, name, and description, and dynamically generates a form for uploading and choosing the suitable arguments. This means that the system can respond to any amount of arguments can be made and called because of the dynamic generation. After saving argument images to a directory using an AJAX call, the user can generate their image on the same page with another AJAX call.

### Notes on Difficulty
I found that the architectural requirements to create this project was non-trivial. Two interfaces of communication and string formats had to be made so the browser and webserver and webserver and go code could pass information safely. A lot of parsing and generating different string formats was done to produce the proper results.