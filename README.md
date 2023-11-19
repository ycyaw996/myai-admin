# HG-Dashboard
黑果仪表盘，轻量简介实用的监控面板，支持多节点监控，多网站状态监控，集成一键测速等实用功能，包括强大的警报告警功能，正在开发中...
## 编译

### Linux
```
SET GOARCH=amd64 && SET GOOS=linux && go build
```

### Windows
```
$env:GOOS="linux"; $env:GOARCH="amd64"; go build
```
## 代码目录说明
```
### 模型 (Models)
- 存放数据库模型和数据访问逻辑。
- 通常包含对数据库的 CRUD（创建、读取、更新、删除）操作。

### 控制器 (Controllers) / 处理器 (Handlers)
- 处理 HTTP 请求和响应。
- 从请求中提取数据，调用模型的方法，并返回响应。

### 路由 (Routes)
- 定义 API 路由和将路由映射到控制器/处理器的函数。

### 服务 (Services)
- 包含业务逻辑。
- 可以被多个控制器调用，以避免代码重复。

### 中间件 (Middlewares)
- 用于处理跨越多个路由的通用逻辑，如身份验证、日志记录等。

### 配置 (Configurations)
- 存放配置文件和环境变量。

### 工具 (Utils) / 库 (Libraries)
- 包含项目中使用的通用函数和辅助工具。

### 测试 (Tests)
- 包含单元测试和集成测试。
```
接口列表：