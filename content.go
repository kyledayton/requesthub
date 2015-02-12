package main

const(
PAGE_CONTENT = `
<html>
<head>
	<title>RequestHub</title>
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
<script>
jQuery(document).ready(function() {
jQuery.get("/requests", function(data) {
        jQuery("#requests").html(JSON.stringify(data));
});
});
</script>

</head>
<body>
<h1>RequestHub</h1>
<div id="requests"></div>
</body>
</html>
`
)
