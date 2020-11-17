package managementzones

// HypervisorTypeComparisionValue The value to compare to.
type HypervisorTypeComparisionValue string

// HypervisorTypeComparisionValues offers the known enum values
var HypervisorTypeComparisionValues = struct {
	Ahv        HypervisorTypeComparisionValue
	HyperV     HypervisorTypeComparisionValue
	Kvm        HypervisorTypeComparisionValue
	Lpar       HypervisorTypeComparisionValue
	Qemu       HypervisorTypeComparisionValue
	VirtualBox HypervisorTypeComparisionValue
	Vmware     HypervisorTypeComparisionValue
	Wpar       HypervisorTypeComparisionValue
	Xen        HypervisorTypeComparisionValue
}{
	"AHV",
	"HYPER_V",
	"KVM",
	"LPAR",
	"QEMU",
	"VIRTUAL_BOX",
	"VMWARE",
	"WPAR",
	"XEN",
}
