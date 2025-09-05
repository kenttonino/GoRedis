# The Role of a Worker

> - https://temporal.talentlms.com/unit/view/id:2118
> - Workers execute your Workflow code.
> - The Worker itself is provided by the Temporal SDK, but your application will include code to configure and run it.
> - When that code executes, the Worker establishes a persistent connection to the Temporal Cluster and begins polling a Task Queue
> - Since Workers execute your code, any Workflows you execute will make no progress unless one Worker is running.

<br />
<br />
<br />



# Worker Initialization Code

```go
import (
    "app"
    "log"
    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"
)

func main() {
    c, err := client.Dial(client.Options{})
    if err != nil {
        log.Fatalln("Unable to create client", err)
    }
    defer c.Close()

    w := worker.New(c, "greeting-tasks", worker.Options{})

    w.RegisterWorkflow(app.GreetSomeone)

    err = w.Run(worker.InterruptCh())
    if err != nil {
        log.Fatalln("Unable to start worker", err)
    }
}
```

<br />
<br />
<br />
