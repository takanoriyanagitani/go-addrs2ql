package netif

import (
	"net"

	gm "github.com/takanoriyanagitani/go-addrs2ql/graph/model"
)

func AddrToModel(a net.Addr) (ret gm.Addr) {
	if nil == a {
		return
	}

	ret.Network = a.Network()
	ret.String = a.String()
	return
}

func AddrsByIfaceIndex(ix string) ([]*gm.Addr, error) {
	iface, e := IfaceByIndex(ix)
	if nil != e {
		return nil, e
	}

	addrs, e := iface.Addrs()
	if nil != e {
		return nil, e
	}

	ret := make([]*gm.Addr, 0, len(addrs))

	for _, addr := range addrs {
		var converted gm.Addr = AddrToModel(addr)
		ret = append(ret, &converted)
	}

	return ret, nil
}
