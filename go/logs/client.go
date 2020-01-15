package logs

// Client is the API client.
type Client struct {
  // URL is the required API endpoint address.
  URL string
}

// AddToken creates a new token.
func (c *Client) AddToken(in AddTokenInput) (*AddTokenOutput, error) {
  var out AddTokenOutput
  return &out, call(c.URL, "add_token", in, &out)
}

// UpdateToken updates an token.
func (c *Client) UpdateToken(in UpdateTokenInput) error {
  return call(c.URL, "update_token", in, nil)
}

// RemoveToken removes an token.
func (c *Client) RemoveToken(in RemoveTokenInput) error {
  return call(c.URL, "remove_token", in, nil)
}

// GetTokens returns all tokens in a project.
func (c *Client) GetTokens(in GetTokensInput) (*GetTokensOutput, error) {
  var out GetTokensOutput
  return &out, call(c.URL, "get_tokens", in, &out)
}

// AddEvents ingests a batch of events.
func (c *Client) AddEvents(in AddEventsInput) error {
  return call(c.URL, "add_events", in, nil)
}

// AddAlert creates a new alert.
func (c *Client) AddAlert(in AddAlertInput) (*AddAlertOutput, error) {
  var out AddAlertOutput
  return &out, call(c.URL, "add_alert", in, &out)
}

// TestAlert test the alert configuration.
func (c *Client) TestAlert(in TestAlertInput) error {
  return call(c.URL, "test_alert", in, nil)
}

// UpdateAlert updates an alert.
func (c *Client) UpdateAlert(in UpdateAlertInput) error {
  return call(c.URL, "update_alert", in, nil)
}

// RemoveAlert removes an alert.
func (c *Client) RemoveAlert(in RemoveAlertInput) error {
  return call(c.URL, "remove_alert", in, nil)
}

// GetAlert returns an alert.
func (c *Client) GetAlert(in GetAlertInput) (*GetAlertOutput, error) {
  var out GetAlertOutput
  return &out, call(c.URL, "get_alert", in, &out)
}

// GetAlerts returns all alerts in a project.
func (c *Client) GetAlerts(in GetAlertsInput) (*GetAlertsOutput, error) {
  var out GetAlertsOutput
  return &out, call(c.URL, "get_alerts", in, &out)
}

// AddNotification creates a new notification.
func (c *Client) AddNotification(in AddNotificationInput) (*AddNotificationOutput, error) {
  var out AddNotificationOutput
  return &out, call(c.URL, "add_notification", in, &out)
}

// UpdateNotification updates a notification.
func (c *Client) UpdateNotification(in UpdateNotificationInput) error {
  return call(c.URL, "update_notification", in, nil)
}

// RemoveNotification removes a notification.
func (c *Client) RemoveNotification(in RemoveNotificationInput) error {
  return call(c.URL, "remove_notification", in, nil)
}

// GetNotification returns a notification.
func (c *Client) GetNotification(in GetNotificationInput) (*GetNotificationOutput, error) {
  var out GetNotificationOutput
  return &out, call(c.URL, "get_notification", in, &out)
}

// GetNotifications returns all notifications.
func (c *Client) GetNotifications(in GetNotificationsInput) (*GetNotificationsOutput, error) {
  var out GetNotificationsOutput
  return &out, call(c.URL, "get_notifications", in, &out)
}

// AddProject creates a new project.
func (c *Client) AddProject(in AddProjectInput) (*AddProjectOutput, error) {
  var out AddProjectOutput
  return &out, call(c.URL, "add_project", in, &out)
}

// UpdateProject updates a project.
func (c *Client) UpdateProject(in UpdateProjectInput) error {
  return call(c.URL, "update_project", in, nil)
}

// RemoveProject removes a project.
func (c *Client) RemoveProject(in RemoveProjectInput) error {
  return call(c.URL, "remove_project", in, nil)
}

// GetProjects returns all projects.
func (c *Client) GetProjects() (*GetProjectsOutput, error) {
  var out GetProjectsOutput
  return &out, call(c.URL, "get_projects", nil, &out)
}

// GetProjectStats returns project statistics.
func (c *Client) GetProjectStats(in GetProjectStatsInput) (*GetProjectStatsOutput, error) {
  var out GetProjectStatsOutput
  return &out, call(c.URL, "get_project_stats", in, &out)
}

// GetDiscoveredFields returns fields discovered in the provided time range.
func (c *Client) GetDiscoveredFields(in GetDiscoveredFieldsInput) (*GetDiscoveredFieldsOutput, error) {
  var out GetDiscoveredFieldsOutput
  return &out, call(c.URL, "get_discovered_fields", in, &out)
}

// GetTimeseries returns a timeseries of event counts in the provided time range.
func (c *Client) GetTimeseries(in GetTimeseriesInput) (*GetTimeseriesOutput, error) {
  var out GetTimeseriesOutput
  return &out, call(c.URL, "get_timeseries", in, &out)
}

// GetNumericFieldStats returns field statistics for a numeric field.
func (c *Client) GetNumericFieldStats(in GetNumericFieldStatsInput) (*GetNumericFieldStatsOutput, error) {
  var out GetNumericFieldStatsOutput
  return &out, call(c.URL, "get_numeric_field_stats", in, &out)
}

// GetStringFieldStats returns field statistics for a string field.
func (c *Client) GetStringFieldStats(in GetStringFieldStatsInput) (*GetStringFieldStatsOutput, error) {
  var out GetStringFieldStatsOutput
  return &out, call(c.URL, "get_string_field_stats", in, &out)
}

// GetBooleanFieldStats returns field statistics for a boolean field.
func (c *Client) GetBooleanFieldStats(in GetBooleanFieldStatsInput) (*GetBooleanFieldStatsOutput, error) {
  var out GetBooleanFieldStatsOutput
  return &out, call(c.URL, "get_boolean_field_stats", in, &out)
}

// Query performs a SQL query against the log events.
func (c *Client) Query(in QueryInput) (*QueryOutput, error) {
  var out QueryOutput
  return &out, call(c.URL, "query", in, &out)
}

// Search performs a search query against the log events.
func (c *Client) Search(in SearchInput) (*SearchOutput, error) {
  var out SearchOutput
  return &out, call(c.URL, "search", in, &out)
}

// GetCount performs a search query against the log events, returning the number of matches.
func (c *Client) GetCount(in GetCountInput) (*GetCountOutput, error) {
  var out GetCountOutput
  return &out, call(c.URL, "get_count", in, &out)
}

