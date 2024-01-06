const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    host: '0.0.0.0',
    port: 8081,
    headers: {
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, PATCH, OPTIONS",
      "Access-Control-Allow-Headers": "X-Requested-With, content-type, Authorization, Accept, Content-Type"
    },
    allowedHosts: ['front-area.lekmax.fr', 'localhost', 'localhost:8080', 'back-area.lekmax.fr'],
    proxy: {
      '/api': {
        target: 'https://back-area.lekmax.fr',
        changeOrigin: true,
        pathRewrite: {
          '^/api': '',
        },
      },
    },
  }
})  
