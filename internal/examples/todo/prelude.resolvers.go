package todo

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"
	"todo/ent"

	"github.com/99designs/gqlgen/graphql/introspection"
)

// Description is the resolver for the description field.
func (r *__FieldResolver) Description(ctx context.Context, obj *ent.Field) (*string, error) {
	panic(fmt.Errorf("not implemented: Description - description"))
}

// Args is the resolver for the args field.
func (r *__FieldResolver) Args(ctx context.Context, obj *ent.Field) ([]*introspection.InputValue, error) {
	panic(fmt.Errorf("not implemented: Args - args"))
}

// Type is the resolver for the type field.
func (r *__FieldResolver) Type(ctx context.Context, obj *ent.Field) (*introspection.Type, error) {
	panic(fmt.Errorf("not implemented: Type - type"))
}

// IsDeprecated is the resolver for the isDeprecated field.
func (r *__FieldResolver) IsDeprecated(ctx context.Context, obj *ent.Field) (bool, error) {
	panic(fmt.Errorf("not implemented: IsDeprecated - isDeprecated"))
}

// DeprecationReason is the resolver for the deprecationReason field.
func (r *__FieldResolver) DeprecationReason(ctx context.Context, obj *ent.Field) (*string, error) {
	panic(fmt.Errorf("not implemented: DeprecationReason - deprecationReason"))
}

// Fields is the resolver for the fields field.
func (r *__TypeResolver) Fields(ctx context.Context, obj *introspection.Type, includeDeprecated bool) ([]*ent.Field, error) {
	panic(fmt.Errorf("not implemented: Fields - fields"))
}

// __Field returns __FieldResolver implementation.
func (r *Resolver) __Field() __FieldResolver { return &__FieldResolver{r} }

// __Type returns __TypeResolver implementation.
func (r *Resolver) __Type() __TypeResolver { return &__TypeResolver{r} }

type __FieldResolver struct{ *Resolver }
type __TypeResolver struct{ *Resolver }
