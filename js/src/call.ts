
import fetch from 'node-fetch'

/**
 * Call method with params via a POST request.
 */

export default function(url: String, method: String, params?: any): Promise<string> {
  return fetch(url + '/' + method, {
    method: 'POST',
    body: JSON.stringify(params),
    headers: {
      'Content-Type': 'application/json'
    }
  }).then(res => res.text())
}