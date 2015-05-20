// peer cache library
// AQ <aq@okaq.com>
// 2015-05-20
package main

import (
    "encoding/json"
    "sync"
)

const (
    PEER_MOTD = "okaq peer cache ok!"
)

var (
    P *Peers
)

type Peers struct {
    Ids []string
    Users int
    Index int
    sync.Mutex
}

func NewPeers() *Peers {
    p := new(Peers)
    p.Users = 8
    p.Index = 0
    p.Ids = make([]string, p.Users)
    return p
}

func (p *Peers) Add(s0 string) {
    p.Lock()
    defer p.Unlock()
    if p.Index <= (p.Users - 1) {
        p.Ids[p.Index] = s0
        if (len(s0) > 0) {
            p.Index = p.Index + 1
        }
    }
}

func (p *Peers) Json() ([]byte, error) {
    p.Lock()
    defer p.Unlock()
    b0, err := json.Marshal(p)
    if err != nil {
        return nil, err
    }
    return b0, nil
}

