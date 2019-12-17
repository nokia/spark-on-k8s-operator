[![Build Status](https://travis-ci.org/GoogleCloudPlatform/spark-on-k8s-operator.svg?branch=master)](https://travis-ci.org/GoogleCloudPlatform/spark-on-k8s-operator.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoogleCloudPlatform/spark-on-k8s-operator)](https://goreportcard.com/report/github.com/GoogleCloudPlatform/spark-on-k8s-operator)

**This is not an officially supported Google product.**

## Community

* Join our [Slack](https://kubernetes.slack.com/messages/CALBDHMTL) channel.
* Check out [who is using the Kubernetes Operator for Apache Spark](docs/who-is-using.md).

## Project Status

**Project status:** *beta* 

**Current API version:** *`v1beta2`*

**If you are currently using the `v1alpha1` or `v1beta1` version of the APIs in your manifests, please update them to use the `v1beta2` version by changing `apiVersion: "sparkoperator.k8s.io/<version>"` to `apiVersion: "sparkoperator.k8s.io/v1beta2"`. You will also need to delete the `previous` version of the CustomResourceDefinitions named `sparkapplications.sparkoperator.k8s.io` and `scheduledsparkapplications.sparkoperator.k8s.io`, and replace them with the `v1beta2` version either by installing the latest version of the operator or by running `kubectl create -f manifest/crds`.**

Customization of Spark pods, e.g., mounting arbitrary volumes and setting pod affinity, is implemented using a Kubernetes [Mutating Admission Webhook](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/), which became beta in Kubernetes 1.9. The mutating admission webhook is enabled by default if you install the operator using the Helm [chart](https://github.com/helm/charts/tree/master/incubator/sparkoperator). Check out the [Quick Start Guide](docs/quick-start-guide.md#using-the-mutating-admission-webhook)  on how to enable the webhook.

## Prerequisites

* Version >= 1.13 of Kubernetes to use the [`subresource` support for CustomResourceDefinitions](https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/#subresources), which became beta in 1.13 and is enabled by default in 1.13 and higher.

## Installation

The easiest way to install the Kubernetes Operator for Apache Spark is to use the Helm [chart](https://github.com/helm/charts/tree/master/incubator/sparkoperator).

```bash
$ helm repo add incubator http://storage.googleapis.com/kubernetes-charts-incubator
$ helm install incubator/sparkoperator --namespace spark-operator
```

For configuration options available in the Helm chart, please refer to [Configuration](https://github.com/helm/charts/tree/master/incubator/sparkoperator#configuration).

## Version Matrix

The following table lists the most recent few versions of the operator.

| Operator Version | API Version | Kubernetes Version | Base Spark Version | Operator Image Tag |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| `latest` (master HEAD) | `v1beta2` | 1.13+ | `2.4.5-SNAPSHOT` | `latest` |
| `v1beta2-1.0.2-2.4.5-SNAPSHOT` | `v1beta2` | 1.13+ | `2.4.5-SNAPSHOT` | `v1beta2-1.0.2-2.4.5-SNAPSHOT` |
| `v1beta2-1.0.1-2.4.4` | `v1beta2` | 1.13+ | `2.4.4` | `v1beta2-1.0.1-2.4.4` | 
| `v1beta2-1.0.0-2.4.4` | `v1beta2` | 1.13+ | `2.4.4` | `v1beta2-1.0.0-2.4.4` |
| `v1beta1-0.9.0` | `v1beta1` | 1.13+ | `2.4.0` | `v2.4.0-v1beta1-0.9.0` |

When installing using the Helm chart, you can choose to use a specific image tag instead of the default one, using the following option:

```
--set operatorVersion=<operator image tag>
```

## Get Started

Get started quickly with the Kubernetes Operator for Apache Spark using the [Quick Start Guide](docs/quick-start-guide.md). 

If you are running the Kubernetes Operator for Apache Spark on Google Kubernetes Engine and want to use Google Cloud Storage (GCS) and/or BigQuery for reading/writing data, also refer to the [GCP guide](docs/gcp.md).

For more information, check the [Design](docs/design.md), [API Specification](docs/api.md) and detailed [User Guide](docs/user-guide.md).

## Overview

The Kubernetes Operator for Apache Spark aims to make specifying and running [Spark](https://github.com/apache/spark) applications as easy and idiomatic as running other workloads on Kubernetes. It uses 
[Kubernetes custom resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) 
for specifying, running, and surfacing status of Spark applications. For a complete reference of the custom resource definitions, please refer to the [API Definition](docs/api.md). For details on its design, please refer to the [design doc](docs/design.md). It requires Spark 2.3 and above that supports Kubernetes as a native scheduler backend.

The Kubernetes Operator for Apache Spark currently supports the following list of features:

* Supports Spark 2.3 and up.
* Enables declarative application specification and management of applications through custom resources. 
* Automatically runs `spark-submit` on behalf of users for each `SparkApplication` eligible for submission.
* Provides native [cron](https://en.wikipedia.org/wiki/Cron) support for running scheduled applications.
* Supports customization of Spark pods beyond what Spark natively is able to do through the mutating admission webhook, e.g., mounting ConfigMaps and volumes, and setting pod affinity/anti-affinity.
* Supports automatic application re-submission for updated `SparkAppliation` objects with updated specification.
* Supports automatic application restart with a configurable restart policy.
* Supports automatic retries of failed submissions with optional linear back-off.
* Supports mounting local Hadoop configuration as a Kubernetes ConfigMap automatically via `sparkctl`.
* Supports automatically staging local application dependencies to Google Cloud Storage (GCS) via `sparkctl`.
* Supports collecting and exporting application-level metrics and driver/executor metrics to Prometheus. 

Nokia CSF Contributions
* Support for SSL enabled driver UI (https://github.com/GoogleCloudPlatform/spark-on-k8s-operator/issues/490)
* Copy hadoop-site xmls present in hadoopConfigMap to operator pod and export HADOOP_CONF_DIR during spark-submit (https://github.com/GoogleCloudPlatform/spark-on-k8s-operator/issues/572)
* Support for Kerberos authentication for sparkapplications launched by operator. Also provides Configurable CDLK and KDC support (https://github.com/GoogleCloudPlatform/spark-on-k8s-operator/issues/302)
  
## Contributing

Please check [CONTRIBUTING.md](CONTRIBUTING.md) and the [Developer Guide](docs/developer-guide.md) out. 
