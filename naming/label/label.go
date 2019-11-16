package label

import (
	"strconv"

	"github.com/gosimple/slug"
)

// Tag represent the tag string
type Label string

func (l Label) String() string {
	return string(l)
}

const (
	prefix = "vultron-"
)

func bindPrefix(clusterName Label, suffix string) string {
	return prefix + slug.Make(string(clusterName))
}

// GenerateServerLabel generate a server name
func GenerateServerLabel(clusterName Label, iteration int) string {
	return bindPrefix(clusterName, "server-"+strconv.Itoa(iteration))
}

// GenerateClientLabel generate a client label
func GenerateClientLabel(clusterName Label, iteration int) string {
	return bindPrefix(clusterName, "client-"+strconv.Itoa(iteration))
}
