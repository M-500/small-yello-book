import { fileURLToPath, URL } from 'node:url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons' // 导入插件
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  // 获取环境变量
  const env = loadEnv(mode, process.cwd())
  return {
    base: '/ybook/',
    plugins: [
      vue(),
      // 配置自定义插件
      createSvgIconsPlugin({
        // Specify the icon folder to be cached
        iconDirs: [path.resolve(process.cwd(), 'src/assets/icons')],
        // Specify symbolId format
        symbolId: 'icon-[dir]-[name]'
      })
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
        // '@': path.resolve(__dirname, 'src')
      }
    },
    // 代理跨域配置
    server: {
      // host: '0.0.0.0',
      proxy: {
        [env.VITE_APP_BASE_API]: {
          // 代理的目标地址
          target: env.VITE_SERVER,
          changeOrigin: true, // 开启代理 代理跨域
          // 重写路径
          rewrite: (path) => path.replace(/^\/api/, '')
        }
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
  }
})
