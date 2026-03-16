package model

import "github.com/HappyLadySauce/Beehive-M/services/user/pb"

type User struct {
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Account       string                 `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Nickname      string                 `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar        string                 `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Phone         string                 `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`       // 可脱敏
	Email         string                 `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`       // 可脱敏
	Gender        int32                  `protobuf:"varint,7,opt,name=gender,proto3" json:"gender,omitempty"`    // 0未知 1男 2女
	Birthday      string                 `protobuf:"bytes,8,opt,name=birthday,proto3" json:"birthday,omitempty"` // YYYY-MM-DD
	Signature     string                 `protobuf:"bytes,9,opt,name=signature,proto3" json:"signature,omitempty"`
	Status        int32                  `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`                                // 1正常 2冻结
	LastLoginAt   int64                  `protobuf:"varint,11,opt,name=last_login_at,json=lastLoginAt,proto3" json:"last_login_at,omitempty"` // 时间戳
	CreatedAt     int64                  `protobuf:"varint,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     int64                  `protobuf:"varint,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *User) ModelToPB() *pb.User {
	if x != nil {
		return &pb.User{
			UserId: x.UserId,
			Account: x.Account,
			Nickname: x.Nickname,
			Avatar: x.Avatar,
			Phone: x.Phone,
			Email: x.Email,
			Gender: x.Gender,
			Birthday: x.Birthday,
			Signature: x.Signature,
			Status: x.Status,
			LastLoginAt: x.LastLoginAt,
			CreatedAt: x.CreatedAt,
			UpdatedAt: x.UpdatedAt,
		}
	}
	return nil
}

func (x *User) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *User) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *User) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *User) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *User) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *User) GetLastLoginAt() int64 {
	if x != nil {
		return x.LastLoginAt
	}
	return 0
}

func (x *User) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *User) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}
