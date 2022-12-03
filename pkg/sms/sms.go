package sms

import (
	"io/ioutil"
	"sort"
	"strings"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func checkCountry(country string, listCountry map[string]string) bool {
	for val := range listCountry {
		if country == val {
			return true
		}
	}
	return false
}

func checkProvider(provider string) bool {
	providers := []string{
		"Topolo",
		"Rond",
		"Kildy"}
	for _, val := range providers {
		if provider == val {
			return true
		}
	}
	return false
}

func GetSmsData(countries map[string]string) []SMSData {
	var result []SMSData
	content, err := ioutil.ReadFile("sms.data")
	if err != nil {
		panic(err)
	}
	splitContent := strings.Split(string(content), "\n")
	for _, i := range splitContent {
		if len(strings.Split(i, ";")) != 4 {
			continue
		} else {
			if !checkProvider(strings.Split(i, ";")[3]) {
				continue
			} else {
				res := checkCountry(strings.Split(i, ";")[0], countries)
				if res {
					smsData := SMSData{
						Country:      strings.Split(i, ";")[0],
						Bandwidth:    strings.Split(i, ";")[1],
						ResponseTime: strings.Split(i, ";")[2],
						Provider:     strings.Split(i, ";")[3],
					}
					result = append(result, smsData)
				}
			}
		}
	}
	return result
}

func GetResultSms() [][]SMSData {
	var result [][]SMSData
	countries := map[string]string{
		"AF": "Afghanistan",
		"AL": "Albania, People's Socialist Republic of",
		"DZ": "Algeria, People's Democratic Republic of",
		"AS": "American Samoa",
		"AD": "Andorra, Principality of",
		"AO": "Angola, Republic of",
		"AI": "Anguilla",
		"AQ": "Antarctica (the territory South of 60 deg S)",
		"AG": "Antigua and Barbuda",
		"AR": "Argentina, Argentine Republic",
		"AM": "Armenia",
		"AW": "Aruba",
		"AU": "Australia, Commonwealth of",
		"AT": "Austria, Republic of",
		"AZ": "Azerbaijan, Republic of",
		"BS": "Bahamas, Commonwealth of the",
		"BH": "Bahrain, Kingdom of",
		"BD": "Bangladesh, People's Republic of",
		"BB": "Barbados",
		"BY": "Belarus",
		"BE": "Belgium, Kingdom of",
		"BZ": "Belize",
		"BJ": "Benin (was Dahomey), People's Republic of",
		"BM": "Bermuda",
		"BT": "Bhutan, Kingdom of",
		"BO": "Bolivia, Republic of",
		"BA": "Bosnia and Herzegovina",
		"BW": "Botswana, Republic of",
		"BV": "Bouvet Island (Bouvetoya)",
		"BR": "Brazil, Federative Republic of",
		"IO": "British Indian Ocean Territory (Chagos Archipelago)",
		"VG": "British Virgin Islands",
		"BN": "Brunei Darussalam",
		"BG": "Bulgaria, People's Republic of",
		"BF": "Burkina Faso (was Upper Volta)",
		"BI": "Burundi, Republic of",
		"KH": "Cambodia, Kingdom of (was Khmer Republic/Kampuchea, Democratic)",
		"CM": "Cameroon, United Republic of",
		"CA": "Canada",
		"CV": "Cape Verde, Republic of",
		"KY": "Cayman Islands",
		"CF": "Central African Republic",
		"TD": "Chad, Republic of",
		"CL": "Chile, Republic of",
		"CN": "China, People's Republic of",
		"CX": "Christmas Island",
		"CC": "Cocos (Keeling) Islands",
		"CO": "Colombia, Republic of",
		"KM": "Comoros, Union of the",
		"CD": "Congo, Democratic Republic of (was Zaire)",
		"CG": "Congo, People's Republic of",
		"CK": "Cook Islands",
		"CR": "Costa Rica, Republic of",
		"CI": "Cote D'Ivoire, Ivory Coast, Republic of the",
		"CU": "Cuba, Republic of",
		"CY": "Cyprus, Republic of",
		"CZ": "Czech Republic",
		"DK": "Denmark, Kingdom of",
		"DJ": "Djibouti, Republic of (was French Afars and Issas)",
		"DM": "Dominica, Commonwealth of",
		"DO": "Dominican Republic",
		"EC": "Ecuador, Republic of",
		"EG": "Egypt, Arab Republic of",
		"SV": "El Salvador, Republic of",
		"GQ": "Equatorial Guinea, Republic of",
		"ER": "Eritrea",
		"EE": "Estonia",
		"ET": "Ethiopia",
		"FO": "Faeroe Islands",
		"FK": "Falkland Islands (Malvinas)",
		"FJ": "Fiji, Republic of the Fiji Islands",
		"FI": "Finland, Republic of",
		"FR": "France, French Republic",
		"GF": "French Guiana",
		"PF": "French Polynesia",
		"TF": "French Southern Territories",
		"GA": "Gabon, Gabonese Republic",
		"GM": "Gambia, Republic of the",
		"GE": "Georgia",
		"DE": "Germany",
		"GH": "Ghana, Republic of",
		"GI": "Gibraltar",
		"GR": "Greece, Hellenic Republic",
		"GL": "Greenland",
		"GD": "Grenada",
		"GP": "Guadaloupe",
		"GU": "Guam",
		"GT": "Guatemala, Republic of",
		"GN": "Guinea, Revolutionary People's Rep'c of",
		"GW": "Guinea-Bissau, Republic of (was Portuguese Guinea)",
		"GY": "Guyana, Republic of",
		"HT": "Haiti, Republic of",
		"HM": "Heard and McDonald Islands",
		"VA": "Holy See (Vatican City State)",
		"HN": "Honduras, Republic of",
		"HK": "Hong Kong, Special Administrative Region of China",
		"HR": "Hrvatska (Croatia)",
		"HU": "Hungary, Hungarian People's Republic",
		"IS": "Iceland, Republic of",
		"IN": "India, Republic of",
		"ID": "Indonesia, Republic of",
		"IR": "Iran, Islamic Republic of",
		"IQ": "Iraq, Republic of",
		"IE": "Ireland",
		"IL": "Israel, State of",
		"IT": "Italy, Italian Republic",
		"JM": "Jamaica",
		"JP": "Japan",
		"JO": "Jordan, Hashemite Kingdom of",
		"KZ": "Kazakhstan, Republic of",
		"KE": "Kenya, Republic of",
		"KI": "Kiribati, Republic of (was Gilbert Islands)",
		"KP": "Korea, Democratic People's Republic of",
		"KR": "Korea, Republic of",
		"KW": "Kuwait, State of",
		"KG": "Kyrgyz Republic",
		"LA": "Lao People's Democratic Republic",
		"LV": "Latvia",
		"LB": "Lebanon, Lebanese Republic",
		"LS": "Lesotho, Kingdom of",
		"LR": "Liberia, Republic of",
		"LY": "Libyan Arab Jamahiriya",
		"LI": "Liechtenstein, Principality of",
		"LT": "Lithuania",
		"LU": "Luxembourg, Grand Duchy of",
		"MO": "Macao, Special Administrative Region of China",
		"MK": "Macedonia, the former Yugoslav Republic of",
		"MG": "Madagascar, Republic of",
		"MW": "Malawi, Republic of",
		"MY": "Malaysia",
		"MV": "Maldives, Republic of",
		"ML": "Mali, Republic of",
		"MT": "Malta, Republic of",
		"MH": "Marshall Islands",
		"MQ": "Martinique",
		"MR": "Mauritania, Islamic Republic of",
		"MU": "Mauritius",
		"YT": "Mayotte",
		"MX": "Mexico, United Mexican States",
		"FM": "Micronesia, Federated States of",
		"MD": "Moldova, Republic of",
		"MC": "Monaco, Principality of",
		"MN": "Mongolia, Mongolian People's Republic",
		"MS": "Montserrat",
		"MA": "Morocco, Kingdom of",
		"MZ": "Mozambique, People's Republic of",
		"MM": "Myanmar (was Burma)",
		"NA": "Namibia",
		"NR": "Nauru, Republic of",
		"NP": "Nepal, Kingdom of",
		"AN": "Netherlands Antilles",
		"NL": "Netherlands, Kingdom of the",
		"NC": "New Caledonia",
		"NZ": "New Zealand",
		"NI": "Nicaragua, Republic of",
		"NE": "Niger, Republic of the",
		"NG": "Nigeria, Federal Republic of",
		"NU": "Niue, Republic of",
		"NF": "Norfolk Island",
		"MP": "Northern Mariana Islands",
		"NO": "Norway, Kingdom of",
		"OM": "Oman, Sultanate of (was Muscat and Oman)",
		"PK": "Pakistan, Islamic Republic of",
		"PW": "Palau",
		"PS": "Palestinian Territory, Occupied",
		"PA": "Panama, Republic of",
		"PG": "Papua New Guinea",
		"PY": "Paraguay, Republic of",
		"PE": "Peru, Republic of",
		"PH": "Philippines, Republic of the",
		"PN": "Pitcairn Island",
		"PL": "Poland, Polish People's Republic",
		"PT": "Portugal, Portuguese Republic",
		"PR": "Puerto Rico",
		"QA": "Qatar, State of",
		"RE": "Reunion",
		"RO": "Romania, Socialist Republic of",
		"RU": "Russian Federation",
		"RW": "Rwanda, Rwandese Republic",
		"SH": "St. Helena",
		"KN": "St. Kitts and Nevis",
		"LC": "St. Lucia",
		"PM": "St. Pierre and Miquelon",
		"VC": "St. Vincent and the Grenadines",
		"WS": "Samoa, Independent State of (was Western Samoa)",
		"SM": "San Marino, Republic of",
		"ST": "Sao Tome and Principe, Democratic Republic of",
		"SA": "Saudi Arabia, Kingdom of",
		"SN": "Senegal, Republic of",
		"CS": "Serbia and Montenegro",
		"SC": "Seychelles, Republic of",
		"SL": "Sierra Leone, Republic of",
		"SG": "Singapore, Republic of",
		"SK": "Slovakia (Slovak Republic)",
		"SI": "Slovenia",
		"SB": "Solomon Islands (was British Solomon Islands)",
		"SO": "Somalia, Somali Republic",
		"ZA": "South Africa, Republic of",
		"GS": "South Georgia and the South Sandwich Islands",
		"ES": "Spain, Spanish State",
		"LK": "Sri Lanka, Democratic Socialist Republic of (was Ceylon)",
		"SD": "Sudan, Democratic Republic of the",
		"SR": "Suriname, Republic of",
		"SJ": "Svalbard & Jan Mayen Islands",
		"SZ": "Swaziland, Kingdom of",
		"SE": "Sweden, Kingdom of",
		"CH": "Switzerland, Swiss Confederation",
		"SY": "Syrian Arab Republic",
		"TW": "Taiwan, Province of China",
		"TJ": "Tajikistan",
		"TZ": "Tanzania, United Republic of",
		"TH": "Thailand, Kingdom of",
		"TL": "Timor-Leste, Democratic Republic of",
		"TG": "Togo, Togolese Republic",
		"TK": "Tokelau (Tokelau Islands)",
		"TO": "Tonga, Kingdom of",
		"TT": "Trinidad and Tobago, Republic of",
		"TN": "Tunisia, Republic of",
		"TR": "Turkey, Republic of",
		"TM": "Turkmenistan",
		"TC": "Turks and Caicos Islands",
		"TV": "Tuvalu (was part of Gilbert & Ellice Islands)",
		"VI": "US Virgin Islands",
		"UG": "Uganda, Republic of",
		"UA": "Ukraine",
		"AE": "United Arab Emirates (was Trucial States)",
		"GB": "United Kingdom of Great Britain & N. Ireland",
		"UM": "United States Minor Outlying Islands",
		"US": "United States of America",
		"UY": "Uruguay, Eastern Republic of",
		"UZ": "Uzbekistan",
		"VU": "Vanuatu (was New Hebrides)",
		"VE": "Venezuela, Bolivarian Republic of",
		"VN": "Viet Nam, Socialist Republic of (was Democratic Republic of & Republic of)",
		"WF": "Wallis and Futuna Islands",
		"EH": "Western Sahara (was Spanish Sahara)",
		"YE": "Yemen",
		"ZM": "Zambia, Republic of",
		"ZW": "Zimbabwe (was Southern Rhodesia)",
	}
	resSMS := GetSmsData(countries)
	sortByName := make([]SMSData, len(resSMS))
	sortByProvider := make([]SMSData, len(resSMS))
	for i, val := range resSMS {
		for short, long := range countries {
			if val.Country == short {
				resSMS[i].Country = long
			}
		}
	}
	copy(sortByName, resSMS)
	copy(sortByProvider, resSMS)

	sort.Slice(sortByName, func(p, q int) bool {
		return sortByName[p].Country < sortByName[q].Country
	})
	sort.Slice(sortByProvider, func(p, q int) bool {
		return sortByProvider[p].Provider < sortByProvider[q].Provider
	})
	result = append(result, sortByName)
	result = append(result, sortByProvider)
	return result
}
