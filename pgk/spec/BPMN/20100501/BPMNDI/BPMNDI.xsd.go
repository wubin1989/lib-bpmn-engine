//	Auto-generated by the "go-xsd" package located at:
//		github.com/UNO-SOFT/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.omg.org/spec/BPMN/20100501/BPMNDI.xsd
package BPMNDI

import (
	dc "github.com/nitram509/golib-bpmn-model/pgk/spec/BPMN/20100501/DC"
	"github.com/nitram509/golib-bpmn-model/pgk/spec/BPMN/20100501/DI"
	xsdt "github.com/nitram509/golib-bpmn-model/pgk/xsdtypes"
)

type XsdGoPkgHasAttr_BpmnElement_XsdtQName_ struct {
	BpmnElement xsdt.QName `xml:"http://www.omg.org/spec/BPMN/20100524/DI bpmnElement,attr"`
}

type TBPMNPlane struct {
	DI.TPlane

	XsdGoPkgHasAttr_BpmnElement_XsdtQName_
}

//	If the WalkHandlers.TBPMNPlane function is not nil (ie. was set by outside code), calls it with this TBPMNPlane instance as the single argument. Then calls the Walk() method on 0/2 embed(s) and 0/0 field(s) belonging to this TBPMNPlane instance.
func (me *TBPMNPlane) Walk() (err error) {
	if fn := WalkHandlers.TBPMNPlane; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_BPMNPlane struct {
	BPMNPlane *TBPMNPlane `xml:"http://www.omg.org/spec/BPMN/20100524/DI BPMNPlane"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_BPMNPlane function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_BPMNPlane instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_BPMNPlane instance.
func (me *XsdGoPkgHasElem_BPMNPlane) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_BPMNPlane; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.BPMNPlane.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type TBPMNLabelStyle struct {
	DI.TStyle

	dc.XsdGoPkgHasElem_Font
}

//	If the WalkHandlers.TBPMNLabelStyle function is not nil (ie. was set by outside code), calls it with this TBPMNLabelStyle instance as the single argument. Then calls the Walk() method on 0/2 embed(s) and 0/0 field(s) belonging to this TBPMNLabelStyle instance.
func (me *TBPMNLabelStyle) Walk() (err error) {
	if fn := WalkHandlers.TBPMNLabelStyle; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_BPMNLabelStyle struct {
	BPMNLabelStyles []*TBPMNLabelStyle `xml:"http://www.omg.org/spec/BPMN/20100524/DI BPMNLabelStyle"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_BPMNLabelStyle function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_BPMNLabelStyle instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_BPMNLabelStyle instance.
func (me *XsdGoPkgHasElems_BPMNLabelStyle) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_BPMNLabelStyle; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.BPMNLabelStyles {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type TBPMNDiagram struct {
	XsdGoPkgHasElem_BPMNPlane

	XsdGoPkgHasElems_BPMNLabelStyle

	DI.TDiagram
}

//	If the WalkHandlers.TBPMNDiagram function is not nil (ie. was set by outside code), calls it with this TBPMNDiagram instance as the single argument. Then calls the Walk() method on 2/3 embed(s) and 0/0 field(s) belonging to this TBPMNDiagram instance.
func (me *TBPMNDiagram) Walk() (err error) {
	if fn := WalkHandlers.TBPMNDiagram; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.XsdGoPkgHasElem_BPMNPlane.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if err = me.XsdGoPkgHasElems_BPMNLabelStyle.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_BPMNDiagram struct {
	BPMNDiagram *TBPMNDiagram `xml:"http://www.omg.org/spec/BPMN/20100524/DI BPMNDiagram"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_BPMNDiagram function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_BPMNDiagram instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_BPMNDiagram instance.
func (me *XsdGoPkgHasElem_BPMNDiagram) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_BPMNDiagram; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.BPMNDiagram.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_BPMNDiagram struct {
	BPMNDiagrams []*TBPMNDiagram `xml:"http://www.omg.org/spec/BPMN/20100524/DI BPMNDiagram"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_BPMNDiagram function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_BPMNDiagram instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_BPMNDiagram instance.
func (me *XsdGoPkgHasElems_BPMNDiagram) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_BPMNDiagram; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.BPMNDiagrams {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_BPMNPlane struct {
	BPMNPlanes []*TBPMNPlane `xml:"http://www.omg.org/spec/BPMN/20100524/DI BPMNPlane"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_BPMNPlane function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_BPMNPlane instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_BPMNPlane instance.
func (me *XsdGoPkgHasElems_BPMNPlane) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_BPMNPlane; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.BPMNPlanes {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_BPMNLabelStyle struct {
	BPMNLabelStyle *TBPMNLabelStyle `xml:"http://www.omg.org/spec/BPMN/20100524/DI BPMNLabelStyle"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_BPMNLabelStyle function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_BPMNLabelStyle instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_BPMNLabelStyle instance.
func (me *XsdGoPkgHasElem_BPMNLabelStyle) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_BPMNLabelStyle; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.BPMNLabelStyle.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasAttr_IsExpanded_XsdtBoolean_ struct {
	IsExpanded xsdt.Boolean `xml:"http://www.omg.org/spec/BPMN/20100524/DI isExpanded,attr"`
}

type XsdGoPkgHasAttr_IsMarkerVisible_XsdtBoolean_ struct {
	IsMarkerVisible xsdt.Boolean `xml:"http://www.omg.org/spec/BPMN/20100524/DI isMarkerVisible,attr"`
}

type XsdGoPkgHasAttr_LabelStyle_XsdtQName_ struct {
	LabelStyle xsdt.QName `xml:"http://www.omg.org/spec/BPMN/20100524/DI labelStyle,attr"`
}

type TBPMNLabel struct {
	DI.TLabel

	XsdGoPkgHasAttr_LabelStyle_XsdtQName_
}

//	If the WalkHandlers.TBPMNLabel function is not nil (ie. was set by outside code), calls it with this TBPMNLabel instance as the single argument. Then calls the Walk() method on 0/2 embed(s) and 0/0 field(s) belonging to this TBPMNLabel instance.
func (me *TBPMNLabel) Walk() (err error) {
	if fn := WalkHandlers.TBPMNLabel; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_BPMNLabel struct {
	BPMNLabel *TBPMNLabel `xml:"http://www.omg.org/spec/BPMN/20100524/DI BPMNLabel"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_BPMNLabel function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_BPMNLabel instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_BPMNLabel instance.
func (me *XsdGoPkgHasElem_BPMNLabel) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_BPMNLabel; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.BPMNLabel.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasAttr_IsMessageVisible_XsdtBoolean_ struct {
	IsMessageVisible xsdt.Boolean `xml:"http://www.omg.org/spec/BPMN/20100524/DI isMessageVisible,attr"`
}

type TParticipantBandKind xsdt.String

//	Since TParticipantBandKind is just a simple String type, this merely sets the current value from the specified string.
func (me *TParticipantBandKind) Set(s string) { (*xsdt.String)(me).Set(s) }

//	This convenience method just performs a simple type conversion to TParticipantBandKind's alias type xsdt.String.
func (me TParticipantBandKind) ToXsdtString() xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TParticipantBandKind is "bottom_initiating".
func (me TParticipantBandKind) IsBottomInitiating() bool { return me.String() == "bottom_initiating" }

//	Returns true if the value of this enumerated TParticipantBandKind is "top_non_initiating".
func (me TParticipantBandKind) IsTopNonInitiating() bool { return me.String() == "top_non_initiating" }

//	Since TParticipantBandKind is just a simple String type, this merely returns the current string value.
func (me TParticipantBandKind) String() string { return xsdt.String(me).String() }

//	Returns true if the value of this enumerated TParticipantBandKind is "top_initiating".
func (me TParticipantBandKind) IsTopInitiating() bool { return me.String() == "top_initiating" }

//	Returns true if the value of this enumerated TParticipantBandKind is "middle_initiating".
func (me TParticipantBandKind) IsMiddleInitiating() bool { return me.String() == "middle_initiating" }

//	Returns true if the value of this enumerated TParticipantBandKind is "middle_non_initiating".
func (me TParticipantBandKind) IsMiddleNonInitiating() bool {
	return me.String() == "middle_non_initiating"
}

//	Returns true if the value of this enumerated TParticipantBandKind is "bottom_non_initiating".
func (me TParticipantBandKind) IsBottomNonInitiating() bool {
	return me.String() == "bottom_non_initiating"
}

type XsdGoPkgHasAttr_ParticipantBandKind_TParticipantBandKind_ struct {
	ParticipantBandKind TParticipantBandKind `xml:"http://www.omg.org/spec/BPMN/20100524/DI participantBandKind,attr"`
}

type XsdGoPkgHasAttr_ChoreographyActivityShape_XsdtQName_ struct {
	ChoreographyActivityShape xsdt.QName `xml:"http://www.omg.org/spec/BPMN/20100524/DI choreographyActivityShape,attr"`
}

type XsdGoPkgHasAttr_IsHorizontal_XsdtBoolean_ struct {
	IsHorizontal xsdt.Boolean `xml:"http://www.omg.org/spec/BPMN/20100524/DI isHorizontal,attr"`
}

type TBPMNShape struct {
	DI.TLabeledShape

	XsdGoPkgHasAttr_IsExpanded_XsdtBoolean_

	XsdGoPkgHasAttr_BpmnElement_XsdtQName_

	XsdGoPkgHasAttr_IsHorizontal_XsdtBoolean_

	XsdGoPkgHasAttr_IsMarkerVisible_XsdtBoolean_

	XsdGoPkgHasElem_BPMNLabel

	XsdGoPkgHasAttr_IsMessageVisible_XsdtBoolean_

	XsdGoPkgHasAttr_ParticipantBandKind_TParticipantBandKind_

	XsdGoPkgHasAttr_ChoreographyActivityShape_XsdtQName_
}

//	If the WalkHandlers.TBPMNShape function is not nil (ie. was set by outside code), calls it with this TBPMNShape instance as the single argument. Then calls the Walk() method on 1/9 embed(s) and 0/0 field(s) belonging to this TBPMNShape instance.
func (me *TBPMNShape) Walk() (err error) {
	if fn := WalkHandlers.TBPMNShape; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.XsdGoPkgHasElem_BPMNLabel.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_BPMNShape struct {
	BPMNShape *TBPMNShape `xml:"http://www.omg.org/spec/BPMN/20100524/DI BPMNShape"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_BPMNShape function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_BPMNShape instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_BPMNShape instance.
func (me *XsdGoPkgHasElem_BPMNShape) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_BPMNShape; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.BPMNShape.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_BPMNShape struct {
	BPMNShapes []*TBPMNShape `xml:"http://www.omg.org/spec/BPMN/20100524/DI BPMNShape"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_BPMNShape function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_BPMNShape instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_BPMNShape instance.
func (me *XsdGoPkgHasElems_BPMNShape) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_BPMNShape; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.BPMNShapes {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_BPMNLabel struct {
	BPMNLabels []*TBPMNLabel `xml:"http://www.omg.org/spec/BPMN/20100524/DI BPMNLabel"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_BPMNLabel function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_BPMNLabel instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_BPMNLabel instance.
func (me *XsdGoPkgHasElems_BPMNLabel) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_BPMNLabel; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.BPMNLabels {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasAttr_TargetElement_XsdtQName_ struct {
	TargetElement xsdt.QName `xml:"http://www.omg.org/spec/BPMN/20100524/DI targetElement,attr"`
}

type TMessageVisibleKind xsdt.String

//	Returns true if the value of this enumerated TMessageVisibleKind is "non_initiating".
func (me TMessageVisibleKind) IsNonInitiating() bool { return me.String() == "non_initiating" }

//	Since TMessageVisibleKind is just a simple String type, this merely sets the current value from the specified string.
func (me *TMessageVisibleKind) Set(s string) { (*xsdt.String)(me).Set(s) }

//	Since TMessageVisibleKind is just a simple String type, this merely returns the current string value.
func (me TMessageVisibleKind) String() string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TMessageVisibleKind's alias type xsdt.String.
func (me TMessageVisibleKind) ToXsdtString() xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TMessageVisibleKind is "initiating".
func (me TMessageVisibleKind) IsInitiating() bool { return me.String() == "initiating" }

type XsdGoPkgHasAttr_MessageVisibleKind_TMessageVisibleKind_ struct {
	MessageVisibleKind TMessageVisibleKind `xml:"http://www.omg.org/spec/BPMN/20100524/DI messageVisibleKind,attr"`
}

type XsdGoPkgHasAttr_SourceElement_XsdtQName_ struct {
	SourceElement xsdt.QName `xml:"http://www.omg.org/spec/BPMN/20100524/DI sourceElement,attr"`
}

type TBPMNEdge struct {
	XsdGoPkgHasAttr_MessageVisibleKind_TMessageVisibleKind_

	DI.TLabeledEdge

	XsdGoPkgHasElem_BPMNLabel

	XsdGoPkgHasAttr_BpmnElement_XsdtQName_

	XsdGoPkgHasAttr_SourceElement_XsdtQName_

	XsdGoPkgHasAttr_TargetElement_XsdtQName_
}

//	If the WalkHandlers.TBPMNEdge function is not nil (ie. was set by outside code), calls it with this TBPMNEdge instance as the single argument. Then calls the Walk() method on 1/6 embed(s) and 0/0 field(s) belonging to this TBPMNEdge instance.
func (me *TBPMNEdge) Walk() (err error) {
	if fn := WalkHandlers.TBPMNEdge; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.XsdGoPkgHasElem_BPMNLabel.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElem_BPMNEdge struct {
	BPMNEdge *TBPMNEdge `xml:"http://www.omg.org/spec/BPMN/20100524/DI BPMNEdge"`
}

//	If the WalkHandlers.XsdGoPkgHasElem_BPMNEdge function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElem_BPMNEdge instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 1/1 field(s) belonging to this XsdGoPkgHasElem_BPMNEdge instance.
func (me *XsdGoPkgHasElem_BPMNEdge) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElem_BPMNEdge; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.BPMNEdge.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasElems_BPMNEdge struct {
	BPMNEdges []*TBPMNEdge `xml:"http://www.omg.org/spec/BPMN/20100524/DI BPMNEdge"`
}

//	If the WalkHandlers.XsdGoPkgHasElems_BPMNEdge function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasElems_BPMNEdge instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasElems_BPMNEdge instance.
func (me *XsdGoPkgHasElems_BPMNEdge) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasElems_BPMNEdge; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		for _, x := range me.BPMNEdges {
			if err = x.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

type XsdGoPkgHasCdata struct {
	XsdGoPkgCDATA string `xml:",chardata"`
}

//	If the WalkHandlers.XsdGoPkgHasCdata function is not nil (ie. was set by outside code), calls it with this XsdGoPkgHasCdata instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XsdGoPkgHasCdata instance.
func (me *XsdGoPkgHasCdata) Walk() (err error) {
	if fn := WalkHandlers.XsdGoPkgHasCdata; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

var (
	//	Set this to false to break a Walk() immediately as soon as the first error is returned by a custom handler function.
	//	If true, Walk() proceeds and accumulates all errors in the WalkErrors slice.
	WalkContinueOnError = true
	//	Contains all errors accumulated during Walk()s. If you're using this, you need to reset this yourself as needed prior to a fresh Walk().
	WalkErrors []error
	//	Your custom error-handling function, if required.
	WalkOnError func(error)
	//	Provides 19 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
	//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
	WalkHandlers = &XsdGoPkgWalkHandlers{}
)

//	Provides 19 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
//	If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
type XsdGoPkgWalkHandlers struct {
	XsdGoPkgHasElems_BPMNDiagram    func(*XsdGoPkgHasElems_BPMNDiagram, bool) error
	XsdGoPkgHasElem_BPMNShape       func(*XsdGoPkgHasElem_BPMNShape, bool) error
	TBPMNDiagram                    func(*TBPMNDiagram, bool) error
	XsdGoPkgHasElem_BPMNPlane       func(*XsdGoPkgHasElem_BPMNPlane, bool) error
	XsdGoPkgHasElems_BPMNPlane      func(*XsdGoPkgHasElems_BPMNPlane, bool) error
	TBPMNLabel                      func(*TBPMNLabel, bool) error
	XsdGoPkgHasElem_BPMNLabel       func(*XsdGoPkgHasElem_BPMNLabel, bool) error
	TBPMNShape                      func(*TBPMNShape, bool) error
	XsdGoPkgHasElem_BPMNEdge        func(*XsdGoPkgHasElem_BPMNEdge, bool) error
	XsdGoPkgHasElems_BPMNEdge       func(*XsdGoPkgHasElems_BPMNEdge, bool) error
	TBPMNPlane                      func(*TBPMNPlane, bool) error
	XsdGoPkgHasCdata                func(*XsdGoPkgHasCdata, bool) error
	XsdGoPkgHasElem_BPMNLabelStyle  func(*XsdGoPkgHasElem_BPMNLabelStyle, bool) error
	XsdGoPkgHasElems_BPMNLabelStyle func(*XsdGoPkgHasElems_BPMNLabelStyle, bool) error
	XsdGoPkgHasElem_BPMNDiagram     func(*XsdGoPkgHasElem_BPMNDiagram, bool) error
	XsdGoPkgHasElems_BPMNShape      func(*XsdGoPkgHasElems_BPMNShape, bool) error
	XsdGoPkgHasElems_BPMNLabel      func(*XsdGoPkgHasElems_BPMNLabel, bool) error
	TBPMNEdge                       func(*TBPMNEdge, bool) error
	TBPMNLabelStyle                 func(*TBPMNLabelStyle, bool) error
}
