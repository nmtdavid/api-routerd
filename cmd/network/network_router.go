// SPDX-License-Identifier: Apache-2.0

package network

import (
	"api-routerd/cmd/network/networkd"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NetworkLinkGet(rw http.ResponseWriter, req *http.Request) {
	link := new(Link)

	r := json.NewDecoder(req.Body).Decode(&link);
	if r != nil {
		log.Error("Failed to find decode json: ", r)
		rw.Write([]byte("500: " + r.Error()))
		return
	}

	switch req.Method {
	case "GET":
		link.GetLink(rw)
		break
	}
}

func NetworkLinkAdd(rw http.ResponseWriter, req *http.Request) {
	link := new(Link)

	r := json.NewDecoder(req.Body).Decode(&link);
	if r != nil {
		log.Error("Failed to find decode json: ", r)
		rw.Write([]byte("500: " + r.Error()))
		return
	}

	switch req.Method {
	case "PUT":
		switch link.Action {
		case "add-link-bridge":
			r := link.LinkCreateBridge()
			if r != nil {
				rw.Write([]byte("500: " + r.Error()))
			}
		}
	}
}

func NetworkLinkDelete(rw http.ResponseWriter, req *http.Request) {
	link := new(Link)

	r := json.NewDecoder(req.Body).Decode(&link);
	if r != nil {
		log.Error("Failed to find decode json: ", r)
		rw.Write([]byte("500: " + r.Error()))
		return
	}

	switch req.Method {
	case "DELETE":
		switch link.Action {
		case "delete-link":
			r := link.LinkDelete()
			if r != nil {
				rw.Write([]byte("500: " + r.Error()))
			}
		}
	}
}

func NetworkLinkSet(rw http.ResponseWriter, req *http.Request) {
	link := new(Link)

	r := json.NewDecoder(req.Body).Decode(&link);
	if r != nil {
		log.Error("Failed to find decode json: ", r)
		rw.Write([]byte("500: " + r.Error()))
		return
	}

	switch req.Method {
	case "POST":
		switch link.Action {
		case "set-link-up", "set-link-down", "set-link-mtu":
			r := link.SetLink()
			if r != nil {
				rw.Write([]byte("500: " + r.Error()))
			}
		}
	}
}

func NetworkGetAddress(rw http.ResponseWriter, req *http.Request) {
	address := new(Address)

	r := json.NewDecoder(req.Body).Decode(&address);
	if r != nil {
		log.Error("Failed to find decode json: ", r)
		rw.Write([]byte("500: " + r.Error()))
		return
	}

	switch req.Method {
	case "GET":
		GetAddress(rw, address)
		break
	}
}

func NetworkAddAddress(rw http.ResponseWriter, req *http.Request) {
	address := new(Address)

	r := json.NewDecoder(req.Body).Decode(&address);
	if r != nil {
		log.Error("Failed to find decode json: ", r)
		rw.Write([]byte("500: " + r.Error()))
		return
	}

	switch req.Method {
	case "PUT":
		switch address.Action {
		case "add-address":
			AddAddress(address)
			break
		}
	}
}

func NetworkDeleteAddres(rw http.ResponseWriter, req *http.Request) {
	address := new(Address)

	r := json.NewDecoder(req.Body).Decode(&address);
	if r != nil {
		log.Error("Failed to find decode json: ", r)
		rw.Write([]byte("500: " + r.Error()))
		return
	}

	switch req.Method {
	case "DELETE":
		DelAddress(address)
		break
	}
}

func NetworkAddRoute(rw http.ResponseWriter, req *http.Request) {
	route := new(Route)

	r := json.NewDecoder(req.Body).Decode(&route);
	if r != nil {
		log.Error("Failed to find decode json: ", r)
		rw.Write([]byte("500:" + r.Error()))
		return
	}

	switch req.Method {
	case "PUT":
		switch route.Action {
		case "add-default-gw":
			AddDefaultGateWay(route)
			break
		}
	}
}

func NetworkDeleteRoute(rw http.ResponseWriter, req *http.Request) {
	route := new(Route)

	r := json.NewDecoder(req.Body).Decode(&route);
	if r != nil {
		log.Error("Failed to find decode json: ", r)
		rw.Write([]byte("500:" + r.Error()))
		return
	}

	switch req.Method {
	case "DELETE":
		switch route.Action {
		case "del-default-gw":
			DelDefaultGateWay(route)
			break
		}
	}
}

func NetworkdConfigureNetwork(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PUT":
		networkd.ConfigureNetworkFile(rw, req)
		break
	}
}

func NetworkdConfigureNetDev(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PUT":
		networkd.ConfigureNetDevFile(rw, req)
		break
	}
}

func NetworkConfigureEthtool(rw http.ResponseWriter, req *http.Request) {
	ethtool := new(Ethtool)

	r := json.NewDecoder(req.Body).Decode(&ethtool);
	if r != nil {
		log.Error("Failed to decode json for ethtool: ", r)
		rw.Write([]byte("500:" + r.Error()))
		return
	}

	switch req.Method {
	case "GET":
		ethtool.GetEthTool(rw)
		break
	}
}

func RegisterRouterNetwork(router *mux.Router) {
	n := router.PathPrefix("/network").Subrouter()
	n.HandleFunc("/link/set", NetworkLinkSet)
	n.HandleFunc("/link/add", NetworkLinkAdd)
	n.HandleFunc("/link/delete", NetworkLinkDelete)
	n.HandleFunc("/link/get", NetworkLinkGet)

	n.HandleFunc("/address/get", NetworkGetAddress)
	n.HandleFunc("/address/add", NetworkAddAddress)

	n.HandleFunc("/route/add", NetworkAddRoute)
	n.HandleFunc("/route/del", NetworkDeleteRoute)

	// systemd-networkd
	networkd.InitNetworkd()
	n.HandleFunc("/networkd/network", NetworkdConfigureNetwork)
	n.HandleFunc("/networkd/netdev", NetworkdConfigureNetDev)

	// ethtool
	n.HandleFunc("/ethtool/get", NetworkConfigureEthtool)
}
