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
