### API 實作

```
1. 取得tasks
2. 新增task
3. 修改task
4. 刪除task
```

---

####

```
Spec:
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
    no need any params, show all of tasks.

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
    no need any params

Response:
StatusCode: 200
    no response
```

---
