<template>
    <Teleport to="body">
      <Transition name="modal-fade">
        <div class="modal-wrapper"
        @click.self="store.closeModal"
        v-if="store.modalState?.component">
          <component :is="store.modalState?.component" v-bind="store.modalState?.props"  />
        </div>
      </Transition>
    </Teleport>
  </template>
  
  <script setup lang="ts">
import useModal from "../stores/useModal";

const store = useModal();


</script>
  
  
  <style scoped>
  .modal-wrapper {
    position: fixed;
    left: 0;
    top: 0;
  
    z-index: 500;
  
    width: 100vw;
    height: 100vh;
    background: rgba(0, 0, 0, 0.2);
  
    display: grid;
    place-items: center;
  }
  
  .modal-fade-enter-from,
  .modal-fade-leave-to {
    opacity: 0;
  }
  
  .modal-fade-enter-active,
  .modal-fade-leave-active {
    transition: 0.25s ease all;
  }
  </style>