# Methods

## Alerts

Alerts define the queries and thresholds used to notify your team of errors and anomalies.

  - [add_alert](./add_alert.md) — creates a new alert.
  - [get_alert](./get_alert.md) — returns an alert.
  - [get_alerts](./get_alerts.md) — returns all alerts in a project.
  - [remove_alert](./remove_alert.md) — removes an alert.
  - [test_alert](./test_alert.md) — test the alert configuration.
  - [update_alert](./update_alert.md) — updates an alert.
  - [update_notification](./update_notification.md) — updates a notification.

## Events

Events represent the log events ingested by one or more programs in your project.

  - [add_events](./add_events.md) — ingests a batch of events.
  - [get_boolean_field_stats](./get_boolean_field_stats.md) — returns field statistics for a boolean field.
  - [get_count](./get_count.md) — performs a search query against the log events, returning the number of matches.
  - [get_discovered_fields](./get_discovered_fields.md) — returns fields discovered in the provided time range.
  - [get_numeric_field_stats](./get_numeric_field_stats.md) — returns field statistics for a numeric field.
  - [get_string_field_stats](./get_string_field_stats.md) — returns field statistics for a string field.
  - [get_timeseries](./get_timeseries.md) — returns a timeseries of event counts in the provided time range.
  - [query](./query.md) — performs a SQL query against the log events.
  - [search](./search.md) — performs a search query against the log events.

## Notifications

Notifications define how alerts are delivered to your team, such as email, Slack, SMS, PagerDuty or Webhook.

  - [add_notification](./add_notification.md) — creates a new notification.
  - [get_notification](./get_notification.md) — returns a notification.
  - [get_notifications](./get_notifications.md) — returns all notifications.
  - [remove_notification](./remove_notification.md) — removes a notification.

## Projects

A project is a distinct set of events, alerts, and notifications. You may create separate projects for various projects, or to separate by environment such as production and staging.

  - [add_project](./add_project.md) — creates a new project.
  - [get_project_stats](./get_project_stats.md) — returns project statistics.
  - [get_projects](./get_projects.md) — returns all projects.
  - [remove_project](./remove_project.md) — removes a project.
  - [update_project](./update_project.md) — updates a project.

