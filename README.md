An example to walk you through how to use [**Azure Event Hubs**](https://azure.microsoft.com/services/event-hubs/?WT.mc_id=devto-blog-abhishgu) integration with [**Dapr**](https://dapr.io/), a set of distributed system building blocks for microservices development. Azure Event Hubs will be used as a "binding" within the Dapr runtime. This will allow services to communicate with Azure Event Hubs without actually knowing about it or being coupled to it directly (via SDK, library etc.), using a simple model defined by the Dapr runtime.

The original code got updated to use the Dapr 0.8 spec of Azure Event Hubs for bindings.
Moreover, it also contains an example on how to use Azure EventHubs as pub/sub instead of binding.

Run it:

```
cd pubsub/
./run-k8s.sh
```

This code was used for the investigation of https://github.com/kyma-project/kyma/issues/8789