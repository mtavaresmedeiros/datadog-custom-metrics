## Overview
This repository as a go package that will export a custom metric with the unix epoch as the value.

### Motivation
The dataDog dont have a function or varible that return the current unix epoch.

    
    []: # Language: go



### Prerequisites:
Kubernetes (https://kubernetes.io/)
DogStatsD (https://docs.datadoghq.com/developers/dogstatsd/?tab=hostagent) 

### Instructions:
In order to use it, you need pass the Environment varible `DOGSTATSD_HOST_IP` with the IP of the DataDog agent.

Ex:

    ```
    env:
      - name: DOGSTATSD_HOST_IP
        value: "datadog-statsd.default.svc.cluster.local"
    ```