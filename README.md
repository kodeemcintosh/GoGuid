# GoGuid
A tiny Guid library for golang

Example:

import(
  "fmt"
  GG "github.com/kvmac/GoGuid"
)

guid type variable assignment
var guid = GG.Guid{}

func main() {

  // guid creation
  guid.New()


  // Guid type needs to be explicitly converted back to a string in order to print
  strGuid := string(guid)


  fmt.Print(strGuid)
  

}
