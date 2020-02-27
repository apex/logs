# get_discovered_fields

The `get_discovered_fields` method returns fields discovered in the provided time range.

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
`fields` | __array__ of [DiscoveredField](../types/DiscoveredField.md) | The fields discovered.
`stats` | [QueryStats](../types/QueryStats.md) | The query statistics.

