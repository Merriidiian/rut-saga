package pb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateReservationRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ReservationId     string                 `protobuf:"bytes,1,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
	ProductId         string                 `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	CustomerId        string                 `protobuf:"bytes,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Quantity          int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ExpirationMinutes int32                  `protobuf:"varint,5,opt,name=expiration_minutes,json=expirationMinutes,proto3" json:"expiration_minutes,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CreateReservationRequest) Reset() {
	*x = CreateReservationRequest{}
	mi := &file_reservation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *CreateReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*CreateReservationRequest) ProtoMessage() {
}
func (x *CreateReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*CreateReservationRequest) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{0}
}
func (x *CreateReservationRequest) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}
func (x *CreateReservationRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}
func (x *CreateReservationRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}
func (x *CreateReservationRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}
func (x *CreateReservationRequest) GetExpirationMinutes() int32 {
	if x != nil {
		return x.ExpirationMinutes
	}
	return 0
}

type ReservationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReservationId string                 `protobuf:"bytes,1,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReservationRequest) Reset() {
	*x = ReservationRequest{}
	mi := &file_reservation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *ReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*ReservationRequest) ProtoMessage() {
}
func (x *ReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*ReservationRequest) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{1}
}
func (x *ReservationRequest) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

type Reservation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReservationId string                 `protobuf:"bytes,1,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
	ProductId     string                 `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	CustomerId    string                 `protobuf:"bytes,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	mi := &file_reservation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*Reservation) ProtoMessage() {
}
func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{2}
}
func (x *Reservation) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}
func (x *Reservation) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}
func (x *Reservation) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}
func (x *Reservation) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}
func (x *Reservation) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_reservation_proto protoreflect.FileDescriptor

const file_reservation_proto_rawDesc = "" + "\n" + "\x11reservation.proto\x12\x0ereservation.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x17validate/validate.proto\"\xff\x01\n" + "\x18CreateReservationRequest\x12/\n" + "\x0ereservation_id\x18\x01 \x01(\tB\b\xfaB\x05r\x03\xb0\x01\x01R\rreservationId\x12'\n" + "\n" + "product_id\x18\x02 \x01(\tB\b\xfaB\x05r\x03\xb0\x01\x01R\tproductId\x12)\n" + "\vcustomer_id\x18\x03 \x01(\tB\b\xfaB\x05r\x03\xb0\x01\x01R\n" + "customerId\x12#\n" + "\bquantity\x18\x04 \x01(\x05B\a\xfaB\x04\x1a\x02 \x00R\bquantity\x129\n" + "\x12expiration_minutes\x18\x05 \x01(\x05B\n" + "\xfaB\a\x1a\x05\x18\xa0\v(\x01R\x11expirationMinutes\"E\n" + "\x12ReservationRequest\x12/\n" + "\x0ereservation_id\x18\x01 \x01(\tB\b\xfaB\x05r\x03\xb0\x01\x01R\rreservationId\"\xa8\x01\n" + "\vReservation\x12%\n" + "\x0ereservation_id\x18\x01 \x01(\tR\rreservationId\x12\x1d\n" + "\n" + "product_id\x18\x02 \x01(\tR\tproductId\x12\x1f\n" + "\vcustomer_id\x18\x03 \x01(\tR\n" + "customerId\x12\x1a\n" + "\bquantity\x18\x04 \x01(\x05R\bquantity\x12\x16\n" + "\x06status\x18\x05 \x01(\tR\x06status2\xb6\x04\n" + "\x12ReservationService\x12{\n" + "\x11CreateReservation\x12(.reservation.v1.CreateReservationRequest\x1a\x1b.reservation.v1.Reservation\"\x1f\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/api/v1/reservations\x12\x80\x01\n" + "\x0eGetReservation\x12\".reservation.v1.ReservationRequest\x1a\x1b.reservation.v1.Reservation\"-\x82\xd3\xe4\x93\x02'\x12%/api/v1/reservations/{reservation_id}\x12\x8f\x01\n" + "\x12ConfirmReservation\x12\".reservation.v1.ReservationRequest\x1a\x1b.reservation.v1.Reservation\"8\x82\xd3\xe4\x93\x022:\x01*\"-/api/v1/reservations/{reservation_id}/confirm\x12\x8d\x01\n" + "\x11CancelReservation\x12\".reservation.v1.ReservationRequest\x1a\x1b.reservation.v1.Reservation\"7\x82\xd3\xe4\x93\x021:\x01*\",/api/v1/reservations/{reservation_id}/cancelB1Z\x1ereservation-service/pkg/gen;pb\xaa\x02\x0eReservation.V1b\x06proto3"

var (
	file_reservation_proto_rawDescOnce sync.Once
	file_reservation_proto_rawDescData []byte
)

func file_reservation_proto_rawDescGZIP() []byte {
	file_reservation_proto_rawDescOnce.Do(func() {
		file_reservation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_reservation_proto_rawDesc), len(file_reservation_proto_rawDesc)))
	})
	return file_reservation_proto_rawDescData
}

var file_reservation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_reservation_proto_goTypes = []any{(*CreateReservationRequest)(nil), (*ReservationRequest)(nil), (*Reservation)(nil)}
var file_reservation_proto_depIdxs = []int32{0, 1, 1, 1, 2, 2, 2, 2, 4, 0, 0, 0, 0}

func init() {
	file_reservation_proto_init()
}
func file_reservation_proto_init() {
	if File_reservation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{File: protoimpl.DescBuilder{GoPackagePath: reflect.TypeOf(x{}).PkgPath(), RawDescriptor: unsafe.Slice(unsafe.StringData(file_reservation_proto_rawDesc), len(file_reservation_proto_rawDesc)), NumEnums: 0, NumMessages: 3, NumExtensions: 0, NumServices: 1}, GoTypes: file_reservation_proto_goTypes, DependencyIndexes: file_reservation_proto_depIdxs, MessageInfos: file_reservation_proto_msgTypes}.Build()
	File_reservation_proto = out.File
	file_reservation_proto_goTypes = nil
	file_reservation_proto_depIdxs = nil
}
