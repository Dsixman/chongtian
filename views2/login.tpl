<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Document</title>
</head>
<body>
  <div>
    <form method="post" action="/">
    	<input type="hidden" name="{{.key}}">
　　　　<input type="text" name="captcha" />
{{.base64str}}
        <img src="{{.base64str}}"><br />
　　　　<input type="submit" value="提交" />
　　</form>
  </div>
</body>
</html>