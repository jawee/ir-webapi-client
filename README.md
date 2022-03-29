# ir-webapi-client

Library for calling the iRacing.com web api for getting data regarding races, members etc.

Basically just a stub for now. Do not use.

#### Usage

```golang
package main

import (
        "github.com/jawee/ir-webapi-client/client"
)

func main() {
    client := client.New()

    err := client.Login("user@email.com", "password")

    resp, err := client.GetMemberRecentRaces(1234))
}
```
