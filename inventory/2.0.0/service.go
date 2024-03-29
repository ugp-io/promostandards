// Code generated by gowsdl DO NOT EDIT.

package service

import (
	"context"
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

//
// Two-letter (alpha-2) ISO 3166-1 code for one of the 243 countries.
//

type ISO3166CountyCode string

const (
	ISO3166CountyCodeAF ISO3166CountyCode = "AF"

	ISO3166CountyCodeAX ISO3166CountyCode = "AX"

	ISO3166CountyCodeAL ISO3166CountyCode = "AL"

	ISO3166CountyCodeDZ ISO3166CountyCode = "DZ"

	ISO3166CountyCodeAS ISO3166CountyCode = "AS"

	ISO3166CountyCodeAD ISO3166CountyCode = "AD"

	ISO3166CountyCodeAO ISO3166CountyCode = "AO"

	ISO3166CountyCodeAI ISO3166CountyCode = "AI"

	ISO3166CountyCodeAQ ISO3166CountyCode = "AQ"

	ISO3166CountyCodeAG ISO3166CountyCode = "AG"

	ISO3166CountyCodeAR ISO3166CountyCode = "AR"

	ISO3166CountyCodeAM ISO3166CountyCode = "AM"

	ISO3166CountyCodeAW ISO3166CountyCode = "AW"

	ISO3166CountyCodeAU ISO3166CountyCode = "AU"

	ISO3166CountyCodeAT ISO3166CountyCode = "AT"

	ISO3166CountyCodeAZ ISO3166CountyCode = "AZ"

	ISO3166CountyCodeBS ISO3166CountyCode = "BS"

	ISO3166CountyCodeBH ISO3166CountyCode = "BH"

	ISO3166CountyCodeBD ISO3166CountyCode = "BD"

	ISO3166CountyCodeBB ISO3166CountyCode = "BB"

	ISO3166CountyCodeBY ISO3166CountyCode = "BY"

	ISO3166CountyCodeBE ISO3166CountyCode = "BE"

	ISO3166CountyCodeBZ ISO3166CountyCode = "BZ"

	ISO3166CountyCodeBJ ISO3166CountyCode = "BJ"

	ISO3166CountyCodeBM ISO3166CountyCode = "BM"

	ISO3166CountyCodeBT ISO3166CountyCode = "BT"

	ISO3166CountyCodeBO ISO3166CountyCode = "BO"

	ISO3166CountyCodeBA ISO3166CountyCode = "BA"

	ISO3166CountyCodeBW ISO3166CountyCode = "BW"

	ISO3166CountyCodeBV ISO3166CountyCode = "BV"

	ISO3166CountyCodeBR ISO3166CountyCode = "BR"

	ISO3166CountyCodeIO ISO3166CountyCode = "IO"

	ISO3166CountyCodeBN ISO3166CountyCode = "BN"

	ISO3166CountyCodeBG ISO3166CountyCode = "BG"

	ISO3166CountyCodeBF ISO3166CountyCode = "BF"

	ISO3166CountyCodeBI ISO3166CountyCode = "BI"

	ISO3166CountyCodeKH ISO3166CountyCode = "KH"

	ISO3166CountyCodeCM ISO3166CountyCode = "CM"

	ISO3166CountyCodeCA ISO3166CountyCode = "CA"

	ISO3166CountyCodeCV ISO3166CountyCode = "CV"

	ISO3166CountyCodeKY ISO3166CountyCode = "KY"

	ISO3166CountyCodeCF ISO3166CountyCode = "CF"

	ISO3166CountyCodeTD ISO3166CountyCode = "TD"

	ISO3166CountyCodeCL ISO3166CountyCode = "CL"

	ISO3166CountyCodeCN ISO3166CountyCode = "CN"

	ISO3166CountyCodeCX ISO3166CountyCode = "CX"

	ISO3166CountyCodeCC ISO3166CountyCode = "CC"

	ISO3166CountyCodeCO ISO3166CountyCode = "CO"

	ISO3166CountyCodeKM ISO3166CountyCode = "KM"

	ISO3166CountyCodeCG ISO3166CountyCode = "CG"

	ISO3166CountyCodeCD ISO3166CountyCode = "CD"

	ISO3166CountyCodeCK ISO3166CountyCode = "CK"

	ISO3166CountyCodeCR ISO3166CountyCode = "CR"

	ISO3166CountyCodeCI ISO3166CountyCode = "CI"

	ISO3166CountyCodeHR ISO3166CountyCode = "HR"

	ISO3166CountyCodeCU ISO3166CountyCode = "CU"

	ISO3166CountyCodeCY ISO3166CountyCode = "CY"

	ISO3166CountyCodeCZ ISO3166CountyCode = "CZ"

	ISO3166CountyCodeDK ISO3166CountyCode = "DK"

	ISO3166CountyCodeDJ ISO3166CountyCode = "DJ"

	ISO3166CountyCodeDM ISO3166CountyCode = "DM"

	ISO3166CountyCodeDO ISO3166CountyCode = "DO"

	ISO3166CountyCodeEC ISO3166CountyCode = "EC"

	ISO3166CountyCodeEG ISO3166CountyCode = "EG"

	ISO3166CountyCodeSV ISO3166CountyCode = "SV"

	ISO3166CountyCodeGQ ISO3166CountyCode = "GQ"

	ISO3166CountyCodeER ISO3166CountyCode = "ER"

	ISO3166CountyCodeEE ISO3166CountyCode = "EE"

	ISO3166CountyCodeET ISO3166CountyCode = "ET"

	ISO3166CountyCodeFK ISO3166CountyCode = "FK"

	ISO3166CountyCodeFO ISO3166CountyCode = "FO"

	ISO3166CountyCodeFJ ISO3166CountyCode = "FJ"

	ISO3166CountyCodeFI ISO3166CountyCode = "FI"

	ISO3166CountyCodeFR ISO3166CountyCode = "FR"

	ISO3166CountyCodeGF ISO3166CountyCode = "GF"

	ISO3166CountyCodePF ISO3166CountyCode = "PF"

	ISO3166CountyCodeTF ISO3166CountyCode = "TF"

	ISO3166CountyCodeGA ISO3166CountyCode = "GA"

	ISO3166CountyCodeGM ISO3166CountyCode = "GM"

	ISO3166CountyCodeGE ISO3166CountyCode = "GE"

	ISO3166CountyCodeDE ISO3166CountyCode = "DE"

	ISO3166CountyCodeGH ISO3166CountyCode = "GH"

	ISO3166CountyCodeGI ISO3166CountyCode = "GI"

	ISO3166CountyCodeGR ISO3166CountyCode = "GR"

	ISO3166CountyCodeGL ISO3166CountyCode = "GL"

	ISO3166CountyCodeGD ISO3166CountyCode = "GD"

	ISO3166CountyCodeGP ISO3166CountyCode = "GP"

	ISO3166CountyCodeGU ISO3166CountyCode = "GU"

	ISO3166CountyCodeGT ISO3166CountyCode = "GT"

	ISO3166CountyCodeGG ISO3166CountyCode = "GG"

	ISO3166CountyCodeGN ISO3166CountyCode = "GN"

	ISO3166CountyCodeGW ISO3166CountyCode = "GW"

	ISO3166CountyCodeGY ISO3166CountyCode = "GY"

	ISO3166CountyCodeHT ISO3166CountyCode = "HT"

	ISO3166CountyCodeHM ISO3166CountyCode = "HM"

	ISO3166CountyCodeVA ISO3166CountyCode = "VA"

	ISO3166CountyCodeHN ISO3166CountyCode = "HN"

	ISO3166CountyCodeHK ISO3166CountyCode = "HK"

	ISO3166CountyCodeHU ISO3166CountyCode = "HU"

	ISO3166CountyCodeIS ISO3166CountyCode = "IS"

	ISO3166CountyCodeIN ISO3166CountyCode = "IN"

	ISO3166CountyCodeID ISO3166CountyCode = "ID"

	ISO3166CountyCodeIR ISO3166CountyCode = "IR"

	ISO3166CountyCodeIQ ISO3166CountyCode = "IQ"

	ISO3166CountyCodeIE ISO3166CountyCode = "IE"

	ISO3166CountyCodeIM ISO3166CountyCode = "IM"

	ISO3166CountyCodeIL ISO3166CountyCode = "IL"

	ISO3166CountyCodeIT ISO3166CountyCode = "IT"

	ISO3166CountyCodeJM ISO3166CountyCode = "JM"

	ISO3166CountyCodeJP ISO3166CountyCode = "JP"

	ISO3166CountyCodeJE ISO3166CountyCode = "JE"

	ISO3166CountyCodeJO ISO3166CountyCode = "JO"

	ISO3166CountyCodeKZ ISO3166CountyCode = "KZ"

	ISO3166CountyCodeKE ISO3166CountyCode = "KE"

	ISO3166CountyCodeKI ISO3166CountyCode = "KI"

	ISO3166CountyCodeKP ISO3166CountyCode = "KP"

	ISO3166CountyCodeKR ISO3166CountyCode = "KR"

	ISO3166CountyCodeKW ISO3166CountyCode = "KW"

	ISO3166CountyCodeKG ISO3166CountyCode = "KG"

	ISO3166CountyCodeLA ISO3166CountyCode = "LA"

	ISO3166CountyCodeLV ISO3166CountyCode = "LV"

	ISO3166CountyCodeLB ISO3166CountyCode = "LB"

	ISO3166CountyCodeLS ISO3166CountyCode = "LS"

	ISO3166CountyCodeLR ISO3166CountyCode = "LR"

	ISO3166CountyCodeLY ISO3166CountyCode = "LY"

	ISO3166CountyCodeLI ISO3166CountyCode = "LI"

	ISO3166CountyCodeLT ISO3166CountyCode = "LT"

	ISO3166CountyCodeLU ISO3166CountyCode = "LU"

	ISO3166CountyCodeMO ISO3166CountyCode = "MO"

	ISO3166CountyCodeMK ISO3166CountyCode = "MK"

	ISO3166CountyCodeMG ISO3166CountyCode = "MG"

	ISO3166CountyCodeMW ISO3166CountyCode = "MW"

	ISO3166CountyCodeMY ISO3166CountyCode = "MY"

	ISO3166CountyCodeMV ISO3166CountyCode = "MV"

	ISO3166CountyCodeML ISO3166CountyCode = "ML"

	ISO3166CountyCodeMT ISO3166CountyCode = "MT"

	ISO3166CountyCodeMH ISO3166CountyCode = "MH"

	ISO3166CountyCodeMQ ISO3166CountyCode = "MQ"

	ISO3166CountyCodeMR ISO3166CountyCode = "MR"

	ISO3166CountyCodeMU ISO3166CountyCode = "MU"

	ISO3166CountyCodeYT ISO3166CountyCode = "YT"

	ISO3166CountyCodeMX ISO3166CountyCode = "MX"

	ISO3166CountyCodeFM ISO3166CountyCode = "FM"

	ISO3166CountyCodeMD ISO3166CountyCode = "MD"

	ISO3166CountyCodeMC ISO3166CountyCode = "MC"

	ISO3166CountyCodeMN ISO3166CountyCode = "MN"

	ISO3166CountyCodeMS ISO3166CountyCode = "MS"

	ISO3166CountyCodeMA ISO3166CountyCode = "MA"

	ISO3166CountyCodeMZ ISO3166CountyCode = "MZ"

	ISO3166CountyCodeMM ISO3166CountyCode = "MM"

	ISO3166CountyCodeNA ISO3166CountyCode = "NA"

	ISO3166CountyCodeNR ISO3166CountyCode = "NR"

	ISO3166CountyCodeNP ISO3166CountyCode = "NP"

	ISO3166CountyCodeNL ISO3166CountyCode = "NL"

	ISO3166CountyCodeAN ISO3166CountyCode = "AN"

	ISO3166CountyCodeNC ISO3166CountyCode = "NC"

	ISO3166CountyCodeNZ ISO3166CountyCode = "NZ"

	ISO3166CountyCodeNI ISO3166CountyCode = "NI"

	ISO3166CountyCodeNE ISO3166CountyCode = "NE"

	ISO3166CountyCodeNG ISO3166CountyCode = "NG"

	ISO3166CountyCodeNU ISO3166CountyCode = "NU"

	ISO3166CountyCodeNF ISO3166CountyCode = "NF"

	ISO3166CountyCodeMP ISO3166CountyCode = "MP"

	ISO3166CountyCodeNO ISO3166CountyCode = "NO"

	ISO3166CountyCodeOM ISO3166CountyCode = "OM"

	ISO3166CountyCodePK ISO3166CountyCode = "PK"

	ISO3166CountyCodePW ISO3166CountyCode = "PW"

	ISO3166CountyCodePS ISO3166CountyCode = "PS"

	ISO3166CountyCodePA ISO3166CountyCode = "PA"

	ISO3166CountyCodePG ISO3166CountyCode = "PG"

	ISO3166CountyCodePY ISO3166CountyCode = "PY"

	ISO3166CountyCodePE ISO3166CountyCode = "PE"

	ISO3166CountyCodePH ISO3166CountyCode = "PH"

	ISO3166CountyCodePN ISO3166CountyCode = "PN"

	ISO3166CountyCodePL ISO3166CountyCode = "PL"

	ISO3166CountyCodePT ISO3166CountyCode = "PT"

	ISO3166CountyCodePR ISO3166CountyCode = "PR"

	ISO3166CountyCodeQA ISO3166CountyCode = "QA"

	ISO3166CountyCodeRE ISO3166CountyCode = "RE"

	ISO3166CountyCodeRO ISO3166CountyCode = "RO"

	ISO3166CountyCodeRU ISO3166CountyCode = "RU"

	ISO3166CountyCodeRW ISO3166CountyCode = "RW"

	ISO3166CountyCodeSH ISO3166CountyCode = "SH"

	ISO3166CountyCodeKN ISO3166CountyCode = "KN"

	ISO3166CountyCodeLC ISO3166CountyCode = "LC"

	ISO3166CountyCodePM ISO3166CountyCode = "PM"

	ISO3166CountyCodeVC ISO3166CountyCode = "VC"

	ISO3166CountyCodeWS ISO3166CountyCode = "WS"

	ISO3166CountyCodeSM ISO3166CountyCode = "SM"

	ISO3166CountyCodeST ISO3166CountyCode = "ST"

	ISO3166CountyCodeSA ISO3166CountyCode = "SA"

	ISO3166CountyCodeSN ISO3166CountyCode = "SN"

	ISO3166CountyCodeCS ISO3166CountyCode = "CS"

	ISO3166CountyCodeSC ISO3166CountyCode = "SC"

	ISO3166CountyCodeSL ISO3166CountyCode = "SL"

	ISO3166CountyCodeSG ISO3166CountyCode = "SG"

	ISO3166CountyCodeSK ISO3166CountyCode = "SK"

	ISO3166CountyCodeSI ISO3166CountyCode = "SI"

	ISO3166CountyCodeSB ISO3166CountyCode = "SB"

	ISO3166CountyCodeSO ISO3166CountyCode = "SO"

	ISO3166CountyCodeZA ISO3166CountyCode = "ZA"

	ISO3166CountyCodeGS ISO3166CountyCode = "GS"

	ISO3166CountyCodeES ISO3166CountyCode = "ES"

	ISO3166CountyCodeLK ISO3166CountyCode = "LK"

	ISO3166CountyCodeSD ISO3166CountyCode = "SD"

	ISO3166CountyCodeSR ISO3166CountyCode = "SR"

	ISO3166CountyCodeSJ ISO3166CountyCode = "SJ"

	ISO3166CountyCodeSZ ISO3166CountyCode = "SZ"

	ISO3166CountyCodeSE ISO3166CountyCode = "SE"

	ISO3166CountyCodeCH ISO3166CountyCode = "CH"

	ISO3166CountyCodeSY ISO3166CountyCode = "SY"

	ISO3166CountyCodeTW ISO3166CountyCode = "TW"

	ISO3166CountyCodeTJ ISO3166CountyCode = "TJ"

	ISO3166CountyCodeTZ ISO3166CountyCode = "TZ"

	ISO3166CountyCodeTH ISO3166CountyCode = "TH"

	ISO3166CountyCodeTL ISO3166CountyCode = "TL"

	ISO3166CountyCodeTG ISO3166CountyCode = "TG"

	ISO3166CountyCodeTK ISO3166CountyCode = "TK"

	ISO3166CountyCodeTO ISO3166CountyCode = "TO"

	ISO3166CountyCodeTT ISO3166CountyCode = "TT"

	ISO3166CountyCodeTN ISO3166CountyCode = "TN"

	ISO3166CountyCodeTR ISO3166CountyCode = "TR"

	ISO3166CountyCodeTM ISO3166CountyCode = "TM"

	ISO3166CountyCodeTC ISO3166CountyCode = "TC"

	ISO3166CountyCodeTV ISO3166CountyCode = "TV"

	ISO3166CountyCodeUG ISO3166CountyCode = "UG"

	ISO3166CountyCodeUA ISO3166CountyCode = "UA"

	ISO3166CountyCodeAE ISO3166CountyCode = "AE"

	ISO3166CountyCodeGB ISO3166CountyCode = "GB"

	ISO3166CountyCodeUS ISO3166CountyCode = "US"

	ISO3166CountyCodeUM ISO3166CountyCode = "UM"

	ISO3166CountyCodeUY ISO3166CountyCode = "UY"

	ISO3166CountyCodeUZ ISO3166CountyCode = "UZ"

	ISO3166CountyCodeVU ISO3166CountyCode = "VU"

	ISO3166CountyCodeVE ISO3166CountyCode = "VE"

	ISO3166CountyCodeVN ISO3166CountyCode = "VN"

	ISO3166CountyCodeVG ISO3166CountyCode = "VG"

	ISO3166CountyCodeVI ISO3166CountyCode = "VI"

	ISO3166CountyCodeWF ISO3166CountyCode = "WF"

	ISO3166CountyCodeEH ISO3166CountyCode = "EH"

	ISO3166CountyCodeYE ISO3166CountyCode = "YE"

	ISO3166CountyCodeZM ISO3166CountyCode = "ZM"

	ISO3166CountyCodeZW ISO3166CountyCode = "ZW"
)

type LabelSizeEnum string

const (
	LabelSizeEnum2XL LabelSizeEnum = "2XL"

	LabelSizeEnum2XS LabelSizeEnum = "2XS"

	LabelSizeEnum3XL LabelSizeEnum = "3XL"

	LabelSizeEnum3XS LabelSizeEnum = "3XS"

	LabelSizeEnum4XL LabelSizeEnum = "4XL"

	LabelSizeEnum4XS LabelSizeEnum = "4XS"

	LabelSizeEnum5XL LabelSizeEnum = "5XL"

	LabelSizeEnum5XS LabelSizeEnum = "5XS"

	LabelSizeEnum6XL LabelSizeEnum = "6XL"

	LabelSizeEnum6XS LabelSizeEnum = "6XS"

	LabelSizeEnumCUSTOM LabelSizeEnum = "CUSTOM"

	LabelSizeEnumL LabelSizeEnum = "L"

	LabelSizeEnumM LabelSizeEnum = "M"

	LabelSizeEnumOSFA LabelSizeEnum = "OSFA"

	LabelSizeEnumS LabelSizeEnum = "S"

	LabelSizeEnumXL LabelSizeEnum = "XL"

	LabelSizeEnumXS LabelSizeEnum = "XS"
)

//
// The type of Quantity UOM
//

type QuantityUomType string

const (
	QuantityUomTypeBX QuantityUomType = "BX"

	QuantityUomTypeCA QuantityUomType = "CA"

	QuantityUomTypeDZ QuantityUomType = "DZ"

	QuantityUomTypeEA QuantityUomType = "EA"

	QuantityUomTypeKT QuantityUomType = "KT"

	QuantityUomTypePK QuantityUomType = "PK"

	QuantityUomTypePR QuantityUomType = "PR"

	QuantityUomTypeRL QuantityUomType = "RL"

	QuantityUomTypeSL QuantityUomType = "SL"

	QuantityUomTypeST QuantityUomType = "ST"

	QuantityUomTypeTH QuantityUomType = "TH"
)

//
// The severity type
//

type SeverityType string

const (
	SeverityTypeError SeverityType = "Error"

	SeverityTypeInformation SeverityType = "Information"

	SeverityTypeWarning SeverityType = "Warning"
)

//
// The severity type
//

type WsVersionType string

const (
	WsVersionType2_0_0 WsVersionType = "2.0.0"
)

type ServiceMessageArray struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ ServiceMessageArray" bson:"-"`

	ServiceMessage []*ServiceMessage `xml:"ServiceMessage,omitempty" json:"ServiceMessage,omitempty" bson:"service_message,omitempty"`
}

type ServiceMessage struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ ServiceMessage" bson:"-"`

	// Response for any message requiring notification to requestor
	//
	Code int32 `xml:"code,omitempty" json:"code,omitempty" bson:"code,omitempty"`

	//
	// Response for any message requiring notification to requestor
	//

	Description string `xml:"description,omitempty" json:"description,omitempty" bson:"description,omitempty"`

	//
	// The severity of the message. Values are enumerated: {Error, Information, Warning}
	//

	Severity *SeverityType `xml:"severity,omitempty" json:"severity,omitempty" bson:"severity,omitempty"`
}

type Filter struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ Filter" bson:"-"`

	PartIdArray *PartIdArray `xml:"partIdArray,omitempty" json:"partIdArray,omitempty" bson:"part_id_array,omitempty"`

	LabelSizeArray *LabelSizeArray `xml:"LabelSizeArray,omitempty" json:"LabelSizeArray,omitempty" bson:"label_size_array,omitempty"`

	PartColorArray *PartColorArray `xml:"PartColorArray,omitempty" json:"PartColorArray,omitempty" bson:"part_color_array,omitempty"`
}

type FutureAvailabilityArray struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ FutureAvailabilityArray" bson:"-"`

	FutureAvailability []*FutureAvailability `xml:"FutureAvailability,omitempty" json:"FutureAvailability,omitempty" bson:"future_availability,omitempty"`
}

type AvailableOn time.Time

type FutureAvailability struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ FutureAvailability" bson:"-"`

	Quantity *Quantity `xml:"Quantity,omitempty" json:"Quantity,omitempty" bson:"quantity,omitempty"`

	AvailableOn *AvailableOn `xml:"availableOn,omitempty" json:"availableOn,omitempty" bson:"available_on,omitempty"`
}

type Inventory struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ Inventory" bson:"-"`

	ProductId *string `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ productId,omitempty" json:"productId,omitempty" bson:"product_id,omitempty"`

	PartInventoryArray *PartInventoryArray `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ PartInventoryArray,omitempty" json:"PartInventoryArray,omitempty" bson:"part_inventory_array,omitempty"`
}

type LabelSizeArray struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ LabelSizeArray" bson:"-"`

	LabelSize []*LabelSize `xml:"labelSize,omitempty" json:"labelSize,omitempty" bson:"label_size,omitempty"`
}

type PartColorArray struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ PartColorArray" bson:"-"`

	PartColor []*PartColor `xml:"partColor,omitempty" json:"partColor,omitempty" bson:"part_color,omitempty"`
}

type PartInventoryArray struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ PartInventoryArray" bson:"-"`

	PartInventory []struct {
		PartId *string `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ partId,omitempty" json:"partId,omitempty" bson:"part_id,omitempty"`

		MainPart *MainPart `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ mainPart,omitempty" json:"mainPart,omitempty" bson:"main_part,omitempty"`

		PartColor *PartColor `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ partColor,omitempty" json:"partColor,omitempty" bson:"part_color,omitempty"`

		LabelSize *LabelSize `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ labelSize,omitempty" json:"labelSize,omitempty" bson:"label_size,omitempty"`

		PartDescription *PartDescription `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ partDescription,omitempty" json:"partDescription,omitempty" bson:"part_description,omitempty"`

		QuantityAvailable *QuantityAvailable `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ quantityAvailable,omitempty" json:"quantityAvailable,omitempty" bson:"quantity_available,omitempty"`

		ManufacturedItem *ManufacturedItem `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ manufacturedItem,omitempty" json:"manufacturedItem,omitempty" bson:"manufactured_item,omitempty"`

		BuyToOrder *BuyToOrder `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ buyToOrder,omitempty" json:"buyToOrder,omitempty" bson:"buy_to_order,omitempty"`

		ReplenishmentLeadTime *ReplenishmentLeadTime `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ replenishmentLeadTime,omitempty" json:"replenishmentLeadTime,omitempty" bson:"replenishment_lead_time,omitempty"`

		AttributeSelection *AttributeSelection `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ attributeSelection,omitempty" json:"attributeSelection,omitempty" bson:"attribute_selection,omitempty"`

		InventoryLocationArray *InventoryLocationArray `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ InventoryLocationArray,omitempty" json:"InventoryLocationArray,omitempty" bson:"inventory_location_array,omitempty"`

		LastModified *LastModified `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ lastModified,omitempty" json:"lastModified,omitempty" bson:"last_modified,omitempty"`
	} `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ PartInventory,omitempty" json:"PartInventory,omitempty" bson:"part_inventory,omitempty"`
}

type QuantityAvailable struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ quantityAvailable" bson:"-"`

	Quantity *Quantity `xml:"Quantity,omitempty" json:"Quantity,omitempty" bson:"quantity,omitempty"`
}

type Quantity struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ Quantity" bson:"-"`

	//
	// The unit of measure; values are enumerated.
	//

	Uom *QuantityUomType `xml:"uom,omitempty" json:"uom,omitempty" bson:"uom,omitempty"`

	//
	// The quantity value
	//
	Value float64 `xml:"value,omitempty" json:"value,omitempty" bson:"value,omitempty"`
}

type PartIdArray struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ partIdArray" bson:"-"`

	PartId []*string `xml:"partId,omitempty" json:"partId,omitempty" bson:"partId,omitempty"`
}

type LastModified time.Time

type InventoryLocationArray struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ InventoryLocationArray" bson:"-"`

	InventoryLocation []*InventoryLocation `xml:"InventoryLocation,omitempty" json:"InventoryLocation,omitempty" bson:"inventory_location,omitempty"`
}

type InventoryLocationQuantity struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ inventoryLocationQuantity" bson:"-"`

	Quantity *Quantity `xml:"Quantity,omitempty" json:"Quantity,omitempty" bson:"quantity,omitempty"`
}

type InventoryLocation struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ InventoryLocation" bson:"-"`

	InventoryLocationId *string `xml:"inventoryLocationId,omitempty" json:"inventoryLocationId,omitempty" bson:"inventory_location_id,omitempty"`

	InventoryLocationName *string `xml:"inventoryLocationName,omitempty" json:"inventoryLocationName,omitempty" bson:"inventory_location_name,omitempty"`

	// The postal code

	PostalCode string `xml:"postalCode,omitempty" json:"postalCode,omitempty" bson:"postal_code,omitempty"`

	// The country

	Country *string `xml:"country,omitempty" json:"country,omitempty" bson:"country,omitempty"`

	InventoryLocationQuantity *InventoryLocationQuantity `xml:"inventoryLocationQuantity,omitempty" json:"inventoryLocationQuantity,omitempty" bson:"inventory_location_quantity,omitempty"`

	FutureAvailabilityArray *FutureAvailabilityArray `xml:"FutureAvailabilityArray,omitempty" json:"FutureAvailabilityArray,omitempty" bson:"future_availability_array,omitempty"`
}

type GetFilterValuesRequest struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/ GetFilterValuesRequest" bson:"-"`

	WsVersion *string `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ wsVersion,omitempty" json:"wsVersion,omitempty"`

	Id *string `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ id,omitempty" json:"id,omitempty"`

	Password *string `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ password,omitempty" json:"password,omitempty"`

	ProductId *string `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ productId,omitempty" json:"productId,omitempty"`
}

type GetFilterValuesResponse struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/ GetFilterValuesResponse" bson:"-"`

	FilterValues struct {
		ProductId *string `xml:"productId,omitempty" json:"productId,omitempty" bson:"product_id,omitempty"`

		Filter *Filter `xml:"Filter,omitempty" json:"Filter,omitempty" bson:"filter,omitempty"`
	} `xml:"FilterValues,omitempty" json:"FilterValues,omitempty" bson:"filter_values,omitempty"`

	ServiceMessageArray *ServiceMessageArray `xml:"ServiceMessageArray,omitempty" json:"ServiceMessageArray,omitempty" bson:"service_message_array,omitempty"`
}

type GetInventoryLevelsRequest struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/ GetInventoryLevelsRequest" bson:"-"`

	WsVersion *string `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ wsVersion,omitempty" json:"wsVersion,omitempty"`

	Id *string `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ id,omitempty" json:"id,omitempty"`

	Password *string `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ password,omitempty" json:"password,omitempty"`

	ProductId *string `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ productId,omitempty" json:"productId,omitempty"`

	Filter *Filter `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/SharedObjects/ Filter,omitempty" json:"Filter,omitempty"`
}

type GetInventoryLevelsResponse struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/Inventory/2.0.0/ GetInventoryLevelsResponse" bson:"-"`

	Inventory *Inventory `xml:"Inventory,omitempty" json:"Inventory,omitempty" bson:"inventory,omitempty"`

	ServiceMessageArray *ServiceMessageArray `xml:"ServiceMessageArray,omitempty" json:"ServiceMessageArray,omitempty" bson:"service_message_array,omitempty"`
}

type InventoryService interface {
	GetFilterValues(request *GetFilterValuesRequest) (*GetFilterValuesResponse, error)

	GetFilterValuesContext(ctx context.Context, request *GetFilterValuesRequest) (*GetFilterValuesResponse, error)

	GetInventoryLevels(request *GetInventoryLevelsRequest) (*GetInventoryLevelsResponse, error)

	GetInventoryLevelsContext(ctx context.Context, request *GetInventoryLevelsRequest) (*GetInventoryLevelsResponse, error)
}

type inventoryService struct {
	client *soap.Client
}

func NewInventoryService(client *soap.Client) InventoryService {
	return &inventoryService{
		client: client,
	}
}

func (service *inventoryService) GetFilterValuesContext(ctx context.Context, request *GetFilterValuesRequest) (*GetFilterValuesResponse, error) {
	response := new(GetFilterValuesResponse)
	err := service.client.CallContext(ctx, "getFilterValues", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *inventoryService) GetFilterValues(request *GetFilterValuesRequest) (*GetFilterValuesResponse, error) {
	return service.GetFilterValuesContext(
		context.Background(),
		request,
	)
}

func (service *inventoryService) GetInventoryLevelsContext(ctx context.Context, request *GetInventoryLevelsRequest) (*GetInventoryLevelsResponse, error) {
	response := new(GetInventoryLevelsResponse)
	err := service.client.CallContext(ctx, "getInventoryLevels", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *inventoryService) GetInventoryLevels(request *GetInventoryLevelsRequest) (*GetInventoryLevelsResponse, error) {
	return service.GetInventoryLevelsContext(
		context.Background(),
		request,
	)
}
