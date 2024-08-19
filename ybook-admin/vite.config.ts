import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

import { createSvgIconsPlugin } from 'vite-plugin-svg-icons' // 导入插件
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),

    // 加上面的代码
    createSvgIconsPlugin({
      // 后续icon都放到 src/assets/icons 这个文件夹下
      iconDirs: [path.resolve(process.cwd(), 'src/assets/icons')],
      // Specify symbolId format
      symbolId: 'icon-[dir]-[name]'
    })
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  // sass全局变量的配置
  css: {
    preprocessorOptions: {
      scss: {
        javascriptEnabled: true,
        additionalData: '@import "./src/styles/variable.scss";'
      }
    }
  }
})
