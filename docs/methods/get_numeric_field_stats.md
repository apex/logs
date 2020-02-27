# get_numeric_field_stats

The `get_numeric_field_stats` method returns field statistics for a numeric field.

  Inputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`field` | __string__ | The field name.
`project_id` | __string__ | The project id.
`query` | __string__ | The search query string.
`start` | __timestamp__ | The start timestamp, events before this time are not included.
`stop` | __timestamp__ | The stop timestamp, events after this time are not included.
`timeout` | __integer__ | A request timeout in seconds, after which a timeout error is returned.

  Outputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`avg` | __float__ | The avg value.
`max` | __float__ | The max value.
`min` | __float__ | The min value.
`stats` | [QueryStats](../types/QueryStats.md) | The query statistics.

