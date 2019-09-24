
# Apex Logs

Client libraries for [Apex Logs](https://apex.sh/logs/).

## Libraries

- [js](./js) — JavaScript client for Node.js and the browser
- [go](./go) — Golang client


# Types

## Alert

  Alert represents configuration for performing alerting.

  - `id` string — the alert id.
  - `project_id` string — the associated project id.
  - `name` string — the name of the alert.
  - `severity` string — the severity of the alert.
  - `query` string — the query performed by the alert.
  - `operator` string — the query performed by the alert.
  - `threshold` int — the threshold for comparison against the selected operator.
  - `limit` int — the maximum number of events in the alert notification.
  - `interval` int — the interval in minutes for performing the alert.
  - `notification_id` string — the notification id for reporting alerts, when omitted the alert will not be run.
  - `description` string — the description of the alert.
  - `updated_at` timestamp — a timestamp indicating when the alert was last updated.
  - `created_at` timestamp — a timestamp indicating when the alert was created.

## DiscoveredField

  DiscoveredField represents a single discovered field.

  - `name` string — the field name.
  - `type` string — the type of discovered field.
  - `count` int — the number of times this field occurred in the sampled events.
  - `percent` float — the percentage of occurrences in the sampled events.

## Event

  Event represents a single log event.

  - `id` string — the event id.
  - `level` string — the severity level.
  - `message` string — the log message.
  - `fields` map — the log fields.
  - `timestamp` timestamp — the creation timestamp.

## LevelTimeseriesPoint

  LevelTimeseriesPoint represents a single point in a level timeseries query.

  - `timestamp` timestamp — the bucket timestamp.
  - `level` int — the severity level.
  - `count` int — the number of events for this bucket.

## Notification

  Notification represents an alert notification.

  - `id` string — the notification id.
  - `project_id` string — the associated project id.
  - `name` string — the name of the notification.
  - `type` string — the type of notification.
  - `slack_webhook_url` string — the Slack webhook URL.
  - `slack_channel` string — the Slack channel name, otherwise the default for the webhook is used.
  - `webhook_url` string — the webhook URL which receives the alert payloads.
  - `sms_numbers` array — the receipients of the alert notifications.
  - `email_addresses` array — the receipients of the alert notifications.
  - `pagerduty_service_key` string — the PagerDuty service key.
  - `updated_at` timestamp — a timestamp indicating when the notification was last updated.
  - `created_at` timestamp — a timestamp indicating when the notification was created.

## Project

  Project represents a customer application.

  - `id` string — the project id.
  - `name` string — the human-friendly project name.
  - `retention` int — the retention of log events in days. When zero the logs do not expire.
  - `location` string — the geographical location where the log events are stored.
  - `description` string — the project description.
  - `updated_at` timestamp — a timestamp indicating when the project was last updated.
  - `created_at` timestamp — a timestamp indicating when the project was created.

## QueryStats

  QueryStats represents query statistics.

  - `total_bytes_processed` int — the total number of bytes processed by the query.
  - `total_bytes_billed` int — the total number of bytes billed by the query.
  - `cache_hit` bool — a boolean indicating if the query was cached.

## StringFieldStat

  StringFieldStat 

  - `value` string — the string value.
  - `count` int — the number of times this field occurred in the sampled events.

## TimeseriesPoint

  TimeseriesPoint represents a single point in a timeseries query.

  - `timestamp` timestamp — the bucket timestamp.
  - `count` int — the number of events for this bucket.

# Methods

  - alerts
    - [add_alert()](#add_alert) — creates a new alert.
    - [get_alert()](#get_alert) — returns an alert.
    - [get_alerts()](#get_alerts) — returns all alerts in a project.
    - [remove_alert()](#remove_alert) — removes an alert.
    - [update_alert()](#update_alert) — updates an alert.
    - [update_notification()](#update_notification) — updates an notification.
  - events
    - [add_events()](#add_events) — ingested a batch of events.
    - [get_discovered_fields()](#get_discovered_fields) — returns fields discovered in the provided time range.
    - [get_level_timeseries()](#get_level_timeseries) — returns a timeseries of event counts in the provided time range grouped by severity level.
    - [get_numeric_field_stats()](#get_numeric_field_stats) — returns field statistics for a numeric field.
    - [get_string_field_stats()](#get_string_field_stats) — returns field statistics for a string field.
    - [get_timeseries()](#get_timeseries) — returns a timeseries of event counts in the provided time range.
    - [query()](#query) — performs a SQL query against the log events.
    - [search()](#search) — performs a search query against the log events.
  - notifications
    - [add_notification()](#add_notification) — creates a new notification.
    - [get_notification()](#get_notification) — returns an notification.
    - [get_notifications()](#get_notifications) — returns all notifications.
    - [remove_notification()](#remove_notification) — removes an notification.
  - projects
    - [add_project()](#add_project) — creates a new project.
    - [get_project_stats()](#get_project_stats) — returns project statistics.
    - [get_projects()](#get_projects) — returns all projects.
    - [remove_project()](#remove_project) — removes a project.
    - [update_project()](#update_project) — updates a project.

## alerts

### add_alert()

  creates a new alert.

  Input:

  - `alert` [Alert](#Alert) — the alert.

  Output:

  - `id` string — the alert id.

### get_alert()

  returns an alert.

  Input:

  - `project_id` string — the project id.
  - `alert_id` string — the alert id.

  Output:

  - `alert` [Alert](#Alert) — the alert.

### get_alerts()

  returns all alerts in a project.

  Input:

  - `project_id` string — the project id.

  Output:

  - `alerts` array — the alerts.

### remove_alert()

  removes an alert.

  Input:

  - `project_id` string — the project id.
  - `alert_id` string — the alert id.

### update_alert()

  updates an alert.

  Input:

  - `alert` [Alert](#Alert) — the alert.

### update_notification()

  updates an notification.

  Input:

  - `notification` [Notification](#Notification) — the notification.

## events

### add_events()

  ingested a batch of events.

  Input:

  - `project_id` string — the project id.
  - `events` array — the batch of events.

### get_discovered_fields()

  returns fields discovered in the provided time range.

  Input:

  - `timeout` int — a request timeout in seconds, after which a timeout error is returned.
  - `project_id` string — the project id.
  - `start` timestamp — the start timestamp, events before this time are not included.
  - `stop` timestamp — the stop timestamp, events after this time are not included.

  Output:

  - `fields` array — the fields discovered.

### get_level_timeseries()

  returns a timeseries of event counts in the provided time range grouped by severity level.

  Input:

  - `timeout` int — a request timeout in seconds, after which a timeout error is returned.
  - `project_id` string — the project id.
  - `start` timestamp — the start timestamp, events before this time are not included.
  - `stop` timestamp — the stop timestamp, events after this time are not included.
  - `query` string — the SQL query string.
  - `max_points` int — the maxmimum number of datapoints to return.

  Output:

  - `points` array — the series of datapoints.

### get_numeric_field_stats()

  returns field statistics for a numeric field.

  Input:

  - `timeout` int — a request timeout in seconds, after which a timeout error is returned.
  - `project_id` string — the project id.
  - `start` timestamp — the start timestamp, events before this time are not included.
  - `stop` timestamp — the stop timestamp, events after this time are not included.
  - `field` string — the field name.

  Output:

  - `min` float — the min value.
  - `avg` float — the avg value.
  - `max` float — The max value.

### get_string_field_stats()

  returns field statistics for a string field.

  Input:

  - `timeout` int — a request timeout in seconds, after which a timeout error is returned.
  - `project_id` string — the project id.
  - `start` timestamp — the start timestamp, events before this time are not included.
  - `stop` timestamp — the stop timestamp, events after this time are not included.
  - `field` string — the field name.

  Output:

  - `values` array — the string values.

### get_timeseries()

  returns a timeseries of event counts in the provided time range.

  Input:

  - `timeout` int — a request timeout in seconds, after which a timeout error is returned.
  - `project_id` string — the project id.
  - `start` timestamp — the start timestamp, events before this time are not included.
  - `stop` timestamp — the stop timestamp, events after this time are not included.
  - `query` string — the SQL query string.
  - `max_points` int — the maxmimum number of datapoints to return.

  Output:

  - `points` array — the series.

### query()

  performs a SQL query against the log events.

  Input:

  - `timeout` int — a request timeout in seconds, after which a timeout error is returned.
  - `project_id` string — the project id.
  - `query` string — the SQL query string.

  Output:

  - `results` map — the query results.
  - `stats` [QueryStats](#QueryStats) — the query statistics.

### search()

  performs a search query against the log events.

  Input:

  - `timeout` int — a request timeout in seconds, after which a timeout error is returned.
  - `project_id` string — the project id.
  - `start` timestamp — the start timestamp, events before this time are not included.
  - `stop` timestamp — the stop timestamp, events after this time are not included.
  - `limit` int — the maxmimum number of events to return.
  - `query` string — the event search query string.

  Output:

  - `results` map — the query results.
  - `stats` [QueryStats](#QueryStats) — the query statistics.

## notifications

### add_notification()

  creates a new notification.

  Input:

  - `notification` [Notification](#Notification) — the notification.

  Output:

  - `id` string — the notification id.

### get_notification()

  returns an notification.

  Input:

  - `project_id` string — the project id.
  - `notification_id` string — the notification id.

  Output:

  - `notification` [Notification](#Notification) — the notification.

### get_notifications()

  returns all notifications.

  Input:

  - `project_id` string — the project id.

  Output:

  - `notifications` array — the notifications.

### remove_notification()

  removes an notification.

  Input:

  - `project_id` string — the project id.
  - `notification_id` string — the notification id.

## projects

### add_project()

  creates a new project.

  Input:

  - `project` [Project](#Project) — the project.

  Output:

  - `id` string — the project id.

### get_project_stats()

  returns project statistics.

  Input:

  - `project_id` string — the project id.

  Output:

  - `bytes_total` int — the total number of bytes stored.
  - `events_total` int — the total number of events stored.

### get_projects()

  returns all projects.

  Output:

  - `projects` array — the projects.

### remove_project()

  removes a project.

  Input:

  - `project_id` string — the project id.

### update_project()

  updates a project.

  Input:

  - `project` [Project](#Project) — the project.

