package response

import "main/model/manage"

type GoodsCategoryResponse struct {
	GoodsCategory manage.MallGoodsCategory `json:"mallGoodsCategory"`
}
