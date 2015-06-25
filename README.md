# RequestHub
Receive HTTP requests, display them in your browser, and forward them to other URLs.

RequestHub is an open source project inspired by [RequestBin](http://requestb.in)

![RequestHub](http://i.imgur.com/pSflmfL.png)

## Overview
I developed this for our organization to maximize our limited pool of public IPs. We can map all of our external service webhooks to one IP, and forward them to numerous internal testing servers. I thought others would have a use for something like this, so I decided to release it as open source software.

## Installation
###### Install
```bash
$ go get github.com/kyledayton/requesthub/...
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
  -config="": YAML Configuration File
  -p=54321: which port to bind to
  -r=256: max requests to store
  -username="": HTTP Basic Auth Username for accessing hub
  -password="": HTTP Basic Auth Password for accessing hub
```

Note: To Enable Basic Auth, you must specify both username and password.

## Usage
Open `http://localhost:54321` in your browser. The index page shows a list of your hubs, and a form for creating a hub. Create a hub and it will redirect you to the hub requests page.

To send requests to the hub, send any HTTP request to `http://localhost:54321/<HUB_NAME>`

The hub requests page shows stored requests sent to the hub. There is a clear button, which will delete all stored requests in the hub. In addition, there is a form for setting the forwarding URL of the hub. Setting a URL and clicking 'Update URL' will forward any incoming requests to the hub into the specified URL.

## Configuration
RequestHub can create default hubs on startup. Simply create a YAML file with the appropriate hub names and forwarding urls, and pass it to the config option.

**config.yml:**
```yaml
hubs:
  test-hub:
    forward_url: 'https://www.example.com/webhook'
  another-hub:
```

```bash
$ requesthub -config config.yml
```
