/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta3

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/resource"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/runtime"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/types"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
)

// Common string formats
// ---------------------
// Many fields in this API have formatting requirements.  The commonly used
// formats are defined here.
//
// C_IDENTIFIER:  This is a string that conforms to the definition of an "identifier"
//     in the C language.  This is captured by the following regex:
//         [A-Za-z_][A-Za-z0-9_]*
//     This defines the format, but not the length restriction, which should be
//     specified at the definition of any field of this type.
//
// DNS_LABEL:  This is a string, no more than 63 characters long, that conforms
//     to the definition of a "label" in RFCs 1035 and 1123.  This is captured
//     by the following regex:
//         [a-z0-9]([-a-z0-9]*[a-z0-9])?
//
// DNS_SUBDOMAIN:  This is a string, no more than 253 characters long, that conforms
//      to the definition of a "subdomain" in RFCs 1035 and 1123.  This is captured
//      by the following regex:
//         [a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*
//     or more simply:
//         DNS_LABEL(\.DNS_LABEL)*

// TypeMeta describes an individual object in an API response or request
// with strings representing the type of the object and its API schema version.
// Structures that are versioned or persisted should inline TypeMeta.
type TypeMeta struct {
	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to.
	Kind string `json:"kind,omitempty" description:"kind of object, in CamelCase; cannot be updated"`

	// APIVersion defines the versioned schema of this representation of an object.
	// Servers should convert recognized schemas to the latest internal value, and
	// may reject unrecognized values.
	APIVersion string `json:"apiVersion,omitempty" description:"version of the schema the object should have"`
}

// ListMeta describes metadata that synthetic resources must have, including lists and
// various status objects.
type ListMeta struct {
	// SelfLink is a URL representing this object.
	SelfLink string `json:"selfLink,omitempty" description:"URL for the object; populated by the system, read-only"`

	// An opaque value that represents the version of this response for use with optimistic
	// concurrency and change monitoring endpoints.  Clients must treat these values as opaque
	// and values may only be valid for a particular resource or set of resources. Only servers
	// will generate resource versions.
	ResourceVersion string `json:"resourceVersion,omitempty" description:"string that identifies the internal version of this object that can be used by clients to determine when objects have changed; populated by the system, read-only; value must be treated as opaque by clients and passed unmodified back to the server: https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#concurrency-control-and-consistency"`
}

// ObjectMeta is metadata that all persisted resources must have, which includes all objects
// users must create.
type ObjectMeta struct {
	// Name is unique within a namespace.  Name is required when creating resources, although
	// some resources may allow a client to request the generation of an appropriate name
	// automatically. Name is primarily intended for creation idempotence and configuration
	// definition.
	Name string `json:"name,omitempty" description:"string that identifies an object. Must be unique within a namespace; cannot be updated"`

	// GenerateName indicates that the name should be made unique by the server prior to persisting
	// it. A non-empty value for the field indicates the name will be made unique (and the name
	// returned to the client will be different than the name passed). The value of this field will
	// be combined with a unique suffix on the server if the Name field has not been provided.
	// The provided value must be valid within the rules for Name, and may be truncated by the length
	// of the suffix required to make the value unique on the server.
	//
	// If this field is specified, and Name is not present, the server will NOT return a 409 if the
	// generated name exists - instead, it will either return 201 Created or 500 with Reason
	// ServerTimeout indicating a unique name could not be found in the time allotted, and the client
	// should retry (optionally after the time indicated in the Retry-After header).
	GenerateName string `json:"generateName,omitempty" description:"an optional prefix to use to generate a unique name; has the same validation rules as name; optional, and is applied only name if is not specified"`

	// Namespace defines the space within which name must be unique. An empty namespace is
	// equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for
	// those objects will be empty.
	Namespace string `json:"namespace,omitempty" description:"namespace of the object; cannot be updated"`

	// SelfLink is a URL representing this object.
	SelfLink string `json:"selfLink,omitempty" description:"URL for the object; populated by the system, read-only"`

	// UID is the unique in time and space value for this object. It is typically generated by
	// the server on successful creation of a resource and is not allowed to change on PUT
	// operations.
	UID types.UID `json:"uid,omitempty" description:"unique UUID across space and time; populated by the system; read-only"`

	// An opaque value that represents the version of this resource. May be used for optimistic
	// concurrency, change detection, and the watch operation on a resource or set of resources.
	// Clients must treat these values as opaque and values may only be valid for a particular
	// resource or set of resources. Only servers will generate resource versions.
	ResourceVersion string `json:"resourceVersion,omitempty" description:"string that identifies the internal version of this object that can be used by clients to determine when objects have changed; populated by the system, read-only; value must be treated as opaque by clients and passed unmodified back to the server: https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#concurrency-control-and-consistency"`

	// CreationTimestamp is a timestamp representing the server time when this object was
	// created. It is not guaranteed to be set in happens-before order across separate operations.
	// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	CreationTimestamp util.Time `json:"creationTimestamp,omitempty" description:"RFC 3339 date and time at which the object was created; populated by the system, read-only; null for lists"`

	// DeletionTimestamp is the time after which this resource will be deleted. This
	// field is set by the server when a graceful deletion is requested by the user, and is not
	// directly settable by a client. The resource will be deleted (no longer visible from
	// resource lists, and not reachable by name) after the time in this field. Once set, this
	// value may not be unset or be set further into the future, although it may be shortened
	// or the resource may be deleted prior to this time. For example, a user may request that
	// a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination
	// signal to the containers in the pod. Once the resource is deleted in the API, the Kubelet
	// will send a hard termination signal to the container.
	DeletionTimestamp *util.Time `json:"deletionTimestamp,omitempty" description:"RFC 3339 date and time at which the object will be deleted; populated by the system when a graceful deletion is requested, read-only; if not set, graceful deletion of the object has not been requested"`

	// Labels are key value pairs that may be used to scope and select individual resources.
	// TODO: replace map[string]string with labels.LabelSet type
	Labels map[string]string `json:"labels,omitempty" description:"map of string keys and values that can be used to organize and categorize objects; may match selectors of replication controllers and services"`

	// Annotations are unstructured key value data stored with a resource that may be set by
	// external tooling. They are not queryable and should be preserved when modifying
	// objects.
	Annotations map[string]string `json:"annotations,omitempty" description:"map of string keys and values that can be used by external tooling to store and retrieve arbitrary metadata about objects"`
}

const (
	// NamespaceDefault means the object is in the default namespace which is applied when not specified by clients
	NamespaceDefault string = "default"
	// NamespaceAll is the default argument to specify on a context when you want to list or filter resources across all namespaces
	NamespaceAll string = ""
)

//
//// ContainerManifest corresponds to the Container Manifest format, documented at:
//// https://developers.google.com/compute/docs/containers/container_vms#container_manifest
//// This is used as the representation of Kubernetes workloads.
//// DEPRECATED: Exists to allow backwards compatible storage for clients accessing etcd
//// directly.
//type ContainerManifest struct {
//	// Required: This must be a supported version string, such as "v1beta1".
//	Version string `json:"version"`
//	// Required: This must be a DNS_SUBDOMAIN.
//	// TODO: ID on Manifest is deprecated and will be removed in the future.
//	ID string `json:"id"`
//	// TODO: UUID on Manifest is deprecated in the future once we are done
//	// with the API refactoring. It is required for now to determine the instance
//	// of a Pod.
//	UUID          types.UID      `json:"uuid,omitempty"`
//	Volumes       []Volume      `json:"volumes"`
//	Containers    []Container   `json:"containers"`
//	RestartPolicy RestartPolicy `json:"restartPolicy,omitempty"`
//}
//
//// ContainerManifestList is used to communicate container manifests to kubelet.
//// DEPRECATED: Exists to allow backwards compatible storage for clients accessing etcd
//// directly.
//type ContainerManifestList struct {
//	TypeMeta `json:",inline"`
//	// ID is the legacy field representing Name
//	ID string `json:"id,omitempty"`
//
//	Items []ContainerManifest `json:"items,omitempty"`
//}

// Volume represents a named volume in a pod that may be accessed by any containers in the pod.
type Volume struct {
	// Required: This must be a DNS_LABEL.  Each volume in a pod must have
	// a unique name.
	Name string `json:"name" description:"volume name; must be a DNS_LABEL and unique within the pod"`
	// Source represents the location and type of a volume to mount.
	// This is optional for now. If not specified, the Volume is implied to be an EmptyDir.
	// This implied behavior is deprecated and will be removed in a future version.
	VolumeSource `json:",inline,omitempty"`
}

// VolumeSource represents the source location of a valume to mount.
// Only one of its members may be specified.
type VolumeSource struct {
	// HostPath represents a pre-existing file or directory on the host
	// machine that is directly exposed to the container. This is generally
	// used for system agents or other privileged things that are allowed
	// to see the host machine. Most containers will NOT need this.
	// TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not
	// mount host directories as read/write.
	HostPath *HostPathVolumeSource `json:"hostPath" description:"pre-existing host file or directory; generally for privileged system daemons or other agents tied to the host"`
	// EmptyDir represents a temporary directory that shares a pod's lifetime.
	EmptyDir *EmptyDirVolumeSource `json:"emptyDir" description:"temporary directory that shares a pod's lifetime"`
	// GCEPersistentDisk represents a GCE Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	GCEPersistentDisk *GCEPersistentDiskVolumeSource `json:"gcePersistentDisk" description:"GCE disk resource attached to the host machine on demand"`
	// GitRepo represents a git repository at a particular revision.
	GitRepo *GitRepoVolumeSource `json:"gitRepo" description:"git repository at a particular revision"`
	// Secret represents a secret that should populate this volume.
	Secret *SecretVolumeSource `json:"secret" description:"secret to populate volume"`
	// NFS represents an NFS mount on the host that shares a pod's lifetime
	NFS *NFSVolumeSource `json:"nfs" description:"NFS volume that will be mounted in the host machine"`
}

// HostPathVolumeSource represents bare host directory volume.
type HostPathVolumeSource struct {
	Path string `json:"path" description:"path of the directory on the host"`
}

type EmptyDirVolumeSource struct {
	// Optional: what type of storage medium should back this directory.
	// The default is "" which means to use the node's default medium.
	Medium StorageType `json:"medium" description:"type of storage used to back the volume; must be an empty string (default) or Memory"`
}

// StorageType defines ways that storage can be allocated to a volume.
type StorageType string

const (
	StorageTypeDefault StorageType = ""       // use whatever the default is for the node
	StorageTypeMemory  StorageType = "Memory" // use memory (tmpfs)
)

// Protocol defines network protocols supported for things like conatiner ports.
type Protocol string

const (
	// ProtocolTCP is the TCP protocol.
	ProtocolTCP Protocol = "TCP"
	// ProtocolUDP is the UDP protocol.
	ProtocolUDP Protocol = "UDP"
)

// GCEPersistentDiskVolumeSource represents a Persistent Disk resource in Google Compute Engine.
//
// A GCE PD must exist and be formatted before mounting to a container.
// The disk must also be in the same GCE project and zone as the kubelet.
// A GCE PD can only be mounted as read/write once.
type GCEPersistentDiskVolumeSource struct {
	// Unique name of the PD resource. Used to identify the disk in GCE
	PDName string `json:"pdName" description:"unique name of the PD resource in GCE"`
	// Required: Filesystem type to mount.
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs"
	// TODO: how do we prevent errors in the filesystem from compromising the machine
	FSType string `json:"fsType,omitempty" description:"file system type to mount, such as ext4, xfs, ntfs"`
	// Optional: Partition on the disk to mount.
	// If omitted, kubelet will attempt to mount the device name.
	// Ex. For /dev/sda1, this field is "1", for /dev/sda, this field is 0 or empty.
	Partition int `json:"partition,omitempty" description:"partition on the disk to mount (e.g., '1' for /dev/sda1); if omitted the plain device name (e.g., /dev/sda) will be mounted"`
	// Optional: Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	ReadOnly bool `json:"readOnly,omitempty" description:"read-only if true, read-write otherwise (false or unspecified)"`
}

// GitRepoVolumeSource represents a volume that is pulled from git when the pod is created.
type GitRepoVolumeSource struct {
	// Repository URL
	Repository string `json:"repository" description:"repository URL"`
	// Commit hash, this is optional
	Revision string `json:"revision" description:"commit hash for the specified revision"`
}

// SecretVolumeSource adapts a Secret into a VolumeSource
//
// https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/design/secrets.md
type SecretVolumeSource struct {
	// Name of the secret in the pod's namespace to use
	SecretName string `json:"secretName" description:"secretName is the name of a secret in the pod's namespace"`
}

// NFSVolumeSource represents an NFS mount that lasts the lifetime of a pod
type NFSVolumeSource struct {
	// Server is the hostname or IP address of the NFS server
	Server string `json:"server" description:"the hostname or IP address of the NFS server"`

	// Path is the exported NFS share
	Path string `json:"path" description:"the path that is exported by the NFS server"`

	// Optional: Defaults to false (read/write). ReadOnly here will force
	// the NFS export to be mounted with read-only permissions
	ReadOnly bool `json:"readOnly,omitempty" description:"forces the NFS export to be mounted with read-only permissions"`
}

// ContainerPort represents a network port in a single container.
type ContainerPort struct {
	// Optional: If specified, this must be a DNS_LABEL.  Each named port
	// in a pod must have a unique name.
	Name string `json:"name,omitempty" description:"name for the port that can be referred to by services; must be a DNS_LABEL and unique without the pod"`
	// Optional: If specified, this must be a valid port number, 0 < x < 65536.
	HostPort int `json:"hostPort,omitempty" description:"number of port to expose on the host; most containers do not need this"`
	// Required: This must be a valid port number, 0 < x < 65536.
	ContainerPort int `json:"containerPort" description:"number of port to expose on the pod's IP address"`
	// Optional: Defaults to "TCP".
	Protocol Protocol `json:"protocol,omitempty" description:"protocol for port; must be UDP or TCP; TCP if unspecified"`
	// Optional: What host IP to bind the external port to.
	HostIP string `json:"hostIP,omitempty" description:"host IP to bind the port to"`
}

// VolumeMount describes a mounting of a Volume within a container.
type VolumeMount struct {
	// Required: This must match the Name of a Volume [above].
	Name string `json:"name" description:"name of the volume to mount"`
	// Optional: Defaults to false (read-write).
	ReadOnly bool `json:"readOnly,omitempty" description:"mounted read-only if true, read-write otherwise (false or unspecified)"`
	// Required.
	MountPath string `json:"mountPath" description:"path within the container at which the volume should be mounted"`
}

// EnvVar represents an environment variable present in a Container.
type EnvVar struct {
	// Required: This must be a C_IDENTIFIER.
	Name string `json:"name" description:"name of the environment variable; must be a C_IDENTIFIER"`
	// Optional: defaults to "".
	Value string `json:"value,omitempty" description:"value of the environment variable; defaults to empty string"`
}

// HTTPGetAction describes an action based on HTTP Get requests.
type HTTPGetAction struct {
	// Optional: Path to access on the HTTP server.
	Path string `json:"path,omitempty" description:"path to access on the HTTP server"`
	// Required: Name or number of the port to access on the container.
	Port util.IntOrString `json:"port,omitempty" description:"number or name of the port to access on the container"`
	// Optional: Host name to connect to, defaults to the pod IP.
	Host string `json:"host,omitempty" description:"hostname to connect to; defaults to pod IP"`
}

// TCPSocketAction describes an action based on opening a socket
type TCPSocketAction struct {
	// Required: Port to connect to.
	Port util.IntOrString `json:"port,omitempty" description:"number of name of the port to access on the container"`
}

// ExecAction describes a "run in container" action.
type ExecAction struct {
	// Command is the command line to execute inside the container, the working directory for the
	// command  is root ('/') in the container's filesystem.  The command is simply exec'd, it is
	// not run inside a shell, so traditional shell instructions ('|', etc) won't work.  To use
	// a shell, you need to explicitly call out to that shell.
	Command []string `json:"command,omitempty" description:"command line to execute inside the container; working directory for the command is root ('/') in the container's file system; the command is exec'd, not run inside a shell; exit status of 0 is treated as live/healthy and non-zero is unhealthy"`
}

// Probe describes a liveness probe to be examined to the container.
type Probe struct {
	// The action taken to determine the health of a container
	Handler `json:",inline"`
	// Length of time before health checking is activated.  In seconds.
	InitialDelaySeconds int64 `json:"initialDelaySeconds,omitempty" description:"number of seconds after the container has started before liveness probes are initiated"`
	// Length of time before health checking times out.  In seconds.
	TimeoutSeconds int64 `json:"timeoutSeconds,omitempty" description:"number of seconds after which liveness probes timeout; defaults to 1 second"`
}

// PullPolicy describes a policy for if/when to pull a container image
type PullPolicy string

const (
	// PullAlways means that kubelet always attempts to pull the latest image.  Container will fail If the pull fails.
	PullAlways PullPolicy = "Always"
	// PullNever means that kubelet never pulls an image, but only uses a local image.  Container will fail if the image isn't present
	PullNever PullPolicy = "Never"
	// PullIfNotPresent means that kubelet pulls if the image isn't present on disk. Container will fail if the image isn't present and the pull fails.
	PullIfNotPresent PullPolicy = "IfNotPresent"
)

// CapabilityType represent POSIX capabilities type
type CapabilityType string

// Capabilities represent POSIX capabilities that can be added or removed to a running container.
type Capabilities struct {
	// Added capabilities
	Add []CapabilityType `json:"add,omitempty" description:"added capabilities"`
	// Removed capabilities
	Drop []CapabilityType `json:"drop,omitempty" description:"droped capabilities"`
}

// ResourceRequirements describes the compute resource requirements.
type ResourceRequirements struct {
	// Limits describes the maximum amount of compute resources required.
	Limits ResourceList `json:"limits,omitempty" description:"Maximum amount of compute resources allowed"`
}

const (
	// TerminationMessagePathDefault means the default path to capture the application termination message running in a container
	TerminationMessagePathDefault string = "/dev/termination-log"
)

// Container represents a single container that is expected to be run on the host.
type Container struct {
	// Required: This must be a DNS_LABEL.  Each container in a pod must
	// have a unique name.
	Name string `json:"name" description:"name of the container; must be a DNS_LABEL and unique within the pod; cannot be updated"`
	// Required.
	Image string `json:"image" description:"Docker image name"`
	// Optional: Defaults to whatever is defined in the image.
	Command []string `json:"command,omitempty" description:"command argv array; not executed within a shell; defaults to entrypoint or command in the image; cannot be updated"`
	// Optional: Defaults to Docker's default.
	WorkingDir     string               `json:"workingDir,omitempty" description:"container's working directory; defaults to image's default; cannot be updated"`
	Ports          []ContainerPort      `json:"ports,omitempty" description:"list of ports to expose from the container; cannot be updated"`
	Env            []EnvVar             `json:"env,omitempty" description:"list of environment variables to set in the container; cannot be updated"`
	Resources      ResourceRequirements `json:"resources,omitempty" description:"Compute Resources required by this container; cannot be updated"`
	VolumeMounts   []VolumeMount        `json:"volumeMounts,omitempty" description:"pod volumes to mount into the container's filesyste; cannot be updated"`
	LivenessProbe  *Probe               `json:"livenessProbe,omitempty" description:"periodic probe of container liveness; container will be restarted if the probe fails; cannot be updated"`
	ReadinessProbe *Probe               `json:"readinessProbe,omitempty" description:"periodic probe of container service readiness; container will be removed from service endpoints if the probe fails; cannot be updated"`
	Lifecycle      *Lifecycle           `json:"lifecycle,omitempty" description:"actions that the management system should take in response to container lifecycle events; cannot be updated"`
	// Optional: Defaults to /dev/termination-log
	TerminationMessagePath string `json:"terminationMessagePath,omitempty" description:"path at which the file to which the container's termination message will be written is mounted into the container's filesystem; message written is intended to be brief final status, such as an assertion failure message; defaults to /dev/termination-log; cannot be updated"`
	// Optional: Default to false.
	Privileged bool `json:"privileged,omitempty" description:"whether or not the container is granted privileged status; defaults to false; cannot be updated"`
	// Optional: Policy for pulling images for this container
	ImagePullPolicy PullPolicy `json:"imagePullPolicy" description:"image pull policy; one of PullAlways, PullNever, PullIfNotPresent; defaults to PullAlways if :latest tag is specified, or PullIfNotPresent otherwise; cannot be updated"`
	// Optional: Capabilities for container.
	Capabilities Capabilities `json:"capabilities,omitempty" description:"capabilities for container; cannot be updated"`
}

// Handler defines a specific action that should be taken
// TODO: pass structured data to these actions, and document that data here.
type Handler struct {
	// One and only one of the following should be specified.
	// Exec specifies the action to take.
	Exec *ExecAction `json:"exec,omitempty" description:"exec-based handler"`
	// HTTPGet specifies the http request to perform.
	HTTPGet *HTTPGetAction `json:"httpGet,omitempty" description:"HTTP-based handler"`
	// TCPSocket specifies an action involving a TCP port.
	// TODO: implement a realistic TCP lifecycle hook
	TCPSocket *TCPSocketAction `json:"tcpSocket,omitempty"  description:"TCP-based handler; TCP hooks not yet supported"`
}

// Lifecycle describes actions that the management system should take in response to container lifecycle
// events.  For the PostStart and PreStop lifecycle handlers, management of the container blocks
// until the action is complete, unless the container process fails, in which case the handler is aborted.
type Lifecycle struct {
	// PostStart is called immediately after a container is created.  If the handler fails, the container
	// is terminated and restarted.
	PostStart *Handler `json:"postStart,omitempty" description:"called immediately after a container is started; if the handler fails, the container is terminated and restarted according to its restart policy; other management of the container blocks until the hook completes"`
	// PreStop is called immediately before a container is terminated.  The reason for termination is
	// passed to the handler.  Regardless of the outcome of the handler, the container is eventually terminated.
	PreStop *Handler `json:"preStop,omitempty" description:"called before a container is terminated; the container is terminated after the handler completes; other management of the container blocks until the hook completes"`
}

type ConditionStatus string

// These are valid condition statuses. "ConditionFull" means a resource is in the condition;
// "ConditionNone" means a resource is not in the condition; "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
const (
	ConditionFull    ConditionStatus = "Full"
	ConditionNone    ConditionStatus = "None"
	ConditionUnknown ConditionStatus = "Unknown"
)

type ContainerStateWaiting struct {
	// Reason could be pulling image,
	Reason string `json:"reason,omitempty" description:"(brief) reason the container is not yet running, such as pulling its image"`
}

type ContainerStateRunning struct {
	StartedAt util.Time `json:"startedAt,omitempty" description:"time at which the container was last (re-)started"`
}

type ContainerStateTerminated struct {
	ExitCode    int       `json:"exitCode" description:"exit status from the last termination of the container"`
	Signal      int       `json:"signal,omitempty" description:"signal from the last termination of the container"`
	Reason      string    `json:"reason,omitempty" description:"(brief) reason from the last termination of the container"`
	Message     string    `json:"message,omitempty" description:"message regarding the last termination of the container"`
	StartedAt   util.Time `json:"startedAt,omitempty" description:"time at which previous execution of the container started"`
	FinishedAt  util.Time `json:"finishedAt,omitempty" description:"time at which the container last terminated"`
	ContainerID string    `json:"containerID,omitempty" description:"container's ID in the format 'docker://<container_id>'"`
}

// ContainerState holds a possible state of container.
// Only one of its members may be specified.
// If none of them is specified, the default one is ContainerStateWaiting.
type ContainerState struct {
	Waiting     *ContainerStateWaiting    `json:"waiting,omitempty" description:"details about a waiting container"`
	Running     *ContainerStateRunning    `json:"running,omitempty" description:"details about a running container"`
	Termination *ContainerStateTerminated `json:"termination,omitempty" description:"details about a terminated container"`
}

type ContainerStatus struct {
	// TODO(dchen1107): Should we rename PodStatus to a more generic name or have a separate states
	// defined for container?
	State                ContainerState `json:"state,omitempty" description:"details about the container's current condition"`
	LastTerminationState ContainerState `json:"lastState,omitempty" description:"details about the container's last termination condition"`
	Ready                bool           `json:"ready" description:"specifies whether the container has passed its readiness probe"`
	// Note that this is calculated from dead containers.  But those containers are subject to
	// garbage collection.  This value will get capped at 5 by GC.
	RestartCount int `json:"restartCount" description:"the number of times the container has been restarted, currently based on the number of dead containers that have not yet been removed"`
	// TODO(dchen1107): Which image the container is running with?
	// The image the container is running
	Image       string `json:"image" description:"image of the container"`
	ImageID     string `json:"imageID" description:"ID of the container's image"`
	ContainerID string `json:"containerID,omitempty" description:"container's ID in the format 'docker://<container_id>'"`
}

// PodPhase is a label for the condition of a pod at the current time.
type PodPhase string

// These are the valid statuses of pods.
const (
	// PodPending means the pod has been accepted by the system, but one or more of the containers
	// has not been started. This includes time before being bound to a node, as well as time spent
	// pulling images onto the host.
	PodPending PodPhase = "Pending"
	// PodRunning means the pod has been bound to a node and all of the containers have been started.
	// At least one container is still running or is in the process of being restarted.
	PodRunning PodPhase = "Running"
	// PodSucceeded means that all containers in the pod have voluntarily terminated
	// with a container exit code of 0, and the system is not going to restart any of these containers.
	PodSucceeded PodPhase = "Succeeded"
	// PodFailed means that all containers in the pod have terminated, and at least one container has
	// terminated in a failure (exited with a non-zero exit code or was stopped by the system).
	PodFailed PodPhase = "Failed"
	// PodUnknown means that for some reason the state of the pod could not be obtained, typically due
	// to an error in communicating with the host of the pod.
	PodUnknown PodPhase = "Unknown"
)

// PodConditionType is a valid value for PodCondition.Type
type PodConditionType string

// These are valid conditions of pod.
const (
	// PodReady means the pod is able to service requests and should be added to the
	// load balancing pools of all matching services.
	PodReady PodConditionType = "Ready"
)

// TODO: add LastTransitionTime, Reason, Message to match NodeCondition api.
type PodCondition struct {
	// Type is the type of the condition
	Type PodConditionType `json:"type" description:"kind of the condition"`
	// Status is the status of the condition
	Status ConditionStatus `json:"status" description:"status of the condition, one of Full, None, Unknown"`
}

// PodInfo contains one entry for every container with available info.
type PodInfo map[string]ContainerStatus

// RestartPolicy describes how the container should be restarted.
// Only one of the following restart policies may be specified.
// If none of the following policies is specified, the default one
// is RestartPolicyAlways.
type RestartPolicy string

const (
	RestartPolicyAlways    RestartPolicy = "Always"
	RestartPolicyOnFailure RestartPolicy = "OnFailure"
	RestartPolicyNever     RestartPolicy = "Never"
)

// DNSPolicy defines how a pod's DNS will be configured.
type DNSPolicy string

const (
	// DNSClusterFirst indicates that the pod should use cluster DNS
	// first, if it is available, then fall back on the default (as
	// determined by kubelet) DNS settings.
	DNSClusterFirst DNSPolicy = "ClusterFirst"

	// DNSDefault indicates that the pod should use the default (as
	// determined by kubelet) DNS settings.
	DNSDefault DNSPolicy = "Default"
)

// PodSpec is a description of a pod
type PodSpec struct {
	Volumes []Volume `json:"volumes" description:"list of volumes that can be mounted by containers belonging to the pod"`
	// Required: there must be at least one container in a pod.
	Containers    []Container   `json:"containers" description:"list of containers belonging to the pod; cannot be updated; containers cannot currently be added or removed; there must be at least one container in a Pod"`
	RestartPolicy RestartPolicy `json:"restartPolicy,omitempty" description:"restart policy for all containers within the pod; one of RestartPolicyAlways, RestartPolicyOnFailure, RestartPolicyNever"`
	// Optional: Set DNS policy.  Defaults to "ClusterFirst"
	DNSPolicy DNSPolicy `json:"dnsPolicy,omitempty" description:"DNS policy for containers within the pod; one of 'ClusterFirst' or 'Default'"`
	// NodeSelector is a selector which must be true for the pod to fit on a node
	NodeSelector map[string]string `json:"nodeSelector,omitempty" description:"selector which must match a node's labels for the pod to be scheduled on that node"`

	// Host is a request to schedule this pod onto a specific host.  If it is non-empty,
	// the the scheduler simply schedules this pod onto that host, assuming that it fits
	// resource requirements.
	Host string `json:"host,omitempty" description:"host requested for this pod"`
}

// PodStatus represents information about the status of a pod. Status may trail the actual
// state of a system.
type PodStatus struct {
	Phase      PodPhase       `json:"phase,omitempty" description:"current condition of the pod."`
	Conditions []PodCondition `json:"Condition,omitempty" description:"current service state of pod"`
	// A human readable message indicating details about why the pod is in this state.
	Message string `json:"message,omitempty" description:"human readable message indicating details about why the pod is in this condition"`

	// Host is the name of the node that this Pod is currently bound to, or empty if no
	// assignment has been done.
	Host   string `json:"host,omitempty" description:"host to which the pod is assigned; empty if not yet scheduled; cannot be updated"`
	HostIP string `json:"hostIP,omitempty" description:"IP address of the host to which the pod is assigned; empty if not yet scheduled"`
	PodIP  string `json:"podIP,omitempty" description:"IP address allocated to the pod; routable at least within the cluster; empty if not yet allocated"`

	// The key of this map is the *name* of the container within the manifest; it has one
	// entry per container in the manifest. The value of this map is currently the output
	// of `docker inspect`. This output format is *not* final and should not be relied
	// upon.
	// TODO: Make real decisions about what our info should look like. Re-enable fuzz test
	// when we have done this.
	Info PodInfo `json:"info,omitempty" description:"map of container name to container status"`
}

// PodStatusResult is a wrapper for PodStatus returned by kubelet that can be encode/decoded
type PodStatusResult struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`
	// Status represents the current information about a pod. This data may not be up
	// to date.
	Status PodStatus `json:"status,omitempty" description:"most recently observed status of the pod; populated by the system, read-only; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`
}

// Pod is a collection of containers that can run on a host. This resource is created
// by clients and scheduled onto hosts.
type Pod struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Spec defines the behavior of a pod.
	Spec PodSpec `json:"spec,omitempty" description:"specification of the desired behavior of the pod; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`

	// Status represents the current information about a pod. This data may not be up
	// to date.
	Status PodStatus `json:"status,omitempty" description:"most recently observed status of the pod; populated by the system, read-only; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`
}

// PodList is a list of Pods.
type PodList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#types-kinds`

	Items []Pod `json:"items" description:"list of pods"`
}

// PodTemplateSpec describes the data a pod should have when created from a template
type PodTemplateSpec struct {
	// Metadata of the pods created from this template.
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Spec defines the behavior of a pod.
	Spec PodSpec `json:"spec,omitempty" description:"specification of the desired behavior of the pod; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`
}

// PodTemplate describes a template for creating copies of a predefined pod.
type PodTemplate struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Spec defines the behavior of a pod.
	Spec PodTemplateSpec `json:"spec,omitempty" description:"specification of the desired behavior of the pod; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`
}

// PodTemplateList is a list of PodTemplates.
type PodTemplateList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	Items []PodTemplate `json:"items" description:"list of pod templates"`
}

// ReplicationControllerSpec is the specification of a replication controller.
type ReplicationControllerSpec struct {
	// Replicas is the number of desired replicas.
	Replicas int `json:"replicas" description:"number of replicas desired"`

	// Selector is a label query over pods that should match the Replicas count.
	Selector map[string]string `json:"selector,omitempty" description:"label keys and values that must match in order to be controlled by this replication controller"`

	// TemplateRef is a reference to an object that describes the pod that will be created if
	// insufficient replicas are detected.
	TemplateRef *ObjectReference `json:"templateRef,omitempty" description:"reference to an object that describes the pod that will be created if insufficient replicas are detected"`

	// Template is the object that describes the pod that will be created if
	// insufficient replicas are detected. This takes precedence over a
	// TemplateRef.
	Template *PodTemplateSpec `json:"template,omitempty" description:"object that describes the pod that will be created if insufficient replicas are detected; takes precendence over templateRef"`
}

// ReplicationControllerStatus represents the current status of a replication
// controller.
type ReplicationControllerStatus struct {
	// Replicas is the number of actual replicas.
	Replicas int `json:"replicas" description:"most recently oberved number of replicas"`
}

// ReplicationController represents the configuration of a replication controller.
type ReplicationController struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Spec defines the desired behavior of this replication controller.
	Spec ReplicationControllerSpec `json:"spec,omitempty" description:"specification of the desired behavior of the replication controller; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`

	// Status is the current status of this replication controller. This data may be
	// out of date by some window of time.
	Status ReplicationControllerStatus `json:"status,omitempty" description:"most recently observed status of the replication controller; populated by the system, read-only; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`
}

// ReplicationControllerList is a collection of replication controllers.
type ReplicationControllerList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	Items []ReplicationController `json:"items" description:"list of replication controllers"`
}

// Session Affinity Type string
type AffinityType string

const (
	// AffinityTypeClientIP is the Client IP based.
	AffinityTypeClientIP AffinityType = "ClientIP"

	// AffinityTypeNone - no session affinity.
	AffinityTypeNone AffinityType = "None"
)

// ServiceStatus represents the current status of a service
type ServiceStatus struct{}

// ServiceSpec describes the attributes that a user creates on a service
type ServiceSpec struct {
	// Port is the TCP or UDP port that will be made available to each pod for connecting to the pods
	// proxied by this service.
	Port int `json:"port" description:"port exposed by the service"`

	// Optional: Supports "TCP" and "UDP".  Defaults to "TCP".
	Protocol Protocol `json:"protocol,omitempty" description:"protocol for port; must be UDP or TCP; TCP if unspecified"`

	// This service will route traffic to pods having labels matching this selector. If null, no endpoints will be automatically created. If empty, all pods will be selected.
	Selector map[string]string `json:"selector" description:"label keys and values that must match in order to receive traffic for this service; if empty, all pods are selected, if not specified, endpoints must be manually specified"`

	// PortalIP is usually assigned by the master.  If specified by the user
	// we will try to respect it or else fail the request.  This field can
	// not be changed by updates.
	// Valid values are None, empty string (""), or a valid IP address
	// None can be specified for headless services when proxying is not required
	PortalIP string `json:"portalIP,omitempty description: IP address of the service; usually assigned by the system; if specified, it will be allocated to the service if unused, and creation of the service will fail otherwise; cannot be updated; 'None' can be specified for a headless service when proxying is not required"`

	// CreateExternalLoadBalancer indicates whether a load balancer should be created for this service.
	CreateExternalLoadBalancer bool `json:"createExternalLoadBalancer,omitempty" description:"set up a cloud-provider-specific load balancer on an external IP"`

	// PublicIPs are used by external load balancers, or can be set by
	// users to handle external traffic that arrives at a node.
	PublicIPs []string `json:"publicIPs,omitempty" description:"externally visible IPs (e.g. load balancers) that should be proxied to this service"`

	// TargetPort is the name or number of the port on the container to direct traffic to.
	// This is useful if the containers the service points to have multiple open ports.
	// Optional: If unspecified, the service port is used (an identity map).
	TargetPort util.IntOrString `json:"targetPort,omitempty" description:"number or name of the port to access on the containers belonging to pods targeted by the service; defaults to the container's first open port"`

	// Optional: Supports "ClientIP" and "None".  Used to maintain session affinity.
	SessionAffinity AffinityType `json:"sessionAffinity,omitempty" description:"enable client IP based session affinity; must be ClientIP or None; defaults to None"`
}

// Service is a named abstraction of software service (for example, mysql) consisting of local port
// (for example 3306) that the proxy listens on, and the selector that determines which pods
// will answer requests sent through the proxy.
type Service struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Spec defines the behavior of a service.
	Spec ServiceSpec `json:"spec,omitempty" description:"specification of the desired behavior of the service; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`

	// Status represents the current status of a service.
	Status ServiceStatus `json:"status,omitempty" description:"most recently observed status of the service; populated by the system, read-only; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`
}

const (
	// PortalIPNone - do not assign a portal IP
	// no proxying required and no environment variables should be created for pods
	PortalIPNone = "None"
)

// ServiceList holds a list of services.
type ServiceList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	Items []Service `json:"items" description:"list of services"`
}

// Endpoints is a collection of endpoints that implement the actual service, for example:
// Name: "mysql", Endpoints: [{"ip": "10.10.1.1", "port": 1909}, {"ip": "10.10.2.2", "port": 8834}]
type Endpoints struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Optional: The IP protocol for these endpoints. Supports "TCP" and
	// "UDP".  Defaults to "TCP".
	Protocol Protocol `json:"protocol,omitempty" description:"IP protocol for endpoint ports; must be UDP or TCP; TCP if unspecified"`

	Endpoints []Endpoint `json:"endpoints,omitempty" description:"list of endpoints corresponding to a service"`
}

// Endpoint is a single IP endpoint of a service.
type Endpoint struct {
	// Required: The IP of this endpoint.
	// TODO: This should allow hostname or IP, see #4447.
	IP string `json:"ip" description:"IP of this endpoint"`

	// Required: The destination port to access.
	Port int `json:"port" description:"destination port of this endpoint"`

	// Optional: The kubernetes object related to the entry point.
	TargetRef *ObjectReference `json:"targetRef,omitempty" description:"reference to object providing the endpoint"`
}

// EndpointsList is a list of endpoints.
type EndpointsList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	Items []Endpoints `json:"items" description:"list of endpoints"`
}

// NodeSpec describes the attributes that a node is created with.
type NodeSpec struct {
	// Capacity represents the available resources of a node.
	// see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/resources.md for more details.
	Capacity ResourceList `json:"capacity,omitempty" description:"compute resource capacity of the node; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/resources.md"`
	// PodCIDR represents the pod IP range assigned to the node
	PodCIDR string `json:"podCIDR,omitempty" description:"pod IP range assigned to the node"`
	// External ID of the node assigned by some machine database (e.g. a cloud provider)
	ExternalID string `json:"externalID,omitempty" description:"external ID assigned to the node by some machine database (e.g. a cloud provider)"`
	// Unschedulable controls node schedulability of new pods. By default node is schedulable.
	Unschedulable bool `json:"unschedulable,omitempty" description:"disable pod scheduling on the node"`
}

// NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
type NodeSystemInfo struct {
	// MachineID is the machine-id reported by the node
	MachineID string `json:"machineID"`
	// SystemUUID is the system-uuid reported by the node
	SystemUUID string `json:"systemUUID"`
}

// NodeStatus is information about the current status of a node.
type NodeStatus struct {
	// NodePhase is the current lifecycle phase of the node.
	Phase NodePhase `json:"phase,omitempty" description:"most recently observed lifecycle phase of the node"`
	// Conditions is an array of current node conditions.
	Conditions []NodeCondition `json:"conditions,omitempty" description:"list of node conditions observed"`
	// Queried from cloud provider, if available.
	Addresses []NodeAddress `json:"addresses,omitempty" description:"list of addresses reachable to the node"`
	// NodeSystemInfo is a set of ids/uuids to uniquely identify the node
	NodeInfo NodeSystemInfo `json:"nodeInfo,omitempty"`
}

// NodeInfo is the information collected on the node.
type NodeInfo struct {
	TypeMeta `json:",inline"`
	// Capacity represents the available resources of a node
	Capacity ResourceList `json:"capacity,omitempty"`
	// NodeSystemInfo is a set of ids/uuids to uniquely identify the node
	NodeSystemInfo `json:",inline,omitempty"`
}

type NodePhase string

// These are the valid phases of node.
const (
	// NodePending means the node has been created/added by the system, but not configured.
	NodePending NodePhase = "Pending"
	// NodeRunning means the node has been configured and has Kubernetes components running.
	NodeRunning NodePhase = "Running"
	// NodeTerminated means the node has been removed from the cluster.
	NodeTerminated NodePhase = "Terminated"
)

type NodeConditionType string

// These are valid conditions of node. Currently, we don't have enough information to decide
// node condition. In the future, we will add more. The proposed set of conditions are:
// NodeReachable, NodeLive, NodeReady, NodeSchedulable, NodeRunnable.
const (
	// NodeReachable means the node can be reached (in the sense of HTTP connection) from node controller.
	NodeReachable NodeConditionType = "Reachable"
	// NodeReady means the node returns StatusOK for HTTP health check.
	NodeReady NodeConditionType = "Ready"
	// NodeSchedulable means the node is ready to accept new pods.
	NodeSchedulable NodeConditionType = "Schedulable"
)

type NodeCondition struct {
	Type               NodeConditionType `json:"type" description:"type of node condition, one of Reachable, Ready"`
	Status             ConditionStatus   `json:"status" description:"status of the condition, one of Full, None, Unknown"`
	LastProbeTime      util.Time         `json:"lastProbeTime,omitempty" description:"last time the condition was probed"`
	LastTransitionTime util.Time         `json:"lastTransitionTime,omitempty" description:"last time the condition transit from one status to another"`
	Reason             string            `json:"reason,omitempty" description:"(brief) reason for the condition's last transition"`
	Message            string            `json:"message,omitempty" description:"human readable message indicating details about last transition"`
}

type NodeAddressType string

// These are valid address type of node.
const (
	NodeHostName   NodeAddressType = "Hostname"
	NodeExternalIP NodeAddressType = "ExternalIP"
	NodeInternalIP NodeAddressType = "InternalIP"
)

type NodeAddress struct {
	Type    NodeAddressType `json:"type"`
	Address string          `json:"address"`
}

// ResourceName is the name identifying various resources in a ResourceList.
type ResourceName string

const (
	// CPU, in cores. (500m = .5 cores)
	ResourceCPU ResourceName = "cpu"
	// Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)
	ResourceMemory ResourceName = "memory"
)

// ResourceList is a set of (resource name, quantity) pairs.
type ResourceList map[ResourceName]resource.Quantity

// Node is a worker node in Kubernetes.
// The name of the node according to etcd is in ID.
type Node struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Spec defines the behavior of a node.
	Spec NodeSpec `json:"spec,omitempty" description:"specification of a node; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`

	// Status describes the current status of a Node
	Status NodeStatus `json:"status,omitempty" description:"most recently observed status of the node; populated by the system, read-only; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`
}

// NodeList is a list of minions.
type NodeList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	Items []Node `json:"items" description:"list of nodes"`
}

// NamespaceSpec describes the attributes on a Namespace
type NamespaceSpec struct {
}

// NamespaceStatus is information about the current status of a Namespace.
type NamespaceStatus struct {
	// Phase is the current lifecycle phase of the namespace.
	Phase NamespacePhase `json:"phase,omitempty" description:"phase is the current lifecycle phase of the namespace"`
}

type NamespacePhase string

// These are the valid phases of a namespace.
const (
	// NamespaceActive means the namespace is available for use in the system
	NamespaceActive NamespacePhase = "Active"
	// NamespaceTerminating means the namespace is undergoing graceful termination
	NamespaceTerminating NamespacePhase = "Terminating"
)

// A namespace provides a scope for Names.
// Use of multiple namespaces is optional
type Namespace struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Spec defines the behavior of the Namespace.
	Spec NamespaceSpec `json:"spec,omitempty" description:"spec defines the behavior of the Namespace; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`

	// Status describes the current status of a Namespace
	Status NamespaceStatus `json:"status,omitempty" description:"status describes the current status of a Namespace; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`
}

// NamespaceList is a list of Namespaces.
type NamespaceList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Items is the list of Namespace objects in the list
	Items []Namespace `json:"items"  description:"items is the list of Namespace objects in the list"`
}

// Binding ties one object to another - for example, a pod is bound to a node by a scheduler.
type Binding struct {
	TypeMeta `json:",inline"`
	// ObjectMeta describes the object that is being bound.
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Target is the object to bind to.
	Target ObjectReference `json:"target" description:"an object to bind to"`
}

// DeleteOptions may be provided when deleting an API object
type DeleteOptions struct {
	TypeMeta `json:",inline"`

	// Optional duration in seconds before the object should be deleted. Value must be non-negative integer.
	// The value zero indicates delete immediately. If this value is nil, the default grace period for the
	// specified type will be used.
	GracePeriodSeconds *int64 `json:"gracePeriodSeconds" description:"the duration in seconds to wait before deleting this object; defaults to a per object value if not specified; zero means delete immediately"`
}

// Status is a return value for calls that don't return other objects.
type Status struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// One of: "Success" or "Failure"
	Status string `json:"status,omitempty" description:"status of the operation; either Success, or Failure"`
	// A human-readable description of the status of this operation.
	Message string `json:"message,omitempty" description:"human-readable description of the status of this operation"`
	// A machine-readable description of why this operation is in the
	// "Failure" status. If this value is empty there
	// is no information available. A Reason clarifies an HTTP status
	// code but does not override it.
	Reason StatusReason `json:"reason,omitempty" description:"machine-readable description of why this operation is in the 'Failure' status; if this value is empty there is no information available; a reason clarifies an HTTP status code but does not override it"`
	// Extended data associated with the reason.  Each reason may define its
	// own extended details. This field is optional and the data returned
	// is not guaranteed to conform to any schema except that defined by
	// the reason type.
	Details *StatusDetails `json:"details,omitempty" description:"extended data associated with the reason; each reason may define its own extended details; this field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type"`
	// Suggested HTTP return code for this status, 0 if not set.
	Code int `json:"code,omitempty" description:"suggested HTTP return code for this status; 0 if not set"`
}

// StatusDetails is a set of additional properties that MAY be set by the
// server to provide additional information about a response. The Reason
// field of a Status object defines what attributes will be set. Clients
// must ignore fields that do not match the defined type of each attribute,
// and should assume that any attribute may be empty, invalid, or under
// defined.
type StatusDetails struct {
	// The ID attribute of the resource associated with the status StatusReason
	// (when there is a single ID which can be described).
	ID string `json:"id,omitempty" description:"the ID attribute of the resource associated with the status StatusReason (when there is a single ID which can be described)"`
	// The kind attribute of the resource associated with the status StatusReason.
	// On some operations may differ from the requested resource Kind.
	Kind string `json:"kind,omitempty" description:"the kind attribute of the resource associated with the status StatusReason; on some operations may differ from the requested resource Kind"`
	// The Causes array includes more details associated with the StatusReason
	// failure. Not all StatusReasons may provide detailed causes.
	Causes []StatusCause `json:"causes,omitempty" description:"the Causes array includes more details associated with the StatusReason failure; not all StatusReasons may provide detailed causes"`
}

// Values of Status.Status
const (
	StatusSuccess = "Success"
	StatusFailure = "Failure"
)

// StatusReason is an enumeration of possible failure causes.  Each StatusReason
// must map to a single HTTP status code, but multiple reasons may map
// to the same HTTP status code.
// TODO: move to apiserver
type StatusReason string

const (
	// StatusReasonUnknown means the server has declined to indicate a specific reason.
	// The details field may contain other information about this error.
	// Status code 500.
	StatusReasonUnknown StatusReason = ""

	// StatusReasonNotFound means one or more resources required for this operation
	// could not be found.
	// Details (optional):
	//   "kind" string - the kind attribute of the missing resource
	//                   on some operations may differ from the requested
	//                   resource.
	//   "id"   string - the identifier of the missing resource
	// Status code 404
	StatusReasonNotFound StatusReason = "NotFound"

	// StatusReasonAlreadyExists means the resource you are creating already exists.
	// Details (optional):
	//   "kind" string - the kind attribute of the conflicting resource
	//   "id"   string - the identifier of the conflicting resource
	// Status code 409
	StatusReasonAlreadyExists StatusReason = "AlreadyExists"

	// StatusReasonConflict means the requested update operation cannot be completed
	// due to a conflict in the operation. The client may need to alter the request.
	// Each resource may define custom details that indicate the nature of the
	// conflict.
	// Status code 409
	StatusReasonConflict StatusReason = "Conflict"

	// StatusReasonInvalid means the requested create or update operation cannot be
	// completed due to invalid data provided as part of the request. The client may
	// need to alter the request. When set, the client may use the StatusDetails
	// message field as a summary of the issues encountered.
	// Details (optional):
	//   "kind" string - the kind attribute of the invalid resource
	//   "id"   string - the identifier of the invalid resource
	//   "causes"      - one or more StatusCause entries indicating the data in the
	//                   provided resource that was invalid.  The code, message, and
	//                   field attributes will be set.
	// Status code 422
	StatusReasonInvalid StatusReason = "Invalid"

	// StatusReasonServerTimeout means the server can be reached and understood the request,
	// but cannot complete the action in a reasonable time. The client should retry the request.
	// This is may be due to temporary server load or a transient communication issue with
	// another server. Status code 500 is used because the HTTP spec provides no suitable
	// server-requested client retry and the 5xx class represents actionable errors.
	// Details (optional):
	//   "kind" string - the kind attribute of the resource being acted on.
	//   "id"   string - the operation that is being attempted.
	// Status code 500
	StatusReasonServerTimeout StatusReason = "ServerTimeout"
)

// StatusCause provides more information about an api.Status failure, including
// cases when multiple errors are encountered.
type StatusCause struct {
	// A machine-readable description of the cause of the error. If this value is
	// empty there is no information available.
	Type CauseType `json:"reason,omitempty" description:"machine-readable description of the cause of the error; if this value is empty there is no information available"`
	// A human-readable description of the cause of the error.  This field may be
	// presented as-is to a reader.
	Message string `json:"message,omitempty" description:"human-readable description of the cause of the error; this field may be presented as-is to a reader"`
	// The field of the resource that has caused this error, as named by its JSON
	// serialization. May include dot and postfix notation for nested attributes.
	// Arrays are zero-indexed.  Fields may appear more than once in an array of
	// causes due to fields having multiple errors.
	// Optional.
	//
	// Examples:
	//   "name" - the field "name" on the current resource
	//   "items[0].name" - the field "name" on the first array entry in "items"
	Field string `json:"field,omitempty" description:"field of the resource that has caused this error, as named by its JSON serialization; may include dot and postfix notation for nested attributes; arrays are zero-indexed; fields may appear more than once in an array of causes due to fields having multiple errors"`
}

// CauseType is a machine readable value providing more detail about what
// occured in a status response. An operation may have multiple causes for a
// status (whether Failure or Success).
type CauseType string

const (
	// CauseTypeFieldValueNotFound is used to report failure to find a requested value
	// (e.g. looking up an ID).
	CauseTypeFieldValueNotFound CauseType = "FieldValueNotFound"
	// CauseTypeFieldValueRequired is used to report required values that are not
	// provided (e.g. empty strings, null values, or empty arrays).
	CauseTypeFieldValueRequired CauseType = "FieldValueRequired"
	// CauseTypeFieldValueDuplicate is used to report collisions of values that must be
	// unique (e.g. unique IDs).
	CauseTypeFieldValueDuplicate CauseType = "FieldValueDuplicate"
	// CauseTypeFieldValueInvalid is used to report malformed values (e.g. failed regex
	// match).
	CauseTypeFieldValueInvalid CauseType = "FieldValueInvalid"
	// CauseTypeFieldValueNotSupported is used to report valid (as per formatting rules)
	// values that can not be handled (e.g. an enumerated string).
	CauseTypeFieldValueNotSupported CauseType = "FieldValueNotSupported"
)

// ObjectReference contains enough information to let you inspect or modify the referred object.
type ObjectReference struct {
	Kind            string    `json:"kind,omitempty" description:"kind of the referent"`
	Namespace       string    `json:"namespace,omitempty" description:"namespace of the referent"`
	Name            string    `json:"name,omitempty" description:"name of the referent"`
	UID             types.UID `json:"uid,omitempty" description:"uid of the referent"`
	APIVersion      string    `json:"apiVersion,omitempty" description:"API version of the referent"`
	ResourceVersion string    `json:"resourceVersion,omitempty" description:"specific resourceVersion to which this reference is made, if any: https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#concurrency-control-and-consistency"`

	// Optional. If referring to a piece of an object instead of an entire object, this string
	// should contain information to identify the sub-object. For example, if the object
	// reference is to a container within a pod, this would take on a value like:
	// "spec.containers{name}" (where "name" refers to the name of the container that triggered
	// the event) or if no container name is specified "spec.containers[2]" (container with
	// index 2 in this pod). This syntax is chosen only to have some well-defined way of
	// referencing a part of an object.
	// TODO: this design is not final and this field is subject to change in the future.
	FieldPath string `json:"fieldPath,omitempty" description:"if referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]"`
}

type EventSource struct {
	// Component from which the event is generated.
	Component string `json:"component,omitempty" description:"component that generated the event"`
	// Host name on which the event is generated.
	Host string `json:"host,omitempty" description:"name of the host where the event is generated"`
}

// Event is a report of an event somewhere in the cluster.
// TODO: Decide whether to store these separately or with the object they apply to.
type Event struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Required. The object that this event is about.
	InvolvedObject ObjectReference `json:"involvedObject,omitempty" description:"object this event is about"`

	// Optional; this should be a short, machine understandable string that gives the reason
	// for this event being generated.
	// TODO: provide exact specification for format.
	Reason string `json:"reason,omitempty" description:"short, machine understandable string that gives the reason for the transition into the object's current status"`

	// Optional. A human-readable description of the status of this operation.
	// TODO: decide on maximum length.
	Message string `json:"message,omitempty" description:"human-readable description of the status of this operation"`

	// Optional. The component reporting this event. Should be a short machine understandable string.
	Source EventSource `json:"source,omitempty" description:"component reporting this event"`

	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	FirstTimestamp util.Time `json:"firstTimestamp,omitempty" description:"the time at which the event was first recorded"`

	// The time at which the most recent occurance of this event was recorded.
	LastTimestamp util.Time `json:"lastTimestamp,omitempty" description:"the time at which the most recent occurance of this event was recorded"`

	// The number of times this event has occurred.
	Count int `json:"count,omitempty" description:"the number of times this event has occurred"`
}

// EventList is a list of events.
type EventList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	Items []Event `json:"items" description:"list of events"`
}

// List holds a list of objects, which may not be known by the server.
type List struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	Items []runtime.RawExtension `json:"items" description:"list of objects"`
}

// A type of object that is limited
type LimitType string

const (
	// Limit that applies to all pods in a namespace
	LimitTypePod LimitType = "Pod"
	// Limit that applies to all containers in a namespace
	LimitTypeContainer LimitType = "Container"
)

// LimitRangeItem defines a min/max usage limit for any resource that matches on kind
type LimitRangeItem struct {
	// Type of resource that this limit applies to
	Type LimitType `json:"type,omitempty" description:"type of resource that this limit applies to"`
	// Max usage constraints on this kind by resource name
	Max ResourceList `json:"max,omitempty" description:"max usage constraints on this kind by resource name"`
	// Min usage constraints on this kind by resource name
	Min ResourceList `json:"min,omitempty" description:"min usage constraints on this kind by resource name"`
}

// LimitRangeSpec defines a min/max usage limit for resources that match on kind
type LimitRangeSpec struct {
	// Limits is the list of LimitRangeItem objects that are enforced
	Limits []LimitRangeItem `json:"limits" description:"limits is the list of LimitRangeItem objects that are enforced"`
}

// LimitRange sets resource usage limits for each kind of resource in a Namespace
type LimitRange struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Spec defines the limits enforced
	Spec LimitRangeSpec `json:"spec,omitempty" description:"spec defines the limits enforced; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`
}

// LimitRangeList is a list of LimitRange items.
type LimitRangeList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Items is a list of LimitRange objects
	Items []LimitRange `json:"items" description:"items is a list of LimitRange objects"`
}

// The following identify resource constants for Kubernetes object types
const (
	// Pods, number
	ResourcePods ResourceName = "pods"
	// Services, number
	ResourceServices ResourceName = "services"
	// ReplicationControllers, number
	ResourceReplicationControllers ResourceName = "replicationcontrollers"
	// ResourceQuotas, number
	ResourceQuotas ResourceName = "resourcequotas"
)

// ResourceQuotaSpec defines the desired hard limits to enforce for Quota
type ResourceQuotaSpec struct {
	// Hard is the set of desired hard limits for each named resource
	Hard ResourceList `json:"hard,omitempty" description:"hard is the set of desired hard limits for each named resource"`
}

// ResourceQuotaStatus defines the enforced hard limits and observed use
type ResourceQuotaStatus struct {
	// Hard is the set of enforced hard limits for each named resource
	Hard ResourceList `json:"hard,omitempty" description:"hard is the set of enforced hard limits for each named resource"`
	// Used is the current observed total usage of the resource in the namespace
	Used ResourceList `json:"used,omitempty" description:"used is the current observed total usage of the resource in the namespace"`
}

// ResourceQuota sets aggregate quota restrictions enforced per namespace
type ResourceQuota struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Spec defines the desired quota
	Spec ResourceQuotaSpec `json:"spec,omitempty" description:"spec defines the desired quota; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`

	// Status defines the actual enforced quota and its current usage
	Status ResourceQuotaStatus `json:"status,omitempty" description:"status defines the actual enforced quota and current usage; https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#spec-and-status"`
}

// ResourceQuotaList is a list of ResourceQuota items
type ResourceQuotaList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Items is a list of ResourceQuota objects
	Items []ResourceQuota `json:"items" description:"items is a list of ResourceQuota objects"`
}

// Secret holds secret data of a certain type.  The total bytes of the values in
// the Data field must be less than MaxSecretSize bytes.
type Secret struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty" description:"standard object metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	// Data contains the secret data.  Each key must be a valid DNS_SUBDOMAIN.
	// The serialized form of the secret data is a base64 encoded string,
	// representing the arbitrary (possibly non-string) data value here.
	Data map[string][]byte `json:"data,omitempty" description:"data contains the secret data.  Each key must be a valid DNS_SUBDOMAIN.  Each value must be a base64 encoded string"`

	// Used to facilitate programatic handling of secret data.
	Type SecretType `json:"type,omitempty" description:"type facilitates programmatic handling of secret data"`
}

const MaxSecretSize = 1 * 1024 * 1024

type SecretType string

const (
	SecretTypeOpaque SecretType = "Opaque" // Default; arbitrary user-defined data
)

type SecretList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty" description:"standard list metadata; see https://github.com/GoogleCloudPlatform/kubernetes/blob/master/docs/api-conventions.md#metadata"`

	Items []Secret `json:"items" description:"items is a list of secret objects"`
}
