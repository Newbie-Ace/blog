<template>
  <el-card>
    <div class="title-search">站内搜索</div>
    <el-input
      v-model="inputSearchValue"
      size="large"
      placeholder="输入关键字..."
      :suffix-icon="Search"
    ></el-input>
    <div class="desc-text">
      <div class="desc-face">（●＞ω＜●）</div>
      <div class="desc-font">这里什么都没有, 要不要试着搜点东西...</div>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Search } from '@element-plus/icons-vue'

const inputSearchValue = ref('')
</script>

<style scoped>
.el-card {
  width: 1000px;
  height: 550px;
  position: relative;
  margin: 0 auto;
}
.title-search {
  font-size: 26px;
}
.el-input {
  width: 800px;
  position: absolute;
  left: 50%;
  top: 30%;
  transform: translateX(-50%);
}
.desc-text {
  position: absolute;
  left: 50%;
  top: 60%;
  transform: translate(-50%, -50%);
  width: 450px;
  height: 150px;
  font-size: 24px;
}
.desc-face {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}
.desc-font {
  position: absolute;
  top: 30px;
}
</style>
