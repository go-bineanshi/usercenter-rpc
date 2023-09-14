package property

// 用户认证服务商
type Provider string

const (
  Username Provider = "username"
  Phone    Provider = "phone"
  Email    Provider = "email"
  Weixin   Provider = "weixin"
)

func (Provider) Values() (kinds []string) {
  for _, s := range []Provider{Username, Phone, Email, Weixin} {
    kinds = append(kinds, string(s))
  }
  return
}
