package GoGuid

import(
	GG "github.com/kvmac/GoGuid"
	"testing"
	"regexp"
)

func TestGuid(t *testing.T) {

	expected := "[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}"

	g := GG.Guid
	g.New()
	actual := string(g)

	if regexp.MatchString(expected, actual) != true {
		fmt.Println(actual)
		t.Error("GoGuid Failed to validate")
	}
}