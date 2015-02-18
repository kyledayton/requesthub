# RequestHub
Receive HTTP requests and display them in your browser.

RequestHub is an open source project inspired by [RequestBin](http://requestb.in)

## Installation
###### Install
```bash
$ go get github.com/kyledayton/requesthub
```

###### Run
```bash
$ export PATH=$PATH:$GOPATH/bin
$ requesthub
```

This will start the server on port 54321.  
There are also a few command line options available:
```bash
$ requesthub -h
Usage of requesthub:
  -p=54321: which port to bind to
  -r=256: max requests to store
```

## Usage
Open `http://localhost:54321` in your browser. The index page shows a list of your hubs, and a form for creating a hub. Create a hub and it will redirect you to the hub requests page.

To send requests to the hub, send any non-GET request to `http://localhost:54321/<HUB_NAME>`

The hub requests page shows stored requests sent to the hub. There is a clear button, which will delete all stored requests in the hub. In addition, there is a form for setting the forwarding URL of the hub. Setting a URL and clicking 'Update URL' will forward any incoming requests to the hub into the specified URL.

## Todo
* Improve the UI
