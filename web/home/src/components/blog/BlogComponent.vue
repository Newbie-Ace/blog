<template>
  <div class="blog-container">
    <div class="blog-left">
      <ArticleComponent></ArticleComponent>
      <ArticleComponent></ArticleComponent>
      <ArticleComponent></ArticleComponent>
      <ArticleComponent></ArticleComponent>
      <ArticleComponent></ArticleComponent>
    </div>
    <div class="blog-right">
      <AuthorComponent></AuthorComponent>
      <RecentArticleComponent></RecentArticleComponent>
    </div>
  </div>
  <el-card class="page-box">
    <el-pagination layout="prev, pager, next" :total="50" />
  </el-card>
</template>

<script setup lang="ts">
import ArticleComponent from './ArticleComponent.vue'
import AuthorComponent from './AuthorComponent.vue'
import RecentArticleComponent from './RecentArticleComponent.vue'
</script>

<style scoped>
.blog-container {
  width: 1000px;
  margin: 0 auto;
  display: flex;
  background-color: #6b8770;
  justify-content: space-between;
}
.blog-left {
  flex: 3;
}
.blog-right {
  flex: 1;
}
.page-box {
  width: 990px;
  margin: 0 auto;
  display: flex;
  justify-content: center;
}
</style>
