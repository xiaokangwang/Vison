package value

type Value interface {
	ListTraits() [][]bytes
	GetTraitInstance(imprint []bytes)TraitInstance
}


type TraitInstance interface {
	ListProp()[][]byte
	GetProp(imprint []byte)Value
}

