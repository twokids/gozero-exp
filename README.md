# gozero-exp

# terminal命令
cd zeroService && go mod init zeroService
cd api && goctl api go -api user.api -dir .
//生成model文件, 目录在api/model里
goctl model mysql datasource -url="zmwb:realize2012@tcp(127.0.0.1:3306)/zero" -table="*"  -dir="./model"
cd rpc && goctl rpc proto -src userService.proto  -dir .

# etcd v3.5和最新版本rpc版本配合使用，不能用v.1.29.1去配合使用