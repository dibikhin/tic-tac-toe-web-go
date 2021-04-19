// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: tttwebgrpc/tictactoe.proto

package tttwebgrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type State int32

const (
	State_UNDEFINED State = 0 // safe default
	State_WAITING   State = 1
	State_IDLE      State = 2
	State_GAME_OVER State = 3
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "UNDEFINED",
		1: "WAITING",
		2: "IDLE",
		3: "GAME_OVER",
	}
	State_value = map[string]int32{
		"UNDEFINED": 0,
		"WAITING":   1,
		"IDLE":      2,
		"GAME_OVER": 3,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_tttwebgrpc_tictactoe_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_tttwebgrpc_tictactoe_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_tttwebgrpc_tictactoe_proto_rawDescGZIP(), []int{0}
}

type For int32

const (
	For_NOTHING For = 0 // safe default
	For_AUTH    For = 1
	For_MARK    For = 2
	For_TURN    For = 3
)

// Enum value maps for For.
var (
	For_name = map[int32]string{
		0: "NOTHING",
		1: "AUTH",
		2: "MARK",
		3: "TURN",
	}
	For_value = map[string]int32{
		"NOTHING": 0,
		"AUTH":    1,
		"MARK":    2,
		"TURN":    3,
	}
)

func (x For) Enum() *For {
	p := new(For)
	*p = x
	return p
}

func (x For) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (For) Descriptor() protoreflect.EnumDescriptor {
	return file_tttwebgrpc_tictactoe_proto_enumTypes[1].Descriptor()
}

func (For) Type() protoreflect.EnumType {
	return &file_tttwebgrpc_tictactoe_proto_enumTypes[1]
}

func (x For) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use For.Descriptor instead.
func (For) EnumDescriptor() ([]byte, []int) {
	return file_tttwebgrpc_tictactoe_proto_rawDescGZIP(), []int{1}
}

type Actions int32

const (
	Actions_NOOP       Actions = 0 // safe default
	Actions_GET_STATUS Actions = 1
	Actions_START_GAME Actions = 2
	Actions_LOG_IN     Actions = 3
	Actions_SET_MARK   Actions = 4
	Actions_DO_TURN    Actions = 5
	Actions_STOP_GAME  Actions = 6
)

// Enum value maps for Actions.
var (
	Actions_name = map[int32]string{
		0: "NOOP",
		1: "GET_STATUS",
		2: "START_GAME",
		3: "LOG_IN",
		4: "SET_MARK",
		5: "DO_TURN",
		6: "STOP_GAME",
	}
	Actions_value = map[string]int32{
		"NOOP":       0,
		"GET_STATUS": 1,
		"START_GAME": 2,
		"LOG_IN":     3,
		"SET_MARK":   4,
		"DO_TURN":    5,
		"STOP_GAME":  6,
	}
)

func (x Actions) Enum() *Actions {
	p := new(Actions)
	*p = x
	return p
}

func (x Actions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Actions) Descriptor() protoreflect.EnumDescriptor {
	return file_tttwebgrpc_tictactoe_proto_enumTypes[2].Descriptor()
}

func (Actions) Type() protoreflect.EnumType {
	return &file_tttwebgrpc_tictactoe_proto_enumTypes[2]
}

func (x Actions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Actions.Descriptor instead.
func (Actions) EnumDescriptor() ([]byte, []int) {
	return file_tttwebgrpc_tictactoe_proto_rawDescGZIP(), []int{2}
}

type Opts int32

const (
	Opts_ZERO   Opts = 0 // safe default
	Opts_ONE    Opts = 1
	Opts_TWO    Opts = 2
	Opts_THREE  Opts = 3
	Opts_FOUR   Opts = 4
	Opts_FIVE   Opts = 5
	Opts_SIX    Opts = 6
	Opts_SEVEN  Opts = 7
	Opts_EIGHT  Opts = 8
	Opts_NINE   Opts = 9
	Opts_O_MARK Opts = 10
	Opts_X_MARK Opts = 11
)

// Enum value maps for Opts.
var (
	Opts_name = map[int32]string{
		0:  "ZERO",
		1:  "ONE",
		2:  "TWO",
		3:  "THREE",
		4:  "FOUR",
		5:  "FIVE",
		6:  "SIX",
		7:  "SEVEN",
		8:  "EIGHT",
		9:  "NINE",
		10: "O_MARK",
		11: "X_MARK",
	}
	Opts_value = map[string]int32{
		"ZERO":   0,
		"ONE":    1,
		"TWO":    2,
		"THREE":  3,
		"FOUR":   4,
		"FIVE":   5,
		"SIX":    6,
		"SEVEN":  7,
		"EIGHT":  8,
		"NINE":   9,
		"O_MARK": 10,
		"X_MARK": 11,
	}
)

func (x Opts) Enum() *Opts {
	p := new(Opts)
	*p = x
	return p
}

func (x Opts) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Opts) Descriptor() protoreflect.EnumDescriptor {
	return file_tttwebgrpc_tictactoe_proto_enumTypes[3].Descriptor()
}

func (Opts) Type() protoreflect.EnumType {
	return &file_tttwebgrpc_tictactoe_proto_enumTypes[3]
}

func (x Opts) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Opts.Descriptor instead.
func (Opts) EnumDescriptor() ([]byte, []int) {
	return file_tttwebgrpc_tictactoe_proto_rawDescGZIP(), []int{3}
}

type Outcome int32

const (
	Outcome_DEFAULT Outcome = 0 // safe default
	Outcome_WON     Outcome = 1
	Outcome_DRAW    Outcome = 2
)

// Enum value maps for Outcome.
var (
	Outcome_name = map[int32]string{
		0: "DEFAULT",
		1: "WON",
		2: "DRAW",
	}
	Outcome_value = map[string]int32{
		"DEFAULT": 0,
		"WON":     1,
		"DRAW":    2,
	}
)

func (x Outcome) Enum() *Outcome {
	p := new(Outcome)
	*p = x
	return p
}

func (x Outcome) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Outcome) Descriptor() protoreflect.EnumDescriptor {
	return file_tttwebgrpc_tictactoe_proto_enumTypes[4].Descriptor()
}

func (Outcome) Type() protoreflect.EnumType {
	return &file_tttwebgrpc_tictactoe_proto_enumTypes[4]
}

func (x Outcome) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Outcome.Descriptor instead.
func (Outcome) EnumDescriptor() ([]byte, []int) {
	return file_tttwebgrpc_tictactoe_proto_rawDescGZIP(), []int{4}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tttwebgrpc_tictactoe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_tttwebgrpc_tictactoe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_tttwebgrpc_tictactoe_proto_rawDescGZIP(), []int{0}
}

type CommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action Actions `protobuf:"varint,1,opt,name=action,proto3,enum=tictactoe.Actions" json:"action,omitempty"`
}

func (x *CommandRequest) Reset() {
	*x = CommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tttwebgrpc_tictactoe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRequest) ProtoMessage() {}

func (x *CommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tttwebgrpc_tictactoe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandRequest.ProtoReflect.Descriptor instead.
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return file_tttwebgrpc_tictactoe_proto_rawDescGZIP(), []int{1}
}

func (x *CommandRequest) GetAction() Actions {
	if x != nil {
		return x.Action
	}
	return Actions_NOOP
}

//
//
//TODO: Fill example
//
type StatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State   State     `protobuf:"varint,1,opt,name=state,proto3,enum=tictactoe.State" json:"state,omitempty"`
	Actions []Actions `protobuf:"varint,2,rep,packed,name=actions,proto3,enum=tictactoe.Actions" json:"actions,omitempty"`
	Opts    []Opts    `protobuf:"varint,3,rep,packed,name=opts,proto3,enum=tictactoe.Opts" json:"opts,omitempty"`
	Message string    `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	For     For       `protobuf:"varint,5,opt,name=for,proto3,enum=tictactoe.For" json:"for,omitempty"` // optional
	// Player player = 6; // optional
	Board   string  `protobuf:"bytes,7,opt,name=board,proto3" json:"board,omitempty"`                             // optional
	Outcome Outcome `protobuf:"varint,8,opt,name=outcome,proto3,enum=tictactoe.Outcome" json:"outcome,omitempty"` // optional
}

func (x *StatusReply) Reset() {
	*x = StatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tttwebgrpc_tictactoe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReply) ProtoMessage() {}

func (x *StatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_tttwebgrpc_tictactoe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReply.ProtoReflect.Descriptor instead.
func (*StatusReply) Descriptor() ([]byte, []int) {
	return file_tttwebgrpc_tictactoe_proto_rawDescGZIP(), []int{2}
}

func (x *StatusReply) GetState() State {
	if x != nil {
		return x.State
	}
	return State_UNDEFINED
}

func (x *StatusReply) GetActions() []Actions {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *StatusReply) GetOpts() []Opts {
	if x != nil {
		return x.Opts
	}
	return nil
}

func (x *StatusReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StatusReply) GetFor() For {
	if x != nil {
		return x.For
	}
	return For_NOTHING
}

func (x *StatusReply) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

func (x *StatusReply) GetOutcome() Outcome {
	if x != nil {
		return x.Outcome
	}
	return Outcome_DEFAULT
}

var File_tttwebgrpc_tictactoe_proto protoreflect.FileDescriptor

var file_tttwebgrpc_tictactoe_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x74, 0x74, 0x77, 0x65, 0x62, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x69, 0x63,
	0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x69,
	0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x3c, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8e,
	0x02, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63,
	0x74, 0x6f, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x4f,
	0x70, 0x74, 0x73, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x66, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x46, 0x6f, 0x72,
	0x52, 0x03, 0x66, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x2c, 0x0a, 0x07, 0x6f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74,
	0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65,
	0x52, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x10, 0x2a,
	0x3c, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45,
	0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x44, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x10, 0x03, 0x2a, 0x30, 0x0a,
	0x03, 0x46, 0x6f, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x54, 0x48, 0x49, 0x4e, 0x47, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x55, 0x54, 0x48, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4d,
	0x41, 0x52, 0x4b, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x55, 0x52, 0x4e, 0x10, 0x03, 0x2a,
	0x69, 0x0a, 0x07, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4f, 0x50, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x47, 0x41,
	0x4d, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x47, 0x5f, 0x49, 0x4e, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x45, 0x54, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x4f, 0x5f, 0x54, 0x55, 0x52, 0x4e, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x53,
	0x54, 0x4f, 0x50, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x10, 0x06, 0x2a, 0x82, 0x01, 0x0a, 0x04, 0x4f,
	0x70, 0x74, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x54, 0x48, 0x52, 0x45, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x4f,
	0x55, 0x52, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x56, 0x45, 0x10, 0x05, 0x12, 0x07,
	0x0a, 0x03, 0x53, 0x49, 0x58, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x45, 0x56, 0x45, 0x4e,
	0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x49, 0x47, 0x48, 0x54, 0x10, 0x08, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x49, 0x4e, 0x45, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x5f, 0x4d, 0x41, 0x52,
	0x4b, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x58, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x10, 0x0b, 0x2a,
	0x29, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x4f, 0x4e, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x52, 0x41, 0x57, 0x10, 0x02, 0x32, 0x79, 0x0a, 0x04, 0x47, 0x61,
	0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x03, 0x52,
	0x75, 0x6e, 0x12, 0x19, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x62, 0x69, 0x6b, 0x68, 0x69, 0x6e, 0x2f, 0x74, 0x69, 0x63,
	0x2d, 0x74, 0x61, 0x63, 0x2d, 0x74, 0x6f, 0x65, 0x2d, 0x77, 0x65, 0x62, 0x2d, 0x67, 0x6f, 0x2f,
	0x74, 0x74, 0x74, 0x77, 0x65, 0x62, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tttwebgrpc_tictactoe_proto_rawDescOnce sync.Once
	file_tttwebgrpc_tictactoe_proto_rawDescData = file_tttwebgrpc_tictactoe_proto_rawDesc
)

func file_tttwebgrpc_tictactoe_proto_rawDescGZIP() []byte {
	file_tttwebgrpc_tictactoe_proto_rawDescOnce.Do(func() {
		file_tttwebgrpc_tictactoe_proto_rawDescData = protoimpl.X.CompressGZIP(file_tttwebgrpc_tictactoe_proto_rawDescData)
	})
	return file_tttwebgrpc_tictactoe_proto_rawDescData
}

var file_tttwebgrpc_tictactoe_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_tttwebgrpc_tictactoe_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tttwebgrpc_tictactoe_proto_goTypes = []interface{}{
	(State)(0),             // 0: tictactoe.State
	(For)(0),               // 1: tictactoe.For
	(Actions)(0),           // 2: tictactoe.Actions
	(Opts)(0),              // 3: tictactoe.Opts
	(Outcome)(0),           // 4: tictactoe.Outcome
	(*Empty)(nil),          // 5: tictactoe.Empty
	(*CommandRequest)(nil), // 6: tictactoe.CommandRequest
	(*StatusReply)(nil),    // 7: tictactoe.StatusReply
}
var file_tttwebgrpc_tictactoe_proto_depIdxs = []int32{
	2, // 0: tictactoe.CommandRequest.action:type_name -> tictactoe.Actions
	0, // 1: tictactoe.StatusReply.state:type_name -> tictactoe.State
	2, // 2: tictactoe.StatusReply.actions:type_name -> tictactoe.Actions
	3, // 3: tictactoe.StatusReply.opts:type_name -> tictactoe.Opts
	1, // 4: tictactoe.StatusReply.for:type_name -> tictactoe.For
	4, // 5: tictactoe.StatusReply.outcome:type_name -> tictactoe.Outcome
	5, // 6: tictactoe.Game.GetStatus:input_type -> tictactoe.Empty
	6, // 7: tictactoe.Game.Run:input_type -> tictactoe.CommandRequest
	7, // 8: tictactoe.Game.GetStatus:output_type -> tictactoe.StatusReply
	7, // 9: tictactoe.Game.Run:output_type -> tictactoe.StatusReply
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tttwebgrpc_tictactoe_proto_init() }
func file_tttwebgrpc_tictactoe_proto_init() {
	if File_tttwebgrpc_tictactoe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tttwebgrpc_tictactoe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_tttwebgrpc_tictactoe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandRequest); i {
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
		file_tttwebgrpc_tictactoe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReply); i {
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
			RawDescriptor: file_tttwebgrpc_tictactoe_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tttwebgrpc_tictactoe_proto_goTypes,
		DependencyIndexes: file_tttwebgrpc_tictactoe_proto_depIdxs,
		EnumInfos:         file_tttwebgrpc_tictactoe_proto_enumTypes,
		MessageInfos:      file_tttwebgrpc_tictactoe_proto_msgTypes,
	}.Build()
	File_tttwebgrpc_tictactoe_proto = out.File
	file_tttwebgrpc_tictactoe_proto_rawDesc = nil
	file_tttwebgrpc_tictactoe_proto_goTypes = nil
	file_tttwebgrpc_tictactoe_proto_depIdxs = nil
}
