package core

type TargetResolver interface {
	Parse(target *TargetStruct)
}
