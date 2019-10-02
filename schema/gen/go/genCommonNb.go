package gengo

import (
	"io"

	ipld "github.com/ipld/go-ipld-prime"
)

type genKindedNbRejections struct {
	TypeIdent string // the identifier in code (sometimes is munged internals like "_Thing__Repr" corresponding to no publicly admitted schema.Type.Name).
	TypeProse string // as will be printed in messages (e.g. can be goosed up a bit, like "Thing.Repr" instead of "_Thing__Repr").
	Kind      ipld.ReprKind
}

func (d genKindedNbRejections) emitNodebuilderMethodCreateMap(w io.Writer) {
	doTemplate(`
		func ({{ .TypeIdent }}) CreateMap() (ipld.MapBuilder, error) {
			return nil, ipld.ErrWrongKind{TypeName: "{{ .TypeProse }}", MethodName: "CreateMap", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: {{ .Kind | ReprKindConst }}}
		}
	`, w, d)
}
func (d genKindedNbRejections) emitNodebuilderMethodAmendMap(w io.Writer) {
	doTemplate(`
		func ({{ .TypeIdent }}) AmendMap() (ipld.MapBuilder, error) {
			return nil, ipld.ErrWrongKind{TypeName: "{{ .TypeProse }}", MethodName: "AmendMap", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: {{ .Kind | ReprKindConst }}}
		}
	`, w, d)
}
func (d genKindedNbRejections) emitNodebuilderMethodCreateList(w io.Writer) {
	doTemplate(`
		func ({{ .TypeIdent }}) CreateList() (ipld.ListBuilder, error) {
			return nil, ipld.ErrWrongKind{TypeName: "{{ .TypeProse }}", MethodName: "CreateList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: {{ .Kind | ReprKindConst }}}
		}
	`, w, d)
}
func (d genKindedNbRejections) emitNodebuilderMethodAmendList(w io.Writer) {
	doTemplate(`
		func ({{ .TypeIdent }}) AmendList() (ipld.ListBuilder, error) {
			return nil, ipld.ErrWrongKind{TypeName: "{{ .TypeProse }}", MethodName: "AmendList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: {{ .Kind | ReprKindConst }}}
		}
	`, w, d)
}
func (d genKindedNbRejections) emitNodebuilderMethodCreateNull(w io.Writer) {
	doTemplate(`
		func ({{ .TypeIdent }}) CreateNull() (ipld.Node, error) {
			return nil, ipld.ErrWrongKind{TypeName: "{{ .TypeProse }}", MethodName: "CreateNull", AppropriateKind: ipld.ReprKindSet_JustNull, ActualKind: {{ .Kind | ReprKindConst }}}
		}
	`, w, d)
}
func (d genKindedNbRejections) emitNodebuilderMethodCreateBool(w io.Writer) {
	doTemplate(`
		func ({{ .TypeIdent }}) CreateBool(bool) (ipld.Node, error) {
			return nil, ipld.ErrWrongKind{TypeName: "{{ .TypeProse }}", MethodName: "CreateBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: {{ .Kind | ReprKindConst }}}
		}
	`, w, d)
}
func (d genKindedNbRejections) emitNodebuilderMethodCreateInt(w io.Writer) {
	doTemplate(`
		func ({{ .TypeIdent }}) CreateInt(int) (ipld.Node, error) {
			return nil, ipld.ErrWrongKind{TypeName: "{{ .TypeProse }}", MethodName: "CreateInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: {{ .Kind | ReprKindConst }}}
		}
	`, w, d)
}
func (d genKindedNbRejections) emitNodebuilderMethodCreateFloat(w io.Writer) {
	doTemplate(`
		func ({{ .TypeIdent }}) CreateFloat(float64) (ipld.Node, error) {
			return nil, ipld.ErrWrongKind{TypeName: "{{ .TypeProse }}", MethodName: "CreateFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: {{ .Kind | ReprKindConst }}}
		}
	`, w, d)
}
func (d genKindedNbRejections) emitNodebuilderMethodCreateString(w io.Writer) {
	doTemplate(`
		func ({{ .TypeIdent }}) CreateString(string) (ipld.Node, error) {
			return nil, ipld.ErrWrongKind{TypeName: "{{ .TypeProse }}", MethodName: "CreateString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: {{ .Kind | ReprKindConst }}}
		}
	`, w, d)
}
func (d genKindedNbRejections) emitNodebuilderMethodCreateBytes(w io.Writer) {
	doTemplate(`
		func ({{ .TypeIdent }}) CreateBytes([]byte) (ipld.Node, error) {
			return nil, ipld.ErrWrongKind{TypeName: "{{ .TypeProse }}", MethodName: "CreateBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: {{ .Kind | ReprKindConst }}}
		}
	`, w, d)
}
func (d genKindedNbRejections) emitNodebuilderMethodCreateLink(w io.Writer) {
	doTemplate(`
		func ({{ .TypeIdent }}) CreateLink(ipld.Link) (ipld.Node, error) {
			return nil, ipld.ErrWrongKind{TypeName: "{{ .TypeProse }}", MethodName: "CreateLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: {{ .Kind | ReprKindConst }}}
		}
	`, w, d)
}

// Embeddable to do all the "nope" methods at once.
type genKindedNbRejections_String struct {
	TypeIdent string // see doc in generateKindedRejections
	TypeProse string // see doc in generateKindedRejections
}

func (gk genKindedNbRejections_String) EmitNodebuilderMethodCreateMap(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_String}.emitNodebuilderMethodCreateMap(w)
}
func (gk genKindedNbRejections_String) EmitNodebuilderMethodAmendMap(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_String}.emitNodebuilderMethodAmendMap(w)
}
func (gk genKindedNbRejections_String) EmitNodebuilderMethodCreateList(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_String}.emitNodebuilderMethodCreateList(w)
}
func (gk genKindedNbRejections_String) EmitNodebuilderMethodAmendList(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_String}.emitNodebuilderMethodAmendList(w)
}
func (gk genKindedNbRejections_String) EmitNodebuilderMethodCreateNull(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_String}.emitNodebuilderMethodCreateNull(w)
}
func (gk genKindedNbRejections_String) EmitNodebuilderMethodCreateBool(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_String}.emitNodebuilderMethodCreateBool(w)
}
func (gk genKindedNbRejections_String) EmitNodebuilderMethodCreateInt(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_String}.emitNodebuilderMethodCreateInt(w)
}
func (gk genKindedNbRejections_String) EmitNodebuilderMethodCreateFloat(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_String}.emitNodebuilderMethodCreateFloat(w)
}
func (gk genKindedNbRejections_String) EmitNodebuilderMethodCreateBytes(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_String}.emitNodebuilderMethodCreateBytes(w)
}
func (gk genKindedNbRejections_String) EmitNodebuilderMethodCreateLink(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_String}.emitNodebuilderMethodCreateLink(w)
}

// Embeddable to do all the "nope" methods at once.
type genKindedNbRejections_Map struct {
	TypeIdent string // see doc in generateKindedRejections
	TypeProse string // see doc in generateKindedRejections
}

func (gk genKindedNbRejections_Map) EmitNodebuilderMethodCreateList(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_Map}.emitNodebuilderMethodCreateList(w)
}
func (gk genKindedNbRejections_Map) EmitNodebuilderMethodAmendList(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_Map}.emitNodebuilderMethodAmendList(w)
}
func (gk genKindedNbRejections_Map) EmitNodebuilderMethodCreateNull(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_Map}.emitNodebuilderMethodCreateNull(w)
}
func (gk genKindedNbRejections_Map) EmitNodebuilderMethodCreateBool(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_Map}.emitNodebuilderMethodCreateBool(w)
}
func (gk genKindedNbRejections_Map) EmitNodebuilderMethodCreateInt(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_Map}.emitNodebuilderMethodCreateInt(w)
}
func (gk genKindedNbRejections_Map) EmitNodebuilderMethodCreateFloat(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_Map}.emitNodebuilderMethodCreateFloat(w)
}
func (gk genKindedNbRejections_Map) EmitNodebuilderMethodCreateString(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_Map}.emitNodebuilderMethodCreateString(w)
}
func (gk genKindedNbRejections_Map) EmitNodebuilderMethodCreateBytes(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_Map}.emitNodebuilderMethodCreateBytes(w)
}
func (gk genKindedNbRejections_Map) EmitNodebuilderMethodCreateLink(w io.Writer) {
	genKindedNbRejections{gk.TypeIdent, gk.TypeProse, ipld.ReprKind_Map}.emitNodebuilderMethodCreateLink(w)
}