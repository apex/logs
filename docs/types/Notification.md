# Notification

The `Notification` represents an alert notification.

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`id` | __string__ | The notification id.
`project_id` | __string__ | The associated project id.
`name` | __string__ | The name of the notification.
`type` | __string__ | The type of notification.
`slack_webhook_url` | __string__ | The Slack webhook URL.
`slack_channel` | __string__ | The Slack channel name, otherwise the default for the webhook is used.
`webhook_url` | __string__ | The webhook URL which receives the alert payloads.
`sms_numbers` | __array__ of __string__ | The receipients of the alert notifications.
`email_addresses` | __array__ of __string__ | The receipients of the alert notifications.
`pagerduty_service_key` | __string__ | The PagerDuty service key.
`updated_at` | __timestamp__ | A timestamp indicating when the notification was last updated.
`created_at` | __timestamp__ | A timestamp indicating when the notification was created.
