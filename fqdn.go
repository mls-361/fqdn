/*
------------------------------------------------------------------------------------------------------------------------
####### fqdn ####### (c) 2020-2021 mls-361 ######################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package fqdn

import (
	"net"
	"os"
	"strings"

	"github.com/mls-361/failure"
)

func fqdn() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	addrs, err := net.LookupHost(hostname)
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		hosts, err := net.LookupAddr(addr)
		if err != nil || len(hosts) == 0 {
			continue
		}

		return strings.TrimSuffix(hosts[0], "."), nil
	}

	return "", nil
}

// FQDN AFAIRE.
func FQDN() (string, error) {
	fqdn, err := fqdn()
	if fqdn == "" || err != nil {
		return "", failure.New(err).
			Msg("impossible to retrieve the FQDN") /////////////////////////////////////////////////////////////////////
	}

	return fqdn, nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
