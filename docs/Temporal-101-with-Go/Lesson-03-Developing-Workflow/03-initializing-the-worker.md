# The Role of a Worker

> - https://temporal.talentlms.com/unit/view/id:2118
> - Workers execute your Workflow code.
> - The Worker itself is provided by the Temporal SDK, but your application will include code to configure and run it.
> - When that code executes, the Worker establishes a persistent connection to the Temporal Cluster and begins polling a Task Queue
> - Since Workers execute your code, any Workflows you execute will make no progress unless one Worker is running.
