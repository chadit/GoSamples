package Mws

import "github.com/svvu/gomws/gmws"

func getXMLStringValue(xmlNode []gmws.XMLNode) string {
	if xmlNode == nil || len(xmlNode) == 0 {
		return ""
	}

	returnValue, _ := xmlNode[0].ToString()
	return returnValue
}

func getXMLBoolFromXMLNode(xmlNode []gmws.XMLNode) bool {
	if xmlNode == nil || len(xmlNode) == 0 {
		return false
	}
	returnValue, _ := xmlNode[0].ToBool()
	return returnValue
}

func getXMLIntValue(xmlNode []gmws.XMLNode) int {
	if xmlNode == nil || len(xmlNode) == 0 {
		return 0
	}

	returnValue, _ := xmlNode[0].ToInt()
	return returnValue
}

func getXMLFloat64Value(xmlNode []gmws.XMLNode) float64 {
	if xmlNode == nil || len(xmlNode) == 0 {
		return 0
	}
	returnValue, _ := xmlNode[0].ToFloat()
	return returnValue
}

func getFloat64FromXMLNode(xmlNode gmws.XMLNode) float64 {
	returnValue, _ := xmlNode.ToFloat()
	return returnValue
}

func getProductDemensions(itemDemensions gmws.XMLNode, pathValue string) float64 {
	productHeight := itemDemensions.FindByPath(pathValue)
	if len(productHeight) > 0 {
		return getFloat64FromXMLNode(productHeight[0])
	}
	return 0
}
