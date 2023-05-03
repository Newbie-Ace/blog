import service from './index'

export const selectBlogList = (page_num: number, page_size: number) => {
  return service({
    url: '/blog',
    method: 'get',
    params: {
      page_num,
      page_size
    }
  })
}
