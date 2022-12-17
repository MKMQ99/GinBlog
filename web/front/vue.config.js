const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: [
    'vuetify'
  ],
  lintOnSave: false,
  assetsDir: 'static',
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = '欢迎来到GinBlog'
      return args
    })
  },
  productionSourceMap: false,
  devServer:{
    host:'42.192.21.110',
  }
})
