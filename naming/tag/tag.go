package tag

import "reflect"

// Tag represent the tag string
type Tag string

// TagProp define a prop string
type TagProp Tag

func (t Tag) String() string {
	return string(t)
}

const (
	TagPropNodeKind TagProp = "node-type"
)

/*

legend :
k = key
v = value

kind of dynamic Tags
====================

vultron:prop:k:v
vultron:prop:v

static Tags
===========

vultron:res:v

vultron:res:cluster
vultron:res:server
vultron:res:client

*/

//const PropTypes = map

const (
	tagPrefix   = "vultron:"
	resourceTag = "res:"
	propTag     = "prop:"
)

func bindPrefix(suffix Tag) string {
	return tagPrefix + suffix
}

func bindResourcePrefix(resourceName Tag) string {
	return bindPrefix(resourceTag + resourceName)
}

func bindPropPrefix(propName TagProp, propValue TagProp) string {
	return bindPrefix(reflect.ValueOf(propTag) + propName + ":" + propValue)
}

// GetClusterResoureTag get cluster resource tag
func GetClusterResoureTag() string {
	return bindResourcePrefix("cluster")
}

// GetPropTag get property tag
func GetPropTag(propName TagProp, propValue string) string {
	return bindPropPrefix(propName, propValue)
}
