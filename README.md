# Kite Layout

基于 [Kite](https://github.com/sllt/kite) 框架的企业级 Go 应用脚手架。

## 特性

- **分层架构** - Handler → Service → Repository → Model，职责清晰
- **依赖注入** - 使用 [Wire](https://github.com/google/wire) 自动生成依赖注入代码
- **数据库支持** - 基于 Kite 的 SQL 抽象，支持事务、连接池
- **JWT 认证** - 开箱即用的 JWT 中间件
- **定时任务** - 集成 gocron 定时任务调度
- **数据库迁移** - 内置迁移工具支持
- **可观测性** - 集成 OpenTelemetry 链路追踪和 Prometheus 指标

