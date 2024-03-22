<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import main from "./assets/main.css";
import ModalWindow from './components/ModalWindow.vue';
</script>

<template>
  <ModalWindow/>
  <!-- <header>

    <div class="wrapper">
   
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to='/blog/2'>Blog</RouterLink>
      </nav>
    </div>
  </header> -->

  <RouterView />
</template>

<style scoped>



</style>
