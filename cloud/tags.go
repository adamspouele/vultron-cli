package cloud

import (
	"strconv"

	"github.com/gosimple/slug"
)

// GenerateServerTag generate a tag name
func GenerateServerTag(clusterName string, number int) string {
	return slug.Make(clusterName) + "-server-" + strconv.Itoa(number)
}

// GenerateClientTag generate a tag name
func GenerateClientTag(clusterName string, number int) string {
	return slug.Make(clusterName) + "-client-" + strconv.Itoa(number)

}
