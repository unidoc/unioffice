// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_TimeNodeList struct {
	// Parallel Time Node
	Par []*CT_TLTimeNodeParallel
	// Sequence Time Node
	Seq []*CT_TLTimeNodeSequence
	// Exclusive
	Excl []*CT_TLTimeNodeExclusive
	// Animate
	Anim []*CT_TLAnimateBehavior
	// Animate Color Behavior
	AnimClr []*CT_TLAnimateColorBehavior
	// Animate Effect
	AnimEffect []*CT_TLAnimateEffectBehavior
	// Animate Motion
	AnimMotion []*CT_TLAnimateMotionBehavior
	// Animate Rotation
	AnimRot []*CT_TLAnimateRotationBehavior
	// Animate Scale
	AnimScale []*CT_TLAnimateScaleBehavior
	// Command
	Cmd []*CT_TLCommandBehavior
	// Set Time Node Behavior
	Set []*CT_TLSetBehavior
	// Audio
	Audio []*CT_TLMediaNodeAudio
	// Video
	Video []*CT_TLMediaNodeVideo
}

func NewCT_TimeNodeList() *CT_TimeNodeList {
	ret := &CT_TimeNodeList{}
	return ret
}

func (m *CT_TimeNodeList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Par != nil {
		separ := xml.StartElement{Name: xml.Name{Local: "p:par"}}
		e.EncodeElement(m.Par, separ)
	}
	if m.Seq != nil {
		seseq := xml.StartElement{Name: xml.Name{Local: "p:seq"}}
		e.EncodeElement(m.Seq, seseq)
	}
	if m.Excl != nil {
		seexcl := xml.StartElement{Name: xml.Name{Local: "p:excl"}}
		e.EncodeElement(m.Excl, seexcl)
	}
	if m.Anim != nil {
		seanim := xml.StartElement{Name: xml.Name{Local: "p:anim"}}
		e.EncodeElement(m.Anim, seanim)
	}
	if m.AnimClr != nil {
		seanimClr := xml.StartElement{Name: xml.Name{Local: "p:animClr"}}
		e.EncodeElement(m.AnimClr, seanimClr)
	}
	if m.AnimEffect != nil {
		seanimEffect := xml.StartElement{Name: xml.Name{Local: "p:animEffect"}}
		e.EncodeElement(m.AnimEffect, seanimEffect)
	}
	if m.AnimMotion != nil {
		seanimMotion := xml.StartElement{Name: xml.Name{Local: "p:animMotion"}}
		e.EncodeElement(m.AnimMotion, seanimMotion)
	}
	if m.AnimRot != nil {
		seanimRot := xml.StartElement{Name: xml.Name{Local: "p:animRot"}}
		e.EncodeElement(m.AnimRot, seanimRot)
	}
	if m.AnimScale != nil {
		seanimScale := xml.StartElement{Name: xml.Name{Local: "p:animScale"}}
		e.EncodeElement(m.AnimScale, seanimScale)
	}
	if m.Cmd != nil {
		secmd := xml.StartElement{Name: xml.Name{Local: "p:cmd"}}
		e.EncodeElement(m.Cmd, secmd)
	}
	if m.Set != nil {
		seset := xml.StartElement{Name: xml.Name{Local: "p:set"}}
		e.EncodeElement(m.Set, seset)
	}
	if m.Audio != nil {
		seaudio := xml.StartElement{Name: xml.Name{Local: "p:audio"}}
		e.EncodeElement(m.Audio, seaudio)
	}
	if m.Video != nil {
		sevideo := xml.StartElement{Name: xml.Name{Local: "p:video"}}
		e.EncodeElement(m.Video, sevideo)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TimeNodeList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TimeNodeList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "par":
				tmp := NewCT_TLTimeNodeParallel()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Par = append(m.Par, tmp)
			case "seq":
				tmp := NewCT_TLTimeNodeSequence()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Seq = append(m.Seq, tmp)
			case "excl":
				tmp := NewCT_TLTimeNodeExclusive()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Excl = append(m.Excl, tmp)
			case "anim":
				tmp := NewCT_TLAnimateBehavior()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Anim = append(m.Anim, tmp)
			case "animClr":
				tmp := NewCT_TLAnimateColorBehavior()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AnimClr = append(m.AnimClr, tmp)
			case "animEffect":
				tmp := NewCT_TLAnimateEffectBehavior()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AnimEffect = append(m.AnimEffect, tmp)
			case "animMotion":
				tmp := NewCT_TLAnimateMotionBehavior()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AnimMotion = append(m.AnimMotion, tmp)
			case "animRot":
				tmp := NewCT_TLAnimateRotationBehavior()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AnimRot = append(m.AnimRot, tmp)
			case "animScale":
				tmp := NewCT_TLAnimateScaleBehavior()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AnimScale = append(m.AnimScale, tmp)
			case "cmd":
				tmp := NewCT_TLCommandBehavior()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cmd = append(m.Cmd, tmp)
			case "set":
				tmp := NewCT_TLSetBehavior()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Set = append(m.Set, tmp)
			case "audio":
				tmp := NewCT_TLMediaNodeAudio()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Audio = append(m.Audio, tmp)
			case "video":
				tmp := NewCT_TLMediaNodeVideo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Video = append(m.Video, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TimeNodeList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TimeNodeList and its children
func (m *CT_TimeNodeList) Validate() error {
	return m.ValidateWithPath("CT_TimeNodeList")
}

// ValidateWithPath validates the CT_TimeNodeList and its children, prefixing error messages with path
func (m *CT_TimeNodeList) ValidateWithPath(path string) error {
	for i, v := range m.Par {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Par[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Seq {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Seq[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Excl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Excl[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Anim {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Anim[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AnimClr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AnimClr[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AnimEffect {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AnimEffect[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AnimMotion {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AnimMotion[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AnimRot {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AnimRot[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AnimScale {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AnimScale[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Cmd {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cmd[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Set {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Set[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Audio {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Audio[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Video {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Video[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
