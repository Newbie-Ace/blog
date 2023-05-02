<template>
  <el-card>
    <div class="header-box">近期文章</div>
    <div class="article-list">
      <div class="article-title">title</div>
      <div class="article-time">time</div>
    </div>
  </el-card>
</template>

<script setup lang="ts"></script>

<style scoped>
.el-card {
  height: 385px;
  margin: 10px 5px;
  position: relative;
}
.header-box {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  font-size: 20px;
  font-weight: 700;
}
.article-list {
  position: absolute;
  top: 50px;
  font-size: 16px;
}
</style>
