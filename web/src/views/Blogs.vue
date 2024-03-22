<template>
  <Header/>

<section class="main">
    <div class="header">
<h2>ბლოგი</h2>
<img :src="assets.blogLogo" alt="blog logo" />
</div>

<div class="filters">
  <div v-for="(filter, index) in filters" 

  :key="filter.id" class="filter"  
  :style="{'background-color': colors[index].background, 'color': colors[index].fontColor}"
  :class="{ active: filter.isActive }"
  @click="() => setActiveFilter(filter)">
{{ filter.name }}

  </div>
</div>

<div class="blogs-wrapper">
  <Blog v-for="blog in filteredBlog" :key="blog.id" :blog="blog"/>
</div>
  </section>
</template>


<script lang="ts" setup>
import Header from "./../components/Header.vue";
import assets from "./../assets";
import {filters} from "./../data/filters";
import {blogsData} from "./../data/blogData";
import {ref, computed} from "vue";
import Blog from "./../components/Blog.vue"


const activeFilters = ref<string[]>([]);

const setActiveFilter = (filter: any) => {
  if (filter.isActive === true){
    activeFilters.value = activeFilters.value.filter((val: string)=> val != filter.name);
    filter.isActive = false;
    return;
  };
  filter.isActive = true;
  activeFilters.value.push(filter.name);
  
  return;

};

const filteredBlog = computed(()=>{
  if(!activeFilters.value.length){
    return blogsData;
  } 

let blogsFiltered: any=[];
for (let i=0; i<activeFilters.value.length; i++){
  blogsData.map((blog: any)=> {
    if (blog.filters.includes(activeFilters.value[i])&& !blogsFiltered.includes(blog)){
      blogsFiltered.push(blog);
    }
  })
}
return blogsFiltered;
})

const colors = [
{background:  "#ffe6e6",
  fontColor: " #ff4d4d",
},
{background:  "#e6ecff",
  fontColor: "#3366ff",
},
{background:  "#ffffff",
  fontColor: "#b3b300",
},
{background:  "#e6fff9",
  fontColor: " #4dffd5",
},
{background:  "#f9e6ff",
  fontColor: "#9c00cc",
},
{background:  " #f9f2ec",
  fontColor: " #ac7339",
},
]

</script>

<style lang="scss" scoped>
.main{
  background-color: #FBFAFF;
  padding: 20px 40px;

  .header{
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .filters{
    display: flex;
    gap: 20px;
    justify-content: center;
    align-items: center;
    margin: 50px 0px;

    .filter {
     border-radius: 20px;
     padding: 5px 20px;
     cursor: pointer;
     border: 1px solid transparent;
    }
    .active{
      border: 1px solid black;
    }
  }

  .blogs-wrapper{
    display: flex;
    justify-content: space-between;
    align-items: stretch; 
    flex-wrap: wrap; 



  }
  .blogs-wrapper > * {
  flex: 0 0 calc(33.333% - 10px);
  margin-bottom: 20px; 

}
}
</style>