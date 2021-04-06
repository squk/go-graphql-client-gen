package types

type IPAddressType int8

const (
	IPAddressType_IPv4 IPAddressType = iota
	IPAddressType_IPv6
)

type DirectiveLocation int8

const (
	DirectiveLocation_QUERY DirectiveLocation = iota
	DirectiveLocation_MUTATION
	DirectiveLocation_SUBSCRIPTION
	DirectiveLocation_FIELD
	DirectiveLocation_FRAGMENTDEFINITION
	DirectiveLocation_FRAGMENTSPREAD
	DirectiveLocation_INLINEFRAGMENT
	DirectiveLocation_SCHEMA
	DirectiveLocation_SCALAR
	DirectiveLocation_OBJECT
	DirectiveLocation_FIELDDEFINITION
	DirectiveLocation_ARGUMENTDEFINITION
	DirectiveLocation_INTERFACE
	DirectiveLocation_UNION
	DirectiveLocation_ENUM
	DirectiveLocation_ENUMVALUE
	DirectiveLocation_INPUTOBJECT
	DirectiveLocation_INPUTFIELDDEFINITION
)

type TypeKind int8

const (
	TypeKind_SCALAR TypeKind = iota
	TypeKind_OBJECT
	TypeKind_INTERFACE
	TypeKind_UNION
	TypeKind_ENUM
	TypeKind_INPUTOBJECT
	TypeKind_LIST
	TypeKind_NONNULL
)
