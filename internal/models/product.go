package models

type Product struct {
	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ImageURL    string `protobuf:"bytes,2,opt,name=imageURL,proto3" json:"imageURL,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price       int64  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Currency    int32  `protobuf:"varint,5,opt,name=currency,proto3" json:"currency,omitempty"` // unmarshal in Unicode
	Discount    uint32 `protobuf:"varint,6,opt,name=discount,proto3" json:"discount,omitempty"`
	ProductURL  string `protobuf:"bytes,7,opt,name=productURL,proto3" json:"productURL,omitempty"`
}
