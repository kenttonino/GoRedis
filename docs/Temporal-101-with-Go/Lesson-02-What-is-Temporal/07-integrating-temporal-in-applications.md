# Integrating Temporal into Other Applications

> - https://temporal.talentlms.com/unit/view/id:2541

<br />
<br />
<br />



# Direct Integration in Application Frontend

![07-application-frontend-integration](./images/01-expense-report-workflow-diagram.png)

> - It is possible to use a Temporal Client from within those applications.
> - It is also possible to issue gRPC requests directly from the applications without using a Temporal Client at all.
> - However, both are atypical approaches.

<br />
<br />
<br />



# Integration through a Backend Application

![08-integration-application-backend-application](./images/08-integration-through-backend-application.png)

> - A more typical approach is to have the end user application make calls to a service, such as a web application that provides a REST endpoint, which acts as an application gateway and uses a Temporal Client to interact with the Cluster.
