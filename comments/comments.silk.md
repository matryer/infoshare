# Comments API

## POST /comments?url=http://infoshare.pl/

```
{
  "body": "I love this conference",
  "author": "John"
}
```

---

### Response

* Status: 201
* Data.body: "I love this conference"
* Data.author: "John"

## GET /comments?url=http://infoshare.pl/
---
* Status: 200
* Data[0].body: "I love this conference"
* Data[0].author: "John"
