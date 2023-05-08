package service

import (
	"go-web-generator/dto"
	"go-web-generator/vo"
)

// QueryPage 用户表分页查询
func QueryPage(param dto.UserPageDTO) []vo.UserVO {

	return []vo.UserVO{}
}

// Add 新增用户表
func Add(param dto.UserAddDTO) bool {

	return true
}

// Update 更新用户表
func Update(param dto.UserAddDTO) bool {

	return true
}

// QueryDetails 查询用户表
func QueryDetails(id int64) vo.UserVO {

	return vo.UserVO{}
}

// Delete 删除用户表
func Delete(id int64) bool {

	return true
}

// DeleteBatch 批量删除用户表
func DeleteBatch(id []int64) bool {

	return true
}
