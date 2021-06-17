package service

import (
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/model"
	"sync"
)

type CategoryService struct {
	wg sync.WaitGroup
}

func (receiver *CategoryService) GetCategories() ([]*pb.Category, error) {

	categories, err := model.GetCategories("1 = 1")
	if err != nil {
		return nil, err
	}
	var categorySlice []*pb.Category
	for _, category := range categories {
		count, err := model.GetArticlesCountByCondition("category_id = ?", category.ID)
		logging.Logger.Debug(count)
		if err != nil {
			return nil, err
		}
		c := &pb.Category{
			Id:           int32(category.ID),
			CategoryName: category.CategoryName,
			ArticleCount: int32(count),
		}
		categorySlice = append(categorySlice, c)
	}

	return categorySlice, err

}