package main

import (
    "context"
    "github.com/DGHeroin/xRPC"
    "log"
)

func main()  {
    dis := xRPC.NewPeer2PeerDiscovery("127.0.0.1:3100", nil)
    client := xRPC.NewClient(dis)
    client.Call(context.Background(), "serviceName", "method", nil, nil)
    log.Println(dis)
}