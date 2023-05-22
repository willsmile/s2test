# s2test: A Simple Smoke Test Tool

## Intentions
- To make smoke tests of API easier and more enjoyable for server-side developers

## Features
- Planing smoke tests by using a JSON style config file
  - Various cookies for test can be prepared
- Reusable test target API information store

## Installation
### Preparation
Please download and install [golang](https://golang.org/dl/)

### Clone from Github
```
$ git clone https://github.com/willsmile/s2test
$ cd s2test
```

### Get the go dependencies (by go modules):
```
$ GO111MODULE=on go get -v -d
```

### Build the tool
```
$ go build -o s2test .
$ ./s2test help
```

## Setup and configuration
### Test plan preparation
To prepare a test plan by using the following format (for example, let's name the file 'plan.json').

```json
{
  "goal": "Have a try on s2test",
  "targetPath": "./api.json",
  "authMethods": {
    "cookieA": {
      "type": "Cookie",
      "name": "cookieName",
      "value": "cookieValue"
    },
    "tokenA": {
      "type": "OAuth 2.0",
      "prefix": "Bearer",
      "value": "tokenValue"
    }
  },
  "tasks": [
    {
      "targetAPI": "GET a sample post",
      "authMethod": "tokenA"
    },
    {
      "targetAPI": "GET a sample todo",
      "authMethod": "tokenA"
    }
  ]
}
```

Also, to write down the information of test target API by using the following format.

```json
{
  "GET a sample post": {
    "description": "GET a sample post",
    "url": "https://jsonplaceholder.typicode.com/posts/1",
    "method": "GET",
    "headers": {
      "Content-type": "application/json; charset=utf-8"
    }
  },
  "GET a sample todo": {
    "description": "GET a sample todo",
    "url": "https://jsonplaceholder.typicode.com/todos/1/",
    "method": "GET",
    "headers": {
      "Content-type": "application/json; charset=utf-8"
    }
  }
}
```

### Test Plan Execution

```
./s2test -p plan.json
```
