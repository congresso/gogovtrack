// Package gogovtrack examples:
//
//    client := gogovtrack.Client{
//      Client: http.DefaultClient,
//    }
//
//     b, err := client.GetBill(1)
//     if err != nil {
//       log.Fatal(err)
//     }
//
//     cs, err := client.GetCommittess()
//     if err != nil {
//       log.Fatal(err)
//     }
//
//     rs, err := client.Filter(gogovtrack.Q{"limit": "20", "id": "1"})
//     if err != nil {
//       log.Fatal(err)
//     }
package gogovtrack
