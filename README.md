# Kite Layout

基于 [Kite](https://github.com/sllt/kite) 框架的企业级 Go 应用脚手架（项目骨架）。

## 特性

- **分层架构** - Handler → Service → Repository → Model
- **依赖注入** - 使用 [Wire](https://github.com/google/wire) 管理依赖
- **数据库支持** - 基于 Kite 的 SQL 抽象，支持事务
- **JWT 认证** - 全局解析令牌 + 受保护路由强制鉴权
- **定时任务** - 集成 gocron 调度
- **数据库迁移** - 内置迁移入口
- **可观测性** - 集成 OpenTelemetry 与 Prometheus

## 目录概览

- `cmd/server`：HTTP + gRPC 服务入口
- `cmd/migration`：数据库迁移入口
- `cmd/task`：定时任务入口
- `internal/handler`：接口层（协议适配）
- `internal/service`：业务逻辑层
- `internal/repository`：数据访问层
- `internal/router`：路由与鉴权包装
- `internal/server`：各类 server 实现
- `pkg/app`：通用应用生命周期管理
- `pkg/config`：环境配置约定

## 快速开始

### 1) 准备配置

```bash
cp configs/.env.example configs/.env
```

然后修改 `configs/.env` 中的关键项（至少包括 `JWT_SECRET`）。

> 注意：`configs/.env` 已被 `.gitignore` 忽略，不应提交真实密钥。

### 2) 初始化依赖工具（可选）

```bash
make init
```

### 3) 执行数据库迁移

```bash
go run ./cmd/migration
```

如需新增迁移模板：

```bash
kite migrate create add_users_index
```

### 4) 启动服务

```bash
go run ./cmd/server
# 或使用 nunu
nunu run ./cmd/server
```

### 5) 启动任务调度

```bash
go run ./cmd/task
```

## 常用命令

```bash
make test     # 运行测试并生成覆盖率报告
make build    # 构建 server 二进制
make swag     # 生成 swagger 文档
make bootstrap
```

`make bootstrap` 会启动 docker-compose、执行迁移并运行服务。

## 配置约定

- 运行时通过 `pkg/config.SetupKiteEnv` 设置 `CONFIGS_DIR=./configs`
- Kite 从 `./configs/.env` 读取配置
- **密钥管理建议**：
  - 本地开发：放 `configs/.env`（不提交）
  - CI/CD：放平台 Secret
  - 生产环境：使用环境变量或密钥管理系统注入

## 鉴权策略

- 全局中间件使用 `NoStrictAuth`：有 token 就解析 claims，无 token 不拦截。
- 需要登录的路由在 `internal/router` 使用 `RouteGroup.UseMiddleware(...)` 做分组强制鉴权。

## 测试

```bash
make test
```

接口冒烟测试（会自动执行迁移、启动服务并验证核心用户接口）：

```bash
bash scripts/smoke.sh
```

当前包含 service/repository 测试；handler 测试依赖 Kite 的集成测试上下文。
