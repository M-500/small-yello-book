import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  // server: {
  //   host: '0.0.0.0',
  //   // port: 4000,
  //   proxy: {
  //     '/api': {
  //       target: 'http://127.0.0.1:8122',
  //       // secure: false, // 请求是否为https
  //       changeOrigin: true,
  //       rewrite:(path)=>path.replace(/^\/api/,'') //api替换为'',
  //     },
  //   },
  // },
  // outDir: 'dist',
  server: {
    // hostname: '0.0.0.0',
    host: "localhost",
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://localhost:8122',
        changeOrigin: true,
        ws: true,
        rewrite: (pathStr) => pathStr.replace(/^\/api/, '')
      },
    },
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src')
    }
  },
  plugins: [vue()],
})
