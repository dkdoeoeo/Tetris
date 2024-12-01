<script setup>
import { Icon } from '@iconify/vue';
import { ref } from 'vue';

const audio = new Audio('/public/ocean.mp3');
audio.loop = true; // Optional: Loop the audio
audio.volume = 0.5; // Set volume (0.0 to 1.0)
const ifAudioPlay = ref(false)

const playPauseAudio = () => {
  if(ifAudioPlay.value) {
    ifAudioPlay.value = !ifAudioPlay.value
    audio.pause().catch(error => {
      console.error('Playback failed:', error);
    });
  }
  else {
    ifAudioPlay.value = !ifAudioPlay.value
    audio.play().catch(error => {
        console.error('Playback failed:', error);
    });
  }
};

</script>

<template>
  <div id="router-view" class="relative flex flex-col justify-start h-full w-full min-h-[100vh] min-w-[100vw]">
    <RouterView></RouterView>
    <button @click="playPauseAudio" class="absolute top-0 right-0 text-gray-600 hover:text-black hover:scale-110 ease-linear duration-[80ms]">
      <Icon v-if="ifAudioPlay" icon="bx:volume-full" width="92" height="92"/>
      <Icon v-else icon="bx:volume-mute" width="92" height="92" />
    </button>

  </div>

</template>

<style scoped></style>
