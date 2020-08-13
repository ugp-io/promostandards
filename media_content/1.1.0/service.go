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

type MediaTypeType string

const (
	MediaTypeTypeImage MediaTypeType = "Image"

	MediaTypeTypeVideo MediaTypeType = "Video"

	MediaTypeTypeAudio MediaTypeType = "Audio"

	MediaTypeTypeDocument MediaTypeType = "Document"
)

type SourceTypeType string

const (
	SourceTypeTypeWeb SourceTypeType = "Web"

	SourceTypeTypeFtp SourceTypeType = "Ftp"
)

type CultureName string

type ErrorMessage struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/MediaService/1.0.0/SharedObjects/ errorMessage"  bson:"-"`

	//
	// Response for any error requiring notification to requestor
	//
	Code int32 `xml:"code,omitempty" json:"code,omitempty"`

	//
	// Response for any error requiring notification to requestor
	//

	Description string `xml:"description,omitempty" json:"description,omitempty"`
}

type ChangeTimeStamp time.Time

type MediaType MediaTypeType

type GetMediaContentRequest struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/MediaService/1.0.0/ GetMediaContentRequest"  bson:"-"`

	WsVersion *string `xml:"wsVersion,omitempty" json:"wsVersion,omitempty"`

	Id *string `xml:"id,omitempty" json:"id,omitempty"`

	Password *string `xml:"password,omitempty" json:"password,omitempty"`

	CultureName *string `xml:"cultureName,omitempty" json:"cultureName,omitempty"`

	MediaType *string `xml:"mediaType,omitempty" json:"mediaType,omitempty"`

	ProductId *string `xml:"productId,omitempty" json:"productId,omitempty"`

	PartId *string `xml:"partId,omitempty" json:"partId,omitempty"`

	// The part ID associated to the product ID.
	//
	ClassType int32 `xml:"classType,omitempty" json:"classType,omitempty"`
}

type MediaContent struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/MediaService/1.0.0/ MediaContent"  bson:"-"`

	ProductId *string `xml:"productId,omitempty" json:"productId,omitempty" bson:"product_id,omitempty"`

	PartId *string `xml:"partId,omitempty" json:"partId,omitempty" bson:"part_id,omitempty"`

	// The URL of the media location.  Any valid URL can be returned including prefixes like http and ftp.

	Url string `xml:"url,omitempty" json:"url,omitempty" bson:"url,omitempty"`

	MediaType *string `xml:"mediaType,omitempty" json:"mediaType,omitempty" bson:"media_yype,omitempty"`

	ClassTypeArray struct {
		ClassType []*ClassType `xml:"ClassType,omitempty" json:"ClassType,omitempty" bson:"class_yype,omitempty"`
	} `xml:"ClassTypeArray,omitempty" json:"ClassTypeArray,omitempty" bson:"class_type_array,omitempty"`

	// The file size
	FileSize float64 `xml:"fileSize,omitempty" json:"fileSize,omitempty" bson:"file_size,omitempty"`

	// Width
	Width float64 `xml:"width,omitempty" json:"width,omitempty" bson:"width,omitempty"`

	// Height
	Height float64 `xml:"height,omitempty" json:"height,omitempty" bson:"height,omitempty"`

	// Dots per inch
	Dpi int32 `xml:"dpi,omitempty" json:"dpi,omitempty" bson:"dpi,omitempty"`

	// The color description

	Color string `xml:"color,omitempty" json:"color,omitempty" bson:"color,omitempty"`

	DecorationArray struct {
		Decoration []*Decoration `xml:"Decoration,omitempty" json:"Decoration,omitempty" bson:"decoration,omitempty"`
	} `xml:"DecorationArray,omitempty" json:"DecorationArray,omitempty" bson:"decoration_array,omitempty"`

	LocationArray struct {
		Location []*Location `xml:"Location,omitempty" json:"Location,omitempty" bson:"location,omitempty"`
	} `xml:"LocationArray,omitempty" json:"LocationArray,omitempty" bson:"location_array,omitempty"`

	// Deprecated.  Use DecorationArray.
	DecorationId int32 `xml:"decorationId,omitempty" json:"decorationId,omitempty" bson:"decoration_id,omitempty"`

	// Information about the media

	Description string `xml:"description,omitempty" json:"description,omitempty" bson:"description,omitempty"`

	// Identifies whether the partId one to one corresponds with the image
	SinglePart bool `xml:"singlePart,omitempty" json:"singlePart,omitempty" bson:"single_part,omitempty"`

	ChangeTimeStamp *ChangeTimeStamp `xml:"changeTimeStamp,omitempty" json:"changeTimeStamp,omitempty" bson:"change_time_stamp,omitempty"`
}

type ClassType struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/MediaService/1.0.0/ ClassType"  bson:"-"`

	// The classification of the media
	ClassTypeId int32 `xml:"classTypeId,omitempty" json:"classTypeId,omitempty" bson:"class_type_id,omitempty"`

	// The classification short name

	ClassTypeName string `xml:"classTypeName,omitempty" json:"classTypeName,omitempty" bson:"class_type_name,omitempty"`
}

type Decoration struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/MediaService/1.0.0/ Decoration"  bson:"-"`

	// The decoration id associated with the media
	DecorationId int32 `xml:"decorationId,omitempty" json:"decorationId,omitempty" bson:"decoration_id,omitempty"`

	// The name of the decoration associated with the id

	DecorationName string `xml:"decorationName,omitempty" json:"decorationName,omitempty" bson:"decoration_name,omitempty"`
}

type Location struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/MediaService/1.0.0/ Location"`

	// The location id associated with the media
	LocationId int32 `xml:"locationId,omitempty" json:"locationId,omitempty" bson:"location_id,omitempty"`

	// The name of the location associated with the id

	LocationName string `xml:"locationName,omitempty" json:"locationName,omitempty" bson:"location_name,omitempty"`
}

type GetMediaContentResponse struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/MediaService/1.0.0/ GetMediaContentResponse"  bson:"-"`

	MediaContentArray struct {
		MediaContent []*MediaContent `xml:"MediaContent,omitempty" json:"MediaContent,omitempty" bson:"media_content,omitempty"`
	} `xml:"MediaContentArray,omitempty" json:"MediaContentArray,omitempty" bson:"media_content_array,omitempty"`

	ErrorMessage *ErrorMessage `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`
}

type GetMediaDateModifiedRequest struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/MediaService/1.0.0/ GetMediaDateModifiedRequest" bson:"-"`

	WsVersion *string `xml:"wsVersion,omitempty" json:"wsVersion,omitempty"`

	Id *string `xml:"id,omitempty" json:"id,omitempty"`

	Password *string `xml:"password,omitempty" json:"password,omitempty"`

	CultureName *string `xml:"cultureName,omitempty" json:"cultureName,omitempty"`

	ChangeTimeStamp *ChangeTimeStamp `xml:"changeTimeStamp,omitempty" json:"changeTimeStamp,omitempty"`
}

type MediaDateModified struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/MediaService/1.0.0/ MediaDateModified" bson:"-"`

	ProductId *string `xml:"productId,omitempty" json:"productId,omitempty"`

	PartId *string `xml:"partId,omitempty" json:"partId,omitempty"`
}

type GetMediaDateModifiedResponse struct {
	XMLName xml.Name `xml:"http://www.promostandards.org/WSDL/MediaService/1.0.0/ GetMediaDateModifiedResponse" bson:"-"`

	MediaDateModifiedArray struct {
		MediaDateModified []*MediaDateModified `xml:"MediaDateModified,omitempty" json:"MediaDateModified,omitempty"`
	} `xml:"MediaDateModifiedArray,omitempty" json:"MediaDateModifiedArray,omitempty"`

	ErrorMessage *ErrorMessage `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`
}

type MediaContentService interface {
	GetMediaContent(request *GetMediaContentRequest) (*GetMediaContentResponse, error)

	GetMediaContentContext(ctx context.Context, request *GetMediaContentRequest) (*GetMediaContentResponse, error)

	GetMediaDateModified(request *GetMediaDateModifiedRequest) (*GetMediaDateModifiedResponse, error)

	GetMediaDateModifiedContext(ctx context.Context, request *GetMediaDateModifiedRequest) (*GetMediaDateModifiedResponse, error)
}

type mediaContentService struct {
	client *soap.Client
}

func NewMediaContentService(client *soap.Client) MediaContentService {
	return &mediaContentService{
		client: client,
	}
}

func (service *mediaContentService) GetMediaContentContext(ctx context.Context, request *GetMediaContentRequest) (*GetMediaContentResponse, error) {
	response := new(GetMediaContentResponse)
	err := service.client.CallContext(ctx, "getMediaContent", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *mediaContentService) GetMediaContent(request *GetMediaContentRequest) (*GetMediaContentResponse, error) {
	return service.GetMediaContentContext(
		context.Background(),
		request,
	)
}

func (service *mediaContentService) GetMediaDateModifiedContext(ctx context.Context, request *GetMediaDateModifiedRequest) (*GetMediaDateModifiedResponse, error) {
	response := new(GetMediaDateModifiedResponse)
	err := service.client.CallContext(ctx, "urn:getMediaDateModified", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *mediaContentService) GetMediaDateModified(request *GetMediaDateModifiedRequest) (*GetMediaDateModifiedResponse, error) {
	return service.GetMediaDateModifiedContext(
		context.Background(),
		request,
	)
}
