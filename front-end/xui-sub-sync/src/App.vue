<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { ref } from 'vue'
const activeIndex = ref('1')
const handleSelect = (key, keyPath) => {
  console.log(key, keyPath)
}

</script>

<template>
  <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      @select="handleSelect"
  >
    <el-menu-item index="1"><RouterLink to="/">Home</RouterLink></el-menu-item>
    <el-menu-item index="2"><RouterLink to="/login">登录</RouterLink></el-menu-item>
  </el-menu>

  <RouterView />
</template>

<style scoped>

</style>
