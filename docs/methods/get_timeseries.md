# get_timeseries

The `get_timeseries` method returns a timeseries of event counts in the provided time range.

  Parameters:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`max_points` | __integer__ | The maxmimum number of datapoints to return.
`project_id` | __string__ | The project id.
`query` | __string__ | The SQL query string.
`start` | __timestamp__ | The start timestamp, events before this time are not included.
`stop` | __timestamp__ | The stop timestamp, events after this time are not included.
`timeout` | __integer__ | A request timeout in seconds, after which a timeout error is returned.

  Results:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`points` | __array__ of [TimeseriesPoint](../types/TimeseriesPoint.md) | The series.
`stats` | [QueryStats](../types/QueryStats.md) | The query statistics.

