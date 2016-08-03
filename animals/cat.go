package animals

import "food"

func PickRightCat() string {
	if (food.PoundsOfMeat() > 25) {
		return Lion();
	} else {
		return Kitty();
	}	
}
