$(document).ready(function () {
	 // Create a socket
    socket = new WebSocket('ws://' + window.location.host + '/ws' );

 	socket.onmessage = function (event) {
 		// TODO: 从event.data中获取数据（即锚点信息#15.line），并刷新页面到锚点(#15.line)
 		var data = event.data
 	 	// window.location.href=direct;
 		// alert(direct);

 	 	window.location.href="/#"+data
 	 	// window.location.href="http:///www.baidu.com"


 	  // location.reload(); 


 	 	// $.get("#frame", function(data){
    //     //data就是b.html的内容，想加在哪就加在哪
    //     $("body").append(data);
// })

	}

	$("[id$='.line']").click(function () {
		// alert(window.location.host)
	var content =  $(this).attr("id")
	// alert( content)
	socket.send(content)
	// TODO：将此line信息通过websocket传给另一个机器
	})
	 
	 // TODO：创建一个函数来接收websocket发来的消息，并将当前的页面跳转到这一行
});

