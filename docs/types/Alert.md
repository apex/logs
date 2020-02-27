# Alert

The `Alert` represents configuration for performing alerting.

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`id` | __string__ | The alert id.
`project_id` | __string__ | The associated project id.
`name` | __string__ | The name of the alert.
`description` | __string__ | The description of the alert.
`severity` | __string__ | The severity of the alert.
`query` | __string__ | The query performed by the alert.
`operator` | __string__ | The query performed by the alert.
`threshold` | __integer__ | The threshold for comparison against the selected operator.
`limit` | __integer__ | The maximum number of events in the alert notification.
`interval` | __integer__ | The interval in minutes for performing the alert.
`notification_id` | __string__ | The notification id for reporting alerts, when omitted the alert will not be run.
`triggered` | __boolean__ | A boolean indicating whether or not the alert is currently triggered.
`muted` | __boolean__ | A boolean used ignore trigger and resolve notifications.
`updated_at` | __timestamp__ | A timestamp indicating when the alert was last updated.
`created_at` | __timestamp__ | A timestamp indicating when the alert was created.

## Examples

An alert configured to notify the team when one or more errors is found.

```json
{
  "created_at": "2019-11-22T13:48:44.21231Z",
  "description": "All production errors",
  "id": "1RHyOWJ8cIlvDb1bcKZHdiISnX6",
  "interval": 5,
  "limit": 100,
  "muted": false,
  "name": "Errors",
  "notification_id": "1Q3bUUxSUZVlreTTwVzdK5z37iX",
  "operator": ">=",
  "project_id": "ping_production",
  "query": "level > warning or message contains \"panic\"",
  "severity": "error",
  "threshold": 1,
  "triggered": false,
  "updated_at": "2019-11-22T13:48:44.21231Z"
}
```

