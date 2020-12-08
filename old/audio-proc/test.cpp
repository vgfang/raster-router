#include <iostream>
#define TSF_IMPLEMENTATION
#include "./TinySoundFont/tsf.h"
int main(){
	SDL_AudioSpec OutputAudioSpec;
	OutputAudioSpec.freq = 44100;
	OutputAudioSpec.format = AUDIO_S16;
	OutputAudioSpec.channels = 2;
	OutputAudioSpec.samples = 4096;
	OutputAudioSpec.callback = AudioCallback;
	
	tsf* TinySoundFont = tsf_load_filename("testfiles/sans.sf2");
	tsf_set_output(TinySoundFont, TSF_MONO, 44100, 0); //sample rate
	tsf_note_on(TinySoundFont, 0, 60, 1.0f); //preset 0, middle C
	short HalfSecond[22050]; //synthesize 0.5 seconds
	tsf_render_short(TinySoundFont, HalfSecond, 22050, 0);
	return 0;
}