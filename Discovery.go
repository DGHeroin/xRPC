package xRPC

type Discovery interface{
    GetConnection() Connection
}
type DiscoveryOption struct {
    server string
}

func defaultDiscoveryOption() *DiscoveryOption {
    return &DiscoveryOption{
    
    }
}

func NewPeer2PeerDiscovery(addr string, o *DiscoveryOption) Discovery {
    if o == nil {
        o = defaultDiscoveryOption()
    }
    o.server = addr
    d := &discoveryP2P{
        option: o,
    }
    
    return d
}
