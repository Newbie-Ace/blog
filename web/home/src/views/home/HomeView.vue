<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <HeaderComponent></HeaderComponent>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
      <el-footer>
        <FooterComponent></FooterComponent>
      </el-footer>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import HeaderComponent from '@/components/home/HeaderComponent.vue'
import FooterComponent from '@/components/home/FooterComponent.vue'
</script>

<style lang="less" scoped>
.common-layout {
  height: 100%;
}
.el-header {
  padding: 0;
}
.el-main {
  background-color: #6b8770;
  height: 100%;
}
.el-footer {
  width: 100%;
  height: 100px;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
