# Installing Charles

{% hint style="info" %}
The installation process was created considering some use cases, for each one of which there is a specific tutorial. If you need to install CharlesCD in a different way, we suggest to check the **custom** section with isolated helm charts.
{% endhint %}

## Introduction

### Components

CharlesCD installation considers these components:

1. Seven specific modules of **Charles' architecture;** 
2. **Keycloak**, used for product authentication and authorization;
3. A **PostgreSQL database** for backend modules \( `charles-application`, `charles-circle-matcher`, deploy and villager\) and Keycloak; 
4. A **Redis** to use the [**Circle Matcher**](https://docs.charlescd.io/referencia/circle-matcher). 

### Continuous Delivery Platform

At this moment, Charles can support two Continuous Delivery \(CD\) platforms:

* **Spinnaker:** if you have your spinnaker already configured, it can be used.  
* **Octopipe:** a native platform created by CharlesCD's team to make easier an installation without previous configurations. 

{% hint style="info" %}
If you want more information about how to configure Spinnaker or Octopipe, check the  [**CD Configuration**](https://docs.charlescd.io/v/v0.2.1-eng/reference/cd-configuration) ****section.
{% endhint %}

## Main installation cases

### Case \#1: Quick installation 

This installation is recommended for those who never used Charles before and just want a **first contact on testing environment**, without looking for scalability or security.

In this case, you will have to:

* Use a yaml file with all the **components**;
* Use a _Load Balancer_ previously configured.

To create this structure, you have to run the files in a configured cluster, such as minikube, GKE, EKS, etc. To execute this, follow the next steps:

```text
kubectl create namespace charles

kubectl apply -f arquivo.yaml
```

At the end of the process, you will have in the `charles`namespace, all the modules of the product, as well as your installed dependencies in a simpler way.

{% hint style="danger" %}
This installation only works for testing environments, we do not recommend this use case for production environments, because it doesn't include database backup, high availability, etc. 
{% endhint %}

### Case \#2:  Installation with helm charts

This installation is recommended for who has already setup your infrastructure due to a more complex environment or have some security or/and scalability limitations, which demands a more complete installation customization from CharlesCD. 

### Prerequisites 

To perform this process, it is necessary to have the following programs: 

* Kubectl
* Helm 

### How does it work?

The most different advantage on this installation is the customization. We will provide a **helm template** with all the available fields for modification, including the database and consumed resources. 

You will find here the **documentation with all the editable fields.**

{% hint style="info" %}
It is important to remember that, if there is no customization, the result will be the same as in the **use case \#1**, that follows a pattern to install PostgreSQL, Redis, Keycloak and Octopipe.

So, don't forget to customize your fields if you want something manageable. 
{% endhint %}

To install, you have to run the following commands after customizing the fields: 

```text
// customize all you need in the file values.yaml before running the following commands
helm install charlescd <repo-folder> -n <namespace>
```
