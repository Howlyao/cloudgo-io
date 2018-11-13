# cloudgo-io

## 处理 web 程序的输入与输出


### 1、概述
设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理。

### 2、内容

* 支持静态文件服务
* 支持简单 js 访问
* 提交表单，并输出一个表格
* 对 /unknown 给出开发中的提示，返回码 5xx

### 3、展示

####  在浏览器输入`http://localhost:8080/static/`,对静态网页的访问，利用js实现对服务器发送GET请求，获取数据将表单初始化

![](https://github.com/Howlyao/cloudgo-io/raw/master/images/1.jpg)

####  输入`http://localhost:8080/register`,向服务器发送GET请求，返回register.html页面信息

![](https://github.com/Howlyao/cloudgo-io/raw/master/images/2.jpg)

![](https://github.com/Howlyao/cloudgo-io/raw/master/images/3.jpg)

####  在register页面表单输入相关信息，点击register按钮,提交表单，对服务器发送POST请求，返回inf.html页面信息，其中inf.html显示一个记录输入信息的表格（模板输出，
{{.Username}}获取服务器后台信息）

![](https://github.com/Howlyao/cloudgo-io/raw/master/images/4.jpg)


####  输入`http://localhost:8080/unknown`，向服务器发送GET请求，返回501信息

![](https://github.com/Howlyao/cloudgo-io/raw/master/images/5.jpg)
