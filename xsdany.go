package gooxml

import (
	"encoding/xml"
	"log"
	"strings"
)

// XSDAny  is used to marshal/unmarshal xsd:any types in the OOXML schema.
type XSDAny struct {
	Attrs  []xml.Attr
	Tokens []xml.Token
}

func cloneToken(tok xml.Token) xml.Token {
	switch el := tok.(type) {
	case xml.CharData:
		cd := xml.CharData{}
		cd = append(cd, el...)
		return cd
	case xml.StartElement, xml.EndElement:
		return tok
	default:
		log.Fatalf("need to suppot %T", el)
	}
	return nil
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (x *XSDAny) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	x.Tokens = append(x.Tokens, cloneToken(start))
lfor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		default:
			x.Tokens = append(x.Tokens, cloneToken(tok))
		case xml.EndElement:
			x.Tokens = append(x.Tokens, cloneToken(tok))
			if el.Name == start.Name {
				break lfor
			}
		}
	}
	return nil
}

type nsSet struct {
	urlToPrefix map[string]string
	prefixToURL map[string]string
	prefixes    []string //required for deterministic output
}

func (n *nsSet) getPrefix(ns string) string {
	// do we have a prefix for this ns?
	if sc, ok := n.urlToPrefix[ns]; ok {
		return sc
	}

	// determine the last path portion of the namespace
	// "urn:schemas-microsoft-com:office:office" = "office"
	// "http://schemas.microsoft.com/office/word/2012/wordml" = "wordml"
	split := strings.Split(ns, "/")
	split = strings.Split(split[len(split)-1], ":")
	// last segment of the namesapce
	last := split[len(split)-1]
	lng := 0
	pfx := []byte{}
	for {
		if lng < len(last) {
			pfx = append(pfx, last[lng])
		} else {
			pfx = append(pfx, '_')
		}
		lng++
		// is this prefix unused?
		if _, ok := n.prefixToURL[string(pfx)]; !ok {
			n.prefixToURL[string(pfx)] = ns
			n.urlToPrefix[ns] = string(pfx)
			n.prefixes = append(n.prefixes, string(pfx))
			return string(pfx)
		}
	}
}

func (n nsSet) applyToSE(se *xml.StartElement) {
	if se.Name.Space == "" {
		return
	}

	pfx := n.getPrefix(se.Name.Space)
	se.Name.Space = ""
	se.Name.Local = pfx + ":" + se.Name.Local
	tmpAttr := se.Attr
	se.Attr = nil
	for _, attr := range tmpAttr {
		// skip these as we create them later
		if attr.Name.Space == "xmlns" {
			continue
		}
		if attr.Name.Space != "" {
			pfx := n.getPrefix(attr.Name.Space)
			attr.Name.Space = ""
			attr.Name.Local = pfx + ":" + attr.Name.Local
		}
		se.Attr = append(se.Attr, attr)
	}
}

func (n nsSet) applyToEE(ee *xml.EndElement) {
	if ee.Name.Space == "" {
		return
	}
	pfx := n.getPrefix(ee.Name.Space)
	ee.Name.Space = ""
	ee.Name.Local = pfx + ":" + ee.Name.Local
}

// MarshalXML implements the xml.Marshaler interface.
func (x *XSDAny) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(x.Tokens) == 0 {
		return nil
	}
	ns := nsSet{urlToPrefix: make(map[string]string),
		prefixToURL: make(map[string]string)}

	// collect the namespaces
	for _, tok := range x.Tokens {
		if se, ok := tok.(xml.StartElement); ok {
			if se.Name.Space != "" {
				ns.getPrefix(se.Name.Space)
			}
			for _, attr := range se.Attr {
				if attr.Name.Space != "" && attr.Name.Space != "xmlns" {
					ns.getPrefix(attr.Name.Space)
				}
			}
		}
	}
	// iniital element must be a StartElement
	se := x.Tokens[0].(xml.StartElement)
	ns.applyToSE(&se)
	// add namespaces to first element
	for _, pfx := range ns.prefixes {
		ns := ns.prefixToURL[pfx]
		se.Attr = append(se.Attr, xml.Attr{
			Name:  xml.Name{Local: "xmlns:" + pfx},
			Value: ns,
		})
	}
	e.EncodeToken(se)

	for _, tok := range x.Tokens[1:] {
		if se, ok := tok.(xml.StartElement); ok {
			ns.applyToSE(&se)
			e.EncodeToken(se)
		} else if ee, ok := tok.(xml.EndElement); ok {
			ns.applyToEE(&ee)
			e.EncodeToken(ee)
		} else if err := e.EncodeToken(tok); err != nil {
			return err
		}
	}
	return nil
}
