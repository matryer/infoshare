# infoshare
Code for infoShare workshop 19 May 2016

* We built a comments API
* Deployed it to App Engine
* http://infoshare-api.appspot.com/comments

GET /comments?url=https://...
[
  {
    "body": "This is my comment",
    "author": "John"
  },
  {
    "body": "This is my comment",
    "author": "John"
  },
  {
    "body": "This is my comment",
    "author": "John"
  }
]

POST /comments
{
  "url": "https://...",
  "body": "This is my comment",
  "author": "John"
}
