<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <div class="logo-box">Nalamal's Blog</div>
        <el-menu
          mode="horizontal"
          background-color="#ededed"
          text-color="#000"
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
          <el-menu-item index="search">
            <el-icon><Search /></el-icon>
            搜索
          </el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
      <el-footer>
        <div>
          <div>Create in May 2023</div>
          <div>Power by Vue and Element-plus</div>
          <div>AQ: xxxxx</div>
        </div>
      </el-footer>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { House, Document, FolderOpened, User, Search } from '@element-plus/icons-vue'
</script>

<style lang="less" scoped>
.common-layout {
  background-color: #6b8770;
  height: 100%;
}
.el-header {
  background-color: #ededed;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 150px;
}
.logo-box {
  color: #6b8770;
  font-size: 20px;
  font-weight: 700;
}
.el-menu--horizontal {
  border-bottom: none;
}
.el-main {
  background-color: #6b8770;
  height: 100%;
}
.el-footer {
  width: 100%;
  height: 100px;
  background-color: #ededed;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
