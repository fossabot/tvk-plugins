package internal

import (
	"os"
	"path"
)

const (
	TVKControlPlaneDeployment    = "k8s-triliovault-control-plane"
	HTTPscheme                   = "http"
	HTTPSscheme                  = "https"
	TriliovaultGroup             = "triliovault.trilio.io"
	TargetKind                   = "Target"
	V1Version                    = "v1"
	APIPath                      = "api"
	LoginPath                    = "login"
	JweToken                     = "jweToken"
	KubeConfigParam              = "kubeconfig"
	ContentType                  = "Content-Type"
	ContentApplicationJSON       = "application/json"
	BackupPlanAPIPath            = "backupplan"
	BackupAPIPath                = "backup"
	MetadataAPIPath              = "metadata"
	ResourceMetadataAPIPath      = "resource-metadata"
	TrilioResourcesAPIPath       = "trilio-resources"
	Results                      = "results"
	FormatYAML                   = "yaml"
	FormatJSON                   = "json"
	FormatWIDE                   = "wide"
	SingleNamespace              = "SingleNamespace"
	MultiNamespace               = "MultiNamespace"
	BackupKind                   = "Backup"
	BackupPlanKind               = "BackupPlan"
	ClusterBackupKind            = "ClusterBackup"
	ClusterBackupPlanKind        = "ClusterBackupPlan"
	IngressKind                  = "Ingress"
	IngressServiceLabel          = "k8s-triliovault-ingress-nginx-controller"
	OldIngressServiceLabel       = "k8s-triliovault-ingress-gateway"
	ServiceTypeNodePort          = "NodePort"
	NamespaceKind                = "Namespace"
	VolumeSnapshotKind           = "VolumeSnapshot"
	VolumeSnapshotClassKind      = "VolumeSnapshotClass"
	PersistentVolumeClaimKind    = "PersistentVolumeClaim"
	CustomResourceDefinitionKind = "CustomResourceDefinition"
	PodKind                      = "Pod"
	DefaultNs                    = "default"
	InstallNamespace             = "INSTALL_NAMESPACE"
	OcpAPIVersion                = "security.openshift.io/v1"

	KubeconfigFlag          = "kubeconfig"
	KubeconfigShorthandFlag = "k"
	KubeconfigUsage         = "Specifies the custom path for your kubeconfig"

	LogLevelFlag            = "log-level"
	LogLevelFlagShorthand   = "l"
	ServiceTypeLoadBalancer = "LoadBalancer"
	ExternalIP              = "ExternalIP"
	LogLevelUsage           = "Set the logging level of the application in the level of" +
		" FATAL, ERROR, WARN, INFO, DEBUG"
	DefaultLogLevel = "INFO"
)

var (
	KubeConfigDefault = path.Join(os.Getenv("HOME"), ".kube", "config")
)
