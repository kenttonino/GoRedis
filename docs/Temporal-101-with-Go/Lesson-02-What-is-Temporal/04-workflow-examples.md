# Expense Report

> - https://temporal.talentlms.com/unit/view/id:2110

<br />

![01-expense-report-example](./images/01-expense-report-workflow-diagram.png)

> - The diagram is a long-running process.
> - Depending on the organization and the number of approvals required, it may take days, weeks, or longer from start to finish.
> - Another is that there's conditional logic; just like a computer program, there are decision points and execution paths that diverge based on the outcome; if your expense report is accepted, reimbursement is the next step, bbut if it's rejected, the next step is sending a notification requesting that you modify and resubmit the report.
> - And that's another characteristic; the workflow can contain cycles; for example, because a rejected report may lead to correction, re-submission, and another review.
> - It's also worth noting that the workflow involves multiple points of human interaction, from the employee, the manager, and the accounting department.
> - It also involves external systems, notably the company's bank, which is the source of the reimbursement, and the employee's bank, which is the target of those funds.
