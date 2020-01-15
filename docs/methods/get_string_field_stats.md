# get_string_field_stats

The `get_string_field_stats` method returns field statistics for a string field.

  Parameters:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`field` | __string__ | The field name.
`limit` | __integer__ | The maximum number of values to return.
`project_id` | __string__ | The project id.
`start` | __timestamp__ | The start timestamp, events before this time are not included.
`stop` | __timestamp__ | The stop timestamp, events after this time are not included.
`timeout` | __integer__ | A request timeout in seconds, after which a timeout error is returned.

  Results:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`stats` | [QueryStats](../types/QueryStats.md) | The query statistics.
`values` | __array__ of [StringFieldStat](../types/StringFieldStat.md) | The string values.

