# add_search

The `add_search` method creates a new saved search.

  Inputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`search` | [Search](../types/Search.md) | The saved search.

  Outputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`id` | __string__ | The saved search id.

## Example

An example of a saved search query for reporting related messages.

Input:

```json
{
  "name": "Weekly reports",
  "project_id": "ping_production",
  "query": "function = \"reporter\" and message in (\"fetching reports\", \"reporting complete\")"
}
```

Output:

```json
{
  "id": "1ZqPnuhRyStpljCI8ahVSsV9txx"
}
```


