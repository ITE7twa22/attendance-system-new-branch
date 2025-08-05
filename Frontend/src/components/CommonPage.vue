<template>
    <div>
      <!-- <img src="../assets/img/background.svg" class=" blur-sm fixed -z-10 origin-top-left -translate-x-1/3 -translate-y-1/3 scale-10 opacity-40 select-none"> -->
      <div class="py-24 sm:py-32">
        <div class="mx-auto max-w-7xl px-6 lg:px-8 flex flex-col items-center">
          <!-- <div>
            <router-link to="/" :class="{ 'active-tab': $route.path === '/' }" class="text-black-10 ml-2 text-sm">داخل المؤسسة</router-link>

            <router-link to="/HomeOutE7twaa" :class="{ 'active-tab': $route.path === '/HomeOutE7twaa' }" class="text-black-10 ml-2 text-sm">خارج المؤسسة</router-link>

</div> -->
<br>

          <!-- <img class="m-4 row-span-2 col-span-2 max-h-24 w-full object-contain lg:col-span-2" src="../assets/img/Logo-Dark-1.png" /> -->
         
          <br>

         
          <!-- <h1 class="text-4xl font-light text ">جمعية كسوة فرح</h1> -->
          <br>
          <p class="text-2xl   mb-10">نظام حساب ساعات التطوع</p>
          <br>
          

          <slot></slot>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  
  </script>
  <style>
  @import url('https://fonts.googleapis.com/css2?family=Mirza:wght@400;500;600;700&display=swap');
  /* Add styles for the active tab */
  .active-tab {
    font-weight: bold;
    font-size:larger;
    font-family: 'Mirza';
   /* background-color: #5daaae31 ; */
   color: black;
   padding: 10px;
   border-radius: 20px;
    /* Add any other styles you want for the active tab */
  }
  .text{
    color: #bfefef;
   /*  --tw-text-opacity: 1;
    color: rgb(156 163 175 / var(--tw-text-opacity)) */
  }
  *{
    font-family: 'Mirza';
  }
  .btn{
  --tw-bg-opacity: 1;
  /* background-color: rgb(191 239 239/ var(--tw-bg-opacity)) ; */
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;

}
body {
 /* /* background-color: #5daaae31; Light Background Color  */
  background-color: #5daaae;
   background-color: #D8D3C9;
}
p {
  color: #14124D; /* كمثال */
}


  </style>