package main

const(
PAGE_CONTENT = `
<html>
<head>
	<title>RequestHub</title>
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
<script>

function fetchRequests() {
	jQuery.get("/requests", function(data) {
		jQuery("#requests").html(JSON.stringify(data));
	});
}

jQuery(document).ready(function() {
	jQuery("#clear").click(function() {
		jQuery.get("/clear", function(d) {
			fetchRequests();
		});
	});

	fetchRequests();
});
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
