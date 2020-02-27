# StringFieldStat

The `StringFieldStat` represents a string field's stats.

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`value` | __string__ | The string value.
`count` | __integer__ | The number of times this field occurred in the sampled events.
`percent` | __float__ | The percentage of occurrences in the sampled events.

## Examples

The field's string value is "downtime" for 87% of the events.

```json
{
  "count": 3413,
  "percent": 0.87,
  "value": "downtime"
}
```

The field's string value is "time_total" for 4% of the events.

```json
{
  "count": 168,
  "percent": 0.04,
  "value": "time_total"
}
```

The field's string value is "time_namelookup" for 0.1% of the events.

```json
{
  "count": 6,
  "percent": 0.001,
  "value": "time_namelookup"
}
```

