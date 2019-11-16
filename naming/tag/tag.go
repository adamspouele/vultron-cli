package tag

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
	tagPrefix   Tag = "vultron:"
	resourceTag Tag = "res:"
	propTag     Tag = "prop:"
)

func bindPrefix(suffix Tag) string {
	return string(tagPrefix + suffix)
}

func bindResourcePrefix(resourceName Tag) string {
	return bindPrefix(resourceTag + resourceName)
}

func bindPropPrefix(propName TagProp, propValue TagProp) string {
	return bindPrefix(propTag + Tag(propName) + ":" + Tag(propValue))
}

// GetClusterResoureTag get cluster resource tag
func GetClusterResoureTag() string {
	return bindResourcePrefix("cluster")
}

// GetPropTag get property tag
func GetPropTag(propName TagProp, propValue TagProp) string {
	return bindPropPrefix(propName, propValue)
}
