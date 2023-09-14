package utils

import (
  "github.com/go-bineanshi/usercenter-rpc/ent/property"
  "regexp"
)

func CheckLoginnameType(loginname string) property.Provider {
  // 正则表达式匹配手机号
  mobileRegexp := `^1[0-9]{10}$`
  if match, _ := regexp.MatchString(mobileRegexp, loginname); match {
    return property.Phone
  }

  // 正则表达式匹配邮箱
  emailRegexp := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
  if match, _ := regexp.MatchString(emailRegexp, loginname); match {
    return property.Email
  }

  // 默认为普通用户名
  return property.Username
}
