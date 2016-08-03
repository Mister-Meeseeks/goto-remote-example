package "zoo"

import "food"

func PickRightCat() {
	if (food.PoundsOfMeat() > 25) {
		Lion();
	} else {
		Kitty();
	}	
}
