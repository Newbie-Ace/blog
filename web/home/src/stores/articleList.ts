import { defineStore } from 'pinia'
import { ElMessage } from 'element-plus'
import { selectBlogList } from '@/utils/service/blogRequest'

interface IArticleList {
  title: string
  time: string
  tags: string
}

const useArticleList = defineStore('articleList', {
  state: (): IArticleList => ({
    title: '',
    time: '',
    tags: ''
  }),
  actions: {
    async getArticleList() {
      const { data: res } = await selectBlogList(5, 1)
      console.log(res.data)
      if (res.code !== 200) {
        return ElMessage.error('获取博客列表失败！')
      }
    }
  }
})

export default useArticleList
