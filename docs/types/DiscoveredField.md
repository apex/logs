# DiscoveredField

The `DiscoveredField` represents a single discovered field.

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`name` | __string__ | The field name.
`type` | __string__ | The type of discovered field.
`count` | __integer__ | The number of times this field occurred in the sampled events.
`percent` | __float__ | The percentage of occurrences in the sampled events.

## Examples

The `aws.log.group` is a user-defined string field, present in 97% of the events.

```json
{
  "count": 14417,
  "name": "aws.log.group",
  "percent": 0.97,
  "type": "string"
}
```

The `duration` field is a user-defined numeric field, present in 11% of the events.

```json
{
  "count": 1770,
  "name": "duration",
  "percent": 0.11,
  "type": "number"
}
```

The `message` field is an example of a built-in field, so it is included in 100% of the events.

```json
{
  "count": 14500,
  "name": "message",
  "percent": 1,
  "type": "string"
}
```

