# Go-Mall

## 📖 项目简介

Go-Mall 是一个基于 [yshop-gin](https://github.com/guchengwuyue/yshop-gin) 进行二次开发的现代化电商系统，采用前后端分离架构，提供完整的电商业务功能。

### ✨ 主要特性

- 🚀 **高性能架构** - 基于 Gin + Gorm + Redis，支持高并发访问
- 💎 **完整功能** - 商品管理、订单系统、支付集成、物流跟踪
- 🔐 **安全可靠** - JWT 认证、RBAC 权限控制、数据加密
- 📱 **多端支持** - 支持 Web、小程序、APP 多端接入
- 🎨 **易于扩展** - 模块化设计，便于二次开发
- 📊 **数据分析** - 内置数据统计和可视化分析

### 🎯 适用场景

- 中小型电商平台
- 企业内部商城系统
- 学习 Go 语言电商项目实战
- 毕业设计或技术研究

---

## 🛠️ 技术栈

### 后端技术

| 技术 | 说明 | 版本 |
|------|------|------|
| Go | 编程语言 | 1.18+ |
| Gin | Web 框架 | Latest |
| Gorm | ORM 框架 | Latest |
| MySQL | 关系型数据库 | 5.7+ / 8.0+ |
| Redis | 缓存数据库 | 5.0+ |
| JWT | 认证授权 | Latest |
| Casbin | 权限管理 | Latest |

### 前端技术（如适用）

- Vue 3.x
- Element Plus
- Vite

---

## 🚀 快速开始

### 环境要求

- Go 1.18+
- MySQL 5.7+ / 8.0+
- Redis 5.0+
- Git

### 克隆项目

```bash
git clone https://github.com/yourusername/go-mall.git
cd go-mall
```

### 配置数据库

1. 创建数据库

```sql
CREATE DATABASE go_mall DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 导入初始数据

```bash
mysql -u root -p go_mall < sql/go_mall.sql
```

### 修改配置

复制配置文件并修改：

```bash
cp config.yaml.example config.yaml
```

编辑 `config.yaml` 文件，修改数据库和 Redis 配置：

```yaml
mysql:
  host: 127.0.0.1
  port: 3306
  database: go_mall
  username: root
  password: your_password

redis:
  host: 127.0.0.1
  port: 6379
  password: ""
  db: 0
```

### 安装依赖

```bash
go mod download
```

### 运行项目

```bash
# 开发模式
go run main.go

# 或者编译后运行
go build -o go-mall
./go-mall
```

访问 `http://localhost:8080`

### Docker 部署（推荐）

```bash
# 使用 docker-compose 一键启动
docker-compose up -d
```

---

## 📁 项目结构

```
go-mall/
├── app/                # 应用核心代码
│   ├── controllers/    # 控制器
│   ├── models/         # 数据模型
│   ├── services/       # 业务逻辑
│   └── middleware/     # 中间件
├── cmd/                # 命令行工具
├── config/             # 配置文件
├── docs/               # 文档
├── pkg/                # 公共包
├── routers/            # 路由定义
├── sql/                # SQL 脚本
├── storage/            # 文件存储
├── main.go             # 入口文件
└── go.mod              # 依赖管理
```

---

## 🎯 核心功能

### 商品管理
- ✅ 商品分类管理
- ✅ 商品信息管理（SPU/SKU）
- ✅ 商品库存管理
- ✅ 商品上下架

### 订单系统
- ✅ 购物车管理
- ✅ 订单创建与支付
- ✅ 订单状态跟踪
- ✅ 退款/售后处理

### 用户系统
- ✅ 用户注册/登录
- ✅ 个人信息管理
- ✅ 收货地址管理
- ✅ 会员等级体系

### 营销功能
- ✅ 优惠券系统
- ✅ 秒杀活动
- ✅ 满减促销
- ✅ 积分系统

### 后台管理
- ✅ 角色权限管理
- ✅ 数据统计分析
- ✅ 系统配置管理
- ✅ 操作日志记录

---

## 🔄 相比原项目的改进

基于 yshop-gin 的基础上，本项目进行了以下改进：

- 🎨 **界面优化** - 重新设计了管理后台界面
- ⚡ **性能提升** - 优化了数据库查询和缓存策略
- 🔧 **功能增强** - 新增了 XXX 功能模块
- 📝 **代码规范** - 改进了代码结构和注释
- 🐛 **Bug 修复** - 修复了原项目的已知问题
- 📚 **文档完善** - 提供了更详细的开发文档

---

## 📸 项目截图

### 前台商城
![商城首页](docs/images/home.png)
![商品详情](docs/images/product.png)

### 后台管理
![管理后台](docs/images/admin.png)
![数据统计](docs/images/dashboard.png)

---

## 🤝 参与贡献

欢迎提交 Issue 和 Pull Request！

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

---

## 📄 开源协议

本项目基于 [yshop-gin](https://github.com/guchengwuyue/yshop-gin) 进行二次开发。

- 原项目采用 Apache-2.0 License
- 本项目同样采用 Apache-2.0 License

详见 [LICENSE](LICENSE) 文件。

---

## 🙏 致谢

- 感谢 [yshop-gin](https://github.com/guchengwuyue/yshop-gin) 项目提供的优秀基础框架
- 感谢所有为本项目做出贡献的开发者

---

## 📮 联系方式

- 项目主页: https://github.com/yourusername/go-mall
- 问题反馈: https://github.com/yourusername/go-mall/issues
- 电子邮件: your.email@example.com

<!-- Updated on 2024-04-05 16:09:00 -->

<!-- Updated on 2024-09-02 06:14:00 -->

<!-- Updated on 2024-03-15 22:01:00 -->

<!-- Updated on 2024-04-05 01:34:00 -->

<!-- Updated on 2024-06-07 05:52:00 -->

<!-- Updated on 2024-11-09 18:03:00 -->

<!-- Updated on 2024-08-21 07:38:00 -->

<!-- Updated on 2024-07-02 00:13:00 -->

<!-- Updated on 2024-10-24 20:14:00 -->

<!-- Updated on 2024-02-01 02:32:00 -->

<!-- Updated on 2024-03-19 19:28:00 -->

<!-- Updated on 2024-04-03 19:36:00 -->

<!-- Updated on 2024-02-15 07:45:00 -->

<!-- Updated on 2024-12-15 17:28:00 -->

<!-- Updated on 2024-01-21 20:36:00 -->

<!-- Updated on 2024-01-07 21:06:00 -->
