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
