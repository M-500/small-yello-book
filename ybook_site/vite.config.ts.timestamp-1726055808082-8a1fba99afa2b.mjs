// vite.config.ts
import { fileURLToPath, URL } from "node:url";
import { defineConfig, loadEnv } from "file:///Users/wulinlin/workspace/github.com/wll/small-yello-book/ybook_site/node_modules/vite/dist/node/index.js";
import vue from "file:///Users/wulinlin/workspace/github.com/wll/small-yello-book/ybook_site/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import { createSvgIconsPlugin } from "file:///Users/wulinlin/workspace/github.com/wll/small-yello-book/ybook_site/node_modules/vite-plugin-svg-icons/dist/index.mjs";
import path from "path";
var __vite_injected_original_import_meta_url = "file:///Users/wulinlin/workspace/github.com/wll/small-yello-book/ybook_site/vite.config.ts";
var vite_config_default = defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd());
  return {
    base: "/ybook/",
    plugins: [
      vue(),
      // 配置自定义插件
      createSvgIconsPlugin({
        // Specify the icon folder to be cached
        iconDirs: [path.resolve(process.cwd(), "src/assets/icons")],
        // Specify symbolId format
        symbolId: "icon-[dir]-[name]"
      })
    ],
    resolve: {
      alias: {
        "@": fileURLToPath(new URL("./src", __vite_injected_original_import_meta_url))
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
          changeOrigin: true,
          // 开启代理 代理跨域
          // 重写路径
          rewrite: (path2) => path2.replace(/^\/api/, "")
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
  };
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCIvVXNlcnMvd3VsaW5saW4vd29ya3NwYWNlL2dpdGh1Yi5jb20vd2xsL3NtYWxsLXllbGxvLWJvb2sveWJvb2tfc2l0ZVwiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9maWxlbmFtZSA9IFwiL1VzZXJzL3d1bGlubGluL3dvcmtzcGFjZS9naXRodWIuY29tL3dsbC9zbWFsbC15ZWxsby1ib29rL3lib29rX3NpdGUvdml0ZS5jb25maWcudHNcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfaW1wb3J0X21ldGFfdXJsID0gXCJmaWxlOi8vL1VzZXJzL3d1bGlubGluL3dvcmtzcGFjZS9naXRodWIuY29tL3dsbC9zbWFsbC15ZWxsby1ib29rL3lib29rX3NpdGUvdml0ZS5jb25maWcudHNcIjtpbXBvcnQgeyBmaWxlVVJMVG9QYXRoLCBVUkwgfSBmcm9tICdub2RlOnVybCdcblxuaW1wb3J0IHsgZGVmaW5lQ29uZmlnLCBsb2FkRW52IH0gZnJvbSAndml0ZSdcbmltcG9ydCB2dWUgZnJvbSAnQHZpdGVqcy9wbHVnaW4tdnVlJ1xuaW1wb3J0IHsgY3JlYXRlU3ZnSWNvbnNQbHVnaW4gfSBmcm9tICd2aXRlLXBsdWdpbi1zdmctaWNvbnMnIC8vIFx1NUJGQ1x1NTE2NVx1NjNEMlx1NEVGNlxuaW1wb3J0IHBhdGggZnJvbSAncGF0aCdcblxuLy8gaHR0cHM6Ly92aXRlanMuZGV2L2NvbmZpZy9cbmV4cG9ydCBkZWZhdWx0IGRlZmluZUNvbmZpZygoeyBtb2RlIH0pID0+IHtcbiAgLy8gXHU4M0I3XHU1M0Q2XHU3M0FGXHU1ODgzXHU1M0Q4XHU5MUNGXG4gIGNvbnN0IGVudiA9IGxvYWRFbnYobW9kZSwgcHJvY2Vzcy5jd2QoKSlcbiAgcmV0dXJuIHtcbiAgICBiYXNlOiAnL3lib29rLycsXG4gICAgcGx1Z2luczogW1xuICAgICAgdnVlKCksXG4gICAgICAvLyBcdTkxNERcdTdGNkVcdTgxRUFcdTVCOUFcdTRFNDlcdTYzRDJcdTRFRjZcbiAgICAgIGNyZWF0ZVN2Z0ljb25zUGx1Z2luKHtcbiAgICAgICAgLy8gU3BlY2lmeSB0aGUgaWNvbiBmb2xkZXIgdG8gYmUgY2FjaGVkXG4gICAgICAgIGljb25EaXJzOiBbcGF0aC5yZXNvbHZlKHByb2Nlc3MuY3dkKCksICdzcmMvYXNzZXRzL2ljb25zJyldLFxuICAgICAgICAvLyBTcGVjaWZ5IHN5bWJvbElkIGZvcm1hdFxuICAgICAgICBzeW1ib2xJZDogJ2ljb24tW2Rpcl0tW25hbWVdJ1xuICAgICAgfSlcbiAgICBdLFxuICAgIHJlc29sdmU6IHtcbiAgICAgIGFsaWFzOiB7XG4gICAgICAgICdAJzogZmlsZVVSTFRvUGF0aChuZXcgVVJMKCcuL3NyYycsIGltcG9ydC5tZXRhLnVybCkpXG4gICAgICAgIC8vICdAJzogcGF0aC5yZXNvbHZlKF9fZGlybmFtZSwgJ3NyYycpXG4gICAgICB9XG4gICAgfSxcbiAgICAvLyBcdTRFRTNcdTc0MDZcdThERThcdTU3REZcdTkxNERcdTdGNkVcbiAgICBzZXJ2ZXI6IHtcbiAgICAgIC8vIGhvc3Q6ICcwLjAuMC4wJyxcbiAgICAgIHByb3h5OiB7XG4gICAgICAgIFtlbnYuVklURV9BUFBfQkFTRV9BUEldOiB7XG4gICAgICAgICAgLy8gXHU0RUUzXHU3NDA2XHU3Njg0XHU3NkVFXHU2ODA3XHU1NzMwXHU1NzQwXG4gICAgICAgICAgdGFyZ2V0OiBlbnYuVklURV9TRVJWRVIsXG4gICAgICAgICAgY2hhbmdlT3JpZ2luOiB0cnVlLCAvLyBcdTVGMDBcdTU0MkZcdTRFRTNcdTc0MDYgXHU0RUUzXHU3NDA2XHU4REU4XHU1N0RGXG4gICAgICAgICAgLy8gXHU5MUNEXHU1MTk5XHU4REVGXHU1Rjg0XG4gICAgICAgICAgcmV3cml0ZTogKHBhdGgpID0+IHBhdGgucmVwbGFjZSgvXlxcL2FwaS8sICcnKVxuICAgICAgICB9XG4gICAgICB9XG4gICAgfSxcbiAgICAvLyBzYXNzXHU1MTY4XHU1QzQwXHU1M0Q4XHU5MUNGXHU3Njg0XHU5MTREXHU3RjZFXG4gICAgY3NzOiB7XG4gICAgICBwcmVwcm9jZXNzb3JPcHRpb25zOiB7XG4gICAgICAgIHNjc3M6IHtcbiAgICAgICAgICBqYXZhc2NyaXB0RW5hYmxlZDogdHJ1ZSxcbiAgICAgICAgICBhZGRpdGlvbmFsRGF0YTogJ0BpbXBvcnQgXCIuL3NyYy9zdHlsZXMvdmFyaWFibGUuc2Nzc1wiOydcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cbiAgfVxufSlcbiJdLAogICJtYXBwaW5ncyI6ICI7QUFBOFgsU0FBUyxlQUFlLFdBQVc7QUFFamEsU0FBUyxjQUFjLGVBQWU7QUFDdEMsT0FBTyxTQUFTO0FBQ2hCLFNBQVMsNEJBQTRCO0FBQ3JDLE9BQU8sVUFBVTtBQUwrTixJQUFNLDJDQUEyQztBQVFqUyxJQUFPLHNCQUFRLGFBQWEsQ0FBQyxFQUFFLEtBQUssTUFBTTtBQUV4QyxRQUFNLE1BQU0sUUFBUSxNQUFNLFFBQVEsSUFBSSxDQUFDO0FBQ3ZDLFNBQU87QUFBQSxJQUNMLE1BQU07QUFBQSxJQUNOLFNBQVM7QUFBQSxNQUNQLElBQUk7QUFBQTtBQUFBLE1BRUoscUJBQXFCO0FBQUE7QUFBQSxRQUVuQixVQUFVLENBQUMsS0FBSyxRQUFRLFFBQVEsSUFBSSxHQUFHLGtCQUFrQixDQUFDO0FBQUE7QUFBQSxRQUUxRCxVQUFVO0FBQUEsTUFDWixDQUFDO0FBQUEsSUFDSDtBQUFBLElBQ0EsU0FBUztBQUFBLE1BQ1AsT0FBTztBQUFBLFFBQ0wsS0FBSyxjQUFjLElBQUksSUFBSSxTQUFTLHdDQUFlLENBQUM7QUFBQTtBQUFBLE1BRXREO0FBQUEsSUFDRjtBQUFBO0FBQUEsSUFFQSxRQUFRO0FBQUE7QUFBQSxNQUVOLE9BQU87QUFBQSxRQUNMLENBQUMsSUFBSSxpQkFBaUIsR0FBRztBQUFBO0FBQUEsVUFFdkIsUUFBUSxJQUFJO0FBQUEsVUFDWixjQUFjO0FBQUE7QUFBQTtBQUFBLFVBRWQsU0FBUyxDQUFDQSxVQUFTQSxNQUFLLFFBQVEsVUFBVSxFQUFFO0FBQUEsUUFDOUM7QUFBQSxNQUNGO0FBQUEsSUFDRjtBQUFBO0FBQUEsSUFFQSxLQUFLO0FBQUEsTUFDSCxxQkFBcUI7QUFBQSxRQUNuQixNQUFNO0FBQUEsVUFDSixtQkFBbUI7QUFBQSxVQUNuQixnQkFBZ0I7QUFBQSxRQUNsQjtBQUFBLE1BQ0Y7QUFBQSxJQUNGO0FBQUEsRUFDRjtBQUNGLENBQUM7IiwKICAibmFtZXMiOiBbInBhdGgiXQp9Cg==
