<!DOCTYPE html>

<html>
<head>
	<title>Beego</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<!-- <link href="/static/css/github.css" rel="stylesheet" type="text/css"> -->
	<script src="/static/js/jquery-1.10.1.min.js"></script>
	 <style>
        #nav {z-index:999; width:50px; height: 20px; border: 1px solid #D4CD49; position:fixed;right:0%;top:30%;background: #000088;padding: 10px }
        .my {color: rgb(0, 119, 170);
            direction: ltr;
            display: inline;
            font-family: Consolas, Monaco, 'Andale Mono', monospace;
            height: auto;
            line-height: 24px;
            text-align: left;
            text-shadow: rgb(255, 255, 255) 0px 1px 0px;
            white-space: pre;
            width: auto;
            word-spacing: 0px;}
    </style>
     <link href="static/css/prism.css" rel="stylesheet"/>
    <script src="static/js/prism.js"></script>
<!-- 
	<script src="/static/js/rainbow-custom.min.js"></script> -->
	<script src="/static/js/inter-action.js"></script>
</head>

<body >
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
{{range $k,$v:=.Code}}<div  id="{{$k}}.line" href="#{{$k}}.line" class="my">{{$k}}</div>  <code  class="language-go">{{$v}}</code>{{end}}
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