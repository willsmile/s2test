# s2test: A Simple Smoke Test Tool

## Intentions
- To make smoke tests of API easier and more enjoyable for server-side developers

## Features
- Planing smoke tests by using a JSON style config file
  - Various cookies for test can be prepared
- Reusable test target API information store

## Usage
### Installation

```
go get github.com/willsmile/s2test
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
      "targetAPI": "/app/settings",
      "usedCookies": ""
    },
    {
      "targetAPI": "/app/version",
      "usedCookies": ""
    },
    {
      "targetAPI": "/error",
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
s2test -p plan.json
```