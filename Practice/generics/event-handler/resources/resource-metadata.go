package resources

import "fmt"

type ObjectMetadata struct {
	Name      string
	Namespace NamespaceName
	Labels    map[string]string
}

type LoggableAndValidatable interface {
	Log() string
	IsValid() (bool, error)
}

type NamespaceName string

const (
	StarOps  NamespaceName = "starops"
	Hitzler  NamespaceName = "hitzler"
	Fichtner NamespaceName = "fichtner"
)

func (n NamespaceName) IsValid() (bool, error) {
	switch n {
	case StarOps, Hitzler, Fichtner:
		return true, nil
	}
	return false, fmt.Errorf("The namespace name %s is not valid", n)
}
