# TimeseriesPoint

The `TimeseriesPoint` represents a single point in a timeseries query.

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`timestamp` | __timestamp__ | The bucket timestamp.
`count` | __integer__ | The number of events for this bucket.

## Examples

A timeseries point with 5 events.

```json
{
  "count": 5,
  "timestamp": "2020-01-15T13:35:38Z"
}
```

