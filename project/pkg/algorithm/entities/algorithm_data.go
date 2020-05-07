package entities

type AlgorithmData struct {
	HandleId  int64 // идентификатор ручки
	AvgIncome int64 // средний доход от ручки
	Count     int64 // сколько раз мы дёрнули за ручку
}
