package filter

type FilterOption func(*GenericFilter)

func WithIncludeKinds(includeKinds []string) FilterOption {
	return func(o *GenericFilter) {
		o.IncludeKinds = includeKinds
	}
}

func WithIncludeKindsStrict(includeKindsStrict bool) FilterOption {
	return func(o *GenericFilter) {
		o.IncludeKindsStrict = includeKindsStrict
	}
}

func WithExcludeKinds(excludeKinds []string) FilterOption {
	return func(o *GenericFilter) {
		o.ExcludeKinds = excludeKinds
	}
}

func WithExcludeKindsStrict(excludeKindsStrict bool) FilterOption {
	return func(o *GenericFilter) {
		o.ExcludeKindsStrict = excludeKindsStrict
	}
}

func WithIncludeApiVersions(includeApiVersions []string) FilterOption {
	return func(o *GenericFilter) {
		o.IncludeApiVersions = includeApiVersions
	}
}

func WithIncludeApiVersionsStrict(includeApiVersionsStrict bool) FilterOption {
	return func(o *GenericFilter) {
		o.IncludeApiVersionsStrict = includeApiVersionsStrict
	}
}

func WithExcludeApiVersions(excludeApiVersions []string) FilterOption {
	return func(o *GenericFilter) {
		o.ExcludeApiVersions = excludeApiVersions
	}
}

func WithExcludeApiVersionsStrict(excludeApiVersionsStrict bool) FilterOption {
	return func(o *GenericFilter) {
		o.ExcludeApiVersionsStrict = excludeApiVersionsStrict
	}
}

func WithIncludeNames(includeNames []string) FilterOption {
	return func(o *GenericFilter) {
		o.IncludeNames = includeNames
	}
}

func WithIncludeNamesStrict(includeNamesStrict bool) FilterOption {
	return func(o *GenericFilter) {
		o.IncludeNamesStrict = includeNamesStrict
	}
}

func WithExcludeNames(excludeNames []string) FilterOption {
	return func(o *GenericFilter) {
		o.ExcludeNames = excludeNames
	}
}

func WithExcludeNamesStrict(excludeNamesStrict bool) FilterOption {
	return func(o *GenericFilter) {
		o.ExcludeNamesStrict = excludeNamesStrict
	}
}

func WithIncludeUnrecognized(includeUnrecognized bool) FilterOption {
	return func(o *GenericFilter) {
		o.IncludeUnrecognized = includeUnrecognized
	}
}

func NewGenericFilter(opts ...FilterOption) *GenericFilter {
	options := &GenericFilter{}

	for _, opt := range opts {
		opt(options)
	}

	return options
}
