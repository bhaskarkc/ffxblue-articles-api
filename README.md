# Simple API built using Go

This API exposes three endpoints.

```txt
| Method | URI                    | Description                                                 |
|--------+------------------------+-------------------------------------------------------------|
| POST   | /articles              | Create/add articles, accepts JSON data.                     |
| GET    | /articles{id}          | Returns the JSON representation of the article              |
| GET    | /tags/{tagName}/{date} | List of articles having this specific tag on the given date |
```

Article detail response JSON looks like following.

```json
{
  "id": "1",
  "title": "latest science shows that potato chips are better for you than sugar",
  "date": "2016-09-22",
  "body": "some text, potentially containing simple markup about how potato chips are great",
  "tags": ["health", "fitness", "science"]
}
```

Tag details response JSON: `/tags/health/20160922`

```json
{
  "tag": "health",
  "count": 17,
  "articles": ["1", "7"],
  "related_tags": ["science", "fitness"]
}
```
