package GoGuid

import(
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

// Guid type declaration
type Guid string


// New Guid method that creates a brand new guid for the Guid type that runs it
func (g *Guid)New() {
	newGuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, newGuid)

	if n != len(newGuid) || err != nil {
		log.Fatal(err)
	} else {

		// Generates variant bits
		newGuid[8] = newGuid[8]&^0xc0 | 0x80

		// Pseudo-random generation of bits
		newGuid[6] = newGuid[6]&^0xf0 | 0x40

		*g = Guid(fmt.Sprintf("%x-%x-%x-%x-%x", newGuid[0:4], newGuid[4:6], newGuid[6:8], newGuid[8:10], newGuid[10:]))
	}
}