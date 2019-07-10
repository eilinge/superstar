# superstar
使用Go语言的iris+xorm框架，实现的球星库
<br/>

## golang下载 1.9.7
<br/>
https://golang.org/dl/
<br/>
https://studygolang.com/dl
<br/>

## go依赖库
<br/>
github.com/kataras/iris
<br/>
github.com/go-xorm/xorm
<br/>
github.com/go-xorm/cmd
<br/>
github.com/go-sql-driver/mysql
<br/>
github.com/gorilla/securecookie
<br/>
github.com/gorilla/websocket
<br/>

## IDE
<br/>
GoLand https://www.jetbrains.com/go/download/
<br/>

## Mysql 5.7.18
<br/>
https://dev.mysql.com/downloads/mysql/5.7.html
<br/>

## 其它参考
<br/>
http://2018.sina.com.cn/ballgamestar/

## 添加其他功能
1. 球员分页(上一页,下一页)
2. 添加功能

## 性能优化和性能对比
1. 模板是否支持每次更新检查:bootstarp.go/htmlEngine.Reload(false)
2. 开启xorm数据缓存:dbhelper.go/xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
3. 不用模板, 不查数据