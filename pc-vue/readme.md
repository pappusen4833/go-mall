[TOC]

# yshop-pc

## 安装依赖

```
npm install
```

如果遇到安装失败查看报错信息，是否为node-sass插件问题，考虑使用`代理`或使用`yarn`来安装依赖

## 启动项目

```
npm run dev
```

- 测试环境如果需要修改服务相关配置，打开config\index.js
  - 修改：`host`、`port`属性
  
  - 修改请求地址`proxyTable`下`target`



## 更改请求地址


1. 封装文件：`src\api\index.js` 	修改`requestConfig.url`



## 画布请求配置

画布请求配置地址：

```
文件路径：src\views\PageIndex\indexAll\Index.vue
266行  var res = await getCanvasData({ terminal: '4' })
```

修改getCanvasData方法里面传入的对象来修改对应画布请求传参，terminal需与画布项目修改的terminal一致



## 打包项目

```
npm run build
```

打包后的内容存放在`dist`文件夹内

