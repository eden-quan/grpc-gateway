// Package generator provides an abstract interface to code generators.
package generator

import (
	"gitlab.lainuoniao.cn/rhinobird/backend/grpc-gateway.git/internal/descriptor"
)

// Generator is an abstraction of code generators.
type Generator interface {
	// Generate generates output files from input .proto files.
	Generate(targets []*descriptor.File) ([]*descriptor.ResponseFile, error)
}
