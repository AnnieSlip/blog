<template>
    <div class="inner">
        <h4>შესვლა</h4>
        <div class="email">
      <p class="lb">ელ-ფოსტა</p>
      <input type="email" id="email" name="email">
      </div>
      <button class="btn">შესვლა</button>
      <button class="close-btn"  @click="store.closeModal">X</button>

    </div>
  </template>
  
  <script setup lang="ts">
  import useModal from "../stores/useModal";

const store = useModal();

</script>
  
  <style lang="scss" scoped>

  .inner {
    background-color: white;
    width: 450px;
    padding: 1.2rem;
    border-radius: 1rem;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 1.4rem;
    position: relative;
  }
  input, .btn{
    width: 100%;
  }
  
  .email {
    width: 100%;


  }
  .lb{
        margin-bottom: 10px;
    }
  input{
    border: none;
    background-color: aliceblue;
    padding: 12px 8px;
    border-radius: 8px;
  }
  .btn{
    color: white;
    background-color: #5D37F3;
    border-radius: 8px;
    padding: 12px 0px;
  }
  .inner > h2 {
    font-size: 3rem;
    font-weight: bold;
  }
  
  .inner > p {
    font-size: 1.8rem;
    font-weight: 300;
  }

  .close-btn{
    position: absolute;
    top: 17px;
    right: 22px;
    color: black;
    font-weight: 600;
    font-size: 14px;
  }
  </style>