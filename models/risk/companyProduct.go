package risk

type CompanyProduct struct {
	ProductId int `json:"productId" form:"productId"`
	ProductName string `json:"productName" form:"productName"`
	ProductShortName string `json:"productShortName" form:"productShortName"`
	ProductCode string `json:"productCode" form:"productCode"`
	BrandId int `json:"brandId" form:"brandId"`
	BrandName string `json:"brandName" form:"brandName"`
	BrandShortName string `json:"brandShortName" form:"brandShortName"`
	ContractProperty int `json:"contractProperty" form:"contractProperty"`
	IsRecommend int `json:"isRecommend" form:"isRecommend"`
	SaleStatus int `json:"saleStatus" form:"saleStatus"`
}
