# dapr-http-cron-handler-template [![Go Report Card](https://goreportcard.com/badge/github.com/dapr-templates/dapr-http-cron-handler-template)](https://goreportcard.com/report/github.com/dapr-templates/dapr-http-cron-handler-template)


### template usage 

* Click "Use this template" above and follow the wizard to select owner and name your new repo
* When done, clone your new repo, and navigate into it
* Initialize your project to set the package names 
  * `go mod init github.com/<your-github-username>/<your-repo-name>`
* Tidy things up and download modules
  * `go mod tidy`
* You're done, now "just" write your subscription event handling logic ;) 


To edit the binding execution schedule, update the [config/cron.yaml](config/cron.yaml) file

```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: run
spec:
  type: bindings.cron
  metadata:
  - name: schedule
    value: "@every 3s"
```

For more information about this binding see the [Dapr docs](https://github.com/dapr/docs/blob/master/reference/specs/bindings/cron.md)

### common operations

Common operations to help you bootstrap a Dapr HTTP services development in `go`:

```shell
$ make
tidy                           Updates the go modules and vendors all dependencies
test                           Tests the entire project
run                            Runs uncompiled code in Dapr
build                          Builds local release binary
image                          Builds and publish docker image
lint                           Lints the entire project
tag                            Creates release tag
clean                          Cleans up generated files
help                           Display available commands
```

This project also includes GitHub actions in [.github/workflows](.github/workflows) that test on each `push` and build images and mark release on each `tag`. Other Dapr GitHub templates to accelerate Dapr development available [here](https://github.com/dapr/go-sdk/tree/master/service).

### License

This software is released under the [MIT](./LICENSE)
