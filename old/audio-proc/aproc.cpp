//g++ -shared -o aproc.so -fPIC aproc.cpp
#define TSF_IMPLEMENTATION
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dlfcn.h>
#include "./TinySoundFont/tsf.h"
#include "minisdl_audio.h"
#define FN_B 255 // filename char buffer

// Input: string containing 
extern "C"
char* Load (char* i) {
	char* str = (char*)malloc(6*sizeof(char));
	strcpy(str, i);
	printf("name = %s\n", str);
	return str;
}

extern "C"
int Add (int a, int b){
	return a + b;
}

extern "C"
int main(){
	tsf* TinySoundFont = tsf_load_filename("testfiles/sans.sf2");
	tsf_set_output(TinySoundFont, TSF_MONO, 44100, 0); //sample rate
	tsf_note_on(TinySoundFont, 0, 60, 1.0f); //preset 0, middle C
	short HalfSecond[22050]; //synthesize 0.5 seconds
	tsf_render_short(TinySoundFont, HalfSecond, 22050, 0);
	return 0;
}