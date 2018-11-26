package linker

import (
	"context"
)

type Linker interface{
	LinkBlueprint(linkTo []byte,inputTraits [][]byte, specification []byte,ctx context.Context)
}
