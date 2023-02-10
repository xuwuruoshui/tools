<template>
  <div id="about">
    <h1>About</h1>
    <el-button @click="flag=!flag">click me</el-button>
    <keep-alive include="A">
      <A v-if="flag"></A>
      <B v-else></B>
    </keep-alive>
  </div>
</template>

<script lang="ts" setup>
import A from '@/components/A.vue'
import B from '@/components/B.vue'
import {ref} from "vue";

let flag = ref<Boolean>(true)
</script>

