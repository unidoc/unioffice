// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
	"strconv"
)

type PivotCacheDefinition struct {
	CT_PivotCacheDefinition
}

func NewPivotCacheDefinition() *PivotCacheDefinition {
	ret := &PivotCacheDefinition{}
	ret.CT_PivotCacheDefinition = *NewCT_PivotCacheDefinition()
	return ret
}
func (m *PivotCacheDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:pivotCacheDefinition"
	return m.CT_PivotCacheDefinition.MarshalXML(e, start)
}
func (m *PivotCacheDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_PivotCacheDefinition = *NewCT_PivotCacheDefinition()
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
		if attr.Name.Local == "invalid" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.InvalidAttr = &parsed
		}
		if attr.Name.Local == "saveData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SaveDataAttr = &parsed
		}
		if attr.Name.Local == "refreshOnLoad" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RefreshOnLoadAttr = &parsed
		}
		if attr.Name.Local == "optimizeMemory" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OptimizeMemoryAttr = &parsed
		}
		if attr.Name.Local == "enableRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EnableRefreshAttr = &parsed
		}
		if attr.Name.Local == "refreshedBy" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefreshedByAttr = &parsed
		}
		if attr.Name.Local == "refreshedDate" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.RefreshedDateAttr = &parsed
		}
		if attr.Name.Local == "refreshedDateIso" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.RefreshedDateIsoAttr = &parsed
		}
		if attr.Name.Local == "backgroundQuery" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BackgroundQueryAttr = &parsed
		}
		if attr.Name.Local == "missingItemsLimit" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.MissingItemsLimitAttr = &pt
		}
		if attr.Name.Local == "createdVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.CreatedVersionAttr = &pt
		}
		if attr.Name.Local == "refreshedVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.RefreshedVersionAttr = &pt
		}
		if attr.Name.Local == "minRefreshableVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			pt := uint8(parsed)
			m.MinRefreshableVersionAttr = &pt
		}
		if attr.Name.Local == "recordCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RecordCountAttr = &pt
		}
		if attr.Name.Local == "upgradeOnRefresh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UpgradeOnRefreshAttr = &parsed
		}
		if attr.Name.Local == "tupleCache" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TupleCacheAttr = &parsed
		}
		if attr.Name.Local == "supportSubquery" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SupportSubqueryAttr = &parsed
		}
		if attr.Name.Local == "supportAdvancedDrill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SupportAdvancedDrillAttr = &parsed
		}
	}
lPivotCacheDefinition:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cacheSource":
				if err := d.DecodeElement(m.CacheSource, &el); err != nil {
					return err
				}
			case "cacheFields":
				if err := d.DecodeElement(m.CacheFields, &el); err != nil {
					return err
				}
			case "cacheHierarchies":
				m.CacheHierarchies = NewCT_CacheHierarchies()
				if err := d.DecodeElement(m.CacheHierarchies, &el); err != nil {
					return err
				}
			case "kpis":
				m.Kpis = NewCT_PCDKPIs()
				if err := d.DecodeElement(m.Kpis, &el); err != nil {
					return err
				}
			case "tupleCache":
				m.TupleCache = NewCT_TupleCache()
				if err := d.DecodeElement(m.TupleCache, &el); err != nil {
					return err
				}
			case "calculatedItems":
				m.CalculatedItems = NewCT_CalculatedItems()
				if err := d.DecodeElement(m.CalculatedItems, &el); err != nil {
					return err
				}
			case "calculatedMembers":
				m.CalculatedMembers = NewCT_CalculatedMembers()
				if err := d.DecodeElement(m.CalculatedMembers, &el); err != nil {
					return err
				}
			case "dimensions":
				m.Dimensions = NewCT_Dimensions()
				if err := d.DecodeElement(m.Dimensions, &el); err != nil {
					return err
				}
			case "measureGroups":
				m.MeasureGroups = NewCT_MeasureGroups()
				if err := d.DecodeElement(m.MeasureGroups, &el); err != nil {
					return err
				}
			case "maps":
				m.Maps = NewCT_MeasureDimensionMaps()
				if err := d.DecodeElement(m.Maps, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lPivotCacheDefinition
		case xml.CharData:
		}
	}
	return nil
}
func (m *PivotCacheDefinition) Validate() error {
	return m.ValidateWithPath("PivotCacheDefinition")
}
func (m *PivotCacheDefinition) ValidateWithPath(path string) error {
	if err := m.CT_PivotCacheDefinition.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
