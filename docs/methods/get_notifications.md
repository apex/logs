# get_notifications

The `get_notifications` method returns all notifications.

  Inputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`project_id` | __string__ | The project id.

  Outputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`notifications` | __array__ of [Notification](../types/Notification.md) | The notifications.

## Example

Input:

```json
{
  "project_id": "ping_production"
}
```

Output:

```json
{
  "notifications": [
    {
      "created_at": "2019-08-28T14:24:02.531785Z",
      "email_addresses": [
        "tj@apex.sh"
      ],
      "id": "1Q3bUUxSUZVlreTTwVzdK5z37iX",
      "name": "Email backend team",
      "project_id": "ping_production",
      "type": "email",
      "updated_at": "2019-08-28T14:24:02.531785Z"
    },
    {
      "created_at": "0001-01-01T00:00:00Z",
      "id": "1RhHf76rlKzNQ25fPhRlI0Qfvf8",
      "name": "Send Slack message",
      "project_id": "ping_production",
      "slack_channel": "#alerts",
      "slack_webhook_url": "https://hooks.slack.com/services/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",
      "type": "slack",
      "updated_at": "2019-12-12T13:01:10.751508Z"
    }
  ]
}
```


