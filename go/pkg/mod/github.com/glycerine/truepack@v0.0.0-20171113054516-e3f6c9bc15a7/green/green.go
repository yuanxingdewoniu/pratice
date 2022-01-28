/*
package green specifies the Greenpack serialization format.
Instead of an IDL file, the Greenpack schema is described
using the same Go source file that holds the Go structs
you wish to serialize. The Go file schema is then compiled
by running `truepack` into msgpack2 (with optional JSON)
in a format we'll call `compiled-schema` format. If
one is not starting in Go, simply write a standalone Go
file that describes your types. See the examples
in `../testdata/my.go`.

The `compiled-schema` is thus type checked upon generation,
and other languages need not parse Go (only msgpack2
or JSON) in order the read and use the compiled
schema to undertand the types on the wire. The
types below desribe those found in
the compiled Greenpack schema files.

The methods that you see below in the godoc are
the autogenerated methods from running
`truepack -file green.go` on this file itself.
They provide an API for reading and writing
compiled Greenpack schema.
*/
package green

// in the generate command, we use -msgp so that
// we serialize the Greenpack schema itself
// as simple msgpack2, rather than in Greenpack format.

//go:generate truepack

const greenSchemaId64 = 0x1a5a94bd49624

// Zkind describes the detailed type of the field.
// Since it also stores the fixed size of a array type,
// it needs to be large. When serialized as msgpack2,
// it will be compressed.
//
// Implentation note: primitives must correspond
// to gen/Primitive, as we will cast/convert them directly.
//
type Zkind uint64

const (

	// primitives.
	// Implementation note: must correspond to gen/Primitive.
	Invalid    Zkind = 0
	Bytes      Zkind = 1 // []byte
	String     Zkind = 2
	Float32    Zkind = 3
	Float64    Zkind = 4
	Complex64  Zkind = 5
	Complex128 Zkind = 6
	Uint       Zkind = 7 // 32 or 64 bit; as in Go, matches native word
	Uint8      Zkind = 8
	Uint16     Zkind = 9
	Uint32     Zkind = 10
	Uint64     Zkind = 11
	Byte       Zkind = 12
	Int        Zkind = 13 // as in Go, matches native word size.
	Int8       Zkind = 14
	Int16      Zkind = 15
	Int32      Zkind = 16
	Int64      Zkind = 17
	Bool       Zkind = 18
	Intf       Zkind = 19 // interface{}
	Time       Zkind = 20 // time.Time
	Ext        Zkind = 21 // extension

	// IDENT means an unrecognized identifier;
	// it typically means a named struct type.
	// The Str field in the Ztype will hold the
	// name of the struct.
	IDENT Zkind = 22

	// compound types
	// implementation note: should correspond to gen/Elem.
	BaseElemCat Zkind = 23
	MapCat      Zkind = 24
	StructCat   Zkind = 25
	SliceCat    Zkind = 26
	ArrayCat    Zkind = 27
	PointerCat  Zkind = 28
)

// Ztype describes any type, be it a BaseElem,
// Map, Struct, Slice, Array, or Pointer.
type Ztype struct {

	// Kind gives the exact type for primitives,
	// and the category for compound types.
	Kind Zkind

	// Str holds the struct name when Kind == 22 (IDENT).
	// Otherwise it typically reflects Kind directly
	// which is useful for human readability.
	Str string `msg:",omitempty"`

	// Domain holds the key type for maps. For
	// pointers and slices it holds the element type.
	// For arrays, it holds the fixed size.
	// Domain is null when Kind is a primitive.
	Domain *Ztype `msg:",omitempty"`

	// Range holds the value type for maps.
	// For arrays (always a fixed size), Range holds
	// the element type.  Otherwise Range is
	// typically null.
	Range *Ztype `msg:",omitempty"`
}

// Schema is the top level container in Greenpack.
// It all starts here.
type Schema struct {

	// SourcePath gives the path to the original Go
	// source file that was parsed to produce this
	// compiled schema.
	SourcePath string

	// SourcePackage notes the original package presented
	// by the SourcePath file.
	SourcePackage string

	// GreenSchemaId is a randomly chosen but stable
	// 53-bit positive integer identifier (see
	// truepack -genid) that can be used to distinguish schemas.
	GreenSchemaId int64

	// Structs holds the collection of the main data
	// descriptor, the Struct. The key is identical
	// to Struct.StructName.
	//
	// This a map rather than a slice in order to:
	// a) insure there are no duplicate struct names; and
	// b) make decoding easy and fast.
	Structs map[string]*Struct

	// Imports archives the imports in the SourcePath
	// to make it possible to understand other package
	// type references.
	Imports []string
}

// Struct represents a single message or struct.
type Struct struct {
	// StructName is the typename of the struct in Go.
	StructName string

	// Fields hold the individual Field descriptors.
	Fields []Field
}

// Field represents fields within a struct.
type Field struct {

	// Zid is the truepack id.
	//
	// Zid numbering detects update collisions
	// when two developers simultaneously add two
	// new fields. Zid numbering allows sane
	// forward/backward data evolution, like protobufs
	// and Cap'nProto.
	//
	// Zid follows Cap'nProto numbering semantics:
	// start at numbering at 0, and strictly/always
	// increase numbers monotically.
	//
	// No gaps and no duplicate Zid are allowed.
	//
	// Duplicate numbers are how collisions (between two
	// developers adding two distinct new fields independently
	// but at the same time) are detected.
	//
	// Therefore this ironclad rule: never delete a field or Zid number,
	// just mark it as deprecated with the `deprecated:"true"`
	// tag. Change its Go type to struct{} as soon as
	// possible so that it becomes skipped; then the Go
	// compiler can help you detect and prevent unwanted use.
	//
	Zid int64

	// the name of the field in the Go schema/source file.
	FieldGoName string

	// optional wire-name designated by a
	// `msg:"tagname"` tag on the struct field.
	FieldTagName string `msg:",omitempty"`

	// =======================
	// type info
	// =======================

	// human readable/Go type description
	FieldTypeStr string `msg:",omitempty"`

	// the broad category of this type. empty if Skip is true
	FieldCategory Zkind `msg:",omitempty"`

	// avail if FieldCategory == BaseElemCat
	FieldPrimitive Zkind `msg:",omitempty"`

	// the machine-parse-able type tree
	FieldFullType *Ztype `msg:",omitempty"`

	// =======================
	// field tag details:
	// =======================

	// if OmitEmpty then we don't serialize
	// the field if it has its zero-value.
	OmitEmpty bool `msg:",omitempty"`

	// Skip means either type struct{} or
	// other unserializable type;
	// or marked  as `msg:"-"`. In any case,
	// if Skip is true: do not serialize
	// this field when writing, and ignore it
	// when reading.
	Skip bool `msg:",omitempty"`

	// Deprecated means tagged with `deprecated:"true"`,
	// or `msg:",deprecated"`.
	// Compilers/libraries should discourage and warn
	// users away from writing to such fields, while
	// not making it impossible to either read or write
	// the field.
	Deprecated bool `msg:",omitempty"`

	// ShowZero means display the field even if
	// it has the zero value. Showzero has no impact
	// on what is transmitted on the wire. Zero
	// valued fields are never transmitted.
	ShowZero bool `msg:",omitempty"`
}