package field_test

import (
	"fbc/ent/field"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	f := field.Int("age").Positive()
	assert.Equal(t, "age", f.Name())
	assert.Equal(t, field.TypeInt, f.Type())
	assert.Len(t, f.Validators(), 1)

	f = field.Int("age").Default(10).Min(10).Max(20)
	assert.True(t, f.HasDefault())
	assert.Equal(t, 10, f.Value())
	assert.Len(t, f.Validators(), 2)

	f = field.Int("age").Range(20, 40).Nullable()
	assert.False(t, f.HasDefault())
	assert.True(t, f.IsNullable())
	assert.Len(t, f.Validators(), 1)
}

func TestFloat(t *testing.T) {
	f := field.Float("age").Positive()
	assert.Equal(t, "age", f.Name())
	assert.Equal(t, field.TypeFloat64, f.Type())
	assert.Len(t, f.Validators(), 1)

	f = field.Float("age").Min(2.5).Max(5)
	assert.Len(t, f.Validators(), 2)
}

func TestBool(t *testing.T) {
	f := field.Bool("active").Default(true)
	assert.Equal(t, "active", f.Name())
	assert.Equal(t, field.TypeBool, f.Type())
	assert.True(t, f.HasDefault())
	assert.Equal(t, true, f.Value())
}

func TestString(t *testing.T) {
	re := regexp.MustCompile("[a-zA-Z0-9]")
	f := field.String("name").Unique().Match(re).Validate(func(string) error { return nil })
	assert.Equal(t, field.TypeString, f.Type())
	assert.Equal(t, "name", f.Name())
	assert.True(t, f.IsUnique())
	assert.Len(t, f.Validators(), 2)
}

func TestTime(t *testing.T) {
	f := field.Time("created_at")
	assert.Equal(t, "created_at", f.Name())
	assert.Equal(t, field.TypeTime, f.Type())
	assert.Equal(t, "time.Time", f.Type().String())
}

func TestField_Tag(t *testing.T) {
	f := field.Bool("expired").StructTag(`json:"expired,omitempty"`)
	require.Equal(t, `json:"expired,omitempty"`, f.Tag())
}