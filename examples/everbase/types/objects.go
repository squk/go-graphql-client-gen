package types

type City struct {
	Continent   Continent    `graphql:"continent" json:"continent,omitempty"`     // continent Continent!
	Country     Country      `graphql:"country" json:"country,omitempty"`         // country Country!
	GeonamesID  int          `graphql:"geonamesID" json:"geonamesID,omitempty"`   // geonamesID Int!
	Id          string       `graphql:"id" json:"id,omitempty"`                   // id String!
	Location    *Coordinates `graphql:"location" json:"location,omitempty"`       // location Coordinates
	Name        string       `graphql:"name" json:"name,omitempty"`               // name String!
	Population  int          `graphql:"population" json:"population,omitempty"`   // population Int!
	TimeZone    *TimeZone    `graphql:"timeZone" json:"timeZone,omitempty"`       // timeZone TimeZone
	TimeZoneDST *TimeZone    `graphql:"timeZoneDST" json:"timeZoneDST,omitempty"` // timeZoneDST TimeZone
}
type Client struct {
	IpAddress IPAddress `graphql:"ipAddress" json:"ipAddress,omitempty"` // ipAddress IPAddress!
	UserAgent string    `graphql:"userAgent" json:"userAgent,omitempty"` // userAgent String!
}
type Continent struct {
	Countries  []Country `graphql:"countries" json:"countries,omitempty"`   // countries [Country!]!
	GeonamesID int       `graphql:"geonamesID" json:"geonamesID,omitempty"` // geonamesID Int!
	Id         string    `graphql:"id" json:"id,omitempty"`                 // id String!
	Name       string    `graphql:"name" json:"name,omitempty"`             // name String!
	Population int       `graphql:"population" json:"population,omitempty"` // population Int!
}
type Coordinates struct {
	Lat  float32 `graphql:"lat" json:"lat,omitempty"`   // lat Float!
	Long float32 `graphql:"long" json:"long,omitempty"` // long Float!
}
type Country struct {
	Alpha2Code   string       `graphql:"alpha2Code" json:"alpha2Code,omitempty"`     // alpha2Code String!
	Alpha3Code   string       `graphql:"alpha3Code" json:"alpha3Code,omitempty"`     // alpha3Code String!
	CallingCodes []string     `graphql:"callingCodes" json:"callingCodes,omitempty"` // callingCodes [String!]!
	Capital      *City        `graphql:"capital" json:"capital,omitempty"`           // capital City
	Cities       []City       `graphql:"cities" json:"cities,omitempty"`             // cities [City!]!
	Continent    Continent    `graphql:"continent" json:"continent,omitempty"`       // continent Continent!
	Currencies   []Currency   `graphql:"currencies" json:"currencies,omitempty"`     // currencies [Currency!]!
	GeonamesID   int          `graphql:"geonamesID" json:"geonamesID,omitempty"`     // geonamesID Int!
	Id           string       `graphql:"id" json:"id,omitempty"`                     // id String!
	Languages    []Language   `graphql:"languages" json:"languages,omitempty"`       // languages [Language!]!
	Location     *Coordinates `graphql:"location" json:"location,omitempty"`         // location Coordinates
	Name         string       `graphql:"name" json:"name,omitempty"`                 // name String!
	Population   int          `graphql:"population" json:"population,omitempty"`     // population Int!
	VatRate      *float32     `graphql:"vatRate" json:"vatRate,omitempty"`           // vatRate Float
}
type Currency struct {
	Convert     *float32  `graphql:"convert" json:"convert,omitempty"`         // convert Float
	Countries   []Country `graphql:"countries" json:"countries,omitempty"`     // countries [Country!]!
	Id          string    `graphql:"id" json:"id,omitempty"`                   // id String!
	IsoCode     string    `graphql:"isoCode" json:"isoCode,omitempty"`         // isoCode String!
	Name        string    `graphql:"name" json:"name,omitempty"`               // name String!
	UnitSymbols []string  `graphql:"unitSymbols" json:"unitSymbols,omitempty"` // unitSymbols [String!]!
}
type DNSRecords struct {
	A     []IPAddress  `graphql:"a" json:"a,omitempty"`         // a [IPAddress!]!
	Aaaa  []IPAddress  `graphql:"aaaa" json:"aaaa,omitempty"`   // aaaa [IPAddress!]!
	Cname []DomainName `graphql:"cname" json:"cname,omitempty"` // cname [DomainName!]!
	Mx    []MXRecord   `graphql:"mx" json:"mx,omitempty"`       // mx [MXRecord!]!
}
type DomainName struct {
	A       []IPAddress  `graphql:"a" json:"a,omitempty"`             // a [IPAddress!]!
	Aaaa    []IPAddress  `graphql:"aaaa" json:"aaaa,omitempty"`       // aaaa [IPAddress!]!
	Cname   []DomainName `graphql:"cname" json:"cname,omitempty"`     // cname [DomainName!]!
	Mx      []MXRecord   `graphql:"mx" json:"mx,omitempty"`           // mx [MXRecord!]!
	Name    string       `graphql:"name" json:"name,omitempty"`       // name String!
	Records DNSRecords   `graphql:"records" json:"records,omitempty"` // records DNSRecords!
}
type EmailAddress struct {
	Address         string                `graphql:"address" json:"address,omitempty"`                 // address String!
	DomainName      DomainName            `graphql:"domainName" json:"domainName,omitempty"`           // domainName DomainName!
	Host            string                `graphql:"host" json:"host,omitempty"`                       // host String!
	Local           string                `graphql:"local" json:"local,omitempty"`                     // local String!
	Ok              bool                  `graphql:"ok" json:"ok,omitempty"`                           // ok Boolean!
	ServiceProvider *EmailServiceProvider `graphql:"serviceProvider" json:"serviceProvider,omitempty"` // serviceProvider EmailServiceProvider
}
type EmailServiceProvider struct {
	Disposable bool       `graphql:"disposable" json:"disposable,omitempty"` // disposable Boolean!
	DomainName DomainName `graphql:"domainName" json:"domainName,omitempty"` // domainName DomainName!
	Free       bool       `graphql:"free" json:"free,omitempty"`             // free Boolean!
	SmtpOk     bool       `graphql:"smtpOk" json:"smtpOk,omitempty"`         // smtpOk Boolean!
}
type HTMLDocument struct {
	All   []HTMLNode `graphql:"all" json:"all,omitempty"`     // all [HTMLNode!]!
	Body  HTMLNode   `graphql:"body" json:"body,omitempty"`   // body HTMLNode!
	First *HTMLNode  `graphql:"first" json:"first,omitempty"` // first HTMLNode
	Html  string     `graphql:"html" json:"html,omitempty"`   // html String!
	Title *string    `graphql:"title" json:"title,omitempty"` // title String
}
type HTMLNode struct {
	All       []HTMLNode `graphql:"all" json:"all,omitempty"`             // all [HTMLNode!]!
	Attribute *string    `graphql:"attribute" json:"attribute,omitempty"` // attribute String
	Children  []HTMLNode `graphql:"children" json:"children,omitempty"`   // children [HTMLNode!]!
	First     *HTMLNode  `graphql:"first" json:"first,omitempty"`         // first HTMLNode
	Html      string     `graphql:"html" json:"html,omitempty"`           // html String!
	Next      *HTMLNode  `graphql:"next" json:"next,omitempty"`           // next HTMLNode
	Parent    *HTMLNode  `graphql:"parent" json:"parent,omitempty"`       // parent HTMLNode
	Previous  *HTMLNode  `graphql:"previous" json:"previous,omitempty"`   // previous HTMLNode
	Text      *string    `graphql:"text" json:"text,omitempty"`           // text String
}
type IPAddress struct {
	Address string        `graphql:"address" json:"address,omitempty"` // address String!
	City    *City         `graphql:"city" json:"city,omitempty"`       // city City
	Country *Country      `graphql:"country" json:"country,omitempty"` // country Country
	Type    IPAddressType `graphql:"type" json:"type,omitempty"`       // type IPAddressType!
}
type Language struct {
	Alpha2Code string    `graphql:"alpha2Code" json:"alpha2Code,omitempty"` // alpha2Code String!
	Countries  []Country `graphql:"countries" json:"countries,omitempty"`   // countries [Country!]!
	Id         string    `graphql:"id" json:"id,omitempty"`                 // id String!
	Name       string    `graphql:"name" json:"name,omitempty"`             // name String!
}
type MXRecord struct {
	Exchange   DomainName `graphql:"exchange" json:"exchange,omitempty"`     // exchange DomainName!
	Preference int        `graphql:"preference" json:"preference,omitempty"` // preference Int!
}
type Markdown struct {
	Html string `graphql:"html" json:"html,omitempty"` // html String!
}
type Query struct {
	Cities       []City       `graphql:"cities" json:"cities,omitempty"`             // cities [City!]!
	Client       Client       `graphql:"client" json:"client,omitempty"`             // client Client!
	Continents   []Continent  `graphql:"continents" json:"continents,omitempty"`     // continents [Continent!]!
	Countries    []Country    `graphql:"countries" json:"countries,omitempty"`       // countries [Country!]!
	Currencies   []Currency   `graphql:"currencies" json:"currencies,omitempty"`     // currencies [Currency!]!
	DomainName   DomainName   `graphql:"domainName" json:"domainName,omitempty"`     // domainName DomainName!
	EmailAddress EmailAddress `graphql:"emailAddress" json:"emailAddress,omitempty"` // emailAddress EmailAddress!
	HtmlDocument HTMLDocument `graphql:"htmlDocument" json:"htmlDocument,omitempty"` // htmlDocument HTMLDocument!
	IpAddress    IPAddress    `graphql:"ipAddress" json:"ipAddress,omitempty"`       // ipAddress IPAddress!
	Languages    []Language   `graphql:"languages" json:"languages,omitempty"`       // languages [Language!]!
	Markdown     Markdown     `graphql:"markdown" json:"markdown,omitempty"`         // markdown Markdown!
	Random       Random       `graphql:"random" json:"random,omitempty"`             // random Random!
	TimeZones    []TimeZone   `graphql:"timeZones" json:"timeZones,omitempty"`       // timeZones [TimeZone!]!
	Url          URL          `graphql:"url" json:"url,omitempty"`                   // url URL!
	Schema       Schema       `graphql:"__schema" json:"__schema,omitempty"`         // __schema __Schema!
	Type         *Type        `graphql:"__type" json:"__type,omitempty"`             // __type __Type
}
type Random struct {
	Float  int    `graphql:"float" json:"float,omitempty"`   // float Int!
	Int    int    `graphql:"int" json:"int,omitempty"`       // int Int!
	String string `graphql:"string" json:"string,omitempty"` // string String!
}
type TimeZone struct {
	Cities []City  `graphql:"cities" json:"cities,omitempty"` // cities [City!]!
	Id     string  `graphql:"id" json:"id,omitempty"`         // id String!
	Name   string  `graphql:"name" json:"name,omitempty"`     // name String!
	Offset float32 `graphql:"offset" json:"offset,omitempty"` // offset Float!
}
type URL struct {
	DomainName   *DomainName   `graphql:"domainName" json:"domainName,omitempty"`     // domainName DomainName
	Host         string        `graphql:"host" json:"host,omitempty"`                 // host String!
	HtmlDocument *HTMLDocument `graphql:"htmlDocument" json:"htmlDocument,omitempty"` // htmlDocument HTMLDocument
	Path         *string       `graphql:"path" json:"path,omitempty"`                 // path String
	Port         *int          `graphql:"port" json:"port,omitempty"`                 // port Int
	Query        *string       `graphql:"query" json:"query,omitempty"`               // query String
	Scheme       string        `graphql:"scheme" json:"scheme,omitempty"`             // scheme String!
	Url          string        `graphql:"url" json:"url,omitempty"`                   // url String!
}
type Directive struct {
	Name        string              `graphql:"name" json:"name,omitempty"`               // name String!
	Description *string             `graphql:"description" json:"description,omitempty"` // description String
	Locations   []DirectiveLocation `graphql:"locations" json:"locations,omitempty"`     // locations [__DirectiveLocation!]!
	Args        []InputValue        `graphql:"args" json:"args,omitempty"`               // args [__InputValue!]!
}
type EnumValue struct {
	Name              string  `graphql:"name" json:"name,omitempty"`                           // name String!
	Description       *string `graphql:"description" json:"description,omitempty"`             // description String
	IsDeprecated      bool    `graphql:"isDeprecated" json:"isDeprecated,omitempty"`           // isDeprecated Boolean!
	DeprecationReason *string `graphql:"deprecationReason" json:"deprecationReason,omitempty"` // deprecationReason String
}
type Field struct {
	Name              string       `graphql:"name" json:"name,omitempty"`                           // name String!
	Description       *string      `graphql:"description" json:"description,omitempty"`             // description String
	Args              []InputValue `graphql:"args" json:"args,omitempty"`                           // args [__InputValue!]!
	Type              Type         `graphql:"type" json:"type,omitempty"`                           // type __Type!
	IsDeprecated      bool         `graphql:"isDeprecated" json:"isDeprecated,omitempty"`           // isDeprecated Boolean!
	DeprecationReason *string      `graphql:"deprecationReason" json:"deprecationReason,omitempty"` // deprecationReason String
}
type InputValue struct {
	Name         string  `graphql:"name" json:"name,omitempty"`                 // name String!
	Description  *string `graphql:"description" json:"description,omitempty"`   // description String
	Type         Type    `graphql:"type" json:"type,omitempty"`                 // type __Type!
	DefaultValue *string `graphql:"defaultValue" json:"defaultValue,omitempty"` // defaultValue String
}
type Schema struct {
	Types            []Type      `graphql:"types" json:"types,omitempty"`                       // types [__Type!]!
	QueryType        Type        `graphql:"queryType" json:"queryType,omitempty"`               // queryType __Type!
	MutationType     *Type       `graphql:"mutationType" json:"mutationType,omitempty"`         // mutationType __Type
	SubscriptionType *Type       `graphql:"subscriptionType" json:"subscriptionType,omitempty"` // subscriptionType __Type
	Directives       []Directive `graphql:"directives" json:"directives,omitempty"`             // directives [__Directive!]!
}
type Type struct {
	Kind          TypeKind      `graphql:"kind" json:"kind,omitempty"`                   // kind __TypeKind!
	Name          *string       `graphql:"name" json:"name,omitempty"`                   // name String
	Description   *string       `graphql:"description" json:"description,omitempty"`     // description String
	Fields        *[]Field      `graphql:"fields" json:"fields,omitempty"`               // fields [__Field!]
	Interfaces    *[]Type       `graphql:"interfaces" json:"interfaces,omitempty"`       // interfaces [__Type!]
	PossibleTypes *[]Type       `graphql:"possibleTypes" json:"possibleTypes,omitempty"` // possibleTypes [__Type!]
	EnumValues    *[]EnumValue  `graphql:"enumValues" json:"enumValues,omitempty"`       // enumValues [__EnumValue!]
	InputFields   *[]InputValue `graphql:"inputFields" json:"inputFields,omitempty"`     // inputFields [__InputValue!]
	OfType        *Type         `graphql:"ofType" json:"ofType,omitempty"`               // ofType __Type
}
