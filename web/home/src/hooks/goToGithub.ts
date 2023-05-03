import { ref } from 'vue'

// github主页地址
const githubUrl = ref('https://github.com/Nalamal')

// 跳转github
export const goToGithub = () => {
  window.location.href = githubUrl.value
}
