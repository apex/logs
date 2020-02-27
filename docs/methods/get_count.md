# get_count

The `get_count` method performs a search query against the log events, returning the number of matches.

  Inputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`project_id` | __string__ | The project id.
`query` | __string__ | The search query string.
`start` | __timestamp__ | The start timestamp, events before this time are not included.
`stop` | __timestamp__ | The stop timestamp, events after this time are not included.
`timeout` | __integer__ | A request timeout in seconds, after which a timeout error is returned.

  Outputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`count` | __integer__ | The query result count.
`stats` | [QueryStats](../types/QueryStats.md) | The query statistics.

