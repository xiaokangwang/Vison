package codestore

type CodeStore interface {
	GetObjectByImprint([]byte)[]byte
	GetObjectHintByImprint([]byte)[]byte
}
