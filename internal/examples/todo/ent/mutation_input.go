// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"todo/ent/todo"
)

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	Text      string
	CreatedAt *time.Time
	Status    *todo.Status
	Priority  *int
	Children  []int
	Parent    *int
	Owner     *int
}

// Mutate applies the CreateTodoInput on the TodoCreate builder.
func (i *CreateTodoInput) Mutate(m *TodoCreate) {
	m.SetText(i.Text)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if ids := i.Children; len(ids) > 0 {
		m.AddChildIDs(ids...)
	}
	if v := i.Parent; v != nil {
		m.SetParentID(*v)
	}
	if v := i.Owner; v != nil {
		m.SetOwnerID(*v)
	}
}

// SetInput applies the change-set in the CreateTodoInput on the create builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c)
	return c
}

// UpdateTodoInput represents a mutation input for updating todos.
type UpdateTodoInput struct {
	Text           *string
	Status         *todo.Status
	Priority       *int
	AddChildIDs    []int
	RemoveChildIDs []int
	Parent         *int
	ClearParent    bool
	Owner          *int
	ClearOwner     bool
}

// Mutate applies the UpdateTodoInput on the TodoMutation.
func (i *UpdateTodoInput) Mutate(m *TodoMutation) {
	if v := i.Text; v != nil {
		m.SetText(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if ids := i.AddChildIDs; len(ids) > 0 {
		m.AddChildIDs(ids...)
	}
	if ids := i.RemoveChildIDs; len(ids) > 0 {
		m.RemoveChildIDs(ids...)
	}
	if i.ClearParent {
		m.ClearParent()
	}
	if v := i.Parent; v != nil {
		m.SetParentID(*v)
	}
	if i.ClearOwner {
		m.ClearOwner()
	}
	if v := i.Owner; v != nil {
		m.SetOwnerID(*v)
	}
}

// SetInput applies the change-set in the UpdateTodoInput on the update builder.
func (u *TodoUpdate) SetInput(i UpdateTodoInput) *TodoUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateTodoInput on the update-one builder.
func (u *TodoUpdateOne) SetInput(i UpdateTodoInput) *TodoUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Name  string
	Todos []int
}

// Mutate applies the CreateUserInput on the UserCreate builder.
func (i *CreateUserInput) Mutate(m *UserCreate) {
	m.SetName(i.Name)
	if ids := i.Todos; len(ids) > 0 {
		m.AddTodoIDs(ids...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the create builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c)
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Name          *string
	AddTodoIDs    []int
	RemoveTodoIDs []int
}

// Mutate applies the UpdateUserInput on the UserMutation.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if ids := i.AddTodoIDs; len(ids) > 0 {
		m.AddTodoIDs(ids...)
	}
	if ids := i.RemoveTodoIDs; len(ids) > 0 {
		m.RemoveTodoIDs(ids...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the update builder.
func (u *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateUserInput on the update-one builder.
func (u *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(u.Mutation())
	return u
}
