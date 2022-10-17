package domain

type ListParameters struct {
	Sorting SortParams
	Offset  uint64
	Limit   uint64
}

type SortParams struct {
	Column string
	Direct string
}
