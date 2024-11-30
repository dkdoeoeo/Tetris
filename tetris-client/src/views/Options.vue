<script setup>
import { useOptionsStore} from '@/stores/Options'
// let settings = ref(useOptionsStore().settings);
// const handleTheme = () => {
// }
</script>

<template>
    <div class="flex flex-col justify-center items-center min-h-svh min-w-[100svw]">
        <h1 id="GameTitle" class="font-jersey text-[15rem]">Tetris</h1>
        <ol class="flex flex-col justify-around items-center font-jersey text-7xl gap-8 m-8 text-gray-800">
            <li class=" hover:scale-[1.2] hover:text-black duration-300 transition ease-out"></li>
            <li class=" hover:scale-[1.2] hover:text-black duration-300 transition ease-out"></li>
            <li class=" hover:scale-[1.2] hover:text-black duration-300 transition ease-out"></li>
            <!-- <li class=" hover:scale-110 duration-300 transition ease-out"><RouterLink class="menu-option">About</RouterLink></li> -->
        </ol>
    </div>
</template>

<style lang="scss" scoped>


</style>