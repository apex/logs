# Search

The `Search` represents a saved search query.

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`id` | __string__ | The saved search id.
`name` | __string__ | The name of the saved search.
`project_id` | __string__ | The associated project id.
`query` | __string__ | The saved search query.
`updated_at` | __timestamp__ | A timestamp indicating when the saved search was last updated.
`created_at` | __timestamp__ | A timestamp indicating when the saved search was created.

## Examples

An example of a saved search query.

```json
{
  "created_at": "2020-03-30T11:23:37.675798+01:00",
  "id": "1ZqPnX3WN2hAGHjKeQpRfEaLYMr",
  "name": "Weekly reports",
  "project_id": "ping_production",
  "query": "function = \"reporter\"  and message in (\"fetching reports\", \"reporting complete\")",
  "updated_at": "2020-03-30T11:30:54.874927+01:00"
}
```

