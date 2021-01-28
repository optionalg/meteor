// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/degenerat3/meteor/core/ent/bot"
	"github.com/degenerat3/meteor/core/ent/host"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// BotCreate is the builder for creating a Bot entity.
type BotCreate struct {
	config
	mutation *BotMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (bc *BotCreate) SetUUID(s string) *BotCreate {
	bc.mutation.SetUUID(s)
	return bc
}

// SetInterval sets the "interval" field.
func (bc *BotCreate) SetInterval(i int) *BotCreate {
	bc.mutation.SetInterval(i)
	return bc
}

// SetDelta sets the "delta" field.
func (bc *BotCreate) SetDelta(i int) *BotCreate {
	bc.mutation.SetDelta(i)
	return bc
}

// SetLastSeen sets the "lastSeen" field.
func (bc *BotCreate) SetLastSeen(i int) *BotCreate {
	bc.mutation.SetLastSeen(i)
	return bc
}

// SetNillableLastSeen sets the "lastSeen" field if the given value is not nil.
func (bc *BotCreate) SetNillableLastSeen(i *int) *BotCreate {
	if i != nil {
		bc.SetLastSeen(*i)
	}
	return bc
}

// SetInfectingID sets the "infecting" edge to the Host entity by ID.
func (bc *BotCreate) SetInfectingID(id int) *BotCreate {
	bc.mutation.SetInfectingID(id)
	return bc
}

// SetNillableInfectingID sets the "infecting" edge to the Host entity by ID if the given value is not nil.
func (bc *BotCreate) SetNillableInfectingID(id *int) *BotCreate {
	if id != nil {
		bc = bc.SetInfectingID(*id)
	}
	return bc
}

// SetInfecting sets the "infecting" edge to the Host entity.
func (bc *BotCreate) SetInfecting(h *Host) *BotCreate {
	return bc.SetInfectingID(h.ID)
}

// Mutation returns the BotMutation object of the builder.
func (bc *BotCreate) Mutation() *BotMutation {
	return bc.mutation
}

// Save creates the Bot in the database.
func (bc *BotCreate) Save(ctx context.Context) (*Bot, error) {
	var (
		err  error
		node *Bot
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BotCreate) SaveX(ctx context.Context) *Bot {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (bc *BotCreate) defaults() {
	if _, ok := bc.mutation.LastSeen(); !ok {
		v := bot.DefaultLastSeen
		bc.mutation.SetLastSeen(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BotCreate) check() error {
	if _, ok := bc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := bc.mutation.Interval(); !ok {
		return &ValidationError{Name: "interval", err: errors.New("ent: missing required field \"interval\"")}
	}
	if _, ok := bc.mutation.Delta(); !ok {
		return &ValidationError{Name: "delta", err: errors.New("ent: missing required field \"delta\"")}
	}
	if _, ok := bc.mutation.LastSeen(); !ok {
		return &ValidationError{Name: "lastSeen", err: errors.New("ent: missing required field \"lastSeen\"")}
	}
	return nil
}

func (bc *BotCreate) sqlSave(ctx context.Context) (*Bot, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BotCreate) createSpec() (*Bot, *sqlgraph.CreateSpec) {
	var (
		_node = &Bot{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bot.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bot.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bot.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := bc.mutation.Interval(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldInterval,
		})
		_node.Interval = value
	}
	if value, ok := bc.mutation.Delta(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldDelta,
		})
		_node.Delta = value
	}
	if value, ok := bc.mutation.LastSeen(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bot.FieldLastSeen,
		})
		_node.LastSeen = value
	}
	if nodes := bc.mutation.InfectingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bot.InfectingTable,
			Columns: []string{bot.InfectingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BotCreateBulk is the builder for creating many Bot entities in bulk.
type BotCreateBulk struct {
	config
	builders []*BotCreate
}

// Save creates the Bot entities in the database.
func (bcb *BotCreateBulk) Save(ctx context.Context) ([]*Bot, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Bot, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BotMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BotCreateBulk) SaveX(ctx context.Context) []*Bot {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}