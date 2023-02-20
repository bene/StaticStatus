# StaticStatus

StatusStatus is the most easy way to spin up a status page. There is no need for
a database or a backend at all. All you need to do is to clone this repo and
deploy it via Netlify.

## Deployment

### Environment variables

All of the following environment variables are required.

| Name                    | Description                                                                                                          | Value                        |
| ----------------------- | -------------------------------------------------------------------------------------------------------------------- | ---------------------------- |
| `CHECK_URL`             | The URL to check the status.                                                                                         | `https://app1.example.com`   |
| `SELF_URL`              | The URL of the status page deployment.                                                                               | `https://status.example.com` |
| `INITIAL_BUILD`         | Must be `true` for first build and then set to `false`. When set to `true` all previous status entries will be lost. | `true` -> `false`            |
| `BUILD_WEBHOOK`         | A build hook URL to redeploy the status page.                                                                        | `https://api.netlify.com`    |
| `AWS_LAMBDA_JS_RUNTIME` | Static value to execute Lambda function in Node 18.                                                                  | `nodejs18.x`                 |

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/bene/StaticStatus)

## Concept

This kinda works like blockchain. A scheduled Netlify functions redeploys the
status page every minute. During the build process not only the current status
is checked but also the status history is fetched from the currently deployed
status page. The new build attaches the new status to the history and deploys
the next history.

## Cost

The price for this 0-backend status page is payed in Netlify build minutes. A
single deploy has an average build time of 9 seconds which results in the
following prices:

| Interval  | Price        |
| --------- | ------------ |
| 1 minute  | 6570 minutes |
| 5 minutes | 1314 minutes |

## Disclaimer

This project has been created to deploy a status page in the quickest, most
easiest way possible. If you want to consistently monitor your app over a longer
period of time, a traditional status page service with a backend is probably the
way to go.

## Development

### Simulate deploy

> go run ./src/main.go

### Build styles

To build the stylesheets use the
[Tailwind CLI](https://tailwindcss.com/blog/standalone-cli).

> tailwindcss -i ./src/main.css -o ./static/main.css
