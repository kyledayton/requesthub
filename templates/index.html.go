package templates

const(
INDEX = `
<html>
<head>
	<title>RequestHub</title>
	<link rel="stylesheet" href="http://cdn.foundation5.zurb.com/foundation.css"/>
	<script src="http://cdn.foundation5.zurb.com/foundation.js"></script>
</head>
<body>
<h1>RequestHub</h1>
<nav class="top-bar" data-topbar role="navigation">
  <ul class="title-area">
    <li class="name">
      <h1><a href="#">My Site</a></h1>
    </li>
     <!-- Remove the class "menu-icon" to get rid of menu icon. Take out "Menu" to just have icon alone -->
    <li class="toggle-topbar menu-icon"><a href="#"><span>Menu</span></a></li>
  </ul>

  <section class="top-bar-section">
    <!-- Right Nav Section -->
    <ul class="right">
      <li class="active"><a href="#">Right Button Active</a></li>
      <li class="has-dropdown">
        <a href="#">Right Button Dropdown</a>
        <ul class="dropdown">
          <li><a href="#">First link in dropdown</a></li>
          <li class="active"><a href="#">Active link in dropdown</a></li>
        </ul>
      </li>
    </ul>

    <!-- Left Nav Section -->
    <ul class="left">
      <li><a href="#">Left Nav Button</a></li>
    </ul>
  </section>
</nav>

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