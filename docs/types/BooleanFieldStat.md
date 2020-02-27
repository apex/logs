# BooleanFieldStat

The `BooleanFieldStat` represents a boolean field's stats.

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`value` | __boolean__ | The boolean value.
`count` | __integer__ | The number of times this field occurred in the sampled events.
`percent` | __float__ | The percentage of occurrences in the sampled events.

## Examples

A total of 95%, or 245 events have the field defined as true.

```json
{
  "count": 245,
  "percent": 0.95,
  "value": true
}
```

A total of 4%, or 12 events have the field defined as false.

```json
{
  "count": 12,
  "percent": 0.04,
  "value": false
}
```

