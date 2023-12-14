package sqlorder

type SQLOrder struct {
	Column string
	By     SQLOrderEnum
}

const SQLOrderASC = "ASC"
const SQLOrderDESC = "DESC"

func (s SQLOrder) IsValid() bool {
	return s.Column != "" && s.By.IsValid()
}

type SQLOrderEnum int

const (
	SQLOrderEnumASC SQLOrderEnum = iota
	SQLOrderEnumDESC
)

func (e SQLOrderEnum) String() string {
	return [...]string{SQLOrderASC, SQLOrderDESC}[e]
}

func (e SQLOrderEnum) IsValid() bool {
	return e == SQLOrderEnumASC || e == SQLOrderEnumDESC
}
