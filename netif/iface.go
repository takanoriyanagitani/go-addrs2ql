package netif

import (
	"iter"
	"net"
	"regexp"
	"slices"
	"strconv"

	gm "github.com/takanoriyanagitani/go-addrs2ql/graph/model"
)

func IfaceToModel(i net.Interface) (ret gm.NetInterface) {
	var ix string = strconv.Itoa(i.Index)
	return gm.NetInterface{
		Index:        ix,
		Mtu:          int32(i.MTU),
		Name:         i.Name,
		HardwareAddr: i.HardwareAddr.String(),
	}
}

func IfaceByIndex(ix string) (*net.Interface, error) {
	i, e := strconv.Atoi(ix)
	if nil != e {
		return nil, e
	}

	return net.InterfaceByIndex(i)
}

func IfaceByName(name string) (*net.Interface, error) {
	return net.InterfaceByName(name)
}

func IfacesByPattern(pat string) ([]net.Interface, error) {
	if 0 == len(pat) {
		return IfacesAll()
	}

	compiled, e := regexp.Compile(pat)
	if nil != e {
		return nil, e
	}

	all, e := IfacesAll()
	if nil != e {
		return nil, e
	}

	var iall iter.Seq[net.Interface] = slices.Values(all)

	var filtered iter.Seq[net.Interface] = func(yield func(net.Interface) bool) {
		for iface := range iall {
			var name string = iface.Name
			var found bool = compiled.MatchString(name)
			if !found {
				continue
			}

			if !yield(iface) {
				return
			}
		}
	}

	return slices.Collect(filtered), nil
}

func IfacesAll() ([]net.Interface, error) {
	return net.Interfaces()
}
