// Code generated by protoc-gen-go. DO NOT EDIT.
// source: channel_server.proto

package proto_server

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LoginRegExtendInfo struct {
	Imei                 string   `protobuf:"bytes,1,opt,name=Imei,proto3" json:"Imei,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=Ip,proto3" json:"Ip,omitempty"`
	PhoneModel           string   `protobuf:"bytes,3,opt,name=PhoneModel,proto3" json:"PhoneModel,omitempty"`
	IdCardNo             string   `protobuf:"bytes,4,opt,name=IdCardNo,proto3" json:"IdCardNo,omitempty"`
	RealName             string   `protobuf:"bytes,5,opt,name=RealName,proto3" json:"RealName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRegExtendInfo) Reset()         { *m = LoginRegExtendInfo{} }
func (m *LoginRegExtendInfo) String() string { return proto.CompactTextString(m) }
func (*LoginRegExtendInfo) ProtoMessage()    {}
func (*LoginRegExtendInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{0}
}

func (m *LoginRegExtendInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRegExtendInfo.Unmarshal(m, b)
}
func (m *LoginRegExtendInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRegExtendInfo.Marshal(b, m, deterministic)
}
func (m *LoginRegExtendInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRegExtendInfo.Merge(m, src)
}
func (m *LoginRegExtendInfo) XXX_Size() int {
	return xxx_messageInfo_LoginRegExtendInfo.Size(m)
}
func (m *LoginRegExtendInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRegExtendInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRegExtendInfo proto.InternalMessageInfo

func (m *LoginRegExtendInfo) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *LoginRegExtendInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *LoginRegExtendInfo) GetPhoneModel() string {
	if m != nil {
		return m.PhoneModel
	}
	return ""
}

func (m *LoginRegExtendInfo) GetIdCardNo() string {
	if m != nil {
		return m.IdCardNo
	}
	return ""
}

func (m *LoginRegExtendInfo) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

//账号注册 Request
type AccountRegOrLoginReq struct {
	IsReg                bool                `protobuf:"varint,1,opt,name=IsReg,proto3" json:"IsReg,omitempty"`
	AccountName          string              `protobuf:"bytes,2,opt,name=AccountName,proto3" json:"AccountName,omitempty"`
	PassWord             string              `protobuf:"bytes,3,opt,name=PassWord,proto3" json:"PassWord,omitempty"`
	Ext                  *LoginRegExtendInfo `protobuf:"bytes,4,opt,name=Ext,proto3" json:"Ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AccountRegOrLoginReq) Reset()         { *m = AccountRegOrLoginReq{} }
func (m *AccountRegOrLoginReq) String() string { return proto.CompactTextString(m) }
func (*AccountRegOrLoginReq) ProtoMessage()    {}
func (*AccountRegOrLoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{1}
}

func (m *AccountRegOrLoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRegOrLoginReq.Unmarshal(m, b)
}
func (m *AccountRegOrLoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRegOrLoginReq.Marshal(b, m, deterministic)
}
func (m *AccountRegOrLoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRegOrLoginReq.Merge(m, src)
}
func (m *AccountRegOrLoginReq) XXX_Size() int {
	return xxx_messageInfo_AccountRegOrLoginReq.Size(m)
}
func (m *AccountRegOrLoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRegOrLoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRegOrLoginReq proto.InternalMessageInfo

func (m *AccountRegOrLoginReq) GetIsReg() bool {
	if m != nil {
		return m.IsReg
	}
	return false
}

func (m *AccountRegOrLoginReq) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

func (m *AccountRegOrLoginReq) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func (m *AccountRegOrLoginReq) GetExt() *LoginRegExtendInfo {
	if m != nil {
		return m.Ext
	}
	return nil
}

//公共请求头部  用于校验登录态
type OpenPre struct {
	OpenId               int32    `protobuf:"varint,1,opt,name=OpenId,proto3" json:"OpenId,omitempty"`
	OpenToken            string   `protobuf:"bytes,2,opt,name=OpenToken,proto3" json:"OpenToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenPre) Reset()         { *m = OpenPre{} }
func (m *OpenPre) String() string { return proto.CompactTextString(m) }
func (*OpenPre) ProtoMessage()    {}
func (*OpenPre) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{2}
}

func (m *OpenPre) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenPre.Unmarshal(m, b)
}
func (m *OpenPre) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenPre.Marshal(b, m, deterministic)
}
func (m *OpenPre) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenPre.Merge(m, src)
}
func (m *OpenPre) XXX_Size() int {
	return xxx_messageInfo_OpenPre.Size(m)
}
func (m *OpenPre) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenPre.DiscardUnknown(m)
}

var xxx_messageInfo_OpenPre proto.InternalMessageInfo

func (m *OpenPre) GetOpenId() int32 {
	if m != nil {
		return m.OpenId
	}
	return 0
}

func (m *OpenPre) GetOpenToken() string {
	if m != nil {
		return m.OpenToken
	}
	return ""
}

//绑定手机请求
type BindPhoneReq struct {
	Head                 *OpenPre `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	PhoneNum             string   `protobuf:"bytes,2,opt,name=PhoneNum,proto3" json:"PhoneNum,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindPhoneReq) Reset()         { *m = BindPhoneReq{} }
func (m *BindPhoneReq) String() string { return proto.CompactTextString(m) }
func (*BindPhoneReq) ProtoMessage()    {}
func (*BindPhoneReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{3}
}

func (m *BindPhoneReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindPhoneReq.Unmarshal(m, b)
}
func (m *BindPhoneReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindPhoneReq.Marshal(b, m, deterministic)
}
func (m *BindPhoneReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindPhoneReq.Merge(m, src)
}
func (m *BindPhoneReq) XXX_Size() int {
	return xxx_messageInfo_BindPhoneReq.Size(m)
}
func (m *BindPhoneReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BindPhoneReq.DiscardUnknown(m)
}

var xxx_messageInfo_BindPhoneReq proto.InternalMessageInfo

func (m *BindPhoneReq) GetHead() *OpenPre {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *BindPhoneReq) GetPhoneNum() string {
	if m != nil {
		return m.PhoneNum
	}
	return ""
}

func (m *BindPhoneReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

//操作类请求 回复状态码和错误信息 状态码为1 表示成功 其他均为失败 错误信息将描述失败原因
type OperationRep struct {
	Res                  int32    `protobuf:"varint,1,opt,name=Res,proto3" json:"Res,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationRep) Reset()         { *m = OperationRep{} }
func (m *OperationRep) String() string { return proto.CompactTextString(m) }
func (*OperationRep) ProtoMessage()    {}
func (*OperationRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{4}
}

func (m *OperationRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationRep.Unmarshal(m, b)
}
func (m *OperationRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationRep.Marshal(b, m, deterministic)
}
func (m *OperationRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationRep.Merge(m, src)
}
func (m *OperationRep) XXX_Size() int {
	return xxx_messageInfo_OperationRep.Size(m)
}
func (m *OperationRep) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationRep.DiscardUnknown(m)
}

var xxx_messageInfo_OperationRep proto.InternalMessageInfo

func (m *OperationRep) GetRes() int32 {
	if m != nil {
		return m.Res
	}
	return 0
}

func (m *OperationRep) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

//实名认证请求
type RealAuthReq struct {
	Head                 *OpenPre `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	IdCardNo             string   `protobuf:"bytes,2,opt,name=IdCardNo,proto3" json:"IdCardNo,omitempty"`
	RealName             string   `protobuf:"bytes,3,opt,name=RealName,proto3" json:"RealName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RealAuthReq) Reset()         { *m = RealAuthReq{} }
func (m *RealAuthReq) String() string { return proto.CompactTextString(m) }
func (*RealAuthReq) ProtoMessage()    {}
func (*RealAuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{5}
}

func (m *RealAuthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealAuthReq.Unmarshal(m, b)
}
func (m *RealAuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealAuthReq.Marshal(b, m, deterministic)
}
func (m *RealAuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealAuthReq.Merge(m, src)
}
func (m *RealAuthReq) XXX_Size() int {
	return xxx_messageInfo_RealAuthReq.Size(m)
}
func (m *RealAuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RealAuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_RealAuthReq proto.InternalMessageInfo

func (m *RealAuthReq) GetHead() *OpenPre {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *RealAuthReq) GetIdCardNo() string {
	if m != nil {
		return m.IdCardNo
	}
	return ""
}

func (m *RealAuthReq) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

//通用的登录接口
type RegOrLoginRep struct {
	Res                  int32     `protobuf:"varint,1,opt,name=Res,proto3" json:"Res,omitempty"`
	Msg                  string    `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	OpenData             *OpenInfo `protobuf:"bytes,3,opt,name=OpenData,proto3" json:"OpenData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegOrLoginRep) Reset()         { *m = RegOrLoginRep{} }
func (m *RegOrLoginRep) String() string { return proto.CompactTextString(m) }
func (*RegOrLoginRep) ProtoMessage()    {}
func (*RegOrLoginRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{6}
}

func (m *RegOrLoginRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegOrLoginRep.Unmarshal(m, b)
}
func (m *RegOrLoginRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegOrLoginRep.Marshal(b, m, deterministic)
}
func (m *RegOrLoginRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegOrLoginRep.Merge(m, src)
}
func (m *RegOrLoginRep) XXX_Size() int {
	return xxx_messageInfo_RegOrLoginRep.Size(m)
}
func (m *RegOrLoginRep) XXX_DiscardUnknown() {
	xxx_messageInfo_RegOrLoginRep.DiscardUnknown(m)
}

var xxx_messageInfo_RegOrLoginRep proto.InternalMessageInfo

func (m *RegOrLoginRep) GetRes() int32 {
	if m != nil {
		return m.Res
	}
	return 0
}

func (m *RegOrLoginRep) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RegOrLoginRep) GetOpenData() *OpenInfo {
	if m != nil {
		return m.OpenData
	}
	return nil
}

type ThirdLoginReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThirdLoginReq) Reset()         { *m = ThirdLoginReq{} }
func (m *ThirdLoginReq) String() string { return proto.CompactTextString(m) }
func (*ThirdLoginReq) ProtoMessage()    {}
func (*ThirdLoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{7}
}

func (m *ThirdLoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThirdLoginReq.Unmarshal(m, b)
}
func (m *ThirdLoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThirdLoginReq.Marshal(b, m, deterministic)
}
func (m *ThirdLoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThirdLoginReq.Merge(m, src)
}
func (m *ThirdLoginReq) XXX_Size() int {
	return xxx_messageInfo_ThirdLoginReq.Size(m)
}
func (m *ThirdLoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ThirdLoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_ThirdLoginReq proto.InternalMessageInfo

//手机号注册或登录 Request
type PhoneRegOrLoginReq struct {
	IsReg                bool                `protobuf:"varint,1,opt,name=IsReg,proto3" json:"IsReg,omitempty"`
	PhoneNum             string              `protobuf:"bytes,2,opt,name=PhoneNum,proto3" json:"PhoneNum,omitempty"`
	Code                 string              `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Ext                  *LoginRegExtendInfo `protobuf:"bytes,4,opt,name=Ext,proto3" json:"Ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PhoneRegOrLoginReq) Reset()         { *m = PhoneRegOrLoginReq{} }
func (m *PhoneRegOrLoginReq) String() string { return proto.CompactTextString(m) }
func (*PhoneRegOrLoginReq) ProtoMessage()    {}
func (*PhoneRegOrLoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{8}
}

func (m *PhoneRegOrLoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneRegOrLoginReq.Unmarshal(m, b)
}
func (m *PhoneRegOrLoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneRegOrLoginReq.Marshal(b, m, deterministic)
}
func (m *PhoneRegOrLoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneRegOrLoginReq.Merge(m, src)
}
func (m *PhoneRegOrLoginReq) XXX_Size() int {
	return xxx_messageInfo_PhoneRegOrLoginReq.Size(m)
}
func (m *PhoneRegOrLoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneRegOrLoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneRegOrLoginReq proto.InternalMessageInfo

func (m *PhoneRegOrLoginReq) GetIsReg() bool {
	if m != nil {
		return m.IsReg
	}
	return false
}

func (m *PhoneRegOrLoginReq) GetPhoneNum() string {
	if m != nil {
		return m.PhoneNum
	}
	return ""
}

func (m *PhoneRegOrLoginReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *PhoneRegOrLoginReq) GetExt() *LoginRegExtendInfo {
	if m != nil {
		return m.Ext
	}
	return nil
}

//账号信息
type OpenInfo struct {
	OpenId               int32    `protobuf:"varint,3,opt,name=OpenId,proto3" json:"OpenId,omitempty"`
	OpenToken            string   `protobuf:"bytes,4,opt,name=OpenToken,proto3" json:"OpenToken,omitempty"`
	RealAuthStatus       int32    `protobuf:"varint,5,opt,name=RealAuthStatus,proto3" json:"RealAuthStatus,omitempty"`
	IsBindPhone          int32    `protobuf:"varint,6,opt,name=IsBindPhone,proto3" json:"IsBindPhone,omitempty"`
	ExtJson              string   `protobuf:"bytes,7,opt,name=ExtJson,proto3" json:"ExtJson,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenInfo) Reset()         { *m = OpenInfo{} }
func (m *OpenInfo) String() string { return proto.CompactTextString(m) }
func (*OpenInfo) ProtoMessage()    {}
func (*OpenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{9}
}

func (m *OpenInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenInfo.Unmarshal(m, b)
}
func (m *OpenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenInfo.Marshal(b, m, deterministic)
}
func (m *OpenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenInfo.Merge(m, src)
}
func (m *OpenInfo) XXX_Size() int {
	return xxx_messageInfo_OpenInfo.Size(m)
}
func (m *OpenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OpenInfo proto.InternalMessageInfo

func (m *OpenInfo) GetOpenId() int32 {
	if m != nil {
		return m.OpenId
	}
	return 0
}

func (m *OpenInfo) GetOpenToken() string {
	if m != nil {
		return m.OpenToken
	}
	return ""
}

func (m *OpenInfo) GetRealAuthStatus() int32 {
	if m != nil {
		return m.RealAuthStatus
	}
	return 0
}

func (m *OpenInfo) GetIsBindPhone() int32 {
	if m != nil {
		return m.IsBindPhone
	}
	return 0
}

func (m *OpenInfo) GetExtJson() string {
	if m != nil {
		return m.ExtJson
	}
	return ""
}

//获取角色信息 Request
type GetAccountsReq struct {
	OpenId               int32    `protobuf:"varint,1,opt,name=OpenId,proto3" json:"OpenId,omitempty"`
	OpenToken            string   `protobuf:"bytes,2,opt,name=OpenToken,proto3" json:"OpenToken,omitempty"`
	GameId               int32    `protobuf:"varint,3,opt,name=GameId,proto3" json:"GameId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountsReq) Reset()         { *m = GetAccountsReq{} }
func (m *GetAccountsReq) String() string { return proto.CompactTextString(m) }
func (*GetAccountsReq) ProtoMessage()    {}
func (*GetAccountsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{10}
}

func (m *GetAccountsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountsReq.Unmarshal(m, b)
}
func (m *GetAccountsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountsReq.Marshal(b, m, deterministic)
}
func (m *GetAccountsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountsReq.Merge(m, src)
}
func (m *GetAccountsReq) XXX_Size() int {
	return xxx_messageInfo_GetAccountsReq.Size(m)
}
func (m *GetAccountsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountsReq proto.InternalMessageInfo

func (m *GetAccountsReq) GetOpenId() int32 {
	if m != nil {
		return m.OpenId
	}
	return 0
}

func (m *GetAccountsReq) GetOpenToken() string {
	if m != nil {
		return m.OpenToken
	}
	return ""
}

func (m *GetAccountsReq) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

//获取角色信息 Response
type GetAccountRep struct {
	Res                  int32      `protobuf:"varint,1,opt,name=Res,proto3" json:"Res,omitempty"`
	Accounts             []*Account `protobuf:"bytes,2,rep,name=Accounts,proto3" json:"Accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAccountRep) Reset()         { *m = GetAccountRep{} }
func (m *GetAccountRep) String() string { return proto.CompactTextString(m) }
func (*GetAccountRep) ProtoMessage()    {}
func (*GetAccountRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{11}
}

func (m *GetAccountRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountRep.Unmarshal(m, b)
}
func (m *GetAccountRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountRep.Marshal(b, m, deterministic)
}
func (m *GetAccountRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountRep.Merge(m, src)
}
func (m *GetAccountRep) XXX_Size() int {
	return xxx_messageInfo_GetAccountRep.Size(m)
}
func (m *GetAccountRep) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountRep.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountRep proto.InternalMessageInfo

func (m *GetAccountRep) GetRes() int32 {
	if m != nil {
		return m.Res
	}
	return 0
}

func (m *GetAccountRep) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

//游戏账号信息
type Account struct {
	UserId               int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	UniqueId             string   `protobuf:"bytes,2,opt,name=UniqueId,proto3" json:"UniqueId,omitempty"`
	ChannelOpenId        string   `protobuf:"bytes,3,opt,name=ChannelOpenId,proto3" json:"ChannelOpenId,omitempty"`
	ChannelUserId        string   `protobuf:"bytes,4,opt,name=ChannelUserId,proto3" json:"ChannelUserId,omitempty"`
	UserToken            string   `protobuf:"bytes,5,opt,name=UserToken,proto3" json:"UserToken,omitempty"`
	RealAuthStatus       int32    `protobuf:"varint,6,opt,name=RealAuthStatus,proto3" json:"RealAuthStatus,omitempty"`
	IsBindPhone          int32    `protobuf:"varint,7,opt,name=IsBindPhone,proto3" json:"IsBindPhone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{12}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Account) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *Account) GetChannelOpenId() string {
	if m != nil {
		return m.ChannelOpenId
	}
	return ""
}

func (m *Account) GetChannelUserId() string {
	if m != nil {
		return m.ChannelUserId
	}
	return ""
}

func (m *Account) GetUserToken() string {
	if m != nil {
		return m.UserToken
	}
	return ""
}

func (m *Account) GetRealAuthStatus() int32 {
	if m != nil {
		return m.RealAuthStatus
	}
	return 0
}

func (m *Account) GetIsBindPhone() int32 {
	if m != nil {
		return m.IsBindPhone
	}
	return 0
}

//发送兑换码 Request
type SendCodeReq struct {
	PhoneNum             string   `protobuf:"bytes,1,opt,name=PhoneNum,proto3" json:"PhoneNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendCodeReq) Reset()         { *m = SendCodeReq{} }
func (m *SendCodeReq) String() string { return proto.CompactTextString(m) }
func (*SendCodeReq) ProtoMessage()    {}
func (*SendCodeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6692a741fe744d, []int{13}
}

func (m *SendCodeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendCodeReq.Unmarshal(m, b)
}
func (m *SendCodeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendCodeReq.Marshal(b, m, deterministic)
}
func (m *SendCodeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendCodeReq.Merge(m, src)
}
func (m *SendCodeReq) XXX_Size() int {
	return xxx_messageInfo_SendCodeReq.Size(m)
}
func (m *SendCodeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendCodeReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendCodeReq proto.InternalMessageInfo

func (m *SendCodeReq) GetPhoneNum() string {
	if m != nil {
		return m.PhoneNum
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRegExtendInfo)(nil), "LoginRegExtendInfo")
	proto.RegisterType((*AccountRegOrLoginReq)(nil), "AccountRegOrLoginReq")
	proto.RegisterType((*OpenPre)(nil), "OpenPre")
	proto.RegisterType((*BindPhoneReq)(nil), "BindPhoneReq")
	proto.RegisterType((*OperationRep)(nil), "OperationRep")
	proto.RegisterType((*RealAuthReq)(nil), "RealAuthReq")
	proto.RegisterType((*RegOrLoginRep)(nil), "RegOrLoginRep")
	proto.RegisterType((*ThirdLoginReq)(nil), "ThirdLoginReq")
	proto.RegisterType((*PhoneRegOrLoginReq)(nil), "PhoneRegOrLoginReq")
	proto.RegisterType((*OpenInfo)(nil), "OpenInfo")
	proto.RegisterType((*GetAccountsReq)(nil), "GetAccountsReq")
	proto.RegisterType((*GetAccountRep)(nil), "GetAccountRep")
	proto.RegisterType((*Account)(nil), "Account")
	proto.RegisterType((*SendCodeReq)(nil), "SendCodeReq")
}

func init() { proto.RegisterFile("channel_server.proto", fileDescriptor_4c6692a741fe744d) }

var fileDescriptor_4c6692a741fe744d = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xad, 0xe3, 0xfc, 0x8e, 0x93, 0xf4, 0xfb, 0xa6, 0x05, 0x59, 0x51, 0x85, 0xaa, 0x55, 0x0b,
	0x54, 0x48, 0x16, 0x0a, 0x37, 0x5c, 0x20, 0xa1, 0x12, 0xaa, 0x62, 0x44, 0x7f, 0xb4, 0x6d, 0x85,
	0x84, 0x10, 0xc8, 0xc4, 0x43, 0x12, 0xb5, 0x59, 0xa7, 0xb6, 0x83, 0x7a, 0xcb, 0x1b, 0xc0, 0x3b,
	0xf0, 0x6a, 0x3c, 0x00, 0x6f, 0x80, 0x76, 0xbd, 0x76, 0x9c, 0x38, 0x40, 0xe0, 0x6e, 0x66, 0x76,
	0x3d, 0x73, 0x76, 0xce, 0x99, 0x31, 0x6c, 0xf6, 0x87, 0x9e, 0x10, 0x74, 0xf5, 0x3e, 0xa2, 0xf0,
	0x13, 0x85, 0xce, 0x24, 0x0c, 0xe2, 0x80, 0x7d, 0x31, 0x00, 0x5f, 0x05, 0x83, 0x91, 0xe0, 0x34,
	0x38, 0xb8, 0x89, 0x49, 0xf8, 0xae, 0xf8, 0x18, 0x20, 0x42, 0xd9, 0x1d, 0xd3, 0xc8, 0x36, 0xb6,
	0x8d, 0xfb, 0x0d, 0xae, 0x6c, 0x6c, 0x43, 0xc9, 0x9d, 0xd8, 0x25, 0x15, 0x29, 0xb9, 0x13, 0xbc,
	0x03, 0x70, 0x3a, 0x0c, 0x04, 0x1d, 0x05, 0x3e, 0x5d, 0xd9, 0xa6, 0x8a, 0xe7, 0x22, 0xd8, 0x81,
	0xba, 0xeb, 0xf7, 0xbc, 0xd0, 0x3f, 0x0e, 0xec, 0xb2, 0x3a, 0xcd, 0x7c, 0x79, 0xc6, 0xc9, 0xbb,
	0x3a, 0xf6, 0xc6, 0x64, 0x57, 0x92, 0xb3, 0xd4, 0x67, 0x5f, 0x0d, 0xd8, 0xdc, 0xef, 0xf7, 0x83,
	0xa9, 0x88, 0x39, 0x0d, 0x4e, 0x42, 0x0d, 0xef, 0x1a, 0x37, 0xa1, 0xe2, 0x46, 0x9c, 0x06, 0x0a,
	0x55, 0x9d, 0x27, 0x0e, 0x6e, 0x83, 0xa5, 0x6f, 0xab, 0x6c, 0x09, 0xbe, 0x7c, 0x48, 0x16, 0x3b,
	0xf5, 0xa2, 0xe8, 0x75, 0x10, 0xfa, 0x1a, 0x66, 0xe6, 0xe3, 0x2e, 0x98, 0x07, 0x37, 0xb1, 0xc2,
	0x67, 0x75, 0x37, 0x9c, 0x62, 0x2b, 0xb8, 0x3c, 0x67, 0x4f, 0xa1, 0x76, 0x32, 0x21, 0x71, 0x1a,
	0x12, 0xde, 0x86, 0xaa, 0x34, 0x5d, 0x5f, 0xc1, 0xa8, 0x70, 0xed, 0xe1, 0x16, 0x34, 0xa4, 0x75,
	0x1e, 0x5c, 0x92, 0xd0, 0x28, 0x66, 0x01, 0xf6, 0x16, 0x9a, 0xcf, 0x46, 0xc2, 0x57, 0xed, 0x91,
	0x6f, 0xd9, 0x82, 0xf2, 0x0b, 0xf2, 0x92, 0x1c, 0x56, 0xb7, 0xee, 0xe8, 0xec, 0x5c, 0x45, 0x15,
	0x62, 0x79, 0xf3, 0x78, 0x3a, 0xd6, 0xa9, 0x32, 0x5f, 0x52, 0xd3, 0x0b, 0x7c, 0xd2, 0x2f, 0x51,
	0x36, 0xeb, 0x42, 0xf3, 0x64, 0x42, 0xa1, 0x17, 0x8f, 0x02, 0xc1, 0x69, 0x82, 0xff, 0x81, 0xc9,
	0x29, 0xd2, 0x00, 0xa5, 0x29, 0x23, 0x47, 0xd1, 0x40, 0x27, 0x93, 0x26, 0xeb, 0x83, 0x25, 0x5b,
	0xbe, 0x3f, 0x8d, 0x87, 0x2b, 0x01, 0xca, 0xb8, 0x2c, 0xfd, 0x86, 0x4b, 0x73, 0x81, 0xcb, 0x37,
	0xd0, 0xca, 0x73, 0xb8, 0x12, 0x32, 0xdc, 0x85, 0xba, 0xac, 0xfe, 0xdc, 0x8b, 0x3d, 0x95, 0xd0,
	0xea, 0x36, 0x14, 0x1c, 0x45, 0x47, 0x76, 0xc4, 0xd6, 0xa1, 0x75, 0x3e, 0x1c, 0x85, 0x7e, 0xaa,
	0x0f, 0xf6, 0xd9, 0x00, 0xd4, 0x0d, 0xfe, 0xb3, 0x6c, 0xfe, 0xb2, 0xc5, 0xab, 0x0a, 0xe5, 0x9b,
	0x91, 0x80, 0x57, 0x53, 0x34, 0x93, 0x8a, 0xf9, 0x6b, 0xa9, 0x94, 0x17, 0xa4, 0x82, 0x77, 0xa1,
	0x9d, 0x12, 0x73, 0x16, 0x7b, 0xf1, 0x34, 0x52, 0x13, 0x52, 0xe1, 0x0b, 0x51, 0x29, 0x7c, 0x37,
	0xca, 0x44, 0x65, 0x57, 0xd5, 0xa5, 0x7c, 0x08, 0x6d, 0xa8, 0x1d, 0xdc, 0xc4, 0x2f, 0xa3, 0x40,
	0xd8, 0x35, 0x55, 0x25, 0x75, 0xd9, 0x3b, 0x68, 0x1f, 0x52, 0xac, 0x87, 0x24, 0x92, 0x5d, 0xfa,
	0x27, 0x59, 0xcb, 0xaf, 0x0e, 0xbd, 0x31, 0xcd, 0x5e, 0x98, 0x78, 0xec, 0x10, 0x5a, 0xb3, 0xfc,
	0xcb, 0x79, 0xdf, 0x81, 0x7a, 0x5a, 0xdf, 0x2e, 0x6d, 0x9b, 0x4a, 0x74, 0xe9, 0x07, 0xd9, 0x09,
	0xfb, 0x61, 0x40, 0x4d, 0x3b, 0xb2, 0xd8, 0x45, 0x44, 0xa1, 0x86, 0x68, 0x72, 0xed, 0x49, 0x2a,
	0x2f, 0xc4, 0xe8, 0x7a, 0x2a, 0x61, 0x68, 0x2a, 0x53, 0x1f, 0x77, 0xa0, 0xd5, 0x4b, 0xf6, 0x5e,
	0x8e, 0x89, 0x06, 0x9f, 0x0f, 0xe6, 0x6e, 0xe9, 0x02, 0xe5, 0xb9, 0x5b, 0xba, 0xce, 0x16, 0x34,
	0xa4, 0x95, 0xb4, 0x22, 0xd9, 0x5a, 0xb3, 0xc0, 0x12, 0xda, 0xaa, 0xab, 0xd0, 0x56, 0x2b, 0xd0,
	0xc6, 0xf6, 0xc0, 0x3a, 0x23, 0xe1, 0x4b, 0xd9, 0x49, 0x66, 0xf2, 0x4a, 0x35, 0xe6, 0x95, 0xda,
	0xfd, 0x5e, 0xca, 0x90, 0x9f, 0xa9, 0xb5, 0x8e, 0xf7, 0x00, 0x7a, 0x43, 0xea, 0x5f, 0x26, 0xa0,
	0xb2, 0x39, 0xee, 0xb4, 0x9c, 0xfc, 0x86, 0x60, 0x6b, 0xf8, 0x18, 0xd6, 0x17, 0x86, 0x05, 0x37,
	0x9c, 0xe2, 0xf8, 0x74, 0xda, 0xce, 0xdc, 0x04, 0xb3, 0x35, 0x7c, 0x02, 0xff, 0x17, 0xf6, 0x33,
	0xde, 0x72, 0x96, 0xed, 0xec, 0x25, 0x5f, 0xef, 0x41, 0x3d, 0x7d, 0x1d, 0x36, 0x9d, 0xdc, 0x43,
	0x8b, 0x10, 0x1f, 0x82, 0x95, 0x53, 0x29, 0xae, 0x3b, 0xf3, 0x9a, 0xed, 0xb4, 0x9d, 0x39, 0x91,
	0xb1, 0x35, 0x7c, 0x00, 0x8d, 0x99, 0xfc, 0x5b, 0x4e, 0x7e, 0xe5, 0x16, 0xd3, 0xef, 0x25, 0x8b,
	0x4b, 0x72, 0x83, 0x4d, 0x27, 0xb7, 0x0c, 0x0b, 0x57, 0x3f, 0x54, 0xd5, 0xdf, 0xf2, 0xd1, 0xcf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x47, 0xcc, 0xbd, 0x45, 0x07, 0x00, 0x00,
}
