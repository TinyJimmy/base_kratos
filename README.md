# Kratos 项目模板

## 安装kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## 安装项目依赖
```
go mod tidy
```
## 修改对应配置
```
export APP_ENV=dev, test, prod
修改`configs/config_{$APP_ENV}.yaml`文件, 
```
## 运行
```
kratos run
```

