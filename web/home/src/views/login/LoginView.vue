<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-avatar">
        <img src="@/assets/img/avatar.jpg" />
      </div>
      <div class="desc-box">一个小众化的博客</div>
      <div class="btn-group">
        <el-button color="#6b8770" round @click="goToHome">进入博客</el-button>
        <el-button color="#6b8770" round @click="goToGithub">GitHub</el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import router from '@/router'

// gitbub网址
const githubUrl = 'https://github.com/Nalamal'

// 点击跳转github按钮，跳转到github首页
const goToGithub = () => {
  window.location.href = githubUrl
}

// 点击按钮，进入博客主页面
const goToHome = () => {
  router.push('/home')
}
</script>

<style lang="less" scoped>
.login-container {
  background-color: #6b8770;
  height: 100%;
}
.login-box {
  width: 300px;
  height: 550px;
  background-color: #ededed;
  border-radius: 10px;
  position: relative;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
.login-avatar {
  width: 150px;
  height: 150px;
  border: 1px solid #6b8770;
  border-radius: 50%;
  box-shadow: 0 0 20px #6b8770;
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
.desc-box {
  width: 100%;
  position: absolute;
  top: 250px;
  color: #6b8770;
  font-size: 20px;
  text-align: center;
}
.btn-group {
  position: absolute;
  top: 350px;
  left: 50%;
  transform: translateX(-50%);
}
.el-button {
  width: 200px;
  margin: 5px auto;
}
</style>
