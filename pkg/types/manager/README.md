# `types/manager`

The `types/manager` package contains interfaces that *manage* cloud resources.

These interfaces are specifically designed in such a way as to *not* have any
dependencies on Kubernetes libraries. This is because the generated
implementations of these interfaces are intended to be used by any client that
wishes to standardize its calling of cloud service APIs using a declarative
resource management strategy but isn't dependent on the Kubernetes
controller-runtime or reconciliation logic.
