
# JS

The js client for Apex Logs supporting Node.js and the browser.

## Example

```js
import Client from 'apex-logs'

const client = new Client({
  url: 'http://localhost:5001'
})

async function run() {
  const { projects } = await client.getProjects()
  // ...
  
  const { alerts } = await client.getAlerts({ project_id: 'production' })
  // ...
}

run()
```

