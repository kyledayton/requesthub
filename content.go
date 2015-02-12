package main

const(
INDEX_PAGE_CONTENT = `
<html>
<head>
	<title>RequestHub</title>
</head>
<body>
<h1>RequestHub</h1>
<form>
	<input type="text" name="hub_name" />
	<input type="submit" value="Create Hub"/>
</form>

<h3>My Hubs:</h3>
<hr/>
</body>
</html>
`

HUB_PAGE_CONTENT = `
<html>
<head>
	<title>RequestHub</title>
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
<script>
(function($) {
  function fetchRequests() {
    $.get("/requests", function(data) {
      var requests = [];
      Object.keys(data.requests).map(function(request) {
        requests.push('<div class="request"><h1> request ' + request + '</h1>' +
          '<div class="headers"><h2>headers</h2><pre><code>' +
          JSON.stringify(data.requests[request].headers, null, 4) +
          '</pre></code></div>' +
          '<div class="body"><h2>body</h2><pre><code>' +
          JSON.stringify(JSON.parse(data.requests[request].body), null, 4) +
          '</pre></code></div></div>');
      });
      $("#requests").html(requests.join(''));
    });
  }

  $(document).ready(function() {
    $("#clear").click(function() {
      $.get("/clear", function() {
        $("#requests").empty();
        fetchRequests();
      });
    });
    fetchRequests();
  });
})(jQuery);

</script>

</head>
<body>
<h1>RequestHub</h1>
<a id="clear" href="#">Clear</a>

<div id="requests"></div>
</body>
</html>
`
)
