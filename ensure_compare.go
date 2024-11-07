package ensure

type EnsureType interface {
	ExactType(have, want any)
}

type EnsureCompare interface {
	Equal(have, want any)
	NotEqual(have, want any)
	Nil(have any)
	NotNil(have any)
	Empty(have any)
	NotEmpty(have any)
}

type EnsureList interface{}

type EnsureMap interface{}
