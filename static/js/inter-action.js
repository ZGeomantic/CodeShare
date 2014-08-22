$(document).ready(function () {
 		$("[id$='.line']").click(function () {
			alert( $(this).attr("id"))
			// TODO：将此line信息通过websocket传给另一个机器
	 	})
	 
	 // TODO：创建一个函数来接收websocket发来的消息，并将当前的页面跳转到这一行
});

