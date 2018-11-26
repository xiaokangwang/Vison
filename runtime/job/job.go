package job

type Runtime interface{
	Invoke(target []byte)
}
