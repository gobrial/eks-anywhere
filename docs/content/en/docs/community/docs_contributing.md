---
title: "Contributing to EKS Anywhere documentation"
weight: 20
description: >
  Guidelines for contributing to EKS Anywhere documentation
---
EKS Anywhere documentation uses the [Hugo](https://gohugo.io/categories/fundamentals) site generator and the [Docsy](https://www.docsy.dev/docs/) theme. To get started contributing:

* View the published EKS Anywhere [user documentation](https://anywhere.eks.amazonaws.com/docs/).
* Fork and clone the [eks-anywhere](https://github.com/aws/eks-anywhere) project.
* See [EKS Anywhere Documentation](https://github.com/aws/eks-anywhere/tree/main/docs) to set up your own docs test site.
* See the [General Guidelines](https://github.com/aws/eks-anywhere/blob/main/docs/content/en/docs/community/contributing.md) for contributing to the EKS Anywhere project
* Create EKS Anywhere documentation [Issues](https://github.com/aws/eks-anywhere/issues) and [Pull Requests](https://github.com/aws/eks-anywhere/pulls).

## Style issues

* **EKS Anywhere**: Always refer to EKS Anywhere as EKS Anywhere and *NOT* EKS-A or EKS-Anywhere.
* **Line breaks**: Put each sentence on its own line and don’t do a line break in the middle of a sentence. 
  We are using a modified [Semantic Line Breaking](https://sembr.org/) in that we are requiring a break at the end of every sentence, but not at commas or other semantic boundaries.
* **Headings**: Use sentence case in headings. So do “Cluster specification reference” and not “Cluster Specification Reference”
* **Cross references**: To cross reference to another doc in the EKS Anywhere docs set, use relref in the link so that Hugo will test it and fail the build for links not found. Also, use relative paths to point to other content in the docs set. For example:
   ```
     [troubleshooting section]({{< relref "../tasks/troubleshoot" >}})
   ```
* **Notes, Warnings, etc.**: You can use this form for notes:
    
    {{% alert title="Note" color="primary" %}}
    <put note here, multiple paragraphs are allowed>
    {{% /alert %}}

* **Embedding content**: If you want to read in content from a separate file, you can use the following format.
  Do this if you think the content might be useful in multiple pages:
  ```
  {{% content "./newfile.md" %}}
  ```
* **General style issues**: Unless otherwise instructed, follow the [Kubernetes Documentation Style Guide](https://kubernetes.io/docs/contribute/style/style-guide/) for formatting and presentation guidance.

## Where to put content

* **Images**: Put all images into the EKS Anywhere GitHub site’s [docs/static/images](https://github.com/aws/eks-anywhere/tree/main/docs/static/images) directory.
* **Yaml examples**: Put full yaml file examples into the EKS Anywhere GitHub site’s [docs/static/manifests](https://github.com/aws/eks-anywhere/tree/main/docs/static/manifests) directory.
  In kubectl examples, you can point to those files using: `https://anywhere.eks.amazonaws.com/manifests/whatever.yaml`
* **Generic instructions for creating a cluster** should go into the [getting started](https://anywhere.eks.amazonaws.com/docs/getting-started/) section in either:
   * [Install EKS Anywhere](https://anywhere.eks.amazonaws.com/docs/getting-started/install/) installation guide: For prerequisites and procedures related to setting up the Administrative machine.
   * [Creating a local cluster](https://anywhere.eks.amazonaws.com/docs/getting-started/local-environment/) or [production cluster](https://anywhere.eks.amazonaws.com/docs/getting-started/production-environment/): For simple instructions for a Docker or vSphere installation, respectively.
* **Instructions that are specific to an EKS Anywhere provider** should go into the appropriate provider section. Currently, [vSphere](https://anywhere.eks.amazonaws.com/docs/reference/vsphere/) is the only supported provider.
  * [Add integrations to cluster]({{< relref "../tasks/cluster/cluster-integrations/" >}}): Add names of suggested third-party tools. Then Link the names of providers to:
    * EKS Anywhere docs instructions for configuring that feature, if instructions are available or
    * Somewhere on the third-party site, if there are no instructions available on the EKS Anywhere site
  * [Compare EKS Anywhere and EKS]({{< relref "../concepts/eksafeatures/" >}}): Add supported third-party solutions to the Amazon EKS Anywhere column.
  Only link to the partner page for now.
* **Workshop content** should contain organized links to existing documentation pages.
  The workshop content should not duplicate existing documentation pages or contain guides that are not part of the main documentation.

## Contributing docs for third-party solutions

To contribute documentation describing how to use third-party software products or projects with EKS Anywhere, follow these guidelines.

### Docs for third-party software in EKS Anywhere

Documentation PRs for EKS Anywhere that describe third-party software that is included in EKS Anywhere are acceptable, provided they meet the quality standards described in the Tips described below. This includes:

* Software bundled with EKS Anywhere (for example, [Cilium docs](https://anywhere.eks.amazonaws.com/docs/tasks/workload/networking-and-security/))
* Supported platforms on which EKS Anywhere runs (for example, [VMware vSphere](https://anywhere.eks.amazonaws.com/docs/reference/vsphere/))
* Curated software that is packaged by the EKS Anywhere project to run EKS Anywhere. This includes documentation for Harbor local registry, Ingress controller, and Prometheus, Grafana, and Fluentd monitoring and logging.

### Docs for third-party software NOT in EKS Anywhere

Documentation for software that is not part of EKS Anywhere software can still be added to EKS Anywhere docs by meeting one of the following criteria:

* **Partners**: Documentation PRs for software from vendors listed on the [EKS Anywhere Partner page](https://aws.amazon.com/eks/eks-anywhere/partners/) can be considered to add to the EKS Anywhere docs.
  Links point to partners from the [Compare EKS Anywhere to EKS](https://anywhere.eks.amazonaws.com/docs/concepts/eksafeatures/) page and other content can be added to EKS Anywhere documentation for features from those partners.
  Contact the AWS container partner team if you are interested in becoming a partner: aws-container-partners@amazon.com
* **Cluster integrations**: Separate, less stringent criteria can be met for a third-party vendor to be listed on the [Add cluster integrations](https://anywhere.eks.amazonaws.com/docs/tasks/cluster/cluster-integrations/) page.

### Tips for contributing third-party docs

The Kubernetes docs project itself describes a similar approach to docs covering third-party software in the [How Docs Handle Third Party and Dual Sourced Content](https://kubernetes.io/blog/2020/05/third-party-dual-sourced-content/) blog.
In line with these general guidelines, we recommend that even acceptable third-party docs contributions to EKS Anywhere:

* **Not be dual-sourced**: The project does not allow content that is already published somewhere else.
  You can provide links to that content, if it is relevant. Heavily rewriting such content to be EKS Anywhere-specific might be acceptable.
* **Not be marketing oriented**. The content shouldn’t sell a third-party products or make vague claims of quality.
* **Not outside the scope of EKS Anywhere**:  Just because some projects or products of a partner are appropriate for EKS Anywhere docs, it doesn’t mean that any project or product by that partner can be documented in EKS Anywhere.
* **Stick to the facts**:  So, for example, docs about third-party software could say: “To set up load balancer ABC, do XYZ” or “Make these modifications to improve speed and efficiency.” It should not make blanket statements like: “ABC load balancer is the best one in the industry.”
* **EKS features**: Features that relate to EKS which runs in AWS or requires an AWS account should link to [the official documentation](https://docs.aws.amazon.com/eks/latest/) as much as possible.
