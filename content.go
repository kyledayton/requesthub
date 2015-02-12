package main

const(
PAGE_CONTENT = `
<html>
<head>
	<title>RequestHub</title>
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
<script>
(function($) {
  $(document).ready(function() {
  $.get("/requests", function(data) {
    var requests = [];
    Object.keys(data.requests).map(function(request) {
      requests.push('<div class="request"><h1> request ' + request + '</h1>' +
        '<div class="headers"><h2>headers</h2><pre><code>' +
        JSON.stringify(data.requests[request].headers, null, 4) +
        '</pre></code></div>' +
        '<div class="body"><h2>body</h2><pre><code>' +
        JSON.stringify(data.requests[request].body, null, 4) +
        '</pre></code></div></div>');
    });
    $("#requests").html(requests.join(''));
  });
  });
})(jQuery);
</script>

</head>
<body>
<h1>RequestHub</h1>
<div id="requests"></div>
</body>
</html>
`
)
