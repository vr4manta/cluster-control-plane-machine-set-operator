// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	machinev1beta1 "github.com/openshift/api/machine/v1beta1"
	internal "github.com/openshift/client-go/machine/applyconfigurations/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// MachineSetApplyConfiguration represents a declarative configuration of the MachineSet type for use
// with apply.
type MachineSetApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *MachineSetSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *MachineSetStatusApplyConfiguration `json:"status,omitempty"`
}

// MachineSet constructs a declarative configuration of the MachineSet type for use with
// apply.
func MachineSet(name, namespace string) *MachineSetApplyConfiguration {
	b := &MachineSetApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("MachineSet")
	b.WithAPIVersion("machine.openshift.io/v1beta1")
	return b
}

// ExtractMachineSet extracts the applied configuration owned by fieldManager from
// machineSet. If no managedFields are found in machineSet for fieldManager, a
// MachineSetApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// machineSet must be a unmodified MachineSet API object that was retrieved from the Kubernetes API.
// ExtractMachineSet provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractMachineSet(machineSet *machinev1beta1.MachineSet, fieldManager string) (*MachineSetApplyConfiguration, error) {
	return extractMachineSet(machineSet, fieldManager, "")
}

// ExtractMachineSetStatus is the same as ExtractMachineSet except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractMachineSetStatus(machineSet *machinev1beta1.MachineSet, fieldManager string) (*MachineSetApplyConfiguration, error) {
	return extractMachineSet(machineSet, fieldManager, "status")
}

func extractMachineSet(machineSet *machinev1beta1.MachineSet, fieldManager string, subresource string) (*MachineSetApplyConfiguration, error) {
	b := &MachineSetApplyConfiguration{}
	err := managedfields.ExtractInto(machineSet, internal.Parser().Type("com.github.openshift.api.machine.v1beta1.MachineSet"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(machineSet.Name)
	b.WithNamespace(machineSet.Namespace)

	b.WithKind("MachineSet")
	b.WithAPIVersion("machine.openshift.io/v1beta1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithKind(value string) *MachineSetApplyConfiguration {
	b.TypeMetaApplyConfiguration.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithAPIVersion(value string) *MachineSetApplyConfiguration {
	b.TypeMetaApplyConfiguration.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithName(value string) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithGenerateName(value string) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithNamespace(value string) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithUID(value types.UID) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithResourceVersion(value string) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithGeneration(value int64) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithCreationTimestamp(value metav1.Time) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *MachineSetApplyConfiguration) WithLabels(entries map[string]string) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Labels == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *MachineSetApplyConfiguration) WithAnnotations(entries map[string]string) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Annotations == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *MachineSetApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.ObjectMetaApplyConfiguration.OwnerReferences = append(b.ObjectMetaApplyConfiguration.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *MachineSetApplyConfiguration) WithFinalizers(values ...string) *MachineSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.ObjectMetaApplyConfiguration.Finalizers = append(b.ObjectMetaApplyConfiguration.Finalizers, values[i])
	}
	return b
}

func (b *MachineSetApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithSpec(value *MachineSetSpecApplyConfiguration) *MachineSetApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *MachineSetApplyConfiguration) WithStatus(value *MachineSetStatusApplyConfiguration) *MachineSetApplyConfiguration {
	b.Status = value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *MachineSetApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Name
}
