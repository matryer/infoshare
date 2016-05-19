# infoshare
Code for infoShare workshop 19 May 2016



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
