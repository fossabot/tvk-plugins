package cmd

const (
	preflightCmdName    = "preflight"
	preflightRunCmdName = "run"
	cleanupCmdName      = "cleanup"

	NamespaceFlag          = "namespace"
	namespaceFlagShorthand = "n"
	namespaceUsage         = "Namespace of the cluster in which the preflight checks will be performed"

	StorageClassFlag  = "storage-class"
	storageClassUsage = "Name of storage class to use for preflight checks"

	SnapshotClassFlag  = "volume-snapshot-class"
	snapshotClassUsage = "Name of volume snapshot class to use for preflight checks"

	LocalRegistryFlag  = "local-registry"
	localRegistryUsage = "Name of the local registry from where the images will be pulled"

	imagePullSecFlag  = "image-pull-secret"
	imagePullSecUsage = "Name of the secret for authentication while pulling the images from the local registry"

	ServiceAccountFlag  = "service-account"
	serviceAccountUsage = "Name of the service account to use for preflight checks and creating preflight resources"

	CleanupOnFailureFlag  = "cleanup-on-failure"
	cleanupOnFailureUsage = "Cleanup the resources on cluster if preflight checks fail. By-default it is false"

	ConfigFileFlag      = "config-file"
	configFileUsage     = "Specify the name of the yaml file for inputs to the preflight Run and Cleanup commands"
	configFlagShorthand = "f"

	PodLimitFlag  = "limits"
	podLimitUsage = "Pod memory and cpu resource limits for DNS and volume snapshot preflight check"

	PodRequestFlag  = "requests"
	podRequestUsage = "Pod memory and cpu resource requests for DNS and volume snapshot preflight check"

	PVCStorageRequestFlag  = "pvc-storage-request"
	pvcStorageRequestUsage = "PVC storage request for volume snapshot preflight check"

	uidFlag  = "uid"
	uidUsage = "UID of the preflight check whose resources must be cleaned"

	preflightLogFilePrefix = "preflight"
	cleanupLogFilePrefix   = "preflight_cleanup"
	preflightUIDLength     = 6

	DefaultPodRequestCPU    = "250m"
	DefaultPodRequestMemory = "64Mi"
	DefaultPodLimitCPU      = "500m"
	DefaultPodLimitMemory   = "128Mi"
	DefaultPVCStorage       = "1Gi"

	filePermission = 0644
)

var (
	kubeconfig        string
	namespace         string
	logLevel          string
	storageClass      string
	snapshotClass     string
	localRegistry     string
	imagePullSecret   string
	serviceAccount    string
	cleanupOnFailure  bool
	inputFileName     string
	podLimits         string
	podRequests       string
	pvcStorageRequest string
	cleanupUID        string
)