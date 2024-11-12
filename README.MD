# 项目说明

## 初始化
```bash
make init
```

## 项目结构
### 全局数据
- 配置文件: `global/global.go`
  - GlobalConfig: 全局配置
  - DB: 数据库连接

## 数据库
### 创建数据模型
```bash
# 创建 User 模型
ent new User

# 生成数据库代码
make gen
```

## 运行项目
```bash
make run
```
