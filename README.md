# ir-webapi-client

Library for calling the iRacing.com web api for getting data regarding races, members etc.

Basically just a stub for now. Do not use.

#### Usage

```golang
package main

import (
        "log"

        "github.com/jawee/ir-webapi-client/client"
)

func main() {
    client := client.New()

    err := client.Login("user@email.com", "password")
    if err != nil {
        panic(err)
    }

    resp, err := client.GetMemberRecentRaces(1234)
    if err != nil {
        panic(err)
    }

    for _, r := range resp.Races {
        log.Printf("%s. %s. %d -> %d\n", r.SessionStartTime, r.SeriesName, r.OldiRating, r.NewiRating);
    }
}
```
