package main

const(
HUB_PAGE_CONTENT = `
<html>
<head>
	<title>RequestHub - {{.Id}}</title>
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
<script>
(function($) {
  function updateForwardURL() {
	$.post("/{{.Id}}/forward", {url: $("#forward_url").val()});
  }

  function fetchRequests() {
    $.get("/{{.Id}}/requests", function(data) {
      var requests = [];
      Object.keys(data).map(function(request) {
	var body = "";
	try {
		body = JSON.stringify(JSON.parse(data[request].body), null, 4);
	} catch(ex) {
		body = data[request].body;
	}

        requests.push('<div class="request"><h1> request ' + request + '</h1>' +
          '<div class="headers"><h2>headers</h2><pre><code>' +
          JSON.stringify(data[request].headers, null, 4) +
          '</pre></code></div>' +
          '<div class="body"><h2>body</h2><pre><code>' + body +
          '</pre></code></div></div>');
      });
      $("#requests").html(requests.join(''));
    });
  }

  $(document).ready(function() {
    $("#clear").click(function() {
      $.get("/{{.Id}}/clear", function() {
        $("#requests").empty();
        fetchRequests();
      });
    });

    $("#update_url").click(function() {
	updateForwardURL();
    });
    fetchRequests();
  });
})(jQuery);

</script>

</head>
<body>
<h1>RequestHub</h1>
<span><h3>{{.Id}}</h3> <a id="clear" href="#">Clear</a></span>

<input id="forward_url" name="url" placeholder="Request Forwarding URL" value="{{.ForwardURL}}"/><a href="#" id="update_url">Update URL</a>

<div id="requests"></div>
</body>
</html>
`
)