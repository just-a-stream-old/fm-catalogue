package model

type Exchange struct {
	ExchangeCode 	string		`json:"exchangeCode" bson:"exchangeCode"`
	Name 			string		`json:"name" bson:"name"`
	MarketIdCode	string		`json:"marketIdCode" bson:"marketIdCode"`
	Country 		string		`json:"country" bson:"country"`
	Currency 		string		`json:"currency" bson:"currency"`
	FinanceAPIs		[]string	`json:"financeApis" bson:"financeApis"`
}