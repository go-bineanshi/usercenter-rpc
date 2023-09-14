// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/go-bineanshi/usercenter-rpc/ent/schema"
	"github.com/go-bineanshi/usercenter-rpc/ent/user"
	"github.com/go-bineanshi/usercenter-rpc/ent/userauth"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userMixin := schema.User{}.Mixin()
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userMixinFields2 := userMixin[2].Fields()
	_ = userMixinFields2
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields1[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields1[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userMixinFields2[0].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(uint8)
	userauthMixin := schema.UserAuth{}.Mixin()
	userauthMixinFields1 := userauthMixin[1].Fields()
	_ = userauthMixinFields1
	userauthMixinFields2 := userauthMixin[2].Fields()
	_ = userauthMixinFields2
	userauthFields := schema.UserAuth{}.Fields()
	_ = userauthFields
	// userauthDescCreatedAt is the schema descriptor for created_at field.
	userauthDescCreatedAt := userauthMixinFields1[0].Descriptor()
	// userauth.DefaultCreatedAt holds the default value on creation for the created_at field.
	userauth.DefaultCreatedAt = userauthDescCreatedAt.Default.(func() time.Time)
	// userauthDescUpdatedAt is the schema descriptor for updated_at field.
	userauthDescUpdatedAt := userauthMixinFields1[1].Descriptor()
	// userauth.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	userauth.DefaultUpdatedAt = userauthDescUpdatedAt.Default.(func() time.Time)
	// userauth.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	userauth.UpdateDefaultUpdatedAt = userauthDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userauthDescStatus is the schema descriptor for status field.
	userauthDescStatus := userauthMixinFields2[0].Descriptor()
	// userauth.DefaultStatus holds the default value on creation for the status field.
	userauth.DefaultStatus = userauthDescStatus.Default.(uint8)
	// userauthDescAccount is the schema descriptor for account field.
	userauthDescAccount := userauthFields[2].Descriptor()
	// userauth.AccountValidator is a validator for the "account" field. It is called by the builders before save.
	userauth.AccountValidator = userauthDescAccount.Validators[0].(func(string) error)
}
