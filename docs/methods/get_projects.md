# get_projects

The `get_projects` method returns all projects.

  Outputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`projects` | __array__ of [Project](../types/Project.md) | The projects.

## Example

Input:

  None.

Output:

```json
{
  "projects": [
    {
      "created_at": "2019-10-30T11:44:26.005127Z",
      "description": "Apex production logs",
      "id": "apex_production",
      "location": "europe-west2",
      "name": "Apex Production",
      "retention": 60,
      "updated_at": "2019-10-30T11:44:26.005127Z"
    },
    {
      "created_at": "2019-11-28T12:33:00.034366Z",
      "description": "Ping production logs.",
      "id": "ping_production",
      "location": "europe-west2",
      "name": "Ping Production",
      "retention": 60,
      "updated_at": "2019-11-28T12:33:00.034366Z"
    },
    {
      "created_at": "2019-11-28T12:34:38.58472Z",
      "description": "Test project logs.",
      "id": "testing",
      "location": "europe-west2",
      "name": "Testing",
      "retention": 15,
      "updated_at": "2019-11-28T12:34:38.58472Z"
    }
  ]
}
```


