package locations

type Country struct {
	Id                 string         `json:id`
	Name               string         `json:name`
	Locate             string         `json:locate`
	CurrencyId         string         `json:currency_id`
	DecimalSeparator   string         `json:decimal_separator`
	ThousandsSeparator string         `json:thousands_separator`
	TimeZone           string         `json:time_zone`
	GeoInformation     GeoInformation `json:geo_information`
	States             []State        `json:states`
}

type GeoInformation struct {
	Location GeoLocation `json:location`
}

type GeoLocation struct {
	Latitude  float64 `json:latitude`
	Longitude float64 `json:longitude`
}

type State struct {
	Id   string `json:id`
	Name string `json:name`
}
