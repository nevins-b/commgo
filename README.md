commgo
======

"command" + "mgo" = "commgo" ( comm-an-go )

This is a collection of structs designed as add-ons to the mgo driver for use with MongoDB's more administrative commands.

<code>

import (
       "github.com/samantharitter/commgo"
       "fmt"
       "labix.org/v2/mgo"
)

func main() {
        session, err := mgo.Dial("localhost:27017")
        if err != nil { panic(err) }

        status := commgo.ServerStatus{}
        err = session.Run("serverStatus", &status)
        if err != nil { panic(err) }

        fmt.Println("\nRunning " + status.Process + " version " + status.Version + "\n")

}

</code>
