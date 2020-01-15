# query

The `query` method performs a SQL query against the log events.

  Parameters:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`project_id` | __string__ | The project id.
`query` | __string__ | The SQL query string.
`timeout` | __integer__ | A request timeout in seconds, after which a timeout error is returned.

  Results:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`results` | __array__ of __object__ | The query results.
`stats` | [QueryStats](../types/QueryStats.md) | The query statistics.

