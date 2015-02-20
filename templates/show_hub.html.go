package templates

const(
SHOW_HUB = `
<html>
<head>
	<title>RequestHub - {{.Id}}</title>
	<link rel="stylesheet" href="/assets/foundation.css"/>
  <script src="/assets/jquery.js"></script>
  <script src="/assets/foundation.js"></script>
  
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
  <script>
  (function($) {
    function updateForwardURL() {
  	$.post("/{{.Id}}/forward", {url: $("#forward_url").val()});
    }

    function fetchRequests() {
      $.get("/{{.Id}}/requests", function(data) {
        var requests = [];
        if( data.length == 0 ) {
          $("#default_content").show();
        } else {
          $("#default_content").hide();
        }

        Object.keys(data).map(function(request) {
      	var body = "";
      	try {
      		body = JSON.stringify(JSON.parse(data[request].body), null, 4);
      	} catch(ex) {
      		body = data[request].body;
      	}

          var headers = [];

          Object.keys(data[request].headers).map(function(h) {
            headers.push(h + ": " + data[request].headers[h].join(','));
          });

          var reqNum = +request + 1;
          requests.push(
            '<hr/><div class="row"><div class="large-1 columns"><h3>' + reqNum + '</h3>' + '</div><div class="large-11 columns">' +
              '<ul class="accordion" data-accordion="req' + reqNum + '">' +
                '<li class="accordion-navigation">' +
                  '<a href="#reqhead' + reqNum + '">Headers</a>' +
                  '<div id="reqhead' + reqNum + '" class="content">' +
                  '<div class="panel"><pre>' + 
                     headers.join('\n') +
                  '</pre></div></div></li>' +
                '<li class="accordion-navigation">' +
                  '<a href="#reqbody' + reqNum + '">Body</a>' +
                  '<div id="reqbody' + reqNum + '" class="content active">' +
                    '<div class="panel"><pre>' + body +
                  '</pre></div></div></li></ul></div></div>');
        });
        $("#requests").html(requests.join(''));
        $(document).foundation('accordion', 'reflow');
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

      $("#forward_form").on('submit', function(e) {
        updateForwardURL();
        e.preventDefault();
      });

      $("#refresh").click(function() {
        $("#requests").html("");
        fetchRequests();
      });

       $(document).foundation({
        accordion: {
          multi_expand: true,
          toggleable: true,
          content_class: 'content',
          active_class: 'active'
        }
      });

      fetchRequests();

    });
  })(jQuery);

  </script>

</head>
<body>
<nav class="top-bar" data-topbar role="navigation">
  <ul class="title-area">
    <li class="name">
      <h1><a href="/">RequestHub</a></h1>
    </li>
  </ul>

  <section class="top-bar-section">
    <ul class="right" style="padding-right: 2%;">
      <li><a id="clear" class="button" href="#">Clear Requests</a></li>
    </ul>
  </section>
</nav>

  <div id="content">
    <div class="row full-width">
      <div class="large-8 columns left">
        <h1><a href="#" id="refresh" style="color: black;">{{.Id}}</a></h1>
      </div>

      <div class="large-4 columns right">
        <form action="#" method="post" id="forward_form">
          <div class="row collapse" style="padding-top: 25px;">
            <div class="large-8 columns">
              <input type="text" name="url" id="forward_url" placeholder="Request Forwarding URL" value="{{.ForwardURL}}"/>
            </div>
            <div class="large-4 columns">
              <a href="#" id="update_url" class="button postfix">Update URL</a>
            </div>
          </div>
        </form>
      </div>
    </div>

    <div class="row full-width" id="requests">
    </div>

    <div id="default_content" class="row full-width hide" >
      <div class="large-12 columns" style="text-align: center; margin-top: 10%;">
        <h1>Your hub is empty!</h1>
        Send some requests to <code>/{{.Id}}</code> and they'll show up here.
      </div>
    </div>
  </div>
</body>
</html>
`)