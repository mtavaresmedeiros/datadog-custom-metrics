## Overview

    []: # Language: go

This repository have a go package that will export a custom metric to the **DataDog** with the current unix epoch time to use to calculate the uptime from a system/process that export a metric with the Start time of the process since unix epoch.

In the dataDog dashboard you will need something like that:

![example](/imgs/dashboard.png)

### Motivation
The dataDog dont have a function or varible that return the current unix epoch.


### Prerequisites:
[Kubernetes](https://kubernetes.io/)

[DogStatsD](https://docs.datadoghq.com/developers/dogstatsd/?tab=hostagent) 

### Instructions:
In order to use it, you need pass the Environment varible `DOGSTATSD_HOST_IP` with the IP of the DataDog agent.

Ex:

    ```
    env:
      - name: DOGSTATSD_HOST_IP
        value: "datadog-statsd.default.svc.cluster.local"
    ```