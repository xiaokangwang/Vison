package node

import (
	"context"
	"github.com/xiaokangwang/Vison/runtime/value"
)

type ExecutionRequest interface{
	GetExecutionTarget() []byte
	GetInputValues() []value.Value
}

type Runtime interface{
	Evaluate(req ExecutionRequest,ctx context.Context) ResultPromise
}

type ResultPromise interface{
	IsReady() bool
	Process() float32
	GetResultValuePromise(position int)ResultValuePromise
}

type ResultValuePromise interface{
	IsReady() bool
	GetValue() value.Value
}