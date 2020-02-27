# get_notification

The `get_notification` method returns a notification.

  Inputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`notification_id` | __string__ | The notification id.
`project_id` | __string__ | The project id.

  Outputs:

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`notification` | [Notification](../types/Notification.md) | The notification.

## Example

Input:

```json
{
  "notification_id": "1Q3bUUxSUZVlreTTwVzdK5z37iX",
  "project_id": "ping_production"
}
```

Output:

```json
{
  "notification": {
    "created_at": "2019-08-28T14:24:02.531785Z",
    "email_addresses": [
      "ops@apex.sh"
    ],
    "id": "1Q3bUUxSUZVlreTTwVzdK5z37iX",
    "name": "Email backend team",
    "project_id": "ping_production",
    "type": "email",
    "updated_at": "2019-08-28T14:24:02.531785Z"
  }
}
```


