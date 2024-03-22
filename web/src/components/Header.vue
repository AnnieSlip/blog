<template>
    <header>
        <img :src="assets.logo" alt="company logo" class="logo"/>
        <button class="btn" @click="handleClick">
            {{ isLoggedIn? "ბლოგის დამატება": "შესვლა" }}
        </button>

    </header>
</template>


<script lang="ts" setup>
import assets from '../assets';
import InfoModalWindow from './InfoModalWindow.vue';
import useModal from '@/stores/useModal';
import {ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();
const store = useModal();

function handleClick() {    
   if( isLoggedIn.value) {
    router.push("/blogs/add")
   }else{
  store.openModal({ component: InfoModalWindow });
   }
}

const isLoggedIn = ref(true);

</script>

<style lang="scss" scoped>
header{
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px 40px;
}
.logo{
    cursor: pointer;
}
.btn{
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 10px 20px;
    border-radius: 8px;
    color: white;
    background-color: #5D37F3;
    cursor: pointer;
    border: none;

}
</style>