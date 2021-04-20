### Important Links
https://github.com/operator-framework/operator-sdk-samples/blob/master/go/demo-operator/controllers/demo_controller.go

https://docs.openshift.com/container-platform/4.7/operators/operator_sdk/golang/osdk-golang-tutorial.html#osdk-golang-tutorial

### Deploy in Cluster
make docker-build IMG=anupamgogoi/demo-operator:latest
make docker-push IMG=anupamgogoi/demo-operator:latest
make deploy IMG=anupamgogoi/demo-operator:latest