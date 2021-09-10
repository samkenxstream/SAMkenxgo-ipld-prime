package schemadmt

type SchemaMap struct {
	Keys   []string
	Values map[string]TypeDefn
}
type Schema struct {
	Types SchemaMap
}
type TypeDefn struct {
	TypeBool   *TypeBool
	TypeString *TypeString
	TypeBytes  *TypeBytes
	TypeInt    *TypeInt
	TypeFloat  *TypeFloat
	TypeMap    *TypeMap
	TypeList   *TypeList
	TypeLink   *TypeLink
	TypeUnion  *TypeUnion
	TypeStruct *TypeStruct
	TypeEnum   *TypeEnum
	TypeCopy   *TypeCopy
}
type TypeTerm struct {
	TypeName   *string
	InlineDefn *InlineDefn
}
type InlineDefn struct {
	TypeMap  *TypeMap
	TypeList *TypeList
}
type TypeBool struct {
}
type TypeString struct {
}
type TypeBytes struct {
}
type TypeInt struct {
}
type TypeFloat struct {
}
type TypeMap struct {
	KeyType        string
	ValueType      TypeTerm
	ValueNullable  bool
	Representation MapRepresentation
}
type MapRepresentation struct {
	MapRepresentation_Map         *MapRepresentation_Map
	MapRepresentation_Stringpairs *MapRepresentation_Stringpairs
	MapRepresentation_Listpairs   *MapRepresentation_Listpairs
}
type MapRepresentation_Map struct {
}
type MapRepresentation_Stringpairs struct {
	InnerDelim string
	EntryDelim string
}
type MapRepresentation_Listpairs struct {
}
type TypeList struct {
	ValueType      TypeTerm
	ValueNullable  bool
	Representation ListRepresentation
}
type ListRepresentation struct {
	ListRepresentation_List *ListRepresentation_List
}
type ListRepresentation_List struct {
}
type TypeUnion struct {
	Members        List__TypeName
	Representation UnionRepresentation
}
type List__TypeName []string
type TypeLink struct {
	ExpectedType *string
}
type UnionRepresentation struct {
	UnionRepresentation_Kinded       *UnionRepresentation_Kinded
	UnionRepresentation_Keyed        *UnionRepresentation_Keyed
	UnionRepresentation_Envelope     *UnionRepresentation_Envelope
	UnionRepresentation_Inline       *UnionRepresentation_Inline
	UnionRepresentation_StringPrefix *UnionRepresentation_StringPrefix
	UnionRepresentation_BytePrefix   *UnionRepresentation_BytePrefix
}
type UnionRepresentation_Kinded struct {
	Keys   []string
	Values map[string]string
}
type UnionRepresentation_Keyed struct {
	Keys   []string
	Values map[string]string
}
type UnionRepresentation_Envelope struct {
	DiscriminantKey   string
	ContentKey        string
	DiscriminantTable Map__String__TypeName
}
type UnionRepresentation_Inline struct {
	DiscriminantKey   string
	DiscriminantTable Map__String__TypeName
}
type UnionRepresentation_StringPrefix struct {
	DiscriminantTable Map__String__TypeName
}
type UnionRepresentation_BytePrefix struct {
	DiscriminantTable Map__TypeName__Int
}
type Map__String__TypeName struct {
	Keys   []string
	Values map[string]string
}
type Map__TypeName__Int struct {
	Keys   []string
	Values map[string]int
}
type TypeStruct struct {
	Fields         Map__FieldName__StructField
	Representation StructRepresentation
}
type Map__FieldName__StructField struct {
	Keys   []string
	Values map[string]StructField
}
type StructField struct {
	Type     TypeTerm
	Optional bool
	Nullable bool
}
type StructRepresentation struct {
	StructRepresentation_Map         *StructRepresentation_Map
	StructRepresentation_Tuple       *StructRepresentation_Tuple
	StructRepresentation_Stringpairs *StructRepresentation_Stringpairs
	StructRepresentation_Stringjoin  *StructRepresentation_Stringjoin
	StructRepresentation_Listpairs   *StructRepresentation_Listpairs
}
type StructRepresentation_Map struct {
	Fields *Map__FieldName__StructRepresentation_Map_FieldDetails
}
type Map__FieldName__StructRepresentation_Map_FieldDetails struct {
	Keys   []string
	Values map[string]StructRepresentation_Map_FieldDetails
}
type StructRepresentation_Map_FieldDetails struct {
	Rename   *string
	Implicit *AnyScalar
}
type StructRepresentation_Tuple struct {
	FieldOrder *List__FieldName
}
type List__FieldName []string
type StructRepresentation_Stringpairs struct {
	InnerDelim string
	EntryDelim string
}
type StructRepresentation_Stringjoin struct {
	Join       string
	FieldOrder *List__FieldName
}
type StructRepresentation_Listpairs struct {
}
type TypeEnum struct {
	Members        Map__EnumValue__Unit
	Representation EnumRepresentation
}
type Map__EnumValue__Unit struct {
	Keys   []string
	Values map[string]Unit
}
type Unit struct {
}
type EnumRepresentation struct {
	EnumRepresentation_String *EnumRepresentation_String
	EnumRepresentation_Int    *EnumRepresentation_Int
}
type EnumRepresentation_String struct {
	Keys   []string
	Values map[string]string
}
type EnumRepresentation_Int struct {
	Keys   []string
	Values map[string]int
}
type TypeCopy struct {
	FromType string
}
type AnyScalar struct {
	Bool   *bool
	String *string
	Bytes  *[]uint8
	Int    *int
	Float  *float64
}