<template>
  <div class="blog-container">
    <div class="blog-left">
      <ArticleComponent
        v-for="item in articleList"
        :key="item.id"
        :title="item.title"
        :time="item.time"
        :tag="item.tag"
      ></ArticleComponent>
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

import { ref } from 'vue'
import { selectBlogList } from '@/utils/service/blogRequest'
import { ElMessage } from 'element-plus'

// 定义Articlelist的类型
interface IArticleList {
  id: number
  title: string
  time: string
  tag: string
}

// 定义博客数据列表
const articleList = ref<IArticleList[]>([])

const getArticleList = async () => {
  const { data: res } = await selectBlogList(5, 1)
  res.data.forEach((item: any) => {})
  if (res.code !== 200) {
    return ElMessage.error('获取博客列表失败！')
  }
}
getArticleList()
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
