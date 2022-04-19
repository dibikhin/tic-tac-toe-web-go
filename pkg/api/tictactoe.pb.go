// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: pkg/api/tictactoe.proto

package api

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GameStatus int32

const (
	GameStatus_UNKNOWN            GameStatus = 0
	GameStatus_NOT_STARTED        GameStatus = 1
	GameStatus_WAITING_P2_JOIN    GameStatus = 2
	GameStatus_WAITING_P1_TO_TURN GameStatus = 3
	GameStatus_WAITING_P2_TO_TURN GameStatus = 4
	GameStatus_WON                GameStatus = 5
	GameStatus_DRAW               GameStatus = 6
	GameStatus_DELETED            GameStatus = 7
)

// Enum value maps for GameStatus.
var (
	GameStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "NOT_STARTED",
		2: "WAITING_P2_JOIN",
		3: "WAITING_P1_TO_TURN",
		4: "WAITING_P2_TO_TURN",
		5: "WON",
		6: "DRAW",
		7: "DELETED",
	}
	GameStatus_value = map[string]int32{
		"UNKNOWN":            0,
		"NOT_STARTED":        1,
		"WAITING_P2_JOIN":    2,
		"WAITING_P1_TO_TURN": 3,
		"WAITING_P2_TO_TURN": 4,
		"WON":                5,
		"DRAW":               6,
		"DELETED":            7,
	}
)

func (x GameStatus) Enum() *GameStatus {
	p := new(GameStatus)
	*p = x
	return p
}

func (x GameStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_api_tictactoe_proto_enumTypes[0].Descriptor()
}

func (GameStatus) Type() protoreflect.EnumType {
	return &file_pkg_api_tictactoe_proto_enumTypes[0]
}

func (x GameStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameStatus.Descriptor instead.
func (GameStatus) EnumDescriptor() ([]byte, []int) {
	return file_pkg_api_tictactoe_proto_rawDescGZIP(), []int{0}
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_tictactoe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_tictactoe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_tictactoe_proto_rawDescGZIP(), []int{0}
}

type GameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerName string `protobuf:"bytes,1,opt,name=playerName,proto3" json:"playerName,omitempty"`
}

func (x *GameRequest) Reset() {
	*x = GameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_tictactoe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameRequest) ProtoMessage() {}

func (x *GameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_tictactoe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameRequest.ProtoReflect.Descriptor instead.
func (*GameRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_tictactoe_proto_rawDescGZIP(), []int{1}
}

func (x *GameRequest) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

type GameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    GameStatus `protobuf:"varint,1,opt,name=status,proto3,enum=tictactoe.GameStatus" json:"status,omitempty"`
	Player1   *Player    `protobuf:"bytes,2,opt,name=player1,proto3" json:"player1,omitempty"`
	Player2   *Player    `protobuf:"bytes,3,opt,name=player2,proto3" json:"player2,omitempty"`
	PlayerWon *Player    `protobuf:"bytes,4,opt,name=playerWon,proto3" json:"playerWon,omitempty"`
	Board     string     `protobuf:"bytes,5,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *GameResponse) Reset() {
	*x = GameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_tictactoe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResponse) ProtoMessage() {}

func (x *GameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_tictactoe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameResponse.ProtoReflect.Descriptor instead.
func (*GameResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_tictactoe_proto_rawDescGZIP(), []int{2}
}

func (x *GameResponse) GetStatus() GameStatus {
	if x != nil {
		return x.Status
	}
	return GameStatus_UNKNOWN
}

func (x *GameResponse) GetPlayer1() *Player {
	if x != nil {
		return x.Player1
	}
	return nil
}

func (x *GameResponse) GetPlayer2() *Player {
	if x != nil {
		return x.Player2
	}
	return nil
}

func (x *GameResponse) GetPlayerWon() *Player {
	if x != nil {
		return x.PlayerWon
	}
	return nil
}

func (x *GameResponse) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

type TurnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerName string `protobuf:"bytes,1,opt,name=playerName,proto3" json:"playerName,omitempty"`
	Turn       string `protobuf:"bytes,2,opt,name=turn,proto3" json:"turn,omitempty"`
}

func (x *TurnRequest) Reset() {
	*x = TurnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_tictactoe_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TurnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TurnRequest) ProtoMessage() {}

func (x *TurnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_tictactoe_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TurnRequest.ProtoReflect.Descriptor instead.
func (*TurnRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_tictactoe_proto_rawDescGZIP(), []int{3}
}

func (x *TurnRequest) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *TurnRequest) GetTurn() string {
	if x != nil {
		return x.Turn
	}
	return ""
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mark string `protobuf:"bytes,1,opt,name=mark,proto3" json:"mark,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_tictactoe_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_tictactoe_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_pkg_api_tictactoe_proto_rawDescGZIP(), []int{4}
}

func (x *Player) GetMark() string {
	if x != nil {
		return x.Mark
	}
	return ""
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_pkg_api_tictactoe_proto protoreflect.FileDescriptor

var file_pkg_api_tictactoe_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63,
	0x74, 0x6f, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x69, 0x63, 0x74, 0x61,
	0x63, 0x74, 0x6f, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xde, 0x01, 0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f,
	0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f,
	0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x31, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x12, 0x2f,
	0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x57, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x57, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x41, 0x0a, 0x0b, 0x54, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x30, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x8f, 0x01, 0x0a, 0x0a, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x41, 0x49, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x50, 0x32, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12,
	0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x31, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x55,
	0x52, 0x4e, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x50, 0x32, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x55, 0x52, 0x4e, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03,
	0x57, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x52, 0x41, 0x57, 0x10, 0x06, 0x12,
	0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x07, 0x32, 0xbb, 0x01, 0x0a,
	0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61,
	0x63, 0x74, 0x6f, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74,
	0x6f, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x04, 0x54, 0x75, 0x72, 0x6e, 0x12, 0x16, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61,
	0x63, 0x74, 0x6f, 0x65, 0x2e, 0x54, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x62, 0x69, 0x6b, 0x68, 0x69,
	0x6e, 0x2f, 0x74, 0x69, 0x63, 0x2d, 0x74, 0x61, 0x63, 0x2d, 0x74, 0x6f, 0x65, 0x2d, 0x77, 0x65,
	0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_tictactoe_proto_rawDescOnce sync.Once
	file_pkg_api_tictactoe_proto_rawDescData = file_pkg_api_tictactoe_proto_rawDesc
)

func file_pkg_api_tictactoe_proto_rawDescGZIP() []byte {
	file_pkg_api_tictactoe_proto_rawDescOnce.Do(func() {
		file_pkg_api_tictactoe_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_tictactoe_proto_rawDescData)
	})
	return file_pkg_api_tictactoe_proto_rawDescData
}

var file_pkg_api_tictactoe_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_api_tictactoe_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_api_tictactoe_proto_goTypes = []interface{}{
	(GameStatus)(0),       // 0: tictactoe.GameStatus
	(*EmptyResponse)(nil), // 1: tictactoe.EmptyResponse
	(*GameRequest)(nil),   // 2: tictactoe.GameRequest
	(*GameResponse)(nil),  // 3: tictactoe.GameResponse
	(*TurnRequest)(nil),   // 4: tictactoe.TurnRequest
	(*Player)(nil),        // 5: tictactoe.Player
}
var file_pkg_api_tictactoe_proto_depIdxs = []int32{
	0, // 0: tictactoe.GameResponse.status:type_name -> tictactoe.GameStatus
	5, // 1: tictactoe.GameResponse.player1:type_name -> tictactoe.Player
	5, // 2: tictactoe.GameResponse.player2:type_name -> tictactoe.Player
	5, // 3: tictactoe.GameResponse.playerWon:type_name -> tictactoe.Player
	2, // 4: tictactoe.Game.GetGame:input_type -> tictactoe.GameRequest
	2, // 5: tictactoe.Game.StartGame:input_type -> tictactoe.GameRequest
	4, // 6: tictactoe.Game.Turn:input_type -> tictactoe.TurnRequest
	3, // 7: tictactoe.Game.GetGame:output_type -> tictactoe.GameResponse
	1, // 8: tictactoe.Game.StartGame:output_type -> tictactoe.EmptyResponse
	1, // 9: tictactoe.Game.Turn:output_type -> tictactoe.EmptyResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_api_tictactoe_proto_init() }
func file_pkg_api_tictactoe_proto_init() {
	if File_pkg_api_tictactoe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_tictactoe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_tictactoe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_tictactoe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_tictactoe_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TurnRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_tictactoe_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_api_tictactoe_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_tictactoe_proto_goTypes,
		DependencyIndexes: file_pkg_api_tictactoe_proto_depIdxs,
		EnumInfos:         file_pkg_api_tictactoe_proto_enumTypes,
		MessageInfos:      file_pkg_api_tictactoe_proto_msgTypes,
	}.Build()
	File_pkg_api_tictactoe_proto = out.File
	file_pkg_api_tictactoe_proto_rawDesc = nil
	file_pkg_api_tictactoe_proto_goTypes = nil
	file_pkg_api_tictactoe_proto_depIdxs = nil
}
