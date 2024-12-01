<script setup>

</script>

<template>
    <div class="flex flex-col justify-center items-center min-h-svh min-w-[100svw]">
        <h1 id="GameTitle" class="font-jersey text-[15rem]">Tetris</h1>
        <ol class="flex flex-col justify-around items-center font-jersey text-7xl gap-8 m-8 text-gray-600">
            <li class=" hover:scale-[1.2] hover:text-black duration-300 transition ease-out"><RouterLink to="/2player" class="menu-option">Start Game</RouterLink></li>
            <li class=" hover:scale-[1.2] hover:text-black duration-300 transition ease-out"><RouterLink to="/Options" class="menu-option">Options</RouterLink></li>
            <li class=" hover:scale-[1.2] hover:text-black duration-300 transition ease-out"><RouterLink to="/" class="menu-option">About</RouterLink></li>
            <!-- <li class=" hover:scale-110 duration-300 transition ease-out"><RouterLink class="menu-option">About</RouterLink></li> -->
        </ol>
    </div>
</template>

<style lang="scss" scoped>

.menu-option {
    white-space: nowrap; /* Prevent text wrapping */
    padding: 5%;   
    text-decoration: none;
}


</style>