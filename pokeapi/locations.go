package pokeapi

func GetLocations(qry string) Location {

	var locations Location

	do("location?"+qry, &locations)
	return locations

}
