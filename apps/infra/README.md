# Silica

This project contains the infrastructure for Silica, e,g, the ingress, postgres, etc. And must be deployed in order for the Silica modules (e.g. [os](../os), [todo](../todo)) to work properly.

## Getting Started

---

### Prerequisites & Dependencies

1. Setup the [global dependencies](../../README.md#prerequisite--dependencies).
2. Install [Helm](https://helm.sh/), which is used to manage the deployment with Kubernetes. Verify with `helm version`.
3. Install [DevSpace](https://www.devspace.sh/). which is used for local Kubernetes development. Verify with `devspace version`.

### Running project in the local cluster

1. Create a local cluster with `npm run init:cluster`.
2. Deploy Silica with `npx nx deploy-dev demo`.
3. Verify the deployment with `helm ls`, you should be able to see a release named `demo`.
4. Try deploying Silica modules, e.g. [os](../os/README.md#running-project-in-the-local-cluster)

### Debugging ingress

The ingress controller [Traefik](https://traefik.io/) dashboard is not exposed in production, but for debug purpose you can run `npx nx start-dev-ingress demo` to access it, which would display the registered routes and the traffic.
