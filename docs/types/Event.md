# Event

The `Event` represents a single log event.

__Name__ | __Type__ | __Description__
--- | --- | --- | 
`id` | __string__ | The event id.
`level` | __string__ | The severity level.
`message` | __string__ | The log message.
`fields` | __object__ | The log fields.
`timestamp` | __timestamp__ | The creation timestamp.

## Examples

An example alert reporting event. This event searchable with `message = "reporting complete"`, or any of the user-defined fields, for example `aws.region = "us-west-2"`.

```json
{
  "fields": {
    "action": "email",
    "action_value": "tj@apex.sh",
    "alert_id": 525,
    "aws": {
      "log": {
        "group": "/aws/lambda/vitals_alert_reporter",
        "stream": "2020/01/14/[59]52bfc9d224d644dc963e3ea1cd7590be"
      },
      "region": "us-west-2"
    },
    "check_id": 21802,
    "function": "alert_reporter",
    "triggered": true,
    "value": 2,
    "version": "59"
  },
  "id": "1WOCGwjnBdxHiD9HEduUAxgVEbC",
  "level": "info",
  "message": "reporting complete",
  "timestamp": "2020-01-14T14:12:26.779999971Z"
}
```

