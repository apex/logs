# get_searches

The `get_searches` method returns all saved searches in a project.

  Inputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`project_id` | __string__ | The project id.

  Outputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`searches` | __array__ of [Search](../types/Search.md) | The saved searches.

## Example

An example listing all of the saved searches for a project.

Input:

```json
{
  "project_id": "ping_production"
}
```

Output:

```json
{
  "searches": [
    {
      "created_at": "2020-03-30T11:23:16.099557+01:00",
      "id": "1ZqPkpH63A5ZpOJEX443ZirtmH6",
      "name": "Alert notifications",
      "project_id": "ping_production",
      "query": "function = \"alert_reporter\" and message = \"reporting complete\"",
      "updated_at": "2020-03-30T11:24:39.560604+01:00"
    },
    {
      "created_at": "2020-03-30T11:23:37.675798+01:00",
      "id": "1ZqPnX3WN2hAGHjKeQpRfEaLYMr",
      "name": "Weekly reports",
      "project_id": "ping_production",
      "query": "function = \"reporter\" and message in (\"fetching reports\", \"reporting complete\")",
      "updated_at": "2020-03-30T11:23:54.874927+01:00"
    },
    {
      "created_at": "2020-03-30T11:23:40.364422+01:00",
      "id": "1ZqPnuhRyStpljCI8ahVSsV9txx",
      "name": "Lambda timeouts",
      "project_id": "ping_production",
      "query": "message contains \"timed out\"",
      "updated_at": "2020-03-30T11:23:45.569638+01:00"
    }
  ]
}
```


