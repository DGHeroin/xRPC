package main

import (
    "github.com/DGHeroin/xRPC"
    "net"
)

func main()  {
    s := xRPC.NewServer()
    ln, _ := net.Listen("tcp", ":1334")
    s.ListenAndSerer(ln)
}
