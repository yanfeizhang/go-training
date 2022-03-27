简单来说分为两步

### 解析模板
可以从字符串中解析，也可以从文件中解析，一般来说文件解析更为常用！

```golang
tmpl, err := template.ParseFiles("1-basic/template/hello.tpl")
```
### 渲染模板
解析完模板以后，就可以把变量传入模板中，开始渲染了。
```golang
sweaters := Service{"TestNamespace", "MyApp", "Func1"}
err = tmpl.Execute(os.Stdout, sweaters)
```

参考： https://www.cnblogs.com/wanghui-garcia/p/10385062.html
