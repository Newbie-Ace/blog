<template>
  <el-card>
    <div class="title-box">title</div>
    <div class="footer-box">
      <div class="pub-time">
        <el-icon><Clock /></el-icon>
        time
      </div>
      <div class="article-tag">
        <el-icon><Promotion /></el-icon>
        tag
      </div>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { Clock, Promotion } from '@element-plus/icons-vue'
</script>

<style scoped>
.el-card {
  height: 150px;
  margin: 10px 5px;
  position: relative;
}
.title-box {
  font-size: 20px;
}
.footer-box {
  position: absolute;
  bottom: 20px;
}
</style>
