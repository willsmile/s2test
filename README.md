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
  "endpoints": "./api.json",
  "ua": "tester",
  "auths": {
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
      "auth": ["tokenA"]
    },
    {
      "targetAPI": "POST a sample post",
      "auth": ["tokenA"],
      "variables": {
        "sample_use_id": "1",
        "sample_post_title": "sample post title"
      }
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
  "POST a sample post": {
    "description": "POST a sample post",
    "url": "https://jsonplaceholder.typicode.com/posts",
    "method": "POST",
    "headers": {
      "Content-type": "application/json; charset=utf-8"
    },
    "body": {
      "title": "#{sample_post_title}",
      "body": "sample post body",
      "userId": "#{sample_use_id}"
    }
  }
}
```
### Test Plan Execution
- Run with both `plan path` and `api path`
```
./s2test -p plan.json -a api.json
```

- Run with only `plan path` (`api path` will be set with the value of `endpoints` in test plan file)
```
./s2test -p plan.json
```
