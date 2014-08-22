<!DOCTYPE html>

<html>
<head>
	<title>Beego</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<link href="/static/css/github.css" rel="stylesheet" type="text/css">
	<script src="/static/js/jquery-1.10.1.min.js"></script>

	<script src="/static/js/rainbow-custom.min.js"></script>
	<script src="/static/js/inter-action.js"></script>

</head>

<body>
	<!-- 
目前，如果有<code data-language="C">"int a"</code>
，就无法用jquery获取节点的id并设置click动作。
<table>
	<tr>
		<td >
			<li id="no1">
				<a href="#line5">连接1</a>
			</li>
		</td>
		<tr>
			<code data-language="C">"int a"</code>
		</tr>
	</tr>

</table>
-->
<pre>
<table>
	{{range $k,$v:=.Code}}
	<tr>
		<td><li id="{{$k}}.line"}><a href="#{{$k}}.line">{{$k}}</a></li></td>

	<!-- <td><code data-language="C">{{$v}}</code></td> -->
		<td>{{$v}}</td>
	</tr>
	{{end}}
</table>
</pre>

<script type="text/javascript">
	$(function () {
		// $(".test ul li").each(function () {
		// 	$(this).click(function(){
		// 		alert($(this).text());
		// 	});
		// });
	$("#no1").click(function(){
 			alert("okdd")

		alert($(this).text())
	});
	$("#line1").click(function () {
 			alert("okdd")
			alert( $(this).attr("id"))
 	})

	});
</script>

</body>
</html>