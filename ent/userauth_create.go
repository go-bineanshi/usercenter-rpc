// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-bineanshi/usercenter-rpc/ent/property"
	"github.com/go-bineanshi/usercenter-rpc/ent/user"
	"github.com/go-bineanshi/usercenter-rpc/ent/userauth"
)

// UserAuthCreate is the builder for creating a UserAuth entity.
type UserAuthCreate struct {
	config
	mutation *UserAuthMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uac *UserAuthCreate) SetCreatedAt(t time.Time) *UserAuthCreate {
	uac.mutation.SetCreatedAt(t)
	return uac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uac *UserAuthCreate) SetNillableCreatedAt(t *time.Time) *UserAuthCreate {
	if t != nil {
		uac.SetCreatedAt(*t)
	}
	return uac
}

// SetUpdatedAt sets the "updated_at" field.
func (uac *UserAuthCreate) SetUpdatedAt(t time.Time) *UserAuthCreate {
	uac.mutation.SetUpdatedAt(t)
	return uac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uac *UserAuthCreate) SetNillableUpdatedAt(t *time.Time) *UserAuthCreate {
	if t != nil {
		uac.SetUpdatedAt(*t)
	}
	return uac
}

// SetStatus sets the "status" field.
func (uac *UserAuthCreate) SetStatus(u uint8) *UserAuthCreate {
	uac.mutation.SetStatus(u)
	return uac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uac *UserAuthCreate) SetNillableStatus(u *uint8) *UserAuthCreate {
	if u != nil {
		uac.SetStatus(*u)
	}
	return uac
}

// SetUserID sets the "user_id" field.
func (uac *UserAuthCreate) SetUserID(u uint64) *UserAuthCreate {
	uac.mutation.SetUserID(u)
	return uac
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uac *UserAuthCreate) SetNillableUserID(u *uint64) *UserAuthCreate {
	if u != nil {
		uac.SetUserID(*u)
	}
	return uac
}

// SetProvider sets the "provider" field.
func (uac *UserAuthCreate) SetProvider(pr property.Provider) *UserAuthCreate {
	uac.mutation.SetProvider(pr)
	return uac
}

// SetAccount sets the "account" field.
func (uac *UserAuthCreate) SetAccount(s string) *UserAuthCreate {
	uac.mutation.SetAccount(s)
	return uac
}

// SetID sets the "id" field.
func (uac *UserAuthCreate) SetID(u uint64) *UserAuthCreate {
	uac.mutation.SetID(u)
	return uac
}

// SetUser sets the "user" edge to the User entity.
func (uac *UserAuthCreate) SetUser(u *User) *UserAuthCreate {
	return uac.SetUserID(u.ID)
}

// Mutation returns the UserAuthMutation object of the builder.
func (uac *UserAuthCreate) Mutation() *UserAuthMutation {
	return uac.mutation
}

// Save creates the UserAuth in the database.
func (uac *UserAuthCreate) Save(ctx context.Context) (*UserAuth, error) {
	uac.defaults()
	return withHooks(ctx, uac.sqlSave, uac.mutation, uac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uac *UserAuthCreate) SaveX(ctx context.Context) *UserAuth {
	v, err := uac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uac *UserAuthCreate) Exec(ctx context.Context) error {
	_, err := uac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uac *UserAuthCreate) ExecX(ctx context.Context) {
	if err := uac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uac *UserAuthCreate) defaults() {
	if _, ok := uac.mutation.CreatedAt(); !ok {
		v := userauth.DefaultCreatedAt()
		uac.mutation.SetCreatedAt(v)
	}
	if _, ok := uac.mutation.UpdatedAt(); !ok {
		v := userauth.DefaultUpdatedAt()
		uac.mutation.SetUpdatedAt(v)
	}
	if _, ok := uac.mutation.Status(); !ok {
		v := userauth.DefaultStatus
		uac.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uac *UserAuthCreate) check() error {
	if _, ok := uac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserAuth.created_at"`)}
	}
	if _, ok := uac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserAuth.updated_at"`)}
	}
	if _, ok := uac.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`ent: missing required field "UserAuth.provider"`)}
	}
	if v, ok := uac.mutation.Provider(); ok {
		if err := userauth.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf(`ent: validator failed for field "UserAuth.provider": %w`, err)}
		}
	}
	if _, ok := uac.mutation.Account(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required field "UserAuth.account"`)}
	}
	if v, ok := uac.mutation.Account(); ok {
		if err := userauth.AccountValidator(v); err != nil {
			return &ValidationError{Name: "account", err: fmt.Errorf(`ent: validator failed for field "UserAuth.account": %w`, err)}
		}
	}
	return nil
}

func (uac *UserAuthCreate) sqlSave(ctx context.Context) (*UserAuth, error) {
	if err := uac.check(); err != nil {
		return nil, err
	}
	_node, _spec := uac.createSpec()
	if err := sqlgraph.CreateNode(ctx, uac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	uac.mutation.id = &_node.ID
	uac.mutation.done = true
	return _node, nil
}

func (uac *UserAuthCreate) createSpec() (*UserAuth, *sqlgraph.CreateSpec) {
	var (
		_node = &UserAuth{config: uac.config}
		_spec = sqlgraph.NewCreateSpec(userauth.Table, sqlgraph.NewFieldSpec(userauth.FieldID, field.TypeUint64))
	)
	if id, ok := uac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uac.mutation.CreatedAt(); ok {
		_spec.SetField(userauth.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uac.mutation.UpdatedAt(); ok {
		_spec.SetField(userauth.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uac.mutation.Status(); ok {
		_spec.SetField(userauth.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := uac.mutation.Provider(); ok {
		_spec.SetField(userauth.FieldProvider, field.TypeEnum, value)
		_node.Provider = value
	}
	if value, ok := uac.mutation.Account(); ok {
		_spec.SetField(userauth.FieldAccount, field.TypeString, value)
		_node.Account = value
	}
	if nodes := uac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userauth.UserTable,
			Columns: []string{userauth.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserAuthCreateBulk is the builder for creating many UserAuth entities in bulk.
type UserAuthCreateBulk struct {
	config
	builders []*UserAuthCreate
}

// Save creates the UserAuth entities in the database.
func (uacb *UserAuthCreateBulk) Save(ctx context.Context) ([]*UserAuth, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uacb.builders))
	nodes := make([]*UserAuth, len(uacb.builders))
	mutators := make([]Mutator, len(uacb.builders))
	for i := range uacb.builders {
		func(i int, root context.Context) {
			builder := uacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserAuthMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uacb *UserAuthCreateBulk) SaveX(ctx context.Context) []*UserAuth {
	v, err := uacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uacb *UserAuthCreateBulk) Exec(ctx context.Context) error {
	_, err := uacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uacb *UserAuthCreateBulk) ExecX(ctx context.Context) {
	if err := uacb.Exec(ctx); err != nil {
		panic(err)
	}
}
