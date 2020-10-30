## Topic idea
REST web api for simple soundfont generation. Given a string and sound, the website will generate and preview an audio file using the given sound synthesized to be pitched-shifted at certain syllables with appropriate pauses to give a mimicry of the english language. For example, if the user entered "Why? I didn't think so.", the audio generated would be 6 syllables of the chosen sound; the question mark would create a pause after and a pitch up of the first syllable and "so" would be in a lower pitch and longer than "I". Pitch ranges would be adjustable from the web interface. The possible uses for this project could be audio for videos or dialogue sounds in a video game without voice acting.

## Programming Languages
1. Python
   + Web backend using Flask
2. Javascript
   + Front-end web logic and event handling
   + HTML/CSS/Jquery included
3. C++
   + Used for performant audio processing
   + TinySoundFont synthesizer library

## Inter-language communication methods
1. Flask (Python to Javascript) REST server
2. SWIG (Python to C++)

## Deployment Technology
- Vagrant VM + Chef