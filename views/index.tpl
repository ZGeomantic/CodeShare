<!DOCTYPE html>

<html>
<head>
	<title>Code Share</title>
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
            .mycon{
            	border-bottom-color: rgb(238, 238, 238);
				border-bottom-style: solid;
				border-bottom-width: 1px;
				box-sizing: border-box;
				color: rgb(51, 51, 51);
				display: block;
				font-family: Consolas, 'Liberation Mono', Courier, monospace;
				font-style: normal;
				font-variant: normal;
				font-weight: normal;
				line-height: 18.200000762939453px;
				margin-bottom: 20px;
				padding-bottom: 20px;
				padding-top: 10px;
				position: relative;
				width: 1287px;
            }
    </style>
     <link href="/static/css/prism.css" rel="stylesheet"/>
    <script src="/static/js/prism.js"></script>
<!-- 
	<script src="/static/js/rainbow-custom.min.js"></script> -->
	<script src="/static/js/inter-action.js"></script>
</head>

<body >
<h1 class="mycon">
	Code Share	
</h1>

{{if .Error}}
<h1 class="mycon">{{.Error}}</h1>
{{else}}
<div id="frame">

<pre>
<p style="font-size: 16px">File Path: {{.File}}</p>{{range $k,$v:=.Code}}<div  id="{{$k}}.line" href="#{{$k}}.line" class="my">{{$k}}</div>  <code  class="language-cpp">{{$v}}</code>{{end}}
</pre>
{{end}}
<h3 class="mycon">Contact me: {{.Email}}</h3>

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
</div>
</body>
</html>