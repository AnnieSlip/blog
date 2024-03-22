<template>
 <section class="wrapper">
 <img :src="props.blog.img" alt="image"  class="img"/>
 <h4>{{ props.blog.author }}</h4>
 <h6>{{ props.blog.created_at }}</h6>
 <h4>{{ props.blog.title }}</h4>

 <div class="filters">
    <div v-for="(filter, index) in props.blog.filters" :key="index" >{{ filter }}</div>
 </div>

 <p>{{ props.blog.text }}</p>

 <button @click="visitBlog" class="btn">სრულად ნახვა</button>

 </section>
</template>

<script lang="ts" setup>
import {useRouter} from "vue-router";

const router = useRouter();

type Blog = {
    id: number,
    title: string,
    created_at: string,
    text: string,
    filters: string[],
    author: string,
    img: string,
    mail: string,

}
const props = defineProps<{
	blog: Blog;
}>();

const visitBlog = () => {
    router.push(`/blogs/${props.blog.id}`)

}
</script>

<style lang="scss" scoped>
.wrapper {
    width: 430px;
}
.img{
    border-radius: 8px;
    width: 430px;
    height: 300px;
}
.filters{
    display: flex;
    gap: 10px;
}
.btn{
    color: purple;
}
</style>