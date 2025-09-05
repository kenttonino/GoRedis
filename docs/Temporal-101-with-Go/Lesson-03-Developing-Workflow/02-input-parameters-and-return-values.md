# Values Must Be Serializable

> - https://temporal.talentlms.com/unit/view/id:2117
> - In order for Temporal to store the Workflow's input and output, data used in input parameters and return values must be serializable.
> - By default, Temporal can handle null or binary values, as well as any data that can be serialized using Go's support for JSON or Protocol Buffers.
> - This means that most of the types you'd typically use in a function, such as integers and floating point numbers, boolean values, and strings, are all handled automatically, as are structs composed from these types, but types such as channels, functions, or unsafe pointers are prohibited as either input parameters or return values.
