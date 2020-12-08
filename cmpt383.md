# Raster Router
Raster Router is a web application for producing templates for algorithmic image editing. The design goal for these templates, or _routes_ is to have a black-box, functional design for creating images where the user inputs arguments, such as text or other images, to receive a modified image as a result. 

## Creating Routes
These routes are easily sharable and and the results will be pipable to other routes, capable of creating an extended chain of sequential image edits. Routes are created by users, through designating ordered steps, with only a single undo available during the route creating process.

## Sharing Routes
Routes will be available for one week after its creation and will be deleted unless on the top 120 of the main leaderboard or top 36 on the weekly leaderboard. Routes can be shared through a single url that asks for inputs, and returns an output image.

## Leaderboard
There will be a leaderboard of routes of the top 120 most used routes, where a vote is equal to 3 uses. These 120 routes are saved in storage indefinately. A weekly leaderboard is made for so new routes have a chance to make it to the top 120. Routes that do not make it to the to 36 of the weekly leaderboard and are not in the top routes, after a week following its creation will be deleted for space.

## Programming Langauges
It is built on the Flask webframework, using JQuery for front-end commands, and Golang for image production. Ctypes is used for calling performant Golang
raster image functions from Python. The library Go Graphics (github.com/fogleman/gg) was used to simplify image editing. Python uses Flask to communicate the client-side Javascript web app.

## Deployment Technology
- Vagrant VM + Chef