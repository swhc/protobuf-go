// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: foot_live_stat.proto

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

//直播 - 直播统计
type FootLiveStatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId int64 `protobuf:"varint,1,opt,name=matchId,proto3" json:"matchId,omitempty"` //比赛id
}

func (x *FootLiveStatRequest) Reset() {
	*x = FootLiveStatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_live_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLiveStatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLiveStatRequest) ProtoMessage() {}

func (x *FootLiveStatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foot_live_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLiveStatRequest.ProtoReflect.Descriptor instead.
func (*FootLiveStatRequest) Descriptor() ([]byte, []int) {
	return file_foot_live_stat_proto_rawDescGZIP(), []int{0}
}

func (x *FootLiveStatRequest) GetMatchId() int64 {
	if x != nil {
		return x.MatchId
	}
	return 0
}

//直播 -直播统计
type FootLiveStatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Odds        []*FootLiveOdds        `protobuf:"bytes,1,rep,name=odds,proto3" json:"odds,omitempty"`                                  //指数
	Event       []*FootLiveEvent       `protobuf:"bytes,2,rep,name=event,proto3" json:"event,omitempty"`                                //事件
	OddsCompany []*FootLiveOddsCompany `protobuf:"bytes,3,rep,name=odds_company,json=oddsCompany,proto3" json:"odds_company,omitempty"` //赔率公司
	IsEnd       bool                   `protobuf:"varint,4,opt,name=is_end,json=isEnd,proto3" json:"is_end,omitempty"`                  //是否结束
	Result      string                 `protobuf:"bytes,5,opt,name=result,proto3" json:"result,omitempty"`                              //结果
}

func (x *FootLiveStatResponse) Reset() {
	*x = FootLiveStatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_live_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLiveStatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLiveStatResponse) ProtoMessage() {}

func (x *FootLiveStatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foot_live_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLiveStatResponse.ProtoReflect.Descriptor instead.
func (*FootLiveStatResponse) Descriptor() ([]byte, []int) {
	return file_foot_live_stat_proto_rawDescGZIP(), []int{1}
}

func (x *FootLiveStatResponse) GetOdds() []*FootLiveOdds {
	if x != nil {
		return x.Odds
	}
	return nil
}

func (x *FootLiveStatResponse) GetEvent() []*FootLiveEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *FootLiveStatResponse) GetOddsCompany() []*FootLiveOddsCompany {
	if x != nil {
		return x.OddsCompany
	}
	return nil
}

func (x *FootLiveStatResponse) GetIsEnd() bool {
	if x != nil {
		return x.IsEnd
	}
	return false
}

func (x *FootLiveStatResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

//指数
type FootLiveOdds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PanId    int64               `protobuf:"varint,1,opt,name=pan_id,json=panId,proto3" json:"pan_id,omitempty"`          //指数名称
	Name     string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                          //指数（1：亚盘，2：欧盘，3：大小球）
	MarkType int64               `protobuf:"varint,3,opt,name=mark_type,json=markType,proto3" json:"mark_type,omitempty"` //指数类型（1：欧盘，2：亚赔，3：大小球，4：角球【待定，暂无数据】）(无用字段，后台标识)
	List     []*FootLiveOddsList `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`                          //指数数据列表
}

func (x *FootLiveOdds) Reset() {
	*x = FootLiveOdds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_live_stat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLiveOdds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLiveOdds) ProtoMessage() {}

func (x *FootLiveOdds) ProtoReflect() protoreflect.Message {
	mi := &file_foot_live_stat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLiveOdds.ProtoReflect.Descriptor instead.
func (*FootLiveOdds) Descriptor() ([]byte, []int) {
	return file_foot_live_stat_proto_rawDescGZIP(), []int{2}
}

func (x *FootLiveOdds) GetPanId() int64 {
	if x != nil {
		return x.PanId
	}
	return 0
}

func (x *FootLiveOdds) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FootLiveOdds) GetMarkType() int64 {
	if x != nil {
		return x.MarkType
	}
	return 0
}

func (x *FootLiveOdds) GetList() []*FootLiveOddsList {
	if x != nil {
		return x.List
	}
	return nil
}

//指数数据列表
type FootLiveOddsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId       int64  `protobuf:"varint,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`                    //公司赔率id
	AddTime         int64  `protobuf:"varint,2,opt,name=add_time,json=addTime,proto3" json:"add_time,omitempty"`                          //指数开始时间（-1：早，-2：未）
	Score           string `protobuf:"bytes,3,opt,name=score,proto3" json:"score,omitempty"`                                              //比分
	InitMainParam   string `protobuf:"bytes,4,opt,name=init_main_param,json=initMainParam,proto3" json:"init_main_param,omitempty"`       //初始主胜（u = 升 | d = 降）
	InitMiddleParam string `protobuf:"bytes,5,opt,name=init_middle_param,json=initMiddleParam,proto3" json:"init_middle_param,omitempty"` //初始盘口（u = 升 | d = 降）
	InitGuestParam  string `protobuf:"bytes,6,opt,name=init_guest_param,json=initGuestParam,proto3" json:"init_guest_param,omitempty"`    //初始客胜（u = 升 | d = 降）
	NewMainParam    string `protobuf:"bytes,7,opt,name=new_main_param,json=newMainParam,proto3" json:"new_main_param,omitempty"`          //即时主胜（u = 升 | d = 降）
	NewMiddleParam  string `protobuf:"bytes,8,opt,name=new_middle_param,json=newMiddleParam,proto3" json:"new_middle_param,omitempty"`    //即时盘口（u = 升 | d = 降）
	NewGuestParam   string `protobuf:"bytes,9,opt,name=new_guest_param,json=newGuestParam,proto3" json:"new_guest_param,omitempty"`       //即时客胜（u = 升 | d = 降）
}

func (x *FootLiveOddsList) Reset() {
	*x = FootLiveOddsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_live_stat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLiveOddsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLiveOddsList) ProtoMessage() {}

func (x *FootLiveOddsList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_live_stat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLiveOddsList.ProtoReflect.Descriptor instead.
func (*FootLiveOddsList) Descriptor() ([]byte, []int) {
	return file_foot_live_stat_proto_rawDescGZIP(), []int{3}
}

func (x *FootLiveOddsList) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *FootLiveOddsList) GetAddTime() int64 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

func (x *FootLiveOddsList) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *FootLiveOddsList) GetInitMainParam() string {
	if x != nil {
		return x.InitMainParam
	}
	return ""
}

func (x *FootLiveOddsList) GetInitMiddleParam() string {
	if x != nil {
		return x.InitMiddleParam
	}
	return ""
}

func (x *FootLiveOddsList) GetInitGuestParam() string {
	if x != nil {
		return x.InitGuestParam
	}
	return ""
}

func (x *FootLiveOddsList) GetNewMainParam() string {
	if x != nil {
		return x.NewMainParam
	}
	return ""
}

func (x *FootLiveOddsList) GetNewMiddleParam() string {
	if x != nil {
		return x.NewMiddleParam
	}
	return ""
}

func (x *FootLiveOddsList) GetNewGuestParam() string {
	if x != nil {
		return x.NewGuestParam
	}
	return ""
}

type FootLiveEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                        //事件id
	IntervalTime string                 `protobuf:"bytes,2,opt,name=interval_time,json=intervalTime,proto3" json:"interval_time,omitempty"` //比赛事件
	Fraction     string                 `protobuf:"bytes,3,opt,name=fraction,proto3" json:"fraction,omitempty"`                             //当前事件比分
	Main         *FootLiveEventDataList `protobuf:"bytes,4,opt,name=main,proto3" json:"main,omitempty"`                                     //主队
	Cust         *FootLiveEventDataList `protobuf:"bytes,5,opt,name=cust,proto3" json:"cust,omitempty"`                                     //客队
}

func (x *FootLiveEvent) Reset() {
	*x = FootLiveEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_live_stat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLiveEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLiveEvent) ProtoMessage() {}

func (x *FootLiveEvent) ProtoReflect() protoreflect.Message {
	mi := &file_foot_live_stat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLiveEvent.ProtoReflect.Descriptor instead.
func (*FootLiveEvent) Descriptor() ([]byte, []int) {
	return file_foot_live_stat_proto_rawDescGZIP(), []int{4}
}

func (x *FootLiveEvent) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FootLiveEvent) GetIntervalTime() string {
	if x != nil {
		return x.IntervalTime
	}
	return ""
}

func (x *FootLiveEvent) GetFraction() string {
	if x != nil {
		return x.Fraction
	}
	return ""
}

func (x *FootLiveEvent) GetMain() *FootLiveEventDataList {
	if x != nil {
		return x.Main
	}
	return nil
}

func (x *FootLiveEvent) GetCust() *FootLiveEventDataList {
	if x != nil {
		return x.Cust
	}
	return nil
}

type FootLiveEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId   int64  `protobuf:"varint,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`      //队伍id
	PersonName string `protobuf:"bytes,2,opt,name=person_name,json=personName,proto3" json:"person_name,omitempty"` //球员名称
	Record     string `protobuf:"bytes,3,opt,name=record,proto3" json:"record,omitempty"`                           //记录（类型）
	TeamId     int64  `protobuf:"varint,4,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`            //球队id
}

func (x *FootLiveEventData) Reset() {
	*x = FootLiveEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_live_stat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLiveEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLiveEventData) ProtoMessage() {}

func (x *FootLiveEventData) ProtoReflect() protoreflect.Message {
	mi := &file_foot_live_stat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLiveEventData.ProtoReflect.Descriptor instead.
func (*FootLiveEventData) Descriptor() ([]byte, []int) {
	return file_foot_live_stat_proto_rawDescGZIP(), []int{5}
}

func (x *FootLiveEventData) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *FootLiveEventData) GetPersonName() string {
	if x != nil {
		return x.PersonName
	}
	return ""
}

func (x *FootLiveEventData) GetRecord() string {
	if x != nil {
		return x.Record
	}
	return ""
}

func (x *FootLiveEventData) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type FootLiveEventDataList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team []*FootLiveEventData `protobuf:"bytes,1,rep,name=team,proto3" json:"team,omitempty"`
}

func (x *FootLiveEventDataList) Reset() {
	*x = FootLiveEventDataList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_live_stat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLiveEventDataList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLiveEventDataList) ProtoMessage() {}

func (x *FootLiveEventDataList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_live_stat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLiveEventDataList.ProtoReflect.Descriptor instead.
func (*FootLiveEventDataList) Descriptor() ([]byte, []int) {
	return file_foot_live_stat_proto_rawDescGZIP(), []int{6}
}

func (x *FootLiveEventDataList) GetTeam() []*FootLiveEventData {
	if x != nil {
		return x.Team
	}
	return nil
}

type FootLiveOddsCompany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`        //指数id
	Name   string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`     //指数类型名称
	Titles []string                   `protobuf:"bytes,3,rep,name=titles,proto3" json:"titles,omitempty"` //标题名称
	List   []*FootLiveOddsCompanyList `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`     //公司列表
}

func (x *FootLiveOddsCompany) Reset() {
	*x = FootLiveOddsCompany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_live_stat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLiveOddsCompany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLiveOddsCompany) ProtoMessage() {}

func (x *FootLiveOddsCompany) ProtoReflect() protoreflect.Message {
	mi := &file_foot_live_stat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLiveOddsCompany.ProtoReflect.Descriptor instead.
func (*FootLiveOddsCompany) Descriptor() ([]byte, []int) {
	return file_foot_live_stat_proto_rawDescGZIP(), []int{7}
}

func (x *FootLiveOddsCompany) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FootLiveOddsCompany) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FootLiveOddsCompany) GetTitles() []string {
	if x != nil {
		return x.Titles
	}
	return nil
}

func (x *FootLiveOddsCompany) GetList() []*FootLiveOddsCompanyList {
	if x != nil {
		return x.List
	}
	return nil
}

type FootLiveOddsCompanyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                            //公司id
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                         //公司名称
	NowData  []string `protobuf:"bytes,3,rep,name=now_data,json=nowData,proto3" json:"now_data,omitempty"`    //即时数据
	InitData []string `protobuf:"bytes,4,rep,name=init_data,json=initData,proto3" json:"init_data,omitempty"` //初始数据
}

func (x *FootLiveOddsCompanyList) Reset() {
	*x = FootLiveOddsCompanyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foot_live_stat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FootLiveOddsCompanyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FootLiveOddsCompanyList) ProtoMessage() {}

func (x *FootLiveOddsCompanyList) ProtoReflect() protoreflect.Message {
	mi := &file_foot_live_stat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FootLiveOddsCompanyList.ProtoReflect.Descriptor instead.
func (*FootLiveOddsCompanyList) Descriptor() ([]byte, []int) {
	return file_foot_live_stat_proto_rawDescGZIP(), []int{8}
}

func (x *FootLiveOddsCompanyList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FootLiveOddsCompanyList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FootLiveOddsCompanyList) GetNowData() []string {
	if x != nil {
		return x.NowData
	}
	return nil
}

func (x *FootLiveOddsCompanyList) GetInitData() []string {
	if x != nil {
		return x.InitData
	}
	return nil
}

var File_foot_live_stat_proto protoreflect.FileDescriptor

var file_foot_live_stat_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x6f, 0x6f, 0x74, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x13, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69,
	0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x22, 0xc7, 0x01, 0x0a, 0x14, 0x46, 0x6f, 0x6f, 0x74,
	0x4c, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x04, 0x6f, 0x64, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x64, 0x64, 0x73, 0x52, 0x04, 0x6f,
	0x64, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x6f, 0x64, 0x64,
	0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x64, 0x64, 0x73, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x0b, 0x6f, 0x64, 0x64, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x45, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x7d, 0x0a, 0x0c, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x64, 0x64,
	0x73, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6d, 0x61, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69,
	0x76, 0x65, 0x4f, 0x64, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x22, 0xd8, 0x02, 0x0a, 0x10, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x64, 0x64,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x69, 0x6e, 0x69, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x2a, 0x0a,
	0x11, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x4d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x6e, 0x69,
	0x74, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x47, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x77,
	0x4d, 0x61, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x77,
	0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x77, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65,
	0x77, 0x47, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0xb8, 0x01, 0x0a, 0x0d,
	0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a,
	0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x46,
	0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x63, 0x75,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4c,
	0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x04, 0x63, 0x75, 0x73, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x46, 0x6f, 0x6f, 0x74, 0x4c,
	0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x15, 0x46,
	0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x22, 0x7f, 0x0a, 0x13,
	0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x64, 0x64, 0x73, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x12,
	0x2c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x64, 0x64, 0x73, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x75, 0x0a,
	0x17, 0x46, 0x6f, 0x6f, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x64, 0x64, 0x73, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x6e, 0x6f, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x69, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x69, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foot_live_stat_proto_rawDescOnce sync.Once
	file_foot_live_stat_proto_rawDescData = file_foot_live_stat_proto_rawDesc
)

func file_foot_live_stat_proto_rawDescGZIP() []byte {
	file_foot_live_stat_proto_rawDescOnce.Do(func() {
		file_foot_live_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_foot_live_stat_proto_rawDescData)
	})
	return file_foot_live_stat_proto_rawDescData
}

var file_foot_live_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_foot_live_stat_proto_goTypes = []interface{}{
	(*FootLiveStatRequest)(nil),     // 0: FootLiveStatRequest
	(*FootLiveStatResponse)(nil),    // 1: FootLiveStatResponse
	(*FootLiveOdds)(nil),            // 2: FootLiveOdds
	(*FootLiveOddsList)(nil),        // 3: FootLiveOddsList
	(*FootLiveEvent)(nil),           // 4: FootLiveEvent
	(*FootLiveEventData)(nil),       // 5: FootLiveEventData
	(*FootLiveEventDataList)(nil),   // 6: FootLiveEventDataList
	(*FootLiveOddsCompany)(nil),     // 7: FootLiveOddsCompany
	(*FootLiveOddsCompanyList)(nil), // 8: FootLiveOddsCompanyList
}
var file_foot_live_stat_proto_depIdxs = []int32{
	2, // 0: FootLiveStatResponse.odds:type_name -> FootLiveOdds
	4, // 1: FootLiveStatResponse.event:type_name -> FootLiveEvent
	7, // 2: FootLiveStatResponse.odds_company:type_name -> FootLiveOddsCompany
	3, // 3: FootLiveOdds.list:type_name -> FootLiveOddsList
	6, // 4: FootLiveEvent.main:type_name -> FootLiveEventDataList
	6, // 5: FootLiveEvent.cust:type_name -> FootLiveEventDataList
	5, // 6: FootLiveEventDataList.team:type_name -> FootLiveEventData
	8, // 7: FootLiveOddsCompany.list:type_name -> FootLiveOddsCompanyList
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_foot_live_stat_proto_init() }
func file_foot_live_stat_proto_init() {
	if File_foot_live_stat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foot_live_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLiveStatRequest); i {
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
		file_foot_live_stat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLiveStatResponse); i {
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
		file_foot_live_stat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLiveOdds); i {
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
		file_foot_live_stat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLiveOddsList); i {
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
		file_foot_live_stat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLiveEvent); i {
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
		file_foot_live_stat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLiveEventData); i {
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
		file_foot_live_stat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLiveEventDataList); i {
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
		file_foot_live_stat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLiveOddsCompany); i {
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
		file_foot_live_stat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FootLiveOddsCompanyList); i {
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
			RawDescriptor: file_foot_live_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foot_live_stat_proto_goTypes,
		DependencyIndexes: file_foot_live_stat_proto_depIdxs,
		MessageInfos:      file_foot_live_stat_proto_msgTypes,
	}.Build()
	File_foot_live_stat_proto = out.File
	file_foot_live_stat_proto_rawDesc = nil
	file_foot_live_stat_proto_goTypes = nil
	file_foot_live_stat_proto_depIdxs = nil
}
