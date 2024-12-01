<script setup>
import { Icon } from '@iconify/vue';

import { useOptionsStore} from '@/stores/Options'
// let settings = ref(useOptionsStore().settings);
// const handleTheme = () => {
// }
</script>

<template>
    <div class="flex flex-col justify-evenly items-center min-h-svh min-w-[100svw]">
        <RouterLink to="/" class="absolute top-0 left-0 text-gray-600 hover:text-black hover:scale-110 ease-linear duration-[80ms]"><Icon icon="line-md:arrow-small-left" width="92" height="92"/></RouterLink>
        <h1 id="GameTitle" class="font-jersey text-[15rem]">Options</h1>
        <ol class="flex flex-col justify-around items-center w-1/2 font-jersey text-7xl gap-8 m-8 text-gray-800">
            <li class="w-full hover:text-black duration-300 transition ease-out 
                        flex justify-between">
                <label for="volumeSlider">Volume</label>
                <input id="volumeSlider" class="" type="range" name="volumeSlider" min="0" max="100" step="1"/>            
            </li>
            <li class="w-full hover:text-black duration-300 transition ease-out 
                        flex justify-between">
                <label>Theme:</label>
                <span id="themeSelectContainer" class="inline-flex flex-center items-center">
                    <select name="Theme" id="themeSelect" class=" text-6xl text-center p-3">
                        <option value="Blue">Blue Horizon</option>
                        <option value="Yellow">Solar Bliss</option>
                        <option value="BlackWhite">Noir blanc</option>
                    </select>
                </span>

            </li>        
        </ol>
    </div>
</template>

<style lang="scss" scoped>

#volumeSlider {
    appearance: none; /* removes browser-specific styling */
    width: 400px; /* width of slider */
    // height: 50px; /* height of slider */
    height: 60px;
    background: #FF5C35 ; /* orange background */
    outline: none; /* remove outline */
    border-radius: 12px; /* round corners */
    padding: 10px;
}

#volumeSlider::-webkit-slider-thumb {
  appearance: none; /* removes browser-specific styling */
  width: 50px; /* handle width */
  height:  50px; /* handle height */
  border-radius: 50%; /* make it circular */
  background: #FFFFFF; /* white color */
  cursor: pointer; /* cursor on hover */
}

#themeSelect {
}
</style>