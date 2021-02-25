// Code generated by gowsdl DO NOT EDIT.

package myservice

import (
	"fmt"
	"context"
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

type CustomTime struct {
    time.Time
}

func (c *CustomTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var shortForm = "2006-01-02T15:04:05Z"
    var v string
	d.DecodeElement(&v, &start)
	if v != "" {
		parse, err := time.Parse(shortForm, v)
		if err != nil {
			var shortForm = "2006-01-02T15:04:05"
			parse, err = time.Parse(shortForm, v)
			if err != nil {
				return err
			}
		}
		*c = CustomTime{parse}
	}
    
    return nil
}

func (c *CustomTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	const shortForm = "2006-01-02T15:04:05Z"
	s := c.Format(shortForm)
	return e.EncodeElement(s, start)
}

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type CultureName string

type ErrorMessage struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/OrderShipmentNotificationService/1.0.0/SharedObjects/ errorMessage"`

	//
	// Response for any error requiring notification to requestor
	//
	Code int32 `xml:"code,omitempty" json:"code,omitempty"`

	//
	// Response for any error requiring notification to requestor
	//

	Description string `xml:"description,omitempty" json:"description,omitempty"`
}

type GetOrderShipmentNotificationRequest struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/OrderShipmentNotificationService/1.0.0/ GetOrderShipmentNotificationRequest"`

	WsVersion *string `xml:"http://www.promostandards.org/WSDL/PO/1.0.0/SharedObjects/ wsVersion,omitempty" json:"wsVersion,omitempty"`

	Id *string `xml:"http://www.promostandards.org/WSDL/PO/1.0.0/SharedObjects/ id,omitempty" json:"id,omitempty"`

	Password *string `xml:"http://www.promostandards.org/WSDL/PO/1.0.0/SharedObjects/ password,omitempty" json:"password,omitempty"`

	// The type of query you wish to perform.
	//

	QueryType int32 `xml:"queryType,omitempty" json:"queryType,omitempty"`

	// The purchase order or sales order number.  Required when the queryType is 1 or 2.
	//

	ReferenceNumber string `xml:"referenceNumber,omitempty" json:"referenceNumber,omitempty"`

	//
	// The earliest date for of shipments to return in UTC.  Required when the queryType is 3.  ISO 8601
	//
	ShipmentDateTimeStamp *CustomTime `xml:"shipmentDateTimeStamp,omitempty" json:"shipmentDateTimeStamp,omitempty"`
}

type ShipmentDestinationTypeType string

const (
	ShipmentDestinationTypeTypeCommercial ShipmentDestinationTypeType = "Commercial"

	ShipmentDestinationTypeTypeResidential ShipmentDestinationTypeType = "Residential"

	ShipmentDestinationTypeTypeNone ShipmentDestinationTypeType = "None"
)

type DimUOMType string

const (
	DimUOMTypeInches DimUOMType = "Inches"

	DimUOMTypeFeet DimUOMType = "Feet"

	DimUOMTypeMm DimUOMType = "mm"

	DimUOMTypeCm DimUOMType = "cm"

	DimUOMTypeMeters DimUOMType = "Meters"
)

type WeightUOMType string

const (
	WeightUOMTypeOunces WeightUOMType = "Ounces"

	WeightUOMTypePounds WeightUOMType = "Pounds"

	WeightUOMTypeGrams WeightUOMType = "Grams"

	WeightUOMTypeKG WeightUOMType = "KG"
)

type GetOrderShipmentNotificationResponse struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/OrderShipmentNotificationService/1.0.0/ GetOrderShipmentNotificationResponse"`

	OrderShipmentNotificationArray struct {
		OrderShipmentNotification []struct {

			// The associated purchase order.
			PurchaseOrderNumber string `xml:"purchaseOrderNumber,omitempty" json:"purchaseOrderNumber,omitempty"`

			// All shipments for this purchase order are complete.
			Complete bool `xml:"complete,omitempty" json:"complete,omitempty"`

			SalesOrderArray struct {
				SalesOrder []struct {

					// The associated sales order.

					SalesOrderNumber string `xml:"salesOrderNumber,omitempty" json:"salesOrderNumber,omitempty"`

					// All shipments for this sales order are complete.
					Complete bool `xml:"complete,omitempty" json:"complete,omitempty"`

					ShipmentLocationArray struct {
						ShipmentLocation []struct {

							// The id of the location.
							Id int32 `xml:"id,omitempty" json:"id,omitempty"`

							// All shipments for this location are complete.
							Complete bool `xml:"complete,omitempty" json:"complete,omitempty"`

							ShipFromAddress struct {
								Address1 *string `xml:"address1,omitempty" json:"address1,omitempty"`

								Address2 *string `xml:"address2,omitempty" json:"address2,omitempty"`

								Address3 *string `xml:"address3,omitempty" json:"address3,omitempty"`

								Address4 *string `xml:"address4,omitempty" json:"address4,omitempty"`

								City *string `xml:"city,omitempty" json:"city,omitempty"`

								Region *string `xml:"region,omitempty" json:"region,omitempty"`

								PostalCode *string `xml:"postalCode,omitempty" json:"postalCode,omitempty"`

								Country *string `xml:"country,omitempty" json:"country,omitempty"`
							} `xml:"ShipFromAddress,omitempty" json:"ShipFromAddress,omitempty"`

							ShipToAddress struct {
								Address1 *string `xml:"address1,omitempty" json:"address1,omitempty"`

								Address2 *string `xml:"address2,omitempty" json:"address2,omitempty"`

								Address3 *string `xml:"address3,omitempty" json:"address3,omitempty"`

								Address4 *string `xml:"address4,omitempty" json:"address4,omitempty"`

								City *string `xml:"city,omitempty" json:"city,omitempty"`

								Region *string `xml:"region,omitempty" json:"region,omitempty"`

								PostalCode *string `xml:"postalCode,omitempty" json:"postalCode,omitempty"`

								Country *string `xml:"country,omitempty" json:"country,omitempty"`
							} `xml:"ShipToAddress,omitempty" json:"ShipToAddress,omitempty"`

							// Used to identify the type of destination for the Ship-To address.
							ShipmentDestinationType *ShipmentDestinationTypeType `xml:"shipmentDestinationType,omitempty" json:"shipmentDestinationType,omitempty"`

							PackageArray struct {
								Package []struct {

									// The id of the package.
									Id int32 `xml:"id,omitempty" json:"id,omitempty"`

									// The tracking number for the package.

									TrackingNumber string `xml:"trackingNumber,omitempty" json:"trackingNumber,omitempty"`

									// The date for the shipment in UTC.  ISO 8601
									ShipmentDate *CustomTime `xml:"shipmentDate,omitempty" json:"shipmentDate,omitempty"`

									// The dimensional unit of measure.
									DimUOM *DimUOMType `xml:"dimUOM,omitempty" json:"dimUOM,omitempty"`

									// The length of the package.
									Length float64 `xml:"length,omitempty" json:"length,omitempty"`

									// The width of the package.
									Width float64 `xml:"width,omitempty" json:"width,omitempty"`

									// The height of the package.
									Height float64 `xml:"height,omitempty" json:"height,omitempty"`

									// The dimensional unit of measure.
									WeightUOM *WeightUOMType `xml:"weightUOM,omitempty" json:"weightUOM,omitempty"`

									// The weight of the package.
									Weight float64 `xml:"weight,omitempty" json:"weight,omitempty"`

									// The carrier delivering the package
									Carrier string `xml:"carrier,omitempty" json:"carrier,omitempty"`

									// The method used for shipping (Ground).

									ShipmentMethod string `xml:"shipmentMethod,omitempty" json:"shipmentMethod,omitempty"`

									// The shipping account used for this shipment.

									ShippingAccount string `xml:"shippingAccount,omitempty" json:"shippingAccount,omitempty"`

									// The terms of the shipment.

									ShipmentTerms string `xml:"shipmentTerms,omitempty" json:"shipmentTerms,omitempty"`

									ItemArray struct {
										Item []struct {

											// The supplier product Id

											SupplierProductId string `xml:"supplierProductId,omitempty" json:"supplierProductId,omitempty"`

											// The supplier part Id associated to the supplier product Id

											SupplierPartId string `xml:"supplierPartId,omitempty" json:"supplierPartId,omitempty"`

											// The distributor product Id

											DistributorProductId string `xml:"distributorProductId,omitempty" json:"distributorProductId,omitempty"`

											// The distributor part Id associated to the supplier product Id

											DistributorPartId string `xml:"distributorPartId,omitempty" json:"distributorPartId,omitempty"`

											// The line number of the item on the purchase order.

											PurchaseOrderLineNumber string `xml:"purchaseOrderLineNumber,omitempty" json:"purchaseOrderLineNumber,omitempty"`

											// The length of the package.
											Quantity float64 `xml:"quantity,omitempty" json:"quantity,omitempty"`
										} `xml:"Item,omitempty" json:"Item,omitempty"`
									} `xml:"ItemArray,omitempty" json:"ItemArray,omitempty"`
								} `xml:"Package,omitempty" json:"Package,omitempty"`
							} `xml:"PackageArray,omitempty" json:"PackageArray,omitempty"`
						} `xml:"ShipmentLocation,omitempty" json:"ShipmentLocation,omitempty"`
					} `xml:"ShipmentLocationArray,omitempty" json:"ShipmentLocationArray,omitempty"`
				} `xml:"SalesOrder,omitempty" json:"SalesOrder,omitempty"`
			} `xml:"SalesOrderArray,omitempty" json:"SalesOrderArray,omitempty"`
		} `xml:"OrderShipmentNotification,omitempty" json:"OrderShipmentNotification,omitempty"`
	} `xml:"OrderShipmentNotificationArray,omitempty" json:"OrderShipmentNotificationArray,omitempty"`

	ErrorMessage *ErrorMessage `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`
}

type OrderShipmentNotificationService interface {
	GetOrderShipmentNotification(request *GetOrderShipmentNotificationRequest) (*GetOrderShipmentNotificationResponse, error)

	GetOrderShipmentNotificationContext(ctx context.Context, request *GetOrderShipmentNotificationRequest) (*GetOrderShipmentNotificationResponse, error)
}

type orderShipmentNotificationService struct {
	client *soap.Client
}

func NewOrderShipmentNotificationService(client *soap.Client) OrderShipmentNotificationService {
	return &orderShipmentNotificationService{
		client: client,
	}
}

func (service *orderShipmentNotificationService) GetOrderShipmentNotificationContext(ctx context.Context, request *GetOrderShipmentNotificationRequest) (*GetOrderShipmentNotificationResponse, error) {
	response := new(GetOrderShipmentNotificationResponse)
	c, _ := xml.Marshal(request)
	fmt.Println(string(c))
	err := service.client.CallContext(ctx, "getOrderShipmentNotification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *orderShipmentNotificationService) GetOrderShipmentNotification(request *GetOrderShipmentNotificationRequest) (*GetOrderShipmentNotificationResponse, error) {
	return service.GetOrderShipmentNotificationContext(
		context.Background(),
		request,
	)
}
