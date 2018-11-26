package interpreterNodeRuntime

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/golang/protobuf/proto"
	"github.com/xiaokangwang/Vison/runtime/codestore"
	"github.com/xiaokangwang/Vison/runtime/hints"
	"github.com/xiaokangwang/Vison/runtime/linker"
	"github.com/xiaokangwang/Vison/runtime/node"

	"github.com/xiaokangwang/VisonC/structure/represent"
)

type Runtime struct{
	codestore codestore.CodeStore
	linker linker.Linker
}

type nodeContext struct{
	perquisiteImpair int
}

func (rt*Runtime) executeNode(req node.ExecutionRequest,ctx context.Context){
	executionTarget := req.GetExecutionTarget()
	describe:=rt.codestore.GetObjectByImprint(executionTarget)

	hint := rt.codestore.GetObjectHintByImprint(executionTarget)
	var nodehint hints.NodeHint
	gob.NewDecoder(bytes.NewReader(hint)).Decode(&nodehint)

	var nodedescribe represent.ImplElab
	proto.Unmarshal(describe,&nodedescribe)

	sum:=len(nodedescribe.Exec)

	// Build runtime structure

	nodeContext:=make([]nodeContext,sum)

	//TODO: correct value for context

	for initializing := range nodedescribe.Exec{
		nodeContext[initializing].perquisiteImpair =
			nodehint.InitialPerquisiteImpair[initializing]
	}

	// Foreach EntryNode and decrease Perquisite Impair








}

