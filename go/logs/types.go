package logs

import (
  "time"
)

// QueryStats represents query statistics.
type QueryStats struct {
  // TotalBytesProcessed is the total number of bytes processed by the query.
  TotalBytesProcessed int `json:"total_bytes_processed"`

  // TotalBytesBilled is the total number of bytes billed by the query.
  TotalBytesBilled int `json:"total_bytes_billed"`

  // CacheHit is a boolean indicating if the query was cached.
  CacheHit bool `json:"cache_hit"`
}

// TimeseriesPoint represents a single point in a timeseries query.
type TimeseriesPoint struct {
  // Timestamp is the bucket timestamp.
  Timestamp time.Time `json:"timestamp"`

  // Count is the number of events for this bucket.
  Count int `json:"count"`
}

// LevelTimeseriesPoint represents a single point in a level timeseries query.
type LevelTimeseriesPoint struct {
  // Timestamp is the bucket timestamp.
  Timestamp time.Time `json:"timestamp"`

  // Level is the severity level.
  Level int `json:"level"`

  // Count is the number of events for this bucket.
  Count int `json:"count"`
}

// DiscoveredField represents a single discovered field.
type DiscoveredField struct {
  // Name is the field name.
  Name string `json:"name"`

  // Type is the type of discovered field.
  Type string `json:"type"`

  // Count is the number of times this field occurred in the sampled events.
  Count int `json:"count"`

  // Percent is the percentage of occurrences in the sampled events.
  Percent float64 `json:"percent"`
}

// StringFieldStat represents a string field's stats.
type StringFieldStat struct {
  // Value is the string value.
  Value string `json:"value"`

  // Count is the number of times this field occurred in the sampled events.
  Count int `json:"count"`
}

// Project represents a customer application.
type Project struct {
  // ID is the project id.
  ID string `json:"id"`

  // Name is the human-friendly project name.
  Name string `json:"name"`

  // Retention is the retention of log events in days. When zero the logs do not expire.
  Retention int `json:"retention"`

  // Location is the geographical location where the log events are stored.
  Location string `json:"location"`

  // Description is the project description.
  Description string `json:"description"`

  // UpdatedAt is a timestamp indicating when the project was last updated.
  UpdatedAt time.Time `json:"updated_at"`

  // CreatedAt is a timestamp indicating when the project was created.
  CreatedAt time.Time `json:"created_at"`
}

// Notification represents an alert notification.
type Notification struct {
  // ID is the notification id.
  ID string `json:"id"`

  // ProjectID is the associated project id.
  ProjectID string `json:"project_id"`

  // Name is the name of the notification.
  Name string `json:"name"`

  // Type is the type of notification.
  Type string `json:"type"`

  // SlackWebhookURL is the Slack webhook URL.
  SlackWebhookURL string `json:"slack_webhook_url"`

  // SlackChannel is the Slack channel name, otherwise the default for the webhook is used.
  SlackChannel string `json:"slack_channel"`

  // WebhookURL is the webhook URL which receives the alert payloads.
  WebhookURL string `json:"webhook_url"`

  // SmsNumbers is the receipients of the alert notifications.
  SmsNumbers []string `json:"sms_numbers"`

  // EmailAddresses is the receipients of the alert notifications.
  EmailAddresses []string `json:"email_addresses"`

  // PagerdutyServiceKey is the PagerDuty service key.
  PagerdutyServiceKey string `json:"pagerduty_service_key"`

  // UpdatedAt is a timestamp indicating when the notification was last updated.
  UpdatedAt time.Time `json:"updated_at"`

  // CreatedAt is a timestamp indicating when the notification was created.
  CreatedAt time.Time `json:"created_at"`
}

// Alert represents configuration for performing alerting.
type Alert struct {
  // ID is the alert id.
  ID string `json:"id"`

  // ProjectID is the associated project id.
  ProjectID string `json:"project_id"`

  // Name is the name of the alert.
  Name string `json:"name"`

  // Description is the description of the alert.
  Description string `json:"description"`

  // Severity is the severity of the alert.
  Severity string `json:"severity"`

  // Query is the query performed by the alert.
  Query string `json:"query"`

  // Operator is the query performed by the alert.
  Operator string `json:"operator"`

  // Threshold is the threshold for comparison against the selected operator.
  Threshold int `json:"threshold"`

  // Limit is the maximum number of events in the alert notification.
  Limit int `json:"limit"`

  // Interval is the interval in minutes for performing the alert.
  Interval int `json:"interval"`

  // NotificationID is the notification id for reporting alerts, when omitted the alert will not be run.
  NotificationID string `json:"notification_id"`

  // Triggered is a boolean indicating whether or not the alert is currently triggered.
  Triggered bool `json:"triggered"`

  // Muted is a boolean used ignore trigger and resolve notifications.
  Muted bool `json:"muted"`

  // UpdatedAt is a timestamp indicating when the alert was last updated.
  UpdatedAt time.Time `json:"updated_at"`

  // CreatedAt is a timestamp indicating when the alert was created.
  CreatedAt time.Time `json:"created_at"`
}

// Event represents a single log event.
type Event struct {
  // ID is the event id.
  ID string `json:"id"`

  // Level is the severity level.
  Level string `json:"level"`

  // Message is the log message.
  Message string `json:"message"`

  // Fields is the log fields.
  Fields map[string]interface{} `json:"fields"`

  // Timestamp is the creation timestamp.
  Timestamp time.Time `json:"timestamp"`
}

// AddEventsInput params.
type AddEventsInput struct {
  // ProjectID is the project id.
  ProjectID string `json:"project_id"`

  // Events is the batch of events.
  Events []Event `json:"events"`
}

// AddAlertInput params.
type AddAlertInput struct {
  // Alert is the alert.
  Alert Alert `json:"alert"`
}

// AddAlertOutput params.
type AddAlertOutput struct {
  // ID is the alert id.
  ID string `json:"id"`
}

// UpdateAlertInput params.
type UpdateAlertInput struct {
  // Alert is the alert.
  Alert Alert `json:"alert"`
}

// RemoveAlertInput params.
type RemoveAlertInput struct {
  // ProjectID is the project id.
  ProjectID string `json:"project_id"`

  // AlertID is the alert id.
  AlertID string `json:"alert_id"`
}

// GetAlertInput params.
type GetAlertInput struct {
  // ProjectID is the project id.
  ProjectID string `json:"project_id"`

  // AlertID is the alert id.
  AlertID string `json:"alert_id"`
}

// GetAlertOutput params.
type GetAlertOutput struct {
  // Alert is the alert.
  Alert Alert `json:"alert"`
}

// GetAlertsInput params.
type GetAlertsInput struct {
  // ProjectID is the project id.
  ProjectID string `json:"project_id"`
}

// GetAlertsOutput params.
type GetAlertsOutput struct {
  // Alerts is the alerts.
  Alerts []Alert `json:"alerts"`
}

// AddNotificationInput params.
type AddNotificationInput struct {
  // Notification is the notification.
  Notification Notification `json:"notification"`
}

// AddNotificationOutput params.
type AddNotificationOutput struct {
  // ID is the notification id.
  ID string `json:"id"`
}

// UpdateNotificationInput params.
type UpdateNotificationInput struct {
  // Notification is the notification.
  Notification Notification `json:"notification"`
}

// RemoveNotificationInput params.
type RemoveNotificationInput struct {
  // ProjectID is the project id.
  ProjectID string `json:"project_id"`

  // NotificationID is the notification id.
  NotificationID string `json:"notification_id"`
}

// GetNotificationInput params.
type GetNotificationInput struct {
  // ProjectID is the project id.
  ProjectID string `json:"project_id"`

  // NotificationID is the notification id.
  NotificationID string `json:"notification_id"`
}

// GetNotificationOutput params.
type GetNotificationOutput struct {
  // Notification is the notification.
  Notification Notification `json:"notification"`
}

// GetNotificationsInput params.
type GetNotificationsInput struct {
  // ProjectID is the project id.
  ProjectID string `json:"project_id"`
}

// GetNotificationsOutput params.
type GetNotificationsOutput struct {
  // Notifications is the notifications.
  Notifications []Notification `json:"notifications"`
}

// AddProjectInput params.
type AddProjectInput struct {
  // Project is the project.
  Project Project `json:"project"`
}

// AddProjectOutput params.
type AddProjectOutput struct {
  // ID is the project id.
  ID string `json:"id"`
}

// UpdateProjectInput params.
type UpdateProjectInput struct {
  // Project is the project.
  Project Project `json:"project"`
}

// RemoveProjectInput params.
type RemoveProjectInput struct {
  // ProjectID is the project id.
  ProjectID string `json:"project_id"`
}

// GetProjectsOutput params.
type GetProjectsOutput struct {
  // Projects is the projects.
  Projects []Project `json:"projects"`
}

// GetProjectStatsInput params.
type GetProjectStatsInput struct {
  // ProjectID is the project id.
  ProjectID string `json:"project_id"`
}

// GetProjectStatsOutput params.
type GetProjectStatsOutput struct {
  // BytesTotal is the total number of bytes stored.
  BytesTotal int `json:"bytes_total"`

  // EventsTotal is the total number of events stored.
  EventsTotal int `json:"events_total"`
}

// GetDiscoveredFieldsInput params.
type GetDiscoveredFieldsInput struct {
  // Timeout is a request timeout in seconds, after which a timeout error is returned.
  Timeout int `json:"timeout"`

  // ProjectID is the project id.
  ProjectID string `json:"project_id"`

  // Start is the start timestamp, events before this time are not included.
  Start time.Time `json:"start"`

  // Stop is the stop timestamp, events after this time are not included.
  Stop time.Time `json:"stop"`
}

// GetDiscoveredFieldsOutput params.
type GetDiscoveredFieldsOutput struct {
  // Fields is the fields discovered.
  Fields []DiscoveredField `json:"fields"`
}

// GetTimeseriesInput params.
type GetTimeseriesInput struct {
  // Timeout is a request timeout in seconds, after which a timeout error is returned.
  Timeout int `json:"timeout"`

  // ProjectID is the project id.
  ProjectID string `json:"project_id"`

  // Start is the start timestamp, events before this time are not included.
  Start time.Time `json:"start"`

  // Stop is the stop timestamp, events after this time are not included.
  Stop time.Time `json:"stop"`

  // Query is the SQL query string.
  Query string `json:"query"`

  // MaxPoints is the maxmimum number of datapoints to return.
  MaxPoints int `json:"max_points"`
}

// GetTimeseriesOutput params.
type GetTimeseriesOutput struct {
  // Points is the series.
  Points []TimeseriesPoint `json:"points"`
}

// GetLevelTimeseriesInput params.
type GetLevelTimeseriesInput struct {
  // Timeout is a request timeout in seconds, after which a timeout error is returned.
  Timeout int `json:"timeout"`

  // ProjectID is the project id.
  ProjectID string `json:"project_id"`

  // Start is the start timestamp, events before this time are not included.
  Start time.Time `json:"start"`

  // Stop is the stop timestamp, events after this time are not included.
  Stop time.Time `json:"stop"`

  // Query is the SQL query string.
  Query string `json:"query"`

  // MaxPoints is the maxmimum number of datapoints to return.
  MaxPoints int `json:"max_points"`
}

// GetLevelTimeseriesOutput params.
type GetLevelTimeseriesOutput struct {
  // Points is the series of datapoints.
  Points []LevelTimeseriesPoint `json:"points"`
}

// GetNumericFieldStatsInput params.
type GetNumericFieldStatsInput struct {
  // Timeout is a request timeout in seconds, after which a timeout error is returned.
  Timeout int `json:"timeout"`

  // ProjectID is the project id.
  ProjectID string `json:"project_id"`

  // Start is the start timestamp, events before this time are not included.
  Start time.Time `json:"start"`

  // Stop is the stop timestamp, events after this time are not included.
  Stop time.Time `json:"stop"`

  // Field is the field name.
  Field string `json:"field"`
}

// GetNumericFieldStatsOutput params.
type GetNumericFieldStatsOutput struct {
  // Min is the min value.
  Min float64 `json:"min"`

  // Avg is the avg value.
  Avg float64 `json:"avg"`

  // Max is The max value.
  Max float64 `json:"max"`
}

// GetStringFieldStatsInput params.
type GetStringFieldStatsInput struct {
  // Timeout is a request timeout in seconds, after which a timeout error is returned.
  Timeout int `json:"timeout"`

  // ProjectID is the project id.
  ProjectID string `json:"project_id"`

  // Start is the start timestamp, events before this time are not included.
  Start time.Time `json:"start"`

  // Stop is the stop timestamp, events after this time are not included.
  Stop time.Time `json:"stop"`

  // Field is the field name.
  Field string `json:"field"`

  // Limit is the maximum number of values to return.
  Limit int `json:"limit"`
}

// GetStringFieldStatsOutput params.
type GetStringFieldStatsOutput struct {
  // Values is the string values.
  Values []StringFieldStat `json:"values"`
}

// QueryInput params.
type QueryInput struct {
  // Timeout is a request timeout in seconds, after which a timeout error is returned.
  Timeout int `json:"timeout"`

  // ProjectID is the project id.
  ProjectID string `json:"project_id"`

  // Query is the SQL query string.
  Query string `json:"query"`
}

// QueryOutput params.
type QueryOutput struct {
  // Results is the query results.
  Results []map[string]interface{} `json:"results"`

  // Stats is the query statistics.
  Stats QueryStats `json:"stats"`
}

// SearchInput params.
type SearchInput struct {
  // Timeout is a request timeout in seconds, after which a timeout error is returned.
  Timeout int `json:"timeout"`

  // ProjectID is the project id.
  ProjectID string `json:"project_id"`

  // Start is the start timestamp, events before this time are not included.
  Start time.Time `json:"start"`

  // Stop is the stop timestamp, events after this time are not included.
  Stop time.Time `json:"stop"`

  // Limit is the maxmimum number of events to return.
  Limit int `json:"limit"`

  // Query is the event search query string.
  Query string `json:"query"`
}

// SearchOutput params.
type SearchOutput struct {
  // Results is the query results.
  Results []map[string]interface{} `json:"results"`

  // Stats is the query statistics.
  Stats QueryStats `json:"stats"`
}

