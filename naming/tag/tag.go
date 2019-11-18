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
	kindTag     Tag = "kind:"
	resourceTag Tag = "res:"
	propTag     Tag = "prop:"
)

func bindPrefix(suffix Tag) string {
	return string(tagPrefix + suffix)
}

func bindResourcePrefix(resourceName Tag) string {
	return bindPrefix(resourceTag + resourceName)
}

func bindKindPrefix(resourceName Tag) string {
	return bindPrefix(kindTag + resourceName)
}

func bindPropPrefix(propName TagProp, propValue TagProp) string {
	return bindPrefix(propTag + Tag(propName) + ":" + Tag(propValue))
}

// GetServerKindTag get server kind tag
func GetServerKindTag() string {
	return bindKindPrefix("server")
}

// GetClientKindTag get client kind tag
func GetClientKindTag() string {
	return bindKindPrefix("client")
}

// GetConsulResourceTag get consul resource tag
func GetConsulResourceTag() string {
	return bindResourcePrefix("consul")
}

// GetNomadResourceTag get nomad resource tag
func GetNomadResourceTag() string {
	return bindResourcePrefix("nomad")
}

// GetClientResourceTag get client resource tag
func GetClientResourceTag() string {
	return bindResourcePrefix("client")
}

// GetPropTag get property tag
func GetPropTag(propName TagProp, propValue TagProp) string {
	return bindPropPrefix(propName, propValue)
}
