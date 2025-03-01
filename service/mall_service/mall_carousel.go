package mall_service

import (
	"main/global"
	"main/model/mall/response"
	"main/model/manage"
)

type MallCarouselService struct {
}

// GetCarouselsForIndex 返回固定数量的轮播图对象(首页调用)
func (m *MallCarouselService) GetCarouselsForIndex(num int) (err error, mallCarousels []manage.MallCarousel, list interface{}) {
	var carouselIndexs []response.MallCarouselIndexResponse
	err = global.GVA_DB.Where("is_deleted = 0").Order("carousel_rank desc").Limit(num).Find(&mallCarousels).Error
	for _, carousel := range mallCarousels {
		carouselIndex := response.MallCarouselIndexResponse{
			CarouselUrl: carousel.CarouselUrl,
			RedirectUrl: carousel.RedirectUrl,
		}
		carouselIndexs = append(carouselIndexs, carouselIndex)
	}
	return err, mallCarousels, carouselIndexs
}
