package ibmcos

import (
	"fmt"
)

// Since there is no API to query these, we have to hard-code them here.

// PowerVSRegion describes resources associated with a region in Power VS.
// We're using a few items from the IBM Cloud VPC offering. The region names
// for VPC are different so another function of this is to correlate those.
type PowerVSRegion struct {
	Description string
	VPCRegion   string
	Zones       []string
}

// PowerVSRegions provides a mapping between Power VS and IBM Cloud VPC regions.
var PowerVSRegions = map[string]PowerVSRegion{
	"dal": {
		Description: "Dallas, USA",
		VPCRegion:   "us-south",
		Zones:       []string{"dal12"},
	},
	"eu-de": {
		Description: "Frankfurt, Germany",
		VPCRegion:   "eu-de",
		Zones: []string{
			"eu-de-1",
			"eu-de-2",
		},
	},
	"lon": {
		Description: "London, UK.",
		VPCRegion:   "eu-gb",
		Zones: []string{
			"lon04",
			"lon06",
		},
	},
	"mon": {
		Description: "Montreal, Canada",
		VPCRegion:   "ca-tor",
		Zones:       []string{"mon01"},
	},
	"osa": {
		Description: "Osaka, Japan",
		VPCRegion:   "jp-osa",
		Zones:       []string{"osa21"},
	},
	"syd": {
		Description: "Sydney, Australia",
		VPCRegion:   "au-syd",
		Zones: []string{
			"syd04",
			"syd05",
		},
	},
	"sao": {
		Description: "São Paulo, Brazil",
		VPCRegion:   "br-sao",
		Zones:       []string{"sao01"},
	},
	"tok": {
		Description: "Tokyo, Japan",
		VPCRegion:   "jp-tok",
		Zones:       []string{"tok04"},
	},
	"us-east": {
		Description: "Washington DC, USA",
		VPCRegion:   "us-east",
		Zones:       []string{"us-east"},
	},
}

// VPCRegionForPowerVSRegion returns the VPC region for the specified PowerVS region.
func VPCRegionForPowerVSRegion(region string) (string, error) {
	if r, ok := PowerVSRegions[region]; ok {
		return r.VPCRegion, nil
	}

	return "", fmt.Errorf("VPC region corresponding to a PowerVS region %s not found ", region)
}
