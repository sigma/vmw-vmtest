package main

import (
	"fmt"

	"github.com/sigma/vmw-guestinfo/rpcvmx"
	"github.com/sigma/vmw-guestinfo/vmcheck"
)

func main() {
	if !vmcheck.IsVirtualWorld() {
		fmt.Println("You are not in a VMware virtual world... :(")
		return
	}

	versions := []string{
		"something weird that seems to act as VMware stuff",
		"VMware Express",
		"VMware ESX(i)",
		"VMware Server",
		"VMware Workstation/Fusion",
		"VMware ACE",
	}

	_, typ := vmcheck.GetVersion()

	v := int(typ)
	if v >= len(versions) {
		v = 0
	}
	fmt.Println("You are running on", versions[v])

	ovfenv := rpcvmx.ConfigGetString("ovfenv", "")
	if ovfenv != "" {
		fmt.Println("An OVF environment is defined:", ovfenv)
	}
}
