// QueryStats represents query statistics.
interface QueryStats {
  // total_bytes_processed is the total number of bytes processed by the query.
  total_bytes_processed: number

  // total_bytes_billed is the total number of bytes billed by the query.
  total_bytes_billed: number

  // cache_hit is a boolean indicating if the query was cached.
  cache_hit: boolean
}

// TimeseriesPoint represents a single point in a timeseries query.
interface TimeseriesPoint {
  // timestamp is the bucket timestamp.
  timestamp: Date

  // count is the number of events for this bucket.
  count: number
}

// LevelTimeseriesPoint represents a single point in a level timeseries query.
interface LevelTimeseriesPoint {
  // timestamp is the bucket timestamp.
  timestamp: Date

  // level is the severity level.
  level: number

  // count is the number of events for this bucket.
  count: number
}

// DiscoveredField represents a single discovered field.
interface DiscoveredField {
  // name is the field name.
  name: string

  // type is the type of discovered field.
  type: string

  // count is the number of times this field occurred in the sampled events.
  count: number

  // percent is the percentage of occurrences in the sampled events.
  percent: number
}

// StringFieldStat represents a string field's stats.
interface StringFieldStat {
  // value is the string value.
  value: string

  // count is the number of times this field occurred in the sampled events.
  count: number

  // percent is the percentage of occurrences in the sampled events.
  percent: number
}

// Project represents a customer application.
interface Project {
  // id is the project id.
  id: string

  // name is the human-friendly project name.
  name: string

  // retention is the retention of log events in days. When zero the logs do not expire.
  retention: number

  // location is the geographical location where the log events are stored.
  location: string

  // description is the project description.
  description?: string

  // updated_at is a timestamp indicating when the project was last updated.
  updated_at: Date

  // created_at is a timestamp indicating when the project was created.
  created_at: Date
}

// Notification represents an alert notification.
interface Notification {
  // id is the notification id.
  id: string

  // project_id is the associated project id.
  project_id: string

  // name is the name of the notification.
  name: string

  // type is the type of notification.
  type: string

  // slack_webhook_url is the Slack webhook URL.
  slack_webhook_url?: string

  // slack_channel is the Slack channel name, otherwise the default for the webhook is used.
  slack_channel?: string

  // webhook_url is the webhook URL which receives the alert payloads.
  webhook_url?: string

  // sms_numbers is the receipients of the alert notifications.
  sms_numbers?: string[]

  // email_addresses is the receipients of the alert notifications.
  email_addresses?: string[]

  // pagerduty_service_key is the PagerDuty service key.
  pagerduty_service_key?: string

  // updated_at is a timestamp indicating when the notification was last updated.
  updated_at: Date

  // created_at is a timestamp indicating when the notification was created.
  created_at: Date
}

// Alert represents configuration for performing alerting.
interface Alert {
  // id is the alert id.
  id: string

  // project_id is the associated project id.
  project_id: string

  // name is the name of the alert.
  name: string

  // description is the description of the alert.
  description: string

  // severity is the severity of the alert.
  severity: string

  // query is the query performed by the alert.
  query: string

  // operator is the query performed by the alert.
  operator: string

  // threshold is the threshold for comparison against the selected operator.
  threshold: number

  // limit is the maximum number of events in the alert notification.
  limit: number

  // interval is the interval in minutes for performing the alert.
  interval: number

  // notification_id is the notification id for reporting alerts, when omitted the alert will not be run.
  notification_id: string

  // triggered is a boolean indicating whether or not the alert is currently triggered.
  triggered: boolean

  // muted is a boolean used ignore trigger and resolve notifications.
  muted: boolean

  // updated_at is a timestamp indicating when the alert was last updated.
  updated_at: Date

  // created_at is a timestamp indicating when the alert was created.
  created_at: Date
}

// Event represents a single log event.
interface Event {
  // id is the event id.
  id: string

  // level is the severity level.
  level: string

  // message is the log message.
  message: string

  // fields is the log fields.
  fields: { [name: string]: any }

  // timestamp is the creation timestamp.
  timestamp: Date
}

// AddEventsInput params.
interface AddEventsInput {
  // project_id is the project id.
  project_id: string

  // events is the batch of events.
  events: Event[]
}

// AddAlertInput params.
interface AddAlertInput {
  // alert is the alert.
  alert: Alert
}

// AddAlertOutput params.
class AddAlertOutput {
  // id is the alert id.
  id: string
}

// TestAlertInput params.
interface TestAlertInput {
  // alert is the alert.
  alert: Alert
}

// UpdateAlertInput params.
interface UpdateAlertInput {
  // alert is the alert.
  alert: Alert
}

// RemoveAlertInput params.
interface RemoveAlertInput {
  // project_id is the project id.
  project_id: string

  // alert_id is the alert id.
  alert_id: string
}

// GetAlertInput params.
interface GetAlertInput {
  // project_id is the project id.
  project_id: string

  // alert_id is the alert id.
  alert_id: string
}

// GetAlertOutput params.
class GetAlertOutput {
  // alert is the alert.
  alert: Alert
}

// GetAlertsInput params.
interface GetAlertsInput {
  // project_id is the project id.
  project_id: string
}

// GetAlertsOutput params.
class GetAlertsOutput {
  // alerts is the alerts.
  alerts: Alert[]
}

// AddNotificationInput params.
interface AddNotificationInput {
  // notification is the notification.
  notification: Notification
}

// AddNotificationOutput params.
class AddNotificationOutput {
  // id is the notification id.
  id: string
}

// UpdateNotificationInput params.
interface UpdateNotificationInput {
  // notification is the notification.
  notification: Notification
}

// RemoveNotificationInput params.
interface RemoveNotificationInput {
  // project_id is the project id.
  project_id: string

  // notification_id is the notification id.
  notification_id: string
}

// GetNotificationInput params.
interface GetNotificationInput {
  // project_id is the project id.
  project_id: string

  // notification_id is the notification id.
  notification_id: string
}

// GetNotificationOutput params.
class GetNotificationOutput {
  // notification is the notification.
  notification: Notification
}

// GetNotificationsInput params.
interface GetNotificationsInput {
  // project_id is the project id.
  project_id: string
}

// GetNotificationsOutput params.
class GetNotificationsOutput {
  // notifications is the notifications.
  notifications: Notification[]
}

// AddProjectInput params.
interface AddProjectInput {
  // project is the project.
  project: Project
}

// AddProjectOutput params.
class AddProjectOutput {
  // id is the project id.
  id: string
}

// UpdateProjectInput params.
interface UpdateProjectInput {
  // project is the project.
  project: Project
}

// RemoveProjectInput params.
interface RemoveProjectInput {
  // project_id is the project id.
  project_id: string
}

// GetProjectsOutput params.
class GetProjectsOutput {
  // projects is the projects.
  projects: Project[]
}

// GetProjectStatsInput params.
interface GetProjectStatsInput {
  // project_id is the project id.
  project_id: string
}

// GetProjectStatsOutput params.
class GetProjectStatsOutput {
  // bytes_total is the total number of bytes stored.
  bytes_total: number

  // events_total is the total number of events stored.
  events_total: number
}

// GetDiscoveredFieldsInput params.
interface GetDiscoveredFieldsInput {
  // timeout is a request timeout in seconds, after which a timeout error is returned.
  timeout: number

  // project_id is the project id.
  project_id: string

  // start is the start timestamp, events before this time are not included.
  start: Date

  // stop is the stop timestamp, events after this time are not included.
  stop: Date

  // query is the event search query string.
  query: string
}

// GetDiscoveredFieldsOutput params.
class GetDiscoveredFieldsOutput {
  // fields is the fields discovered.
  fields: DiscoveredField[]

  // stats is the query statistics.
  stats: QueryStats
}

// GetTimeseriesInput params.
interface GetTimeseriesInput {
  // timeout is a request timeout in seconds, after which a timeout error is returned.
  timeout: number

  // project_id is the project id.
  project_id: string

  // start is the start timestamp, events before this time are not included.
  start: Date

  // stop is the stop timestamp, events after this time are not included.
  stop: Date

  // query is the SQL query string.
  query: string

  // max_points is the maxmimum number of datapoints to return.
  max_points: number
}

// GetTimeseriesOutput params.
class GetTimeseriesOutput {
  // points is the series.
  points: TimeseriesPoint[]

  // stats is the query statistics.
  stats: QueryStats
}

// GetLevelTimeseriesInput params.
interface GetLevelTimeseriesInput {
  // timeout is a request timeout in seconds, after which a timeout error is returned.
  timeout: number

  // project_id is the project id.
  project_id: string

  // start is the start timestamp, events before this time are not included.
  start: Date

  // stop is the stop timestamp, events after this time are not included.
  stop: Date

  // query is the SQL query string.
  query: string

  // max_points is the maxmimum number of datapoints to return.
  max_points: number
}

// GetLevelTimeseriesOutput params.
class GetLevelTimeseriesOutput {
  // points is the series of datapoints.
  points: LevelTimeseriesPoint[]

  // stats is the query statistics.
  stats: QueryStats
}

// GetNumericFieldStatsInput params.
interface GetNumericFieldStatsInput {
  // timeout is a request timeout in seconds, after which a timeout error is returned.
  timeout: number

  // project_id is the project id.
  project_id: string

  // start is the start timestamp, events before this time are not included.
  start: Date

  // stop is the stop timestamp, events after this time are not included.
  stop: Date

  // field is the field name.
  field: string
}

// GetNumericFieldStatsOutput params.
class GetNumericFieldStatsOutput {
  // min is the min value.
  min: number

  // avg is the avg value.
  avg: number

  // max is The max value.
  max: number

  // stats is the query statistics.
  stats: QueryStats
}

// GetStringFieldStatsInput params.
interface GetStringFieldStatsInput {
  // timeout is a request timeout in seconds, after which a timeout error is returned.
  timeout: number

  // project_id is the project id.
  project_id: string

  // start is the start timestamp, events before this time are not included.
  start: Date

  // stop is the stop timestamp, events after this time are not included.
  stop: Date

  // field is the field name.
  field: string

  // limit is the maximum number of values to return.
  limit: number
}

// GetStringFieldStatsOutput params.
class GetStringFieldStatsOutput {
  // values is the string values.
  values: StringFieldStat[]

  // stats is the query statistics.
  stats: QueryStats
}

// QueryInput params.
interface QueryInput {
  // timeout is a request timeout in seconds, after which a timeout error is returned.
  timeout: number

  // project_id is the project id.
  project_id: string

  // query is the SQL query string.
  query: string
}

// QueryOutput params.
class QueryOutput {
  // results is the query results.
  results: { [name: string]: any }[]

  // stats is the query statistics.
  stats: QueryStats
}

// SearchInput params.
interface SearchInput {
  // timeout is a request timeout in seconds, after which a timeout error is returned.
  timeout: number

  // project_id is the project id.
  project_id: string

  // start is the start timestamp, events before this time are not included.
  start: Date

  // stop is the stop timestamp, events after this time are not included.
  stop: Date

  // limit is the maxmimum number of events to return.
  limit: number

  // query is the event search query string.
  query: string
}

// SearchOutput params.
class SearchOutput {
  // results is the query results.
  results: { [name: string]: any }[]

  // stats is the query statistics.
  stats: QueryStats
}

// GetCountInput params.
interface GetCountInput {
  // timeout is a request timeout in seconds, after which a timeout error is returned.
  timeout: number

  // project_id is the project id.
  project_id: string

  // start is the start timestamp, events before this time are not included.
  start: Date

  // stop is the stop timestamp, events after this time are not included.
  stop: Date

  // query is the event search query string.
  query: string
}

// GetCountOutput params.
class GetCountOutput {
  // count is the query result count.
  count: number

  // stats is the query statistics.
  stats: QueryStats
}


import call from './call'

/**
 * Client is the API client.
 */

export class Client {

  private url: string

  /**
   * Initialize.
   */

  constructor(params: { url: string }) {
    this.url = params.url
  }

  /**
   * addEvents: ingested a batch of events.
   */

  async addEvents(params: AddEventsInput) {
    await call(this.url, 'add_events', params)
  }

  /**
   * addAlert: creates a new alert.
   */

  async addAlert(params: AddAlertInput): Promise<AddAlertOutput> {
    let res = await call(this.url, 'add_alert', params)
    let out: AddAlertOutput = JSON.parse(res)
    return out
  }

  /**
   * testAlert: test the alert configuration.
   */

  async testAlert(params: TestAlertInput) {
    await call(this.url, 'test_alert', params)
  }

  /**
   * updateAlert: updates an alert.
   */

  async updateAlert(params: UpdateAlertInput) {
    await call(this.url, 'update_alert', params)
  }

  /**
   * removeAlert: removes an alert.
   */

  async removeAlert(params: RemoveAlertInput) {
    await call(this.url, 'remove_alert', params)
  }

  /**
   * getAlert: returns an alert.
   */

  async getAlert(params: GetAlertInput): Promise<GetAlertOutput> {
    let res = await call(this.url, 'get_alert', params)
    let out: GetAlertOutput = JSON.parse(res)
    return out
  }

  /**
   * getAlerts: returns all alerts in a project.
   */

  async getAlerts(params: GetAlertsInput): Promise<GetAlertsOutput> {
    let res = await call(this.url, 'get_alerts', params)
    let out: GetAlertsOutput = JSON.parse(res)
    return out
  }

  /**
   * addNotification: creates a new notification.
   */

  async addNotification(params: AddNotificationInput): Promise<AddNotificationOutput> {
    let res = await call(this.url, 'add_notification', params)
    let out: AddNotificationOutput = JSON.parse(res)
    return out
  }

  /**
   * updateNotification: updates a notification.
   */

  async updateNotification(params: UpdateNotificationInput) {
    await call(this.url, 'update_notification', params)
  }

  /**
   * removeNotification: removes a notification.
   */

  async removeNotification(params: RemoveNotificationInput) {
    await call(this.url, 'remove_notification', params)
  }

  /**
   * getNotification: returns a notification.
   */

  async getNotification(params: GetNotificationInput): Promise<GetNotificationOutput> {
    let res = await call(this.url, 'get_notification', params)
    let out: GetNotificationOutput = JSON.parse(res)
    return out
  }

  /**
   * getNotifications: returns all notifications.
   */

  async getNotifications(params: GetNotificationsInput): Promise<GetNotificationsOutput> {
    let res = await call(this.url, 'get_notifications', params)
    let out: GetNotificationsOutput = JSON.parse(res)
    return out
  }

  /**
   * addProject: creates a new project.
   */

  async addProject(params: AddProjectInput): Promise<AddProjectOutput> {
    let res = await call(this.url, 'add_project', params)
    let out: AddProjectOutput = JSON.parse(res)
    return out
  }

  /**
   * updateProject: updates a project.
   */

  async updateProject(params: UpdateProjectInput) {
    await call(this.url, 'update_project', params)
  }

  /**
   * removeProject: removes a project.
   */

  async removeProject(params: RemoveProjectInput) {
    await call(this.url, 'remove_project', params)
  }

  /**
   * getProjects: returns all projects.
   */

  async getProjects(): Promise<GetProjectsOutput> {
    let res = await call(this.url, 'get_projects')
    let out: GetProjectsOutput = JSON.parse(res)
    return out
  }

  /**
   * getProjectStats: returns project statistics.
   */

  async getProjectStats(params: GetProjectStatsInput): Promise<GetProjectStatsOutput> {
    let res = await call(this.url, 'get_project_stats', params)
    let out: GetProjectStatsOutput = JSON.parse(res)
    return out
  }

  /**
   * getDiscoveredFields: returns fields discovered in the provided time range.
   */

  async getDiscoveredFields(params: GetDiscoveredFieldsInput): Promise<GetDiscoveredFieldsOutput> {
    let res = await call(this.url, 'get_discovered_fields', params)
    let out: GetDiscoveredFieldsOutput = JSON.parse(res)
    return out
  }

  /**
   * getTimeseries: returns a timeseries of event counts in the provided time range.
   */

  async getTimeseries(params: GetTimeseriesInput): Promise<GetTimeseriesOutput> {
    let res = await call(this.url, 'get_timeseries', params)
    let out: GetTimeseriesOutput = JSON.parse(res)
    return out
  }

  /**
   * getLevelTimeseries: returns a timeseries of event counts in the provided time range grouped by severity level.
   */

  async getLevelTimeseries(params: GetLevelTimeseriesInput): Promise<GetLevelTimeseriesOutput> {
    let res = await call(this.url, 'get_level_timeseries', params)
    let out: GetLevelTimeseriesOutput = JSON.parse(res)
    return out
  }

  /**
   * getNumericFieldStats: returns field statistics for a numeric field.
   */

  async getNumericFieldStats(params: GetNumericFieldStatsInput): Promise<GetNumericFieldStatsOutput> {
    let res = await call(this.url, 'get_numeric_field_stats', params)
    let out: GetNumericFieldStatsOutput = JSON.parse(res)
    return out
  }

  /**
   * getStringFieldStats: returns field statistics for a string field.
   */

  async getStringFieldStats(params: GetStringFieldStatsInput): Promise<GetStringFieldStatsOutput> {
    let res = await call(this.url, 'get_string_field_stats', params)
    let out: GetStringFieldStatsOutput = JSON.parse(res)
    return out
  }

  /**
   * query: performs a SQL query against the log events.
   */

  async query(params: QueryInput): Promise<QueryOutput> {
    let res = await call(this.url, 'query', params)
    let out: QueryOutput = JSON.parse(res)
    return out
  }

  /**
   * search: performs a search query against the log events.
   */

  async search(params: SearchInput): Promise<SearchOutput> {
    let res = await call(this.url, 'search', params)
    let out: SearchOutput = JSON.parse(res)
    return out
  }

  /**
   * getCount: performs a search query against the log events, returning the number of matches.
   */

  async getCount(params: GetCountInput): Promise<GetCountOutput> {
    let res = await call(this.url, 'get_count', params)
    let out: GetCountOutput = JSON.parse(res)
    return out
  }

}
