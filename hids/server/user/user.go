package user

import "github.com/1uLang/zhiannet-api/hids/model/user"

func Add() (uint64, error) {
	return user.Add()
}
