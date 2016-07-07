// Package gogovtrack offers an api wrapper for the govtrack http api.
//
// The API is based around resources. Govtrack has the following resources `bill`, `committee`, `committee_member`, `role`, `person`, `vote`, `vote_voter`, `cosponsorhip`
// You must initialize the resource you want to use then make your request.
//
// Examples of how to use the api:
//
//     api := gogovtrack.NewAPI(http.DefaultClient, gogovtrack.BaseURL)
//
//Examples of how to get one bill, all bills, and filtered bills:
//
//     api.Br().One("id")
//     api.Br().All()
//     api.Br().Filter(gogovtrack.Q{"limit": "15", "order_by": "-current_status_date""}).All()
//
// Examples of how to get one committes, all committes, filtered committes
//
//     api.Cr().One("id")
//     api.Cr().All()
//     api.Cr().Filter(gogovtrack.Q{"limit": "10"}).All()
package gogovtrack
