package calc

var opers map[string]OperationInterface

func init() {
	opers = make(map[string]OperationInterface, 0)
	opers["+"] = &OperationAdd{}
	opers["-"] = &OperationSub{}
	opers["*"] = &OperationMul{}
	opers["/"] = &OperationDiv{}

}

func OperationFactory(ope string) OperationInterface {
	return opers[ope]
}