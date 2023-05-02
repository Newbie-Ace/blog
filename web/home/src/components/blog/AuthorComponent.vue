<template>
  <el-card>
    <div class="author-info">站长信息</div>
    <div class="avatar-box">
      <img src="@/assets/img/avatar.jpg" />
    </div>
    <div class="author-name">Nalamal</div>
    <div class="cate-box">
      <div>文章</div>
      <div>标签</div>
      <div>归档</div>
    </div>
    <div class="cate-count">
      <div>xx</div>
      <div>xx</div>
      <div>xx</div>
    </div>
    <div class="github-icon">
      <el-icon @click="goToGithub"><IceTea /></el-icon>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { IceTea } from '@element-plus/icons-vue'

// gitbub网址
const githubUrl = 'https://github.com/Nalamal'

// 点击跳转github按钮，跳转到github首页
const goToGithub = () => {
  window.location.href = githubUrl
}
</script>

<style lang="less" scoped>
.el-card {
  height: 400px;
  margin: 10px 5px;
  position: relative;
}
.author-info {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  font-size: 20px;
  font-weight: 700;
}
.avatar-box {
  width: 150px;
  height: 150px;
  border-radius: 50%;
  position: absolute;
  left: 50%;
  top: 50px;
  transform: translateX(-50%);
  img {
    width: 100%;
    height: 100%;
    border-radius: 50%;
  }
}
.author-name {
  position: absolute;
  left: 50%;
  top: 220px;
  transform: translateX(-50%);
  font-weight: 600;
}
.cate-box {
  position: absolute;
  left: 50%;
  top: 270px;
  transform: translateX(-50%);
  width: 100%;
  display: flex;
  justify-content: space-around;
}
.cate-count {
  position: absolute;
  left: 50%;
  top: 290px;
  transform: translateX(-50%);
  width: 100%;
  display: flex;
  justify-content: space-around;
}
.github-icon {
  position: absolute;
  left: 50%;
  bottom: 30px;
  transform: translateX(-50%);
  font-size: 30px;

  :hover {
    color: #6b8770;
    cursor: pointer;
  }
}
</style>
