package xRPC

type discoveryP2P struct {
    option *DiscoveryOption
}

func (d discoveryP2P) GetConnection() Connection {
    panic("implement me")
}

