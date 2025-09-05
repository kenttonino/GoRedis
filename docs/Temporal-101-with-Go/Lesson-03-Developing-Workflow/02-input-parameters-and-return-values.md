# Values Must Be Serializable

> - https://temporal.talentlms.com/unit/view/id:2117
> - In order for Temporal to store the Workflow's input and output, data used in input parameters and return values must be serializable.
> - By default, Temporal can handle null or binary values, as well as any data that can be serialized using Go's support for JSON or Protocol Buffers.
> - This means that most of the types you'd typically use in a function, such as integers and floating point numbers, boolean values, and strings, are all handled automatically, as are structs composed from these types, but types such as channels, functions, or unsafe pointers are prohibited as either input parameters or return values.

<br />
<br />
<br />



# Data Confidentiality

> - Although the input parameters and return values are stored as part of the Event History of your Workflow Executions, you can create a custom Data Converter to encrypt the data as it enters the Temporal Cluster and decrypt it upon exit, thereby maintaining the confidentiality of any sensitive data used as input or output of your applications.

<br />
<br />
<br />



# Avoid Passing Large Amounts of Data

> - Because the Event History contains the input and output, which is also sent across the network from the application of the Temporal Cluster, you'll have better performance if you limit the amount of data sent.
> - For example, imagine you've created a Workflow that will convert audio files from one format to another.
> - It would be much better to pass the path or URL for the files as input than to pass the content of the files.
> - To protect against unexpected failures caused by sending or storing too much data, the Temporal Server imposes various limits beyond which it will emit warnings or errors, depending on the severity.
