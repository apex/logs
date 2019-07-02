
# Go

The Go client for Apex Logs.

## Example

Error handling is omitted for brevity.

```go
c := logs.New("http://localhost:5001")

res, _ := c.GetProjects()

for _, project := range res.Projects {
  stats, _ := c.GetProjectStats(logs.GetProjectStatsInput{
    ProjectID: project.ID,
  })

  fmt.Printf("%s has %d logs\n", project.Name, stats.EventsTotal)
}
```

