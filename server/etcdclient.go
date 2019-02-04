package main

import (
	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/clientv3"
)

const (
	ETCD2 = "v2" // etcd version 2
	ETCD3 = "v3" // etcd version 3
)

type clientManager struct{
	clients map[string]client.Client
	v3clients map[string]clientv3.Client

}

type clientConfig struct{
	Name string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	URL string `json:"url,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	CA string `json:"ca,omitempty"`
	Cert string `json:"cert,omitempty"`
	CertKey string `json:"cert_key,omitempty"`
	Timeout int `json:"timeout,omitempty"`
}
