package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Feature holds cluster-wide information about feature gates.  The canonical name is `cluster`
type Feature struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user settable values for configuration
	// +required
	Spec FeatureSpec `json:"spec"`
	// status holds observed values from the cluster. They may not be overridden.
	// +optional
	Status FeatureStatus `json:"status"`
}

type FeatureSet string

var (
	// Default feature set that allows upgrades.
	Default FeatureSet = ""

	// TechPreviewNoUpgrade turns on tech preview features that are not part of the normal supported platform. Turning
	// this feature set on CANNOT BE UNDONE and PREVENTS UPGRADES.
	TechPreviewNoUpgrade FeatureSet = "TechPreviewNoUpgrade"
)

type FeatureSpec struct {
	// featureSet changes the list of features in the cluster.  The default is empty.  Be very careful adjusting this setting.
	// Turning on or off features may cause irreversible changes in your cluster which cannot be undone.
	FeatureSet FeatureSet `json:"featureSet,omitempty"`
}

type FeatureStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FeatureList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ListMeta `json:"metadata"`
	Items           []Feature `json:"items"`
}

type FeatureEnabledDisabled struct {
	Enabled  []string
	Disabled []string
}

// FeatureSets Contains a map of Feature names to Enabled/Disabled Feature.
//
// NOTE: The caller needs to make sure to check for the existence of the value
// using golang's existence field. A possible scenario is an upgrade where new
// FeatureSets are added and a controller has not been upgraded with a newer
// version of this file. In this upgrade scenario the map could return nil.
//
// example:
//   if featureSet, ok := FeatureSets["SomeNewFeature"]; ok { }
//
// If you put an item in either of these lists, put your area and name on it so we can find owners.
var FeatureSets = map[FeatureSet]*FeatureEnabledDisabled{
	Default: &FeatureEnabledDisabled{
		Enabled: []string{
			"ExperimentalCriticalPodAnnotation", // sig-pod, sjenning
			"RotateKubeletServerCertificate",    // sig-pod, sjenning
		},
		Disabled: []string{
			"LocalStorageCapacityIsolation", // sig-pod, sjenning
			"PersistentLocalVolumes",        // sig-storage, hekumar@redhat.com
		},
	},
	TechPreviewNoUpgrade: &FeatureEnabledDisabled{
		Enabled: []string{
			"ExperimentalCriticalPodAnnotation", // sig-pod, sjenning
			"RotateKubeletServerCertificate",    // sig-pod, sjenning
		},
		Disabled: []string{
			"LocalStorageCapacityIsolation", // sig-pod, sjenning
			"PersistentLocalVolumes",        // sig-storage, hekumar@redhat.com
		},
	},
}
