 ## model_auto [中文版](readme_cn.md)
 ##### Secondary packaging based on GORM (https://github.com/jinzhu/gorm) library, Relatively simple operation on data table objects, including the commonly used Create, Get, Update operations (the CRUD operation of the database, generally delete uses soft delete.) 
 
 ## Install
    Open file Makefile, change GOPATH param in your local "model_auto" path. 
    
    Execute the command as follows:
    1、make vgo
    
    2、make vendor
    
    3、make build
    
    4、make install
    
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
    
#### model_auto command example:
```shell
./model_auto -import=github.com/model_auto/model -dbName=db_name -user=user_name -password=user_password -host=localhost -port=3306
```

#### [Application example](https://github.com/koalaone/model_auto_example)
	https://github.com/koalaone/model_auto_example