
Go client for [Apex Logs](https://apex.sh/logs/).

## Example

Error handling is omitted for brevity.

```go
c := logs.Client{
  URL: "<ENDPOINT>",
  AuthToken: "<TOKEN>",
}

res, _ := c.GetProjects()

for _, project := range res.Projects {
  stats, _ := c.GetProjectStats(logs.GetProjectStatsInput{
    ProjectID: project.ID,
  })

  fmt.Printf("%s has %d logs\n", project.Name, stats.EventsTotal)
}
```

## Resources

To learn more about Apex Logs visit the [documentation](https://apex.sh/docs/logs/), and to contribute to this client visit the [github.com/apex/rpc](https://github.com/apex/rpc/) project which is used to generate this client.
