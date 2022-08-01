import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Go语言学习笔记',
  base: '/my-go-tour/',
  themeConfig: {
    sidebar: [
      {
        text: 'Go语言基础',
        items:[
          {
            text: '基础',
            link: '/basic/go-basic'
          }
        ]
      }
    ]
  }
})