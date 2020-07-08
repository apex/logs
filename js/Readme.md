
# JS

The JavaScript client for Apex Logs supporting Node.js, Deno, and the browser.

## Example

```js
import Client from 'apex-logs'

const client = new Client({
  url: '<ENDPOINT>',
  authToken: '<TOKEN>'
})

async function run() {
  const { projects } = await client.getProjects()
  // ...
  
  const { alerts } = await client.getAlerts({ project_id: 'production' })
  // ...
}

run()
```

