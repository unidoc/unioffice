// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"time"
)

type CT_PivotCacheDefinition struct {
	IdAttr *string
	// Invalid Cache
	InvalidAttr *bool
	// Save Pivot Records
	SaveDataAttr *bool
	// Refresh On Load
	RefreshOnLoadAttr *bool
	// Optimize Cache for Memory
	OptimizeMemoryAttr *bool
	// Enable PivotCache Refresh
	EnableRefreshAttr *bool
	// Last Refreshed By
	RefreshedByAttr *string
	// PivotCache Last Refreshed Date
	RefreshedDateAttr *float64
	// PivotCache Last Refreshed Date ISO
	RefreshedDateIsoAttr *time.Time
	// Background Query
	BackgroundQueryAttr *bool
	// Missing Items Limit
	MissingItemsLimitAttr *uint32
	// PivotCache Created Version
	CreatedVersionAttr *uint8
	// PivotCache Last Refreshed Version
	RefreshedVersionAttr *uint8
	// Minimum Version Required for Refresh
	MinRefreshableVersionAttr *uint8
	// PivotCache Record Count
	RecordCountAttr *uint32
	// Upgrade PivotCache on Refresh
	UpgradeOnRefreshAttr *bool
	// Tuple Cache
	TupleCacheAttr *bool
	// Supports Subqueries
	SupportSubqueryAttr *bool
	// Supports Attribute Drilldown
	SupportAdvancedDrillAttr *bool
	// PivotCache Source Description
	CacheSource *CT_CacheSource
	// PivotCache Fields
	CacheFields *CT_CacheFields
	// PivotCache Hierarchies
	CacheHierarchies *CT_CacheHierarchies
	// OLAP KPIs
	Kpis *CT_PCDKPIs
	// Tuple Cache
	TupleCache *CT_TupleCache
	// Calculated Items
	CalculatedItems *CT_CalculatedItems
	// Calculated Members
	CalculatedMembers *CT_CalculatedMembers
	// OLAP Dimensions
	Dimensions *CT_Dimensions
	// OLAP Measure Groups
	MeasureGroups *CT_MeasureGroups
	// OLAP Measure Group
	Maps *CT_MeasureDimensionMaps
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_PivotCacheDefinition() *CT_PivotCacheDefinition {
	ret := &CT_PivotCacheDefinition{}
	ret.CacheSource = NewCT_CacheSource()
	ret.CacheFields = NewCT_CacheFields()
	return ret
}
func (m *CT_PivotCacheDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.InvalidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "invalid"},
			Value: fmt.Sprintf("%v", *m.InvalidAttr)})
	}
	if m.SaveDataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "saveData"},
			Value: fmt.Sprintf("%v", *m.SaveDataAttr)})
	}
	if m.RefreshOnLoadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refreshOnLoad"},
			Value: fmt.Sprintf("%v", *m.RefreshOnLoadAttr)})
	}
	if m.OptimizeMemoryAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "optimizeMemory"},
			Value: fmt.Sprintf("%v", *m.OptimizeMemoryAttr)})
	}
	if m.EnableRefreshAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "enableRefresh"},
			Value: fmt.Sprintf("%v", *m.EnableRefreshAttr)})
	}
	if m.RefreshedByAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refreshedBy"},
			Value: fmt.Sprintf("%v", *m.RefreshedByAttr)})
	}
	if m.RefreshedDateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refreshedDate"},
			Value: fmt.Sprintf("%v", *m.RefreshedDateAttr)})
	}
	if m.RefreshedDateIsoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refreshedDateIso"},
			Value: fmt.Sprintf("%v", *m.RefreshedDateIsoAttr)})
	}
	if m.BackgroundQueryAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "backgroundQuery"},
			Value: fmt.Sprintf("%v", *m.BackgroundQueryAttr)})
	}
	if m.MissingItemsLimitAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "missingItemsLimit"},
			Value: fmt.Sprintf("%v", *m.MissingItemsLimitAttr)})
	}
	if m.CreatedVersionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "createdVersion"},
			Value: fmt.Sprintf("%v", *m.CreatedVersionAttr)})
	}
	if m.RefreshedVersionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refreshedVersion"},
			Value: fmt.Sprintf("%v", *m.RefreshedVersionAttr)})
	}
	if m.MinRefreshableVersionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minRefreshableVersion"},
			Value: fmt.Sprintf("%v", *m.MinRefreshableVersionAttr)})
	}
	if m.RecordCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "recordCount"},
			Value: fmt.Sprintf("%v", *m.RecordCountAttr)})
	}
	if m.UpgradeOnRefreshAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "upgradeOnRefresh"},
			Value: fmt.Sprintf("%v", *m.UpgradeOnRefreshAttr)})
	}
	if m.TupleCacheAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tupleCache"},
			Value: fmt.Sprintf("%v", *m.TupleCacheAttr)})
	}
	if m.SupportSubqueryAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "supportSubquery"},
			Value: fmt.Sprintf("%v", *m.SupportSubqueryAttr)})
	}
	if m.SupportAdvancedDrillAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "supportAdvancedDrill"},
			Value: fmt.Sprintf("%v", *m.SupportAdvancedDrillAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	secacheSource := xml.StartElement{Name: xml.Name{Local: "x:cacheSource"}}
	e.EncodeElement(m.CacheSource, secacheSource)
	secacheFields := xml.StartElement{Name: xml.Name{Local: "x:cacheFields"}}
	e.EncodeElement(m.CacheFields, secacheFields)
	if m.CacheHierarchies != nil {
		secacheHierarchies := xml.StartElement{Name: xml.Name{Local: "x:cacheHierarchies"}}
		e.EncodeElement(m.CacheHierarchies, secacheHierarchies)
	}
	if m.Kpis != nil {
		sekpis := xml.StartElement{Name: xml.Name{Local: "x:kpis"}}
		e.EncodeElement(m.Kpis, sekpis)
	}
	if m.TupleCache != nil {
		setupleCache := xml.StartElement{Name: xml.Name{Local: "x:tupleCache"}}
		e.EncodeElement(m.TupleCache, setupleCache)
	}
	if m.CalculatedItems != nil {
		secalculatedItems := xml.StartElement{Name: xml.Name{Local: "x:calculatedItems"}}
		e.EncodeElement(m.CalculatedItems, secalculatedItems)
	}
	if m.CalculatedMembers != nil {
		secalculatedMembers := xml.StartElement{Name: xml.Name{Local: "x:calculatedMembers"}}
		e.EncodeElement(m.CalculatedMembers, secalculatedMembers)
	}
	if m.Dimensions != nil {
		sedimensions := xml.StartElement{Name: xml.Name{Local: "x:dimensions"}}
		e.EncodeElement(m.Dimensions, sedimensions)
	}
	if m.MeasureGroups != nil {
		semeasureGroups := xml.StartElement{Name: xml.Name{Local: "x:measureGroups"}}
		e.EncodeElement(m.MeasureGroups, semeasureGroups)
	}
	if m.Maps != nil {
		semaps := xml.StartElement{Name: xml.Name{Local: "x:maps"}}
		e.EncodeElement(m.Maps, semaps)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PivotCacheDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CacheSource = NewCT_CacheSource()
	m.CacheFields = NewCT_CacheFields()
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
			// SPECIAL
			pt := uint32(parsed)
			m.MissingItemsLimitAttr = &pt
		}
		if attr.Name.Local == "createdVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint8(parsed)
			m.CreatedVersionAttr = &pt
		}
		if attr.Name.Local == "refreshedVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint8(parsed)
			m.RefreshedVersionAttr = &pt
		}
		if attr.Name.Local == "minRefreshableVersion" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 8)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint8(parsed)
			m.MinRefreshableVersionAttr = &pt
		}
		if attr.Name.Local == "recordCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
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
lCT_PivotCacheDefinition:
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
			break lCT_PivotCacheDefinition
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PivotCacheDefinition) Validate() error {
	return m.ValidateWithPath("CT_PivotCacheDefinition")
}
func (m *CT_PivotCacheDefinition) ValidateWithPath(path string) error {
	if err := m.CacheSource.ValidateWithPath(path + "/CacheSource"); err != nil {
		return err
	}
	if err := m.CacheFields.ValidateWithPath(path + "/CacheFields"); err != nil {
		return err
	}
	if m.CacheHierarchies != nil {
		if err := m.CacheHierarchies.ValidateWithPath(path + "/CacheHierarchies"); err != nil {
			return err
		}
	}
	if m.Kpis != nil {
		if err := m.Kpis.ValidateWithPath(path + "/Kpis"); err != nil {
			return err
		}
	}
	if m.TupleCache != nil {
		if err := m.TupleCache.ValidateWithPath(path + "/TupleCache"); err != nil {
			return err
		}
	}
	if m.CalculatedItems != nil {
		if err := m.CalculatedItems.ValidateWithPath(path + "/CalculatedItems"); err != nil {
			return err
		}
	}
	if m.CalculatedMembers != nil {
		if err := m.CalculatedMembers.ValidateWithPath(path + "/CalculatedMembers"); err != nil {
			return err
		}
	}
	if m.Dimensions != nil {
		if err := m.Dimensions.ValidateWithPath(path + "/Dimensions"); err != nil {
			return err
		}
	}
	if m.MeasureGroups != nil {
		if err := m.MeasureGroups.ValidateWithPath(path + "/MeasureGroups"); err != nil {
			return err
		}
	}
	if m.Maps != nil {
		if err := m.Maps.ValidateWithPath(path + "/Maps"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
