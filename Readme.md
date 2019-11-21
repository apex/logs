
### Types

  - [Alert](#Alert) — Represents configuration for performing alerting.
  - [DiscoveredField](#DiscoveredField) — Represents a single discovered field.
  - [Event](#Event) — Represents a single log event.
  - [LevelTimeseriesPoint](#LevelTimeseriesPoint) — Represents a single point in a level timeseries query.
  - [Notification](#Notification) — Represents an alert notification.
  - [Project](#Project) — Represents a customer application.
  - [QueryStats](#QueryStats) — Represents query statistics.
  - [StringFieldStat](#StringFieldStat) — Represents a string field's stats.
  - [TimeseriesPoint](#TimeseriesPoint) — Represents a single point in a timeseries query.
### Methods

  - [Alerts](#Alerts) — Define queries and thresholds used to notify your team.
  - [Events](#Events) — Ingesting and querying log events.
  - [Notifications](#Notifications) — Define how alerts are delivered to your team.
  - [Projects](#Projects) — Projects contain isolated events, alerts, and notifications.

# Alerts

  Alerts define the queries and thresholds used to notify your team of errors and anomalies. 

  - [add_alert()](#add_alert) — creates a new alert.
  - [get_alert()](#get_alert) — returns an alert.
  - [get_alerts()](#get_alerts) — returns all alerts in a project.
  - [remove_alert()](#remove_alert) — removes an alert.
  - [test_alert()](#test_alert) — test the alert configuration.
  - [update_alert()](#update_alert) — updates an alert.
  - [update_notification()](#update_notification) — updates a notification.
## add_alert

  Creates a new alert.
  - `alert` [Alert](#Alert) — The alert.

  Returns:

  - `id` __string__ — The alert id.

## get_alert

  Returns an alert.
  - `project_id` __string__ — The project id.
  - `alert_id` __string__ — The alert id.

  Returns:

  - `alert` [Alert](#Alert) — The alert.

## get_alerts

  Returns all alerts in a project.
  - `project_id` __string__ — The project id.

  Returns:

  - `alerts` __array__ — The alerts.

## remove_alert

  Removes an alert.
  - `project_id` __string__ — The project id.
  - `alert_id` __string__ — The alert id.

## test_alert

  Test the alert configuration.
  - `alert` [Alert](#Alert) — The alert.

## update_alert

  Updates an alert.
  - `alert` [Alert](#Alert) — The alert.

## update_notification

  Updates a notification.
  - `notification` [Notification](#Notification) — The notification.

# Events

  Events represent the log events ingested by one or more programs in your project. 

  - [add_events()](#add_events) — ingested a batch of events.
  - [count()](#count) — performs a search query against the log events, returning the number of matches.
  - [get_discovered_fields()](#get_discovered_fields) — returns fields discovered in the provided time range.
  - [get_level_timeseries()](#get_level_timeseries) — returns a timeseries of event counts in the provided time range grouped by severity level.
  - [get_numeric_field_stats()](#get_numeric_field_stats) — returns field statistics for a numeric field.
  - [get_string_field_stats()](#get_string_field_stats) — returns field statistics for a string field.
  - [get_timeseries()](#get_timeseries) — returns a timeseries of event counts in the provided time range.
  - [query()](#query) — performs a SQL query against the log events.
  - [search()](#search) — performs a search query against the log events.
## add_events

  Ingested a batch of events.
  - `project_id` __string__ — The project id.
  - `events` __array__ — The batch of events.

## count

  Performs a search query against the log events, returning the number of matches.
  - `timeout` __int__ — A request timeout in seconds, after which a timeout error is returned.
  - `project_id` __string__ — The project id.
  - `start` __timestamp__ — The start timestamp, events before this time are not included.
  - `stop` __timestamp__ — The stop timestamp, events after this time are not included.
  - `query` __string__ — The event search query string.

  Returns:

  - `count` __int__ — The query result count.
  - `stats` [QueryStats](#QueryStats) — The query statistics.

## get_discovered_fields

  Returns fields discovered in the provided time range.
  - `timeout` __int__ — A request timeout in seconds, after which a timeout error is returned.
  - `project_id` __string__ — The project id.
  - `start` __timestamp__ — The start timestamp, events before this time are not included.
  - `stop` __timestamp__ — The stop timestamp, events after this time are not included.
  - `query` __string__ — The event search query string.

  Returns:

  - `fields` __array__ — The fields discovered.
  - `stats` [QueryStats](#QueryStats) — The query statistics.

## get_level_timeseries

  Returns a timeseries of event counts in the provided time range grouped by severity level.
  - `timeout` __int__ — A request timeout in seconds, after which a timeout error is returned.
  - `project_id` __string__ — The project id.
  - `start` __timestamp__ — The start timestamp, events before this time are not included.
  - `stop` __timestamp__ — The stop timestamp, events after this time are not included.
  - `query` __string__ — The SQL query string.
  - `max_points` __int__ — The maxmimum number of datapoints to return.

  Returns:

  - `points` __array__ — The series of datapoints.
  - `stats` [QueryStats](#QueryStats) — The query statistics.

## get_numeric_field_stats

  Returns field statistics for a numeric field.
  - `timeout` __int__ — A request timeout in seconds, after which a timeout error is returned.
  - `project_id` __string__ — The project id.
  - `start` __timestamp__ — The start timestamp, events before this time are not included.
  - `stop` __timestamp__ — The stop timestamp, events after this time are not included.
  - `field` __string__ — The field name.

  Returns:

  - `min` __float__ — The min value.
  - `avg` __float__ — The avg value.
  - `max` __float__ — The max value.
  - `stats` [QueryStats](#QueryStats) — The query statistics.

## get_string_field_stats

  Returns field statistics for a string field.
  - `timeout` __int__ — A request timeout in seconds, after which a timeout error is returned.
  - `project_id` __string__ — The project id.
  - `start` __timestamp__ — The start timestamp, events before this time are not included.
  - `stop` __timestamp__ — The stop timestamp, events after this time are not included.
  - `field` __string__ — The field name.
  - `limit` __int__ — The maximum number of values to return.

  Returns:

  - `values` __array__ — The string values.
  - `stats` [QueryStats](#QueryStats) — The query statistics.

## get_timeseries

  Returns a timeseries of event counts in the provided time range.
  - `timeout` __int__ — A request timeout in seconds, after which a timeout error is returned.
  - `project_id` __string__ — The project id.
  - `start` __timestamp__ — The start timestamp, events before this time are not included.
  - `stop` __timestamp__ — The stop timestamp, events after this time are not included.
  - `query` __string__ — The SQL query string.
  - `max_points` __int__ — The maxmimum number of datapoints to return.

  Returns:

  - `points` __array__ — The series.
  - `stats` [QueryStats](#QueryStats) — The query statistics.

## query

  Performs a SQL query against the log events.
  - `timeout` __int__ — A request timeout in seconds, after which a timeout error is returned.
  - `project_id` __string__ — The project id.
  - `query` __string__ — The SQL query string.

  Returns:

  - `results` __array__ — The query results.
  - `stats` [QueryStats](#QueryStats) — The query statistics.

## search

  Performs a search query against the log events.
  - `timeout` __int__ — A request timeout in seconds, after which a timeout error is returned.
  - `project_id` __string__ — The project id.
  - `start` __timestamp__ — The start timestamp, events before this time are not included.
  - `stop` __timestamp__ — The stop timestamp, events after this time are not included.
  - `limit` __int__ — The maxmimum number of events to return.
  - `query` __string__ — The event search query string.

  Returns:

  - `results` __array__ — The query results.
  - `stats` [QueryStats](#QueryStats) — The query statistics.

# Notifications

  Notifications define how alerts are delivered to your team, such as email, Slack, SMS, PagerDuty or Webhook. 

  - [add_notification()](#add_notification) — creates a new notification.
  - [get_notification()](#get_notification) — returns a notification.
  - [get_notifications()](#get_notifications) — returns all notifications.
  - [remove_notification()](#remove_notification) — removes a notification.
## add_notification

  Creates a new notification.
  - `notification` [Notification](#Notification) — The notification.

  Returns:

  - `id` __string__ — The notification id.

## get_notification

  Returns a notification.
  - `project_id` __string__ — The project id.
  - `notification_id` __string__ — The notification id.

  Returns:

  - `notification` [Notification](#Notification) — The notification.

## get_notifications

  Returns all notifications.
  - `project_id` __string__ — The project id.

  Returns:

  - `notifications` __array__ — The notifications.

## remove_notification

  Removes a notification.
  - `project_id` __string__ — The project id.
  - `notification_id` __string__ — The notification id.

# Projects

  A project is a distinct set of events, alerts, and notifications. You may create separate projects for various projects, or to separate by environment such as production and staging. 

  - [add_project()](#add_project) — creates a new project.
  - [get_project_stats()](#get_project_stats) — returns project statistics.
  - [get_projects()](#get_projects) — returns all projects.
  - [remove_project()](#remove_project) — removes a project.
  - [update_project()](#update_project) — updates a project.
## add_project

  Creates a new project.
  - `project` [Project](#Project) — The project.

  Returns:

  - `id` __string__ — The project id.

## get_project_stats

  Returns project statistics.
  - `project_id` __string__ — The project id.

  Returns:

  - `bytes_total` __int__ — The total number of bytes stored.
  - `events_total` __int__ — The total number of events stored.

## get_projects

  Returns all projects.

  Returns:

  - `projects` __array__ — The projects.

## remove_project

  Removes a project.
  - `project_id` __string__ — The project id.

## update_project

  Updates a project.
  - `project` [Project](#Project) — The project.

# Types

## Alert

  The `Alert` type represents configuration for performing alerting.

  - `id` __string__ — The alert id.
  - `project_id` __string__ — The associated project id.
  - `name` __string__ — The name of the alert.
  - `description` __string__ — The description of the alert.
  - `severity` __string__ — The severity of the alert.
  - `query` __string__ — The query performed by the alert.
  - `operator` __string__ — The query performed by the alert.
  - `threshold` __int__ — The threshold for comparison against the selected operator.
  - `limit` __int__ — The maximum number of events in the alert notification.
  - `interval` __int__ — The interval in minutes for performing the alert.
  - `notification_id` __string__ — The notification id for reporting alerts, when omitted the alert will not be run.
  - `triggered` __bool__ — A boolean indicating whether or not the alert is currently triggered.
  - `muted` __bool__ — A boolean used ignore trigger and resolve notifications.
  - `updated_at` __timestamp__ — A timestamp indicating when the alert was last updated.
  - `created_at` __timestamp__ — A timestamp indicating when the alert was created.

## DiscoveredField

  The `DiscoveredField` type represents a single discovered field.

  - `name` __string__ — The field name.
  - `type` __string__ — The type of discovered field.
  - `count` __int__ — The number of times this field occurred in the sampled events.
  - `percent` __float__ — The percentage of occurrences in the sampled events.

## Event

  The `Event` type represents a single log event.

  - `id` __string__ — The event id.
  - `level` __string__ — The severity level.
  - `message` __string__ — The log message.
  - `fields` __map__ — The log fields.
  - `timestamp` __timestamp__ — The creation timestamp.

## LevelTimeseriesPoint

  The `LevelTimeseriesPoint` type represents a single point in a level timeseries query.

  - `timestamp` __timestamp__ — The bucket timestamp.
  - `level` __int__ — The severity level.
  - `count` __int__ — The number of events for this bucket.

## Notification

  The `Notification` type represents an alert notification.

  - `id` __string__ — The notification id.
  - `project_id` __string__ — The associated project id.
  - `name` __string__ — The name of the notification.
  - `type` __string__ — The type of notification.
  - `slack_webhook_url` __string__ — The Slack webhook URL.
  - `slack_channel` __string__ — The Slack channel name, otherwise the default for the webhook is used.
  - `webhook_url` __string__ — The webhook URL which receives the alert payloads.
  - `sms_numbers` __array__ — The receipients of the alert notifications.
  - `email_addresses` __array__ — The receipients of the alert notifications.
  - `pagerduty_service_key` __string__ — The PagerDuty service key.
  - `updated_at` __timestamp__ — A timestamp indicating when the notification was last updated.
  - `created_at` __timestamp__ — A timestamp indicating when the notification was created.

## Project

  The `Project` type represents a customer application.

  - `id` __string__ — The project id.
  - `name` __string__ — The human-friendly project name.
  - `retention` __int__ — The retention of log events in days. When zero the logs do not expire.
  - `location` __string__ — The geographical location where the log events are stored.
  - `description` __string__ — The project description.
  - `updated_at` __timestamp__ — A timestamp indicating when the project was last updated.
  - `created_at` __timestamp__ — A timestamp indicating when the project was created.

## QueryStats

  The `QueryStats` type represents query statistics.

  - `total_bytes_processed` __int__ — The total number of bytes processed by the query.
  - `total_bytes_billed` __int__ — The total number of bytes billed by the query.
  - `cache_hit` __bool__ — A boolean indicating if the query was cached.

## StringFieldStat

  The `StringFieldStat` type represents a string field's stats.

  - `value` __string__ — The string value.
  - `count` __int__ — The number of times this field occurred in the sampled events.
  - `percent` __float__ — The percentage of occurrences in the sampled events.

## TimeseriesPoint

  The `TimeseriesPoint` type represents a single point in a timeseries query.

  - `timestamp` __timestamp__ — The bucket timestamp.
  - `count` __int__ — The number of events for this bucket.

