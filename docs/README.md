# Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

    helm repo add rabbitping https://udhos.github.io/rabbitping

Update files from repo:

    helm repo update

Search rabbitping:

    helm search repo rabbitping -l --version ">=0.0.0"
    NAME                    CHART VERSION   APP VERSION DESCRIPTION
    rabbitping/rabbitping   0.1.0           0.1.0       Install rabbitping helm chart into kubernetes.

To install the charts:

    helm install my-rabbitping rabbitping/rabbitping
    #            ^             ^          ^
    #            |             |           \_______ chart
    #            |             |
    #            |              \__________________ repo
    #            |
    #             \________________________________ release (chart instance installed in cluster)

To uninstall the charts:

    helm uninstall my-rabbitping

# Source

<https://github.com/udhos/rabbitping>
