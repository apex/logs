# search

The `search` method performs a search query against the log events.

  Inputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`limit` | __integer__ | The maxmimum number of events to return.
`project_id` | __string__ | The project id.
`query` | __string__ | The search query string.
`start` | __timestamp__ | The start timestamp, events before this time are not included.
`stop` | __timestamp__ | The stop timestamp, events after this time are not included.
`timeout` | __integer__ | A request timeout in seconds, after which a timeout error is returned.

  Outputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`results` | __array__ of __object__ | The query results.
`stats` | [QueryStats](../types/QueryStats.md) | The query statistics.

## Example

Input:

```json
{
  "limit": 500,
  "project_id": "app_production",
  "query": "function = \"alert_reporter\" and message = \"reporting complete\"",
  "start": "2020-01-14T13:56:05.961Z",
  "stop": "2020-01-15T13:56:05.961Z"
}
```

Output:

```json
{
  "results": [
    {
      "fields": "{\"action\":\"slack\",\"alert_id\":3637,\"aws\":{\"log\":{\"group\":\"/aws/lambda/vitals_alert_reporter\",\"stream\":\"2020/01/15/[61]6865a073c8864f1f820ba002c0cea12e\"},\"region\":\"us-west-2\"},\"check_id\":17142,\"function\":\"alert_reporter\",\"triggered\":true,\"value\":1,\"version\":\"61\"}",
      "id": "1WQyis5UQYQeIYMFAbNjkvKLjsn",
      "level": "info",
      "message": "reporting complete",
      "timestamp": "2020-01-15T13:50:26.628999948Z"
    },
    {
      "fields": "{\"action\":\"slack\",\"alert_id\":5322,\"aws\":{\"log\":{\"group\":\"/aws/lambda/vitals_alert_reporter\",\"stream\":\"2020/01/15/[61]6865a073c8864f1f820ba002c0cea12e\"},\"region\":\"us-west-2\"},\"check_id\":21055,\"function\":\"alert_reporter\",\"triggered\":true,\"value\":5,\"version\":\"61\"}",
      "id": "1WQxxVw35oaoCcfCXGYj9jx1zmn",
      "level": "info",
      "message": "reporting complete",
      "timestamp": "2020-01-15T13:44:10.963000059Z"
    }
  ],
  "stats": {
    "cache_hit": false,
    "total_bytes_billed": 1493172224,
    "total_bytes_processed": 1492717808
  }
}
```


