package ratings

import "Restobook/entities"

type RatingsInterface interface {
	Create(entities.Rating) (entities.Rating, error)
	Update(entities.Rating) (entities.Rating, error)
	Delete(userId, restaurantId int) (entities.Rating, error)
	IsCanGiveRating(userId, restaurantId int) (bool, error)
}
