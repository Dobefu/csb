package utils

import (
	"testing"

	"github.com/Dobefu/csb/cmd/database/structs"
	"github.com/Dobefu/csb/cmd/init_env"
	"github.com/stretchr/testify/assert"
)

func TestConstructWhere(t *testing.T) {
	init_env.Main("../../../.env.test")

	where, args := ConstructWhere([]structs.QueryWhere{
		{
			Name:     "test",
			Value:    "test",
			Operator: structs.EQUALS,
		},
	})

	assert.Equal(t, where, "WHERE test = ?")
	assert.Equal(t, args, []interface{}{"test"})

	where, args = ConstructWhere([]structs.QueryWhere{
		{
			Name:     "test1",
			Value:    "test 1",
			Operator: structs.EQUALS,
		},
		{
			Name:     "test2",
			Value:    "test 2",
			Operator: structs.NOT_EQUALS,
		},
	})

	assert.Equal(t, where, "WHERE test1 = ? AND test2 <> ?")
	assert.Equal(t, args, []interface{}{"test 1", "test 2"})
}
