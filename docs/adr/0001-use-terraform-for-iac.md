Status
accepted

Date
2026-10-22

Tags
terraform,infrastructure-as-code,aws

Driver
Chris Aspinwall

Deciders
Chris Aspinwall

Consulted
None (solo capstone project)

Informed
None (solo capstone project)

Use Terraform for infrastructure as code
Context and Problem Statement
This platform needs repeatable, reviewable provisioning of AWS networking, compute, IAM, and EKS. Infrastructure must be versioned in Git, planned before it is applied, and destroyable on demand to control cost. Which infrastructure as code tool should manage everything below the Kubernetes layer?

Decision Drivers
Reproducibility: the full environment must be rebuildable from Git alone
Reviewability: every change should produce a plan that can be reviewed in a pull request before it is applied
Cost control: the environment must be easy to tear down and recreate on demand
Portability: the skills and code should carry over to other clouds or on prem providers
Career alignment: the tool should match what target platform engineering and DoD/federal roles ask for
Considered Options
Terraform
AWS CloudFormation / CDK
Pulumi
Manual console / CLI scripts
Decision Outcome
Chosen option: "Terraform", because it is the only option that meets every decision driver: it provides a plan/apply review workflow, works across clouds and providers, supports full teardown with terraform destroy, and is the tool most commonly named in target job postings.

Implementation details:

All infrastructure in this repo is managed by Terraform
State is stored in an S3 backend with state locking
Reusable pieces are organized as modules
terraform plan runs in CI before any apply
Consequences
🟩 Good, because one tool covers AWS now and other clouds or providers later
🟩 Good, because plans are reviewable in pull requests before anything changes
🟩 Good, because it matches the tooling named in target platform engineering roles
🟥 Bad, because state must be protected; it can contain secrets
🟥 Bad, because drift is possible if anyone changes resources in the console
🟥 Bad, because Kubernetes manifests will still be managed separately (Helm, Flux)
Confirmation
CI runs terraform fmt -check, terraform validate, and terraform plan on every pull request that touches terraform/
No AWS resources in this project are created outside Terraform; a scheduled terraform plan that shows unexpected changes indicates drift
The S3 state bucket is encrypted, versioned, and blocks public access
Pros and Cons of the Options
Terraform
Declarative HCL with a provider for nearly every cloud and SaaS platform.

🟩 Good, because it is multi cloud with a large provider ecosystem
🟩 Good, because the plan/apply workflow makes every change reviewable
🟩 Good, because it is widely used in DoD and federal programs
🟥 Bad, because the state file must be secured and locked
🟥 Bad, because HCL is not a general purpose language
AWS CloudFormation / CDK
AWS native templates (CloudFormation) or code that synthesizes them (CDK).

🟩 Good, because it is native to AWS
🟩 Good, because there is no state file to manage
🟥 Bad, because it is AWS only
🟥 Bad, because the feedback loop is slower
🟥 Bad, because it is less portable to Azure, GCP, or on prem
Pulumi
Infrastructure as code written in general purpose languages such as Go, Python, or TypeScript.

🟩 Good, because it uses real programming languages
🟩 Good, because it has a strong testing story
🟥 Bad, because the community is smaller
🟥 Bad, because it is less common in target job postings
Manual console / CLI scripts
Creating resources by hand in the AWS console or with ad hoc scripts.

🟩 Good, because it is fast for one offs
🟥 Bad, because it is not reproducible
🟥 Bad, because there is no drift detection
🟥 Bad, because there is no review process
More Information
Follow up: ADR 0002, EKS vs RKE2 (Week 14)
Revisit this decision if the platform moves to a provider without a mature Terraform provider, or if Kubernetes resources start being managed from Terraform instead of Helm and Flux