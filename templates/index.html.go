package templates

const(
INDEX = `
<html>
<head>
	<title>RequestHub</title>
	<link rel="stylesheet" href="/assets/foundation.css"/>
	<script src="/assets/jquery.js"></script>
	<script src="/assets/foundation.js"></script>
	<script src="/assets/modernizr.js"></script>

  <style>
    #content {
      width: 90%;
      margin: auto;
      margin-top: 2%;
    }

    .full-width {
       width: 100%;
       margin-left: auto;
       margin-right: auto;
       max-width: initial;
    }
  </style>
</head>
<body>

<nav class="top-bar" data-topbar role="navigation">
  <ul class="title-area">
    <li class="name">
      <h1><a href="/">RequestHub</a></h1>
    </li>
  </ul>
  </section>
</nav>

  <div id="content">

  <div class="row full-width">
    <div class="large-8 columns left">
      <h1>Hubs</h1>
    </div>

    <div class="large-4 columns right">

    <form method="POST">
        <div class="row collapse" style="padding-top: 25px;">
          <div class="large-8 columns">
              <input type="text" name="hub_name" placeholder="New Hub Name"/>
          </div>
          <div class="large-4 columns">
            <a href="#" id="submit" onclick="document.forms[0].submit();" class="button postfix">Create</a>
          </div>
        </div>
    </form>

  </div>

  <hr/>
  <ul>
  {{range .}}
    <li><a href="/show/{{.Id}}">{{.Id}}</a> (<a href="/{{.Id}}/delete">delete</a>)</li>
  {{end}}
  </ul>


</div>

</body>
</html>
`)
