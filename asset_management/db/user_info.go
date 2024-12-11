package db

import "fmt"

type UserInfo struct {
	UserID       string `gorm:"primaryKey;column:user_id"` // 用户ID
	UserPassword string `gorm:"column:user_password"`      // 用户登录密码
	UserType     string `gorm:"column:user_type"`          // 用户类型
	UserName     string `gorm:"column:user_name"`          // 用户名
	SourceID     string `gorm:"column:source_id"`          // 认证来源ID
	RegisterTime string `gorm:"column:register_time"`      // 注册时间
	IdentityTime string `gorm:"column:identity_time"`      // 认证时间
}

func (u *UserInfo) TableName() string {
	return "user_info"
}

func CreateUserInfo(UserInfo *UserInfo) error {
	result := DB.Create(UserInfo)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("insert failed,affected rows:%d", result.RowsAffected)
	}
	return nil
}

// Login 实现基于用户名和密码的登录
func Login(username, password string) (*UserInfo, error) {
	var user UserInfo
	err := DB.Where("user_name = ? and user_password = ?", username, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	if user.UserPassword != password {
		return nil, fmt.Errorf("password not match")
	}
	return &user, nil
}

// GetUserInfoByID 查询用户信息
func GetUserInfoByID(id string) (*UserInfo, error) {
	var UserInfo UserInfo
	err := DB.Model(&UserInfo).First(&UserInfo, "user_id = ?", id).Error
	return &UserInfo, err
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(id string, updatedFields interface{}) error {
	return DB.Model(&UserInfo{}).Where("user_id = ?", id).Updates(updatedFields).Error
}

// DeleteUserInfo 删除用户信息
func DeleteUserInfo(id string) error {
	return DB.Delete(&UserInfo{}, "user_id = ?", id).Error
}
func GetAllUserInfoWithConditions(conditions map[string]interface{}, page, pageSize int) ([]UserInfo, error) {
	var asset []UserInfo

	// 计算 OFFSET
	offset := (page - 1) * pageSize

	// 使用 Limit 和 Offset 进行分页查询，并根据条件构建 WHERE 子句
	err := DB.Model(&asset).Where(conditions).Limit(pageSize).Offset(offset).Find(&asset).Error
	return asset, err
}
