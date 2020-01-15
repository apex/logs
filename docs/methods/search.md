# search

The `search` method performs a search query against the log events.

  Parameters:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`limit` | __integer__ | The maxmimum number of events to return.
`project_id` | __string__ | The project id.
`query` | __string__ | The event search query string.
`start` | __timestamp__ | The start timestamp, events before this time are not included.
`stop` | __timestamp__ | The stop timestamp, events after this time are not included.
`timeout` | __integer__ | A request timeout in seconds, after which a timeout error is returned.

  Results:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`results` | __array__ of __object__ | The query results.
`stats` | [QueryStats](../types/QueryStats.md) | The query statistics.

