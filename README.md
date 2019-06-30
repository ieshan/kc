# KC
Basic kubernates commands in a easier way

## Install KC

### Prerequisite

- Need to set `$GOPATH` in your computer
- Add `$GOPATH/bin` in your `$PATH` variable

### Install

```
go get -u github.com/ieshan/kc
```

### Check installation
Run the following command
```
kc help
```
The output should be
```
Basic kubernetes commands

Usage:
  kc [command]

Available Commands:
  context     Switch context
  deploy      Returns a list of deployments
  describe    Describe pod
  dp          Delete pod
  help        Help about any command
  logs        Pod log
  ns          Work with kubernates namespace
  pf          Forward port: example: kc pf {{pod}} {{local_port}}:{{pod_port}}
  pods        Returns a list of pods
  sh          Access pod shell
  svc         Returns a list of svc

Flags:
  -h, --help               help for kc
  -n, --namespace string   namespace (default "aunty")

Use "kc [command] --help" for more information about a command.
```

## Usage

### Switch context
* Run the following command
```
kc context
```
* Use the arrow keys on the keyboard to select context when prompted with list of available contexts
```
Use the arrow keys to navigate: ↓ ↑ → ←
? Select Default Context (minikube)::
  ▸ dev-context
    stating-context
    production-context
    minikube
```

### Manage namespaces
#### Add a namespace
Run the following command, replace the `{{namespace}}` with the actual namespace.
```
kc ns add {{namespace}}
```
#### Remove a namespace
* Run the following command.
```
kc ns delete
```
* Use the arrow keys on the keyboard to select the namespace you want to delete from the list of prompted namespaces.
```
Use the arrow keys to navigate: ↓ ↑ → ←
? Select namespace to delete::
  ▸ default
    namespace-1
    namespace-2
```

#### Select a namespace as default
* Run the following command.
```
kc ns select
```
* Use the arrow keys on the keyboard to select the namespace you want to mask as default so that you don't have to pass namespace flag with every other command.
```
Use the arrow keys to navigate: ↓ ↑ → ←
? Select Default Namespace (default)::
  ▸ default
    namespace-1
    namespace-2
```

### Get the list of deployments
* Run the following commands to get the list of deployments
```
kc deploy
```
* It will output the list of deployment under selected namespace
```
NAME                       READY   UP-TO-DATE   AVAILABLE   AGE
app-1                      3/3     3            3           16d
app-2                      0/0     0            0           64d
app-3                      1/1     1            1           1d
```

### Get the list of services
* Run the following command, replace the `{{pod}}` with the actual pod name.
```
kc svc {{pod}}
```
* It will output the list of services under selected namespace.
```
NAME                        TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
service-1                   ClusterIP      10.4.167.55    <none>        80/TCP           6d
service-2                   ClusterIP      10.4.168.21    <none>        8088/TCP         7d
service-3                   LoadBalancer   10.4.160.66    20.168.3.5    80:30053/TCP     8d
```

### Get the list of pods
* Run the following command.
```
kc pods
```
* It will output the list of pods under selected namespace.
```
NAME                                     READY   STATUS    RESTARTS   AGE
pod-0                                    1/1     Running   0          81d
pod-1                                    1/1     Running   1          8d
pod-2-2a29ff1-bfg12                      1/1     Running   5          5d
pod-2-2a29ff1-2rkgr                      1/1     Running   3          7d
```

### Describe a pod
* Run the following command, replace the `{{pod}}` with the actual pod name.
```
kc describe {{pod}}
```

### Delete a pod
* Run the following command, replace the `{{pod}}` with the actual pod name.
```
kc dp {{pod}}
```

### Get logs of a pod
* Run the following command, replace the `{{pod}}` with the actual pod name.
```
kc logs {{pod}}
```

### Forward port
Run the following command, replace the `{{pod}}` with the actual pod name, `{{local_port}}` with local machine port number and `{{pod_port}}` with the pod's port number.
```
kc pf {{pod}} {{local_port}}:{{pod_port}}
```

### Access a pod's shell
Run the following command, replace the `{{pod}}` with the actual pod name.
```
kc sh {{pod}}
```
