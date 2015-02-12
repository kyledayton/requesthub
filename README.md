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

## Endpoints
### Receive Request
Any `non-GET` request made to "/" will be stored in memory. By default, the newest 256 requests are stored. This can be changed by using the `-r` command line option.

### View Requests
A `GET` request to "/" shows the stored requests, sorted from newest to oldest.

### Clear Requests
Sending any request to "/clear" will remove all stored requests.

## TODO
* Add a web UI for viewing requests
