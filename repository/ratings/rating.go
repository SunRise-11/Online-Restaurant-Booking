package ratings

import (
	"Restobook/entities"

	"gorm.io/gorm"
)

type RatingRepository struct {
	db *gorm.DB
}

func NewRatingsRepo(db *gorm.DB) *RatingRepository {
	return &RatingRepository{db: db}
}

func (rr RatingRepository) Create(rating entities.Rating) (entities.Rating, error) {
	if err := rr.db.Create(&rating).Error; err != nil {
		return rating, err
	}

	var r entities.Rating

	rr.db.Preload("User").First(&r, "user_id = ? AND restaurant_detail_id = ?", &rating.UserID, &rating.RestaurantDetailID)

	return r, nil
}

func (rr RatingRepository) IsCanGiveRating(userId, restaurantId int) (bool, error) {
	var transaction entities.Transaction

	const SUCCESS_STATUS = "SUCCESS"

	if err := rr.db.Where("user_id = ? AND restaurant_id = ? AND status = ?", userId, restaurantId, SUCCESS_STATUS).First(&transaction).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (rr *RatingRepository) Update(rating entities.Rating) (entities.Rating, error) {
	var r entities.Rating

	if err := rr.db.First(&r, "user_id = ? AND restaurant_detail_id = ?", rating.UserID, rating.RestaurantDetailID).Error; err != nil {
		return r, err
	}

	rr.db.Model(&r).Updates(rating)

	rr.db.Preload("User").First(&r, "user_id = ? AND restaurant_detail_id = ?", &rating.UserID, &rating.RestaurantDetailID)

	return r, nil
}

func (rr *RatingRepository) Delete(userId, restaurantId int) (entities.Rating, error) {
	rating := entities.Rating{}

	if err := rr.db.First(&rating, "user_id = ? AND restaurant_detail_id = ?", userId, restaurantId).Error; err != nil {
		return rating, err
	}

	rr.db.Delete(&rating)

	return rating, nil
}
