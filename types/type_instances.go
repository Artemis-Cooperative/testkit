package types

type TypeInstance struct {
	Value any
	Tags  []string
}

var typeInstances = []TypeInstance{
	{
		Value: true,
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"ordered",
		},
	},
	{
		Value: "",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"ordered",
		},
	},
	{
		Value: 0,
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Value: int8(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Value: int16(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Value: int32(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Value: int64(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Value: uint(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Value: uint8(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Value: uint16(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Value: uint32(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Value: uint64(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Value: uintptr(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"ordered",
		},
	},
	{
		Value: float32(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"ordered",
			"numeric",
		},
	},
	{
		Value: float64(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"ordered",
			"numeric",
		},
	},
	{
		Value: complex64(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"numeric",
		},
	},
	{
		Value: complex128(0),
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"numeric",
		},
	},
	{
		Value: []any{},
		Tags: []string{
			"element",
		},
	},
	{
		Value: [0]any{},
		Tags: []string{
			"comparable",
			"element",
			"key",
		},
	},
	{
		Value: map[any]any{},
		Tags: []string{
			"element",
		},
	},
	{
		Value: struct{}{},
		Tags: []string{
			"comparable",
			"element",
			"key",
		},
	},
	{
		Value: make(chan any),
		Tags: []string{
			"comparable",
			"element",
			"key",
		},
	},
	{
		Value: func() {},
		Tags: []string{
			"element",
		},
	},
	{
		Value: (*any)(nil),
		Tags: []string{
			"comparable",
			"element",
			"key",
		},
	},
	{
		Value: nil,
		Tags:  []string{},
	},
}
