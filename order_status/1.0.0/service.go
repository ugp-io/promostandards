// Code generated by gowsdl DO NOT EDIT.

package myservice

import (
	// "fmt"
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
			var shortForm = "2006-01-02T15:04:05Z"
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

type GetOrderStatusDetailsRequest struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/OrderStatusService/1.0.0/ GetOrderStatusDetailsRequest"`

	// The Standard Version of the Web Service being referenced

	WsVersion *string `xml:"wsVersion,omitempty" json:"wsVersion,omitempty"`

	// The customerID or any other agreed upon ID.

	Id *string `xml:"id,omitempty" json:"id,omitempty"`

	//
	// The password associated with the customerID.
	//

	Password *string `xml:"password,omitempty" json:"password,omitempty"`

	//
	// The type of query you wish to perform. 1 = PO Number, 2 = Sales Order Number, 3 = Last update, 4 = All Open
	//
	QueryType int32 `xml:"queryType,omitempty" json:"queryType,omitempty"`

	//
	// The purchase order number associated with the Order. If you pass in a purchaseOrderNumber then you will expect to a response of an array with a SINGLE order status for that purchase order.  If this field is left blank it will assume that you are requesting multiple order statuses.
	//

	ReferenceNumber string `xml:"referenceNumber,omitempty" json:"referenceNumber,omitempty"`

	//
	// Beginning date time since last status change. This field will indicate the change date time for any status change GREATER the date time stamp provided.
	//
	StatusTimeStamp *CustomTime `xml:"statusTimeStamp,omitempty" json:"statusTimeStamp,omitempty"`
}

type GetOrderStatusDetailsResponse struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/OrderStatusService/1.0.0/ GetOrderStatusDetailsResponse"`

	OrderStatusArray struct {
		OrderStatus []struct {

			//
			// The associated Purchase Order Number from the customer
			//

			PurchaseOrderNumber string `xml:"purchaseOrderNumber,omitempty" json:"purchaseOrderNumber,omitempty"`

			OrderStatusDetailArray struct {
				OrderStatusDetail []struct {

					//
					// The associated factory sales order number
					// (This is used in the example of internal PO splitting at the factory level)
					//

					FactoryOrderNumber string `xml:"factoryOrderNumber,omitempty" json:"factoryOrderNumber,omitempty"`

					//
					// Paired standard order status from allowed values
					//

					StatusID int32 `xml:"statusID,omitempty" json:"statusID,omitempty"`

					//
					// Textual description of statusID
					//

					StatusName string `xml:"statusName,omitempty" json:"statusName,omitempty"`

					//
					// The expected ship date for the purchase order
					//
					ExpectedShipDate *CustomTime `xml:"expectedShipDate,omitempty" json:"expectedShipDate,omitempty"`

					//
					// The expected date the order should arrive at customer also known as the “in hands date”
					//
					ExpectedDeliveryDate *CustomTime `xml:"expectedDeliveryDate,omitempty" json:"expectedDeliveryDate,omitempty"`

					ResponseToArray struct {
						RespondTo []struct {

							//
							// The name of the person to respond to
							//

							Name string `xml:"name,omitempty" json:"name,omitempty"`

							//
							// A monitored email address that a recipient can send their response to
							//

							EmailAddress string `xml:"emailAddress,omitempty" json:"emailAddress,omitempty"`

							//
							// The phone number of the person to respond to
							//

							PhoneNumber string `xml:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
						} `xml:"RespondTo,omitempty" json:"RespondTo,omitempty"`
					} `xml:"ResponseToArray,omitempty" json:"ResponseToArray,omitempty"`

					//
					// This allows further detail about the status
					//

					AdditionalExplanation string `xml:"additionalExplanation,omitempty" json:"additionalExplanation,omitempty"`

					//
					// True or False to determine if it is necessary to provide information back to the supplier to complete the order
					//
					ResponseRequired bool `xml:"responseRequired,omitempty" json:"responseRequired,omitempty"`

					//
					// Time of order status
					//
					ValidTimestamp *CustomTime `xml:"validTimestamp,omitempty" json:"validTimestamp,omitempty"`
				} `xml:"OrderStatusDetail,omitempty" json:"OrderStatusDetail,omitempty"`
			} `xml:"OrderStatusDetailArray,omitempty" json:"OrderStatusDetailArray,omitempty"`
		} `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty"`
	} `xml:"OrderStatusArray,omitempty" json:"OrderStatusArray,omitempty"`

	//
	// Response for any error requiring notification to requestor
	//

	ErrorMessage string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`
}

type GetOrderStatusTypesRequest struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/OrderStatusService/1.0.0/ GetOrderStatusTypesRequest"`

	// The Standard Version of the Web Service being referenced

	WsVersion *string `xml:"wsVersion,omitempty" json:"wsVersion,omitempty"`

	// The customerID or any other agreed upon ID.

	Id *string `xml:"id,omitempty" json:"id,omitempty"`

	//
	// The password associated with the customerID.
	//

	Password *string `xml:"password,omitempty" json:"password,omitempty"`
}

type GetOrderStatusTypesResponse struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/OrderStatusService/1.0.0/ GetOrderStatusTypesResponse"`

	StatusArray struct {
		Status []struct {

			//
			// The numerical value of the order status
			//
			Id int32 `xml:"id,omitempty" json:"id,omitempty"`

			//
			// The string status name from possible values
			//

			Name string `xml:"name,omitempty" json:"name,omitempty"`
		} `xml:"Status,omitempty" json:"Status,omitempty"`
	} `xml:"StatusArray,omitempty" json:"StatusArray,omitempty"`

	//
	// Response for any error requiring notification to requestor
	//

	ErrorMessage string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`
}

type OrderStatusService interface {
	GetOrderStatusDetails(request *GetOrderStatusDetailsRequest) (*GetOrderStatusDetailsResponse, error)

	GetOrderStatusDetailsContext(ctx context.Context, request *GetOrderStatusDetailsRequest) (*GetOrderStatusDetailsResponse, error)

	GetOrderStatusTypes(request *GetOrderStatusTypesRequest) (*GetOrderStatusTypesResponse, error)

	GetOrderStatusTypesContext(ctx context.Context, request *GetOrderStatusTypesRequest) (*GetOrderStatusTypesResponse, error)
}

type orderStatusService struct {
	client *soap.Client
}

func NewOrderStatusService(client *soap.Client) OrderStatusService {
	return &orderStatusService{
		client: client,
	}
}

func (service *orderStatusService) GetOrderStatusDetailsContext(ctx context.Context, request *GetOrderStatusDetailsRequest) (*GetOrderStatusDetailsResponse, error) {
	response := new(GetOrderStatusDetailsResponse)
	// c, _ := xml.Marshal(request)
	// fmt.Println(string(c))
	err := service.client.CallContext(ctx, "getOrderStatusDetails", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *orderStatusService) GetOrderStatusDetails(request *GetOrderStatusDetailsRequest) (*GetOrderStatusDetailsResponse, error) {
	return service.GetOrderStatusDetailsContext(
		context.Background(),
		request,
	)
}

func (service *orderStatusService) GetOrderStatusTypesContext(ctx context.Context, request *GetOrderStatusTypesRequest) (*GetOrderStatusTypesResponse, error) {
	response := new(GetOrderStatusTypesResponse)
	err := service.client.CallContext(ctx, "getOrderStatusTypes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *orderStatusService) GetOrderStatusTypes(request *GetOrderStatusTypesRequest) (*GetOrderStatusTypesResponse, error) {
	return service.GetOrderStatusTypesContext(
		context.Background(),
		request,
	)
}
