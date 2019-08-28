require.config({
	waitSeconds :15,
	baseUrl : "/mp/assets/js/",
	paths: {
    'zepto' : MYENV+'/mylib/zepto',
		'util' : MYENV+'/mylib/util'
	}
});

require(
	['util', 'zepto'],
	function (util){
    var msg = util.getUrlParameter("msg");

    $('#errbox').text(msg);
    if (msg.indexOf("重新登录") == -1){
      $('#btnGoback').parent().removeClass('hide')
      $('#icon').addClass("weui-icon-warn weui-icon_msg")
      $('#txtTitle').text("出现了一个错误")

    }else{
      // 重新登录
      $('#icon').addClass("weui-icon-warn weui-icon_msg-primary")
      $('#txtTitle').text("登录验证失效")
      var corp = util.getUrlParameter("corp"),
        agent = util.getUrlParameter("agent");

      $('#btnGoMain').attr("href", mainURL+'/oauth/'+corp+'/'+agent).parent().removeClass('hide')
    }

    return {}
});
