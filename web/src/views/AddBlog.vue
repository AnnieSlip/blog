<template>
    <header>
        <img :src="assets.logo" alt="logo">
    </header>
    <section class="wrapper">
        <a class="back" href="/">back</a>
        <div class="container">
        <h2>ბლოგის დამატება</h2>
        <h4>ატვირთეთ ფოტო</h4>
        <div class="input uploaded" v-if="uploadedImage?.name">{{ uploadedImage.name }}</div>
        <label v-else for="input-file" id="drop-area" @dragover="dragImage" @drop="dropImage">
        <input  
         @change="upload" 
         type="file" accept="image/*" id="input-file" hidden>
        <img :src="assets.upload" alt="upload logo">
        <p>ჩააგდეთ ფაილი აქ ან აირჩიეთ ფაილი</p>
  
    </label>
  
        <div class="inputs-wrapper">
            <div>
            <h4>ავტორი*</h4>
            <input placeholder="შეიყვანეთ ავტორი" :class="`${isAuthorValid ? 'input success' : 'input'}`" v-model="author">
            <ul>
                <li class="rules">მინიმუმ 4 სიმბოლო</li>
                <li class="rules">მინიმუმ 2 სიტყვა</li>
                <li class="rules">მხოლოდ ქართული სიმბოლოები</li>

            </ul>
            </div>
            <div>
                <h4>სათაური*</h4>
                <input placeholder="შეიყვანეთ სათაური"  :class="`${isTitleValid ? 'input success' : 'input'}`" v-model="title">
                <p class="rules">მინიმუმ ორი სიმბოლო</p>
            </div>
        </div>
        <h4>აღწერა *</h4>
        <textarea type="text" placeholder="შეიყვანეთ აღწერა" :class="`${isDescriptionValid ? 'input textarea success' : 'input textarea'}`"  v-model="description"></textarea>
        <p class="rules">მინიმუმ 2 სიმბოლო</p>
        <div class="inputs-wrapper">
            <div>
                <h4>გამოქვეყნების თარიღი*</h4>
                <input type="date" :class="`${isCreatedAtValid ? 'input success' : 'input'}`"  v-model="createdAt">
            </div>
            <div>
                <h4>კატეგორია*</h4>
                <div class="dropdown-wrapper">
                <div :class="`${isSelectedCategoriesValid ? 'input success dropdown' : 'input dropdown'}`"  @click="()=>isDropDownOpen=!isDropDownOpen">{{`${selectedCategories.length? selectedCategories: "შეიყვანეთ სათაური"}`}}</div>
                <div class="opened" v-if="isDropDownOpen">
                <div v-for="(category, index) in categories" :key="index" class="category" @click="()=>selectCategory(category)">
                    {{  category }}
                </div>
                </div>
                </div>
            </div>
        </div>
        <div class="inputs-wrapper">
            <div>
        <h4>ელ-ფოსტა</h4>
        <input type="email" placeholder="შეიყვანეთ მეილი" :class="`${isEmailValid ? 'input success' : 'input'}`" v-model="email">
        </div>
        <div><!--No Content--></div>
        </div>

        <div class="inputs-wrapper">
            <div><!--No Content--></div>
        <div>
            <button :class="`${isSubmitBtnDisabled ? 'disabled publish-btn' : 'publish-btn'}`"
         :disabled="isSubmitBtnDisabled">გამოქვეყნება</button>
        </div>
        </div>
        </div>
    </section>

</template>

<script lang="ts" setup>
import assets from '../assets';
import {ref, computed} from "vue";

const categories = ["UX/UI", "მარკეტი", "აპლიკაცია", "კვლევა", "FIGMA"];

const uploadedImage = ref();
const author= ref<string>();
const title = ref<string>();
const description = ref<string>();
const email = ref<string>("");
const createdAt = ref<Date>();
const selectedCategories = ref<string[]>([]);
const isDropDownOpen = ref<boolean>(false);

const dragImage = (e: any)=> {
    e.preventDefault();
}
const upload = (e: any) => {
    uploadedImage.value= e.target?.files[0];
};
const dropImage = (e: any) => {
    e.preventDefault();
    const files = e.dataTransfer.files;

    if (files.length > 0) {
        const uploadedFile = files[0];
        uploadedImage.value = uploadedFile;
    }
}

const selectCategory = (category: string)=>{
    if (!selectedCategories.value.includes(category)){
    selectedCategories.value.push(category);
    }
    isDropDownOpen.value=false;

}

// validations
const isSubmitBtnDisabled = computed(() => {
    if (isAuthorValid.value && isTitleValid.value && isDescriptionValid.value && isSelectedCategoriesValid.value && isCreatedAtValid.value && isEmailValid.value){
        return false;
    }
    return true;
})

const isAuthorValid = computed (() => {
    const georgianLettersRegex = /^[\u10D0-\u10F0\s]+$/;
    if(!author.value){
        return false;
    }
    if (author.value.length > 4 && author.value.split(/\W+/).length > 1 && georgianLettersRegex.test(author.value)  ){
        return true;
    }
    return false;
});

const isTitleValid = computed (()=>{
    if (!title.value || title.value.length < 2){
        return false;
    }
    return true;
})

const isDescriptionValid = computed (()=> {
    if (!description.value || description.value.length < 2){
        return false;
    }
    return true;
})

const isSelectedCategoriesValid = computed(()=>{
    if(selectedCategories.value.length){
        return true;
    }
    return false;
})

const isCreatedAtValid = computed(()=>{
    if (createdAt.value){
        return true;
    }
    return false;
})

const isEmailValid = computed(() => {
  const emailRegex = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
  return emailRegex.test(email.value);
});
</script>
<style lang="scss" scoped>

header{
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 15px 0px;


    img{
        cursor: pointer;
    }
}

.wrapper{
    background-color: #FBFAFF;
    padding: 40px 0px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    position: relative;

    a{
       position: absolute;
       top: 60px;
       left: 100px;
    }

    .container{
        width: 45%;

        .inputs-wrapper{
            display: flex;
            justify-content: space-between;
            gap: 30px;
            div{
                flex: 1;
            }
        }
       .uploaded{
        border: none !important;
        background-color: #F2F2F2;
        color: gray;
        display: flex;
        align-items: center;
        padding: 0px 13px;
       }

        .input{
            border-radius: 8px;
            width: 100%;
            border: 1px solid grey;
            height: 40px;
            padding: 5px 10px;
            
        }
  
        #drop-area{
            width: 100%;
            height: 180px;
            border-radius: 8px;
            background-color: #e6e6fa;
            border: 1px dashed gray;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;

        }
        .textarea{
            height: 170px;
            padding-top: 13px;
        }
        .rules{
            color: grey;
            font-size: 12px;
        }
        ul{

            padding: 0px 18px;
        }

        .dropdown-wrapper{
            position: relative;
        .dropdown{
            cursor: pointer;
            display: flex;
            align-items: center;
            background-color: #fff;
            
        }
        .opened{
            background-color: #fff;
            position: absolute;
            width: 100%;
            top: 52px;
            border-radius: 8px;
            .category{
                padding: 7px;
                display: flex;
                align-items: center;
                cursor: pointer;
                margin-top: 10px;
                font-size: 12px;
                border-radius: 12px;
                
            }

            :hover{
                .category {
                    background-color: gray;
                }
            }

         }
    }

 .publish-btn{
     display: flex;
    justify-content: center;
    width: 100%;
    align-items: center;
    padding: 10px 20px;
    border-radius: 8px;
    color: white;
    background-color: #5D37F3;
    cursor: pointer;
    border: none;
    margin-top: 60px;
        }
    }

 .disabled{
     background-color: gray !important;
     cursor: not-allowed !important;  

 }

 .success{
    border: 1px solid #72CC51 !important;
    background-color: #EDF9E9 !important;
 }

}

</style>