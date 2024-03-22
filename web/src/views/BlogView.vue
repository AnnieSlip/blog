<template>
  <Header/>
  <section class="blog" v-if="blog">
    <a class="back" href="/">back</a>
    <div>
    <img :src="blog.img" alt="image" class="img"/>
    <h6>{{ blog.author }}</h6>
    <div class="info">
      <p>{{ blog.created_at }}</p>
      <p>{{ blog.mail }}</p>
    </div>

    <h2>{{ blog.title }}</h2>
    <div class="filters">
      <div v-for="(filter, index) in blog.filters" :key="index">{{ filter }}
      </div>
    </div>
    <p>{{ blog.text }}</p>
    </div>
  </section>
</template>


<script lang="ts" setup>
import { blogsData } from '@/data/blogData';
import { useRoute } from 'vue-router';
import {onMounted, ref} from 'vue';
import Header from '../components/Header.vue'

const route = useRoute();
const blog = ref();
const index = ref();



onMounted(() => {
  index.value = parseInt(route.params.id as string);
  blog.value = blogsData[index.value-1];
  
});

</script>
<style lang="scss" scoped>

.blog{
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
  padding: 30px 0px;

  .back{
    position: absolute;
    top: 30px;
    left: 30px;
  }

  div{
    width: 60%;
  }
  .filters{
    display: flex;
    gap: 10px;
    font-size: 12px;
  }
}
.info{
  color: gray;
  display: flex;
  gap: 10px;
}

.img{
  width: 100%;
  height: 400px;
  border-radius: 8px;

}
</style>
