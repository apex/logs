# Project

The `Project` represents a customer application.

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`id` | __string__ | The project id.
`name` | __string__ | The human-friendly project name.
`retention` | __integer__ | The retention of log events in days. When zero the logs do not expire.
`location` | __string__ | The geographical location where the log events are stored.
`description` | __string__ | The project description.
`updated_at` | __timestamp__ | A timestamp indicating when the project was last updated.
`created_at` | __timestamp__ | A timestamp indicating when the project was created.

## Examples

A project configured for a production environment with 60 days of log retention.

```json
{
  "created_at": "2019-10-30T11:44:26.005127Z",
  "description": "Apex production logs",
  "id": "apex_production",
  "location": "europe-west2",
  "name": "Apex Production",
  "retention": 60,
  "updated_at": "2019-10-30T11:44:26.005127Z"
}
```

