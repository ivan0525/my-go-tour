import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Go语言学习笔记',
  base: '/my-go-tour/',
  themeConfig: {
    sidebar: [
      {
        text: 'Go语言学习',
        items: [
          {
            text: '起步',
            link: '/start'
          },
          {
            text: '基础',
            link: '/basic/go-basic'
          }
        ]
      }
    ]
  }
})