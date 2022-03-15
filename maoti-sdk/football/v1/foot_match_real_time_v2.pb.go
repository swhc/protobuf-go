// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_match_real_time_v2.proto

package v1

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

//比赛日历
type FootMatchRealTimeRequestV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int64   `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`            //即时的状态 0全部,1：未开始，2为开始，3为结束
	DateTime string  `protobuf:"bytes,3,opt,name=dateTime,proto3" json:"dateTime,omitempty"`         // 时间范围
	EventIds []int64 `protobuf:"varint,2,rep,packed,name=eventIds,proto3" json:"eventIds,omitempty"` //比赛筛选
}

func (x *FootMatchRealTimeRequestV2) Reset() {
	*x = FootMatchRealTimeRequestV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_real_time_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootMatchRealTimeRequestV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootMatchRealTimeRequestV2) ProtoMessage() {}

func (x *FootMatchRealTimeRequestV2) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_real_time_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootMatchRealTimeRequestV2.ProtoReflect.Descriptor instead.
func (*FootMatchRealTimeRequestV2) Descriptor() ([]byte, []int) {
	return file_foot_match_real_time_v2_proto_rawDescGZIP(), []int{0}
}

func (x *FootMatchRealTimeRequestV2) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FootMatchRealTimeRequestV2) GetDateTime() string {
	if x != nil {
		return x.DateTime
	}
	return ""
}

func (x *FootMatchRealTimeRequestV2) GetEventIds() []int64 {
	if x != nil {
		return x.EventIds
	}
	return nil
}

type FootMatchRealTimeResponseV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C    int64                      `protobuf:"varint,1,opt,name=c,proto3" json:"c,omitempty"`        //count
	CurT int64                      `protobuf:"varint,24,opt,name=curT,proto3" json:"curT,omitempty"` //服务器当前时间戳
	List []*FootMatchRealTimeInfoV2 `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FootMatchRealTimeResponseV2) Reset() {
	*x = FootMatchRealTimeResponseV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_real_time_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootMatchRealTimeResponseV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootMatchRealTimeResponseV2) ProtoMessage() {}

func (x *FootMatchRealTimeResponseV2) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_real_time_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootMatchRealTimeResponseV2.ProtoReflect.Descriptor instead.
func (*FootMatchRealTimeResponseV2) Descriptor() ([]byte, []int) {
	return file_foot_match_real_time_v2_proto_rawDescGZIP(), []int{1}
}

func (x *FootMatchRealTimeResponseV2) GetC() int64 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *FootMatchRealTimeResponseV2) GetCurT() int64 {
	if x != nil {
		return x.CurT
	}
	return 0
}

func (x *FootMatchRealTimeResponseV2) GetList() []*FootMatchRealTimeInfoV2 {
	if x != nil {
		return x.List
	}
	return nil
}

//message FootMatchRealTimeListV2 {
//  int64 stm = 2;  //      当前比赛开始时间戳   startTime
//  repeated  FootMatchRealTimeInfoV2 list = 3;//  数据列表
//  int64 st = 6;    //后端用来排序前端不用管  eventStatusResult
//  int64 tId = 7;    //当前联赛id   //tournamentId
//}
type FootMatchRealTimeInfoV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mcid      int64  `protobuf:"varint,99,opt,name=mcid,proto3" json:"mcid,omitempty"`                            //当前记录ID 赛事iD用于收藏传递
	Mid       int64  `protobuf:"varint,3,opt,name=mid,proto3" json:"mid,omitempty"`                               //主队ID   teamMainId
	Cid       int64  `protobuf:"varint,5,opt,name=cid,proto3" json:"cid,omitempty"`                               //客队ID   teamCustId
	Sc        string `protobuf:"bytes,98,opt,name=sc,proto3" json:"sc,omitempty"`                                 //主客总比分   mainCustScore
	St        int64  `protobuf:"varint,13,opt,name=st,proto3" json:"st,omitempty"`                                //当前比赛状态 1：未开始，2为开始，3为结束 4特殊状况临时中断比赛  5、延迟  6、取消  eventStatusResult
	Pt        int64  `protobuf:"varint,14,opt,name=pt,proto3" json:"pt,omitempty"`                                //播放类型（1:直播，2：3D）  playbackType
	Url       string `protobuf:"bytes,15,opt,name=url,proto3" json:"url,omitempty"`                               //跳转地址   playUrl
	Hsc       string `protobuf:"bytes,17,opt,name=hsc,proto3" json:"hsc,omitempty"`                               //当前半场得分状况   HalfCourt
	Ck        string `protobuf:"bytes,11,opt,name=ck,proto3" json:"ck,omitempty"`                                 //角球
	Yc        string `protobuf:"bytes,12,opt,name=yc,proto3" json:"yc,omitempty"`                                 //黄牌
	Rc        string `protobuf:"bytes,51,opt,name=rc,proto3" json:"rc,omitempty"`                                 //红牌
	Rk        string `protobuf:"bytes,52,opt,name=rk,proto3" json:"rk,omitempty"`                                 //排名
	Re        string `protobuf:"bytes,22,opt,name=re,proto3" json:"re,omitempty"`                                 //备注说明：什么原因中断  reason
	Stid      int64  `protobuf:"varint,23,opt,name=stid,proto3" json:"stid,omitempty"`                            //当前比赛eapi的原有状态   eventStatusId
	Esc       string `protobuf:"bytes,35,opt,name=esc,proto3" json:"esc,omitempty"`                               //加时得分   ExtratimeScore
	Psc       string `protobuf:"bytes,36,opt,name=psc,proto3" json:"psc,omitempty"`                               //点球得分   PenaltyshootoutScore
	Ou        string `protobuf:"bytes,1,opt,name=ou,proto3" json:"ou,omitempty"`                                  //欧盘|中间参数（格式：1.70 +0/0.5 2.00）【这个app需要舍弃，h5在使用】只需要显示下面两对应 middleParam
	Ya        string `protobuf:"bytes,2,opt,name=ya,proto3" json:"ya,omitempty"`                                  //亚赔的即时指数【格式：0.5|u,+2.5,0.25|d】  oddsUp
	Ov        string `protobuf:"bytes,73,opt,name=ov,proto3" json:"ov,omitempty"`                                 //大小球的即时指数【格式：0.5|u,+2.5,0.25|d】  oddsDown
	Round     string `protobuf:"bytes,42,opt,name=round,proto3" json:"round,omitempty"`                           //场次 | 编号  round
	Ctm       string `protobuf:"bytes,45,opt,name=ctm,proto3" json:"ctm,omitempty"`                               //分类用的时间   classifyTime
	Stm       int64  `protobuf:"varint,21,opt,name=stm,proto3" json:"stm,omitempty"`                              //      当前比赛开始时间戳   startTime(data_match.match_time)
	Tid       int64  `protobuf:"varint,7,opt,name=tid,proto3" json:"tid,omitempty"`                               //当前联赛id   //tournamentId
	CurT      int64  `protobuf:"varint,24,opt,name=curT,proto3" json:"curT,omitempty"`                            //服务器当前时间戳
	El        int64  `protobuf:"varint,31,opt,name=el,proto3" json:"el,omitempty"`                                //事件发生时的比赛时长  elapsed
	IsVisible int64  `protobuf:"varint,32,opt,name=is_visible,json=isVisible,proto3" json:"is_visible,omitempty"` //是否可见：0隐藏 1可见
	Atm       int64  `protobuf:"varint,33,opt,name=atm,proto3" json:"atm,omitempty"`                              //data_match.start_time
}

func (x *FootMatchRealTimeInfoV2) Reset() {
	*x = FootMatchRealTimeInfoV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_match_real_time_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootMatchRealTimeInfoV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootMatchRealTimeInfoV2) ProtoMessage() {}

func (x *FootMatchRealTimeInfoV2) ProtoReflect() protoreflect.Message {
	mi := &file_foot_match_real_time_v2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootMatchRealTimeInfoV2.ProtoReflect.Descriptor instead.
func (*FootMatchRealTimeInfoV2) Descriptor() ([]byte, []int) {
	return file_foot_match_real_time_v2_proto_rawDescGZIP(), []int{2}
}

func (x *FootMatchRealTimeInfoV2) GetMcid() int64 {
	if x != nil {
		return x.Mcid
	}
	return 0
}

func (x *FootMatchRealTimeInfoV2) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *FootMatchRealTimeInfoV2) GetCid() int64 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *FootMatchRealTimeInfoV2) GetSc() string {
	if x != nil {
		return x.Sc
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetSt() int64 {
	if x != nil {
		return x.St
	}
	return 0
}

func (x *FootMatchRealTimeInfoV2) GetPt() int64 {
	if x != nil {
		return x.Pt
	}
	return 0
}

func (x *FootMatchRealTimeInfoV2) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetHsc() string {
	if x != nil {
		return x.Hsc
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetCk() string {
	if x != nil {
		return x.Ck
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetYc() string {
	if x != nil {
		return x.Yc
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetRc() string {
	if x != nil {
		return x.Rc
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetRk() string {
	if x != nil {
		return x.Rk
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetRe() string {
	if x != nil {
		return x.Re
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetStid() int64 {
	if x != nil {
		return x.Stid
	}
	return 0
}

func (x *FootMatchRealTimeInfoV2) GetEsc() string {
	if x != nil {
		return x.Esc
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetPsc() string {
	if x != nil {
		return x.Psc
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetOu() string {
	if x != nil {
		return x.Ou
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetYa() string {
	if x != nil {
		return x.Ya
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetOv() string {
	if x != nil {
		return x.Ov
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetRound() string {
	if x != nil {
		return x.Round
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetCtm() string {
	if x != nil {
		return x.Ctm
	}
	return ""
}

func (x *FootMatchRealTimeInfoV2) GetStm() int64 {
	if x != nil {
		return x.Stm
	}
	return 0
}

func (x *FootMatchRealTimeInfoV2) GetTid() int64 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *FootMatchRealTimeInfoV2) GetCurT() int64 {
	if x != nil {
		return x.CurT
	}
	return 0
}

func (x *FootMatchRealTimeInfoV2) GetEl() int64 {
	if x != nil {
		return x.El
	}
	return 0
}

func (x *FootMatchRealTimeInfoV2) GetIsVisible() int64 {
	if x != nil {
		return x.IsVisible
	}
	return 0
}

func (x *FootMatchRealTimeInfoV2) GetAtm() int64 {
	if x != nil {
		return x.Atm
	}
	return 0
}

var File_foot_match_real_time_v2_proto protoreflect.FileDescriptor

var file_foot_match_real_time_v2_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x61,
	0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6c, 0x0a, 0x1a, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x61, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x32, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x22, 0x6d, 0x0a,
	0x1b, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x32, 0x12, 0x0c, 0x0a, 0x01,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x75,
	0x72, 0x54, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x75, 0x72, 0x54, 0x12, 0x2c,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x46,
	0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x56, 0x32, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xfe, 0x03, 0x0a,
	0x17, 0x46, 0x6f, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x63, 0x69, 0x64,
	0x18, 0x63, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6d, 0x63, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x69, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x73, 0x63, 0x18, 0x62, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x63,
	0x12, 0x0e, 0x0a, 0x02, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x70, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x70, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x73, 0x63, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x68, 0x73, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x79, 0x63, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x79, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x63, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x72, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x6b, 0x18, 0x34, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x74, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x73, 0x63, 0x18,
	0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x73, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x73,
	0x63, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x73, 0x63, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x75, 0x12, 0x0e, 0x0a, 0x02,
	0x79, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x79, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x76, 0x18, 0x49, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x76, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x74, 0x6d, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x63, 0x74, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x6d, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x73, 0x74, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x75, 0x72, 0x54,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x75, 0x72, 0x54, 0x12, 0x0e, 0x0a, 0x02,
	0x65, 0x6c, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x74, 0x6d, 0x18, 0x21, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x74, 0x6d, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_match_real_time_v2_proto_rawDescOnce sync.Once
	file_foot_match_real_time_v2_proto_rawDescData = file_foot_match_real_time_v2_proto_rawDesc
)

func file_foot_match_real_time_v2_proto_rawDescGZIP() []byte {
	file_foot_match_real_time_v2_proto_rawDescOnce.Do(func() {
		file_foot_match_real_time_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_match_real_time_v2_proto_rawDescData)
	})
	return file_foot_match_real_time_v2_proto_rawDescData
}

var file_foot_match_real_time_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_foot_match_real_time_v2_proto_goTypes = []interface{}{
	(*FootMatchRealTimeRequestV2)(nil),  // 0: FootMatchRealTimeRequestV2
	(*FootMatchRealTimeResponseV2)(nil), // 1: FootMatchRealTimeResponseV2
	(*FootMatchRealTimeInfoV2)(nil),     // 2: FootMatchRealTimeInfoV2
}
var file_foot_match_real_time_v2_proto_depIdxs = []int32{
	2, // 0: FootMatchRealTimeResponseV2.list:type_name -> FootMatchRealTimeInfoV2
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_foot_match_real_time_v2_proto_init() }
func file_foot_match_real_time_v2_proto_init() {
	if File_foot_match_real_time_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_match_real_time_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootMatchRealTimeRequestV2); i {
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
		file_foot_match_real_time_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootMatchRealTimeResponseV2); i {
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
		file_foot_match_real_time_v2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootMatchRealTimeInfoV2); i {
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
			RawDescriptor: file_foot_match_real_time_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_match_real_time_v2_proto_goTypes,
		DependencyIndexes: file_foot_match_real_time_v2_proto_depIdxs,
		MessageInfos:      file_foot_match_real_time_v2_proto_msgTypes,
	}.Build()
	File_foot_match_real_time_v2_proto = out.File
	file_foot_match_real_time_v2_proto_rawDesc = nil
	file_foot_match_real_time_v2_proto_goTypes = nil
	file_foot_match_real_time_v2_proto_depIdxs = nil
}
