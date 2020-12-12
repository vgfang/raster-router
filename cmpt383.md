# Raster Router
Raster Router is a web application for producing templates for algorithmic image editing. The design goal for these templates, or _routes_ is to have a black-box, functional design for creating images where the user inputs arguments (text or other images), to receive a modified image as a result. 

## Programming Langauges
It is built on the Flask webframework, using JQuery for front-end commands, and Golang for image production. Ctypes is used for calling performant Golang
raster image functions from Python. Python uses Flask as a webserver to communicate the client-side Javascript web app.

### External Libraries
- github.com/fogleman/gg
   + A golang graphics library was used to provide simple image manipulation functions
- github.com/nfnt/resize
   + A golang library for resizing images
- github.com/eligrey/FileSaver.js
   + A Javascript file was used to handle client-side downloading of a JS string for the route creation

### Deployment
- Vagrant VM and Chef are used to deploy a Virtual Machine.
1. Create and Start the Vagrant VM
2. Start the Flask server: `cd project; python3 app.py`
3. The local IP will be: `192.168.33.10:5000`

## How it works
### Routestrings
Routestrings (`.route`) are a custom plaintext markup format to represent all of the data needed to create a _route_. They are designed to be human readable and contain the information for a route's name, description, arguments, and commands. To make these processable with user input, three break characters had to be made. These are the newline(`\sn\n`), space(\s: ), and argument(\arg: ).

During the process of creating routes, the browser client is creating a routestring in the background using Javascript. The image generation functionality is achieved though AJAX calls from Javascript to Flask, where the file manipulation is done, which calls the Golang binary to generate a new image. Lastly, Javascript updates the browser client image.

### Commandstrings
A Commandstring is a single line of plaintext which represents a single command. These can be ran sequentially and running a routestring executes a list of commands to produce the final image. These strings can be parsed by the Golang code, to call different functions and are passed from Javascript to Python to the Golang shared library.

### Argumentstrings
Argumentstrings are a plaintext file format, similar to a Routestring in syntax, which defines the user's arguments during _Image routing_. These are used provide a linking mechanism between arguments and the routestring commands.

## How to use
### Creating Routes
1. The user generates a blank canvas of the size that they want.
2. The user then can add arguments and commands. This is rendered in a preview image which only shows the difference of a single command. To use multiple commands, the user must sequentially write one command and continue.
3. Arguments are used by first adding an argument, then referencing it in either the text of an `ADD_TEXT` command or the imagefile of an `ADD_IMAGE` command, using the prefix `\arg: ` before the argument name.
4. Downloading the route serves a `.route` extention Routestring. This can be used to _route_ an image.

### Using a Route
1. To use a route, the user must upload a valid Routestring file.
2. Users are then redirected to a page where they can choose their arguments and read the route name and description. A dynamic form is generated using the Routestring's argument fields.
3. To give an image argument, the user must upload a valid image file, then upload them to the server, using the upload button.
4. To give a text argument, the user can just type into a textarea element.
5. To generate the final image, pressing the "Generate Image" button, and waiting will render the image below on the same page. 


