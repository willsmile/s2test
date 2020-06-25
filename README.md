# s2test: A Simple Smoke Test Tool

## Intentions
- To make smoke tests of API easier and more enjoyable for server-side developers

## Features
- Planing smoke tests by using a JSON style config file
  - Various cookies for test can be prepared
- Reusable test target API information store

## Usage
### Preparation
#### Install Golang
Please download and install [golang](https://golang.org/dl/)

### Installation
#### Compiling from source
##### Clone from Github:
```
$ git clone https://github.com/willsmile/s2test
$ cd s2test
```

##### Get the go dependencies (by go modules):
```
$ go mod tidy
```

#### Build the tool
```
$ go build -o s2test .
$ ./s2test help
```

### Test plan preparation
To prepare a test plan by using the following format (for example, let's name the file 'plan.json').

```json
{
  "goal": "Have a try on s2test",
  "targetPath": "./api.json",
  "preparedCookies": {
    "myCookie": {
      "cookieName": "cookieValue"
    }
  },
  "tasks": [
    {
      "targetAPI": "GET date",
      "usedCookies": ""
    },
    {
      "targetAPI": "GET key&value",
      "usedCookies": ""
    }
  ]
}
```

Also, to write down the information of test target API by using the following format.

```json
{
  "GET date": {
    "description": "Get current date from jsontest.com",
    "url": "http://date.jsontest.com",
    "method": "GET",
    "headers": {
      "Content-type": "application/json; charset=utf-8"
    }
  },
  "GET key&value": {
    "description": "Get key and value from jsontest.com",
    "url": "http://echo.jsontest.com/tool/s2test/",
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