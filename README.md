# dapr-http-cron-handler-template

[![Test](https://github.com/mchmarny/dapr-http-cron-handler-template/workflows/Test/badge.svg)](https://github.com/mchmarny/dapr-http-cron-handler-template/actions?query=workflow%3ATest) ![Release](https://github.com/mchmarny/dapr-http-cron-handler-template/workflows/Release/badge.svg?query=workflow%3ARelease) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mchmarny/dapr-http-cron-handler-template) [![Go Report Card](https://goreportcard.com/badge/github.com/mchmarny/dapr-http-cron-handler-template)](https://goreportcard.com/report/github.com/mchmarny/dapr-http-cron-handler-template)


## template usage 

* Click "Use this template" above and follow the wizard to select owner and name your new repo
* Clone and navigate to your new repo (`git clone git@github.com:<USERNAME>/<REPO-NAME>.git && cd <REPO-NAME>`)
* Initialize your project to set the package names and update imports (`make init`)
* Write your subscription event handling logic 

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
$ make help
tidy                           Updates the go modules and vendors all dependencies
test                           Tests the entire project
build                          Builds local release binary
debug                          Runs uncompiled code it in Dapr in debug mode
run                            Builds binary and runs it in Dapr
image                          Builds and publish docker image
lint                           Lints the entire project
tag                            Creates release tag
clean                          Cleans up generated files
reset                          Resets go modules
help                           Display available commands
```

This project also includes GitHub actions in [.github/workflows](.github/workflows) that test on each `push` and build images and mark release on each `tag`. Other Dapr GitHub templates to accelerate Dapr development available [here](https://github.com/dapr/go-sdk/tree/master/service).

## License

This software is released under the [MIT](./LICENSE)
