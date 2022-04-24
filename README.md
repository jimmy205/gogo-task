### API 實作

```
default:
    port: 8000
    db  : in-memory
```

```
command:
    run server
        - go run command/main.go
    run task test
        - go test -v usecase/task/*
    run docker
        - docker build -t gogo-task
        - docker run -p 8000:8000 gogo-task
```

概覽

```
可以使用heroku串接api
( 初次呼叫時需等待約10~20秒鐘待機器從待機轉為上線 )

heroku domain:
    https://gogo-task.herokuapp.com/

1. 取得tasks
2. 新增task
3. 修改task
4. 刪除task
```

---

####

```
API Spec:
    {
        name: string
        status: boolean
            - 0 = incomplete
            - 1 = complete
    }
```

#### 取得 tasks

```=
Path: /tasks
Method: GET

Request:
    不需輸入參數，回傳全部。

Response:
    {
        "result":[
            {"id":1,"name":"buy dinner","status":0}
        ]
    }
```

#### 新增 task

```=
Path: /task
Method: POST

Request:
    {
        "name":"buy dinner"
    }

Response:
StatusCode: 201
    {
        "result":[
            {"id":1,"name":"buy dinner","status":0}
        ]
    }
```

#### 修改 task

```=
Path: /task/id
Method: PUT

Request:
    {
        "name":"buy dinner",
        "status": 1,
        "id": 1
    }

Response:
StatusCode: 200
    {
        "result":[
            {"id":1,"name":"buy dinner","status":1}
        ]
    }
```

#### 刪除

```=
Path: /task/id
Method: DELETE

Request:
    不需參數。

Response:
StatusCode: 200
    no response
```

---
