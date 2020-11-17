package managementzones

// SimpleHostTechType Predefined technology, if technology is not predefined, then the verbatim type must be set
type SimpleHostTechType string

// SimpleHostTechTypes offers the known enum values
var SimpleHostTechTypes = struct {
	Apparmor            SimpleHostTechType
	Bosh                SimpleHostTechType
	Boshbpm             SimpleHostTechType
	CloudFoundry        SimpleHostTechType
	Containerd          SimpleHostTechType
	Crio                SimpleHostTechType
	DiegoCell           SimpleHostTechType
	Docker              SimpleHostTechType
	Garden              SimpleHostTechType
	Grsecurity          SimpleHostTechType
	Kubernetes          SimpleHostTechType
	Openshift           SimpleHostTechType
	OpenStackCompute    SimpleHostTechType
	OpenStackController SimpleHostTechType
	Selinux             SimpleHostTechType
}{
	"APPARMOR",
	"BOSH",
	"BOSHBPM",
	"CLOUDFOUNDRY",
	"CONTAINERD",
	"CRIO",
	"DIEGO_CELL",
	"DOCKER",
	"GARDEN",
	"GRSECURITY",
	"KUBERNETES",
	"OPENSHIFT",
	"OPENSTACK_COMPUTE",
	"OPENSTACK_CONTROLLER",
	"SELINUX",
}
