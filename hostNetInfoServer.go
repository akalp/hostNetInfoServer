package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
)

var (
	port     = flag.String("port", "8080", "the port to listen on")
	ifPrefix = flag.String("if-prefix", "", "the prefix to match for network interface names")
)

type AddressInfo struct {
	IPv4 []string `json:"ipv4"`
	IPv6 []string `json:"ipv6"`
}

type InterfaceInfo struct {
	Name      string      `json:"name"`
	MTU       int         `json:"mtu"`
	Flags     string      `json:"flags"`
	Addresses AddressInfo `json:"addresses"`
}

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ifaces, err := net.Interfaces()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var ifacesInfo []InterfaceInfo

		for _, iface := range ifaces {
			if *ifPrefix != "" && !startsWith(iface.Name, *ifPrefix) {
				continue
			}

			ifaceInfo := InterfaceInfo{
				Name:  iface.Name,
				MTU:   iface.MTU,
				Flags: iface.Flags.String(),
			}

			addrs, err := iface.Addrs()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var v4Addrs []string
			var v6Addrs []string

			for _, addr := range addrs {
				ip, _, err := net.ParseCIDR(addr.String())
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				if ip.To4() != nil {
					v4Addrs = append(v4Addrs, ip.String())
				} else {
					v6Addrs = append(v6Addrs, ip.String())
				}
			}

			ifaceInfo.Addresses = AddressInfo{
				IPv4: v4Addrs,
				IPv6: v6Addrs,
			}

			ifacesInfo = append(ifacesInfo, ifaceInfo)
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(ifacesInfo)
		if err != nil {
			return
		}
	})

	addr := ":" + *port
	fmt.Printf("Serving on port %s\n", *port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func init() {
	if p := os.Getenv("HNIS_PORT"); p != "" {
		port = &p
	}
	if ifp := os.Getenv("HNIS_IF_PREFIX"); ifp != "" {
		ifPrefix = &ifp
	}
}
