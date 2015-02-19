package main

const(
INDEX_PAGE_CONTENT = `
<html>
<head>
	<title>RequestHub</title>
</head>
<body>
<h1>RequestHub</h1>
<form method="POST">
	<input type="text" name="hub_name" />
	<input type="submit" value="Create Hub"/>
</form>

<h3>My Hubs:</h3>
<hr/>
<ul>
{{range .}}
	<li><a href="/{{.Id}}">{{.Id}}</a> (<a href="/{{.Id}}/delete">delete</a>)</li>
{{end}}
</body>
</html>
`)