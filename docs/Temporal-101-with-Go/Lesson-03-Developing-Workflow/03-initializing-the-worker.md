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



# Worker Initialization Explanation

There are typically three things you need in order to configure a Worker:
#

```plaintext
1. Temporal Client
    - Used to communicate with the Temporal Cluster.

2. Name of Task Queue
    - Maintained by the Temporal Server and polled by the Worker.

3. Fully Qualified Name of the Workflow Definition Function
    - Used in the call to RegisterWorkflow.
```

> - Every Workflow Definition function must be registered with at least one Worker for execution to proceed, but you may register multiple of theses functions with any given Worker.
> - Once you call the Run function of Worker, the Worker will then begin a "long poll" on the specified task queue.
> - If you start the Worker from a terminal using a program like the one shown above, don't be surprised if you see nothing more than a few lines of output.
> - It is the expected behavior and the program isn't stuck, it's just busy polling the task queue and working on the tasks that it has accepted from the Temporal Client.
