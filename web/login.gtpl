<html>
<head>
<title></title>
</head>
<body>
<form action="http://127.0.0.1:9090/login?username=astaxie" method="post">
用户名 :<input type="text" name="username">
密码 :<input type="password" name="password">
<input type="submit" value=" 登陆 ">
<input type="hidden" name="token" value="{{.}}">
</form>
</body>
</html>