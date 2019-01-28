 ##model_auto [English](readme.md)
 #####基于GORM（https://github.com/jinzhu/gorm）库的二次封装，对数据表对象操作相对简单，包括常用的Create，Get，Update操作（数据库的CRUD操作，一般删除使用软删除）。

 ##安装步骤
    打开 Makefile 文件, 修改 GOPATH 参数为model_auto包的本地当前目录. 
    
    执行目录顺序如下:
    1、make vgo
    
    2、make install
    
    3、make vendor
    
    4、make build
    
    5、./model_auto -help
    
```shell
NAME:
   model-auto - provide model-auto tools for koalaone@163.com

USAGE:
   model_auto [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --import value     model package import name [$APP_CONFIG]
   --dbName value     db name [$APP_CONFIG]
   --user value       db user name [$APP_CONFIG]
   --password value   db user password [$APP_CONFIG]
   --host value       db server host [$APP_CONFIG]
   --port value       db server port [$APP_CONFIG]
   --tableName value  db server port [$APP_CONFIG]
   --help, -h         show help
   --version, -v      print the version
```
    
#### model_auto 命令示例:
```shell
./model_auto -import=github.com/model_auto/model -dbName=db_name -user=user_name -password=user_password -host=localhost -port=3306
```
