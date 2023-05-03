<template>
  <div class="header-box">
    <div class="logo-box">Nalamal's Blog</div>
    <el-menu
      :default-active="activeIndex"
      mode="horizontal"
      background-color="#fff"
      text-color="#000"
      active-text-color="#6b8770"
      :ellipsis="false"
      router
    >
      <el-menu-item index="login">
        <el-icon><House /></el-icon>
        首页
      </el-menu-item>
      <el-menu-item index="home">
        <el-icon><Document /></el-icon>
        博客
      </el-menu-item>
      <el-sub-menu index="cate">
        <template #title>
          <el-icon><FolderOpened /></el-icon>
          归档
        </template>
        <el-menu-item index="tech">技术</el-menu-item>
        <el-menu-item index="live">生活</el-menu-item>
      </el-sub-menu>
      <el-menu-item index="about">
        <el-icon><User /></el-icon>
        关于
      </el-menu-item>
      <el-menu-item index="write">
        <el-icon><Tickets /></el-icon>
        写作
      </el-menu-item>
      <el-menu-item index="search">
        <el-icon><Search /></el-icon>
        搜索
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { House, Document, FolderOpened, User, Tickets, Search } from '@element-plus/icons-vue'

const activeIndex = ref('home')
</script>

<style scoped>
.header-box {
  width: 100%;
  display: flex;
  justify-content: space-around;
  align-items: center;
}
.logo-box {
  color: #6b8770;
  font-size: 20px;
  font-weight: 700;
}
</style>
