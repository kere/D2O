package baseinfo

import "sort"

// AreasSort array
type AreasSort []Area

// Less sort series
func (s AreasSort) Len() int {
	return len(s)
}

// Less sort series
func (s AreasSort) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

// Swap sort series
func (s AreasSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Area class
type Area struct {
	ID int    `json:"id"`
	CN string `json:"cn"`
	EN string `json:"en"`
}

func init() {
	sort.Sort(Areas)
}

// Areas data
var Areas = AreasSort([]Area{
	Area{1, "阿布哈兹（格鲁吉亚）", "Abkhazia"},
	Area{2, "阿富汗", "Afghanistan"},
	Area{3, "阿尔巴尼亚", "Albania"},
	Area{4, "阿尔及利亚", "Algeria"},
	// Area{5, "安道尔", "Andorra"},
	Area{6, "安哥拉", "Angola"},
	// Area{7, "安提瓜和巴布达", "Antigua and Barbuda"},
	Area{8, "阿根廷", "Argentina"},
	Area{9, "亚美尼亚", "Armenia"},
	Area{10, "澳大利亚", "Australia"},
	Area{11, "奥地利", "Austria"},
	Area{12, "阿塞拜疆", "Azerbaijan"},
	//B编辑
	Area{13, "巴哈马", "Commonwealth oftheBahamas"},
	Area{14, "巴林", "Bahrain"},
	Area{15, "孟加拉国", "Bangladesh"},
	// Area{16, "巴巴多斯", "Barbados"},
	Area{17, "白俄罗斯", "Belarus"},
	Area{18, "比利时", "Belgium"},
	// Area{19, "伯利兹", "Belize"},
	Area{20, "贝宁", "Benin"},
	Area{21, "不丹", "Bhutan"},
	Area{22, "玻利维亚", "Bolivia"},
	Area{23, "波黑", "Bosnia and Herzegovina"},
	Area{24, "博茨瓦纳", "Botswana"},
	Area{25, "巴西", "Brazil"},
	Area{26, "文莱", "Brunei"},
	Area{27, "保加利亚", "Bulgaria"},
	Area{28, "布基纳法索", "Burkina Faso"},
	Area{29, "布隆迪", "Burundi"},
	//C编辑
	Area{30, "柬埔寨", "Cambodia"},
	Area{31, "喀麦隆", "Cameroon"},
	Area{32, "加拿大", "Canada"},
	// Area{33, "佛得角", "Cape Verde"},
	Area{34, "加泰罗尼亚（西班牙）", "Catalen"},
	Area{35, "中非共和国", "Central African Republic"},
	Area{36, "乍得", "Chad"},
	Area{37, "智利", "Chile"},
	Area{38, "中国", "China"},
	Area{39, "哥伦比亚", "Colombia"},
	Area{40, "科摩罗", "Comoros"},
	Area{41, "刚果共和国", "Congo (Brazzaville)"},
	Area{42, "刚果民主共和国", "Congo (Kinshasa)"},
	Area{43, "库克群岛（新西兰）", "Cook Islands"},
	Area{44, "哥斯达黎加", "Costa Rica"},
	Area{45, "科特迪瓦", "Côte d'Ivoire"},
	Area{46, "克罗地亚", "Croatia"},
	Area{47, "古巴", "Cuba"},
	Area{48, "塞浦路斯", "Cyprus"},
	Area{49, "捷克", "Czech Republic"},
	//D编辑

	Area{50, "丹麦", "Denmark"},
	Area{51, "吉布提", "Djibouti"},
	Area{52, "顿涅茨克", "Donetsk People's Republic"},
	Area{53, "多米尼克", "Dominica"},
	Area{54, "多米尼加", "Dominican Republic"},
	//E编辑

	Area{55, "厄瓜多尔", "Ecuador"},
	Area{56, "埃及", "Egypt"},
	Area{57, "萨尔瓦多", "El Salvador"},
	Area{58, "赤道几内亚", "Equatorial Guinea"},
	Area{59, "厄立特里亚", "Eritrea"},
	Area{60, "爱沙尼亚", "Estonia"},
	Area{61, "埃塞俄比亚", "Ethiopia"},
	//F编辑

	Area{62, "斐济", "Fiji"},
	Area{63, "芬兰", "Finland"},
	Area{64, "法国", "France"},
	//G编辑

	Area{65, "加蓬", "Gabon"},
	Area{66, "冈比亚", "Gambia, The"},
	Area{67, "格鲁吉亚", "Georgia"},
	Area{68, "德国", "Germany"},
	Area{69, "加纳", "Ghana"},
	Area{70, "希腊", "Greece"},
	Area{71, "格林纳达", "Grenada"},
	Area{72, "危地马拉", "Guatemala"},
	Area{73, "几内亚", "Guinea"},
	Area{74, "几内亚比绍", "Guinea-Bissau"},
	Area{75, "圭亚那", "Guyana"},
	//H编辑

	Area{76, "海地", "Haiti"},
	Area{77, "洪都拉斯", "Honduras"},
	Area{78, "匈牙利", "Hungary"},
	//I编辑

	Area{79, "冰岛", "Iceland"},
	Area{80, "印度", "India"},
	Area{81, "印度尼西亚", "Indonesia"},
	Area{82, "伊朗", "Iran"},
	Area{83, "伊拉克", "Iraq"},
	Area{84, "爱尔兰", "Ireland"},
	Area{85, "以色列", "Israel"},
	Area{86, "意大利", "Italy"},

	Area{500, "伊斯兰国", "ISIS"},
	//J编辑

	Area{87, "牙买加", "Jamaica"},
	Area{88, "日本", "Japan"},
	Area{89, "约旦", "Jordan"},
	//K编辑

	Area{90, "哈萨克斯坦", "Kazakhstan"},
	Area{91, "肯尼亚", "Kenya"},
	Area{92, "基里巴斯", "Kiribati"},
	Area{93, "韩国", "Korea, South"},
	Area{94, "科索沃", "Kosovo"},
	Area{95, "科威特", "Kuwait"},
	Area{96, "吉尔吉斯斯坦", "Kyrgyzstan"},
	//L编辑

	Area{97, "老挝", "Laos"},
	Area{98, "拉脱维亚", "Latvia"},
	Area{99, "黎巴嫩", "Lebanon"},
	Area{100, "莱索托", "Lesotho"},
	Area{101, "利比里亚", "Liberia"},
	Area{102, "利比亚", "Libya"},
	Area{103, "列支敦士登", "Liechtenstein"},
	Area{104, "立陶宛", "Lithuania"},
	Area{105, "卢森堡", "Luxembourg"},
	//M编辑

	Area{106, "马达加斯加", "Madagascar"},
	Area{107, "马拉维", "Malawi"},
	Area{108, "马来西亚", "Malaysia"},
	Area{109, "马尔代夫", "Maldives"},
	Area{110, "马耳他骑士团", "Maltese Knights"},
	Area{111, "马里", "Mali"},
	Area{112, "马耳他", "Malta"},
	Area{113, "马绍尔群岛", "Marshall Islands"},
	Area{114, "毛里塔尼亚", "Mauritania"},
	Area{115, "毛里求斯", "Mauritius"},
	Area{116, "墨西哥", "Mexico"},
	Area{117, "密克罗尼西亚联邦", "Micronesia"},
	Area{118, "摩尔多瓦", "Moldova"},
	Area{119, "摩纳哥", "Monaco"},
	Area{120, "蒙古国", "Mongolia"},
	Area{121, "黑山", "Montenegro"},
	Area{122, "摩洛哥", "Morocco"},
	Area{123, "莫桑比克", "Mozambique"},
	Area{124, "缅甸", "Myanmar"},
	//N编辑

	Area{125, "纳戈尔诺-卡拉巴赫", "Nagorno-Karabakh"},
	Area{126, "纳米比亚", "Namibia"},
	Area{127, "瑙鲁", "Nauru"},
	Area{128, "尼泊尔", "Nepal"},
	Area{129, "荷兰", "Netherlands"},
	Area{130, "新西兰", "New Zealand"},
	Area{131, "尼加拉瓜", "Nicaragua"},
	Area{132, "尼日尔", "Niger"},
	Area{133, "尼日利亚", "Nigeria"},
	Area{134, "纽埃（新西兰）", "Niue"},
	Area{135, "北塞浦路斯", "Northern Cyprus"},
	Area{136, "北马其顿", "North Macedonia"},
	Area{137, "挪威", "Norway"},
	//O编辑

	Area{138, "阿曼", "Oman"},

	//P编辑
	Area{139, "巴基斯坦", "Pakistan"},
	Area{140, "帕劳", "Palau"},
	Area{141, "巴勒斯坦", "Palestine"},
	Area{142, "巴拿马", "Panama"},
	Area{143, "巴布亚新几内亚", "Papua New Guinea"},
	Area{144, "巴拉圭", "Paraguay"},
	Area{145, "朝鲜", "People's Republic of Korea"},
	Area{146, "秘鲁", "Peru"},
	Area{147, "菲律宾", "Philippines"},
	Area{148, "波兰", "Poland"},
	Area{149, "葡萄牙", "Portugal"},
	Area{150, "德涅斯特河沿岸", "Pridnestrovie"},
	Area{151, "邦特兰", "Puntland"},

	//Q编辑
	Area{152, "卡塔尔", "Qatar"},
	//R编辑

	Area{153, "罗马尼亚", "Romania"},
	Area{154, "俄罗斯", "Russia"},
	Area{155, "卢旺达", "Rwanda"},
	//S编辑

	Area{156, "圣基茨和尼维斯", "Saint Christopher and Nevis"},
	// Area{157, "圣卢西亚", "Saint Lucia"},
	// Area{158, "圣文森特和格林纳丁斯", "Saint Vincent and the Grenadines"},
	Area{159, "萨摩亚", "Samoa"},
	Area{160, "圣马力诺", "San Marino"},
	// Area{161, "圣多美和普林西比", "São Tomé and Príncipe"},
	Area{162, "沙特阿拉伯", "Saudi Arabia"},
	Area{163, "塞内加尔", "Senegal"},
	Area{164, "塞尔维亚", "Serbia"},
	Area{165, "塞舌尔", "Seychelles"},
	// Area{166, "塞拉利昂", "Sierra Leone"},
	Area{167, "新加坡", "Singapore"},
	Area{168, "斯洛伐克", "Slovakia"},
	Area{169, "斯洛文尼亚", "Slovenia"},
	Area{170, "所罗门群岛", "Solomon Islands"},
	Area{171, "索马里", "Somali"},
	// Area{172, "索马里兰", "Somaliland"},
	Area{173, "南非", "South Africa"},
	Area{174, "南奥塞梯", "South Ossetia"},
	Area{175, "南苏丹", "South Sudan"},
	Area{176, "西班牙", "Spain"},
	Area{177, "斯里兰卡", "Sri Lanka"},
	Area{178, "苏丹", "Sudan"},
	// Area{179, "苏里南", "Suriname"},
	// Area{180, "斯威士兰", "Swaziland"},
	Area{181, "瑞典", "Sweden"},
	Area{182, "瑞士", "Switzerland"},
	Area{183, "叙利亚", "Syria"},
	//T编辑

	Area{184, "塔吉克斯坦", "Tajikistan"},
	Area{185, "坦桑尼亚", "Tanzania"},
	Area{186, "泰国", "Thailand"},
	Area{187, "东帝汶", "Timor-Leste"},
	// Area{188, "多哥", "Togo"},
	Area{189, "汤加", "Tonga"},
	Area{190, "特立尼达和多巴哥", "Trinidad and Tobago"},
	Area{191, "突尼斯", "Tunisia"},
	Area{192, "土耳其", "Turkey"},
	Area{193, "土库曼斯坦", "Turkmenistan"},
	Area{194, "图瓦卢", "Tuvalu"},
	//U编辑

	Area{195, "乌干达", "Uganda"},
	Area{196, "乌克兰", "Ukraine"},
	Area{197, "阿联酋", "United Arab Emirates"},
	Area{198, "英国", "United Kingdom"},
	Area{199, "美国", "United States"},
	Area{200, "乌拉圭", "Uruguay"},
	Area{201, "乌兹别克斯坦", "Uzbekistan"},
	//V编辑

	Area{202, "瓦努阿图", "Vanuatu"},
	Area{203, "梵蒂冈", "Vatican city（the Holy see)"},
	Area{204, "委内瑞拉", "Venezuela"},
	Area{205, "越南", "Vietnam"},
	//W编辑

	Area{206, "西撒哈拉", "Western Sahara"},
	//Y编辑

	Area{207, "也门", "Yemen"},
	//Z编辑

	Area{208, "赞比亚", "Zambia"},
	Area{209, "津巴布韦", "Zimbabwe"},
	//地区编辑

	Area{210, "阿布穆萨岛", "AbuMusa"},
	Area{211, "布雷米", "Al Buraymi"},
	Area{212, "爱丽丝浅滩", "Alice Shoal"},
	Area{213, "南极", "Antarctica"},
	Area{214, "阿维斯岛", "Aves Island"},
	Area{215, "巴得梅", "Badme"},
	Area{216, "比尔泰维勒", "Bir Tawil"},
	Area{217, "休达", "Ceuta"},
	Area{218, "独岛", "Dokdo"},
	Area{219, "埃塞奎博河地区", "Essequibo River"},
	Area{220, "斐迪南迪亚岛", "Ferdinandea"},
	Area{221, "格罗里奥索群岛", "Glorioso Islands"},
	Area{222, "哈拉伊卜三角区", "Hala'ib Triangle"},
	Area{223, "汉斯岛", "Hans Island"},
	Area{224, "伊米亚岛", "Imia"},
	Area{225, "卡恩格瓦尼", "Kangwane"},
	Area{226, "克什米尔", "Kashmir"},
	Area{227, "林梦地区", "Lin Meng area"},
	Area{228, "梅利利亚", "Melilla"},
	Area{229, "欧加登", "Ogaden"},
	Area{230, "佩雷吉尔岛", "Perejil island"},
	Area{231, "卢卡万兹岛", "Rukwanzi Island"},
	Area{232, "索科特拉岛", "Socotra"},
	Area{233, "南设得兰群岛", "South Shetland Islands"},
	Area{234, "北方四岛", "The Kuril Islands"},

	// 海外省
	Area{235, "瓜德罗普（法国）", "Guadeloupe"},
	Area{236, "法属圭亚那（法国）", "Guyane française"},
	Area{237, "马提尼克（法国）", "Martinique"},
	Area{238, "留尼旺（法国）", "Réunion"},
	Area{239, "奥兰群岛（芬兰）", "Åland Islands"},
	Area{240, "伏伊伏丁那（塞尔维亚）", "Vojvodina"},
	// 海外属地

	Area{241, "美属萨摩亚（美国）", "American Samoa"},
	Area{242, "安圭拉（英国）", "Anguilla"},
	Area{243, "阿鲁巴（荷兰）", "Aruba"},
	Area{244, "亚速尔群岛（葡萄牙）", "Azores"},
	Area{245, "百慕大（英国）", "Bermuda"},
	Area{246, "博奈尔岛（荷兰）", "Bonaire"},
	Area{247, "英属印度洋领地（英国）", "British Indian Ocean Territory"},
	Area{248, "加那利群岛（西班牙）", "Canary Islands"},
	Area{249, "开曼群岛（英国）", "Cayman Islands"},
	Area{250, "圣诞岛（澳大利亚）", "Christmas lsland"},
	Area{251, "科科斯群岛（澳大利亚）", "Cocos (Keeling) Islands"},
	Area{252, "库拉索（荷兰）", "Curaçao"},
	Area{253, "复活节岛（智利）", "Easter Island"},
	Area{254, "福克兰群岛英国、阿根廷争议）", "Falkland Islands (Islas Malvinas)"},
	Area{255, "法罗群岛（丹麦）", "Faroe Islands"},
	Area{256, "法属波利尼西亚（法国）", "French Polynesia"},
	Area{257, "直布罗陀", "Gibraltar"},
	Area{258, "格陵兰（丹麦）", "Greenland"},
	Area{259, "关岛（美国）", "Guam"},
	Area{260, "根西岛（英国）", "Guernsey"},
	Area{261, "扬马延岛（挪威）", "Jan Mayen"},
	Area{262, "泽西岛（英国）", "Jersey"},
	Area{263, "马恩岛（英国）", "Isle of Man"},
	Area{264, "马约特", "Mayotte"},
	Area{265, "马德拉群岛（葡萄牙）", "Madeira"},
	Area{266, "中途岛（美国）", "Midway Island"},
	Area{267, "蒙特塞拉特岛（英国）", "Montserrat"},
	Area{268, "纳弗沙岛（美国）", "Navassa Island"},
	Area{269, "新喀里多尼亚（法国）", "New Caledonia"},
	Area{270, "诺福克岛（澳大利亚）", "Norfolk Island"},
	Area{271, "北马里亚纳群岛（美国）", "Northern Mariana Islands"},
	Area{272, "皮特凯恩群岛（英国）", "Pitcairn Islands"},
	Area{273, "波多黎各（美国）", "Puerto Rico"},
	Area{274, "萨巴岛（荷兰）", "Saba Island"},
	Area{275, "圣赫勒拿（英国）", "Saint Helena"},
	Area{276, "法属圣马丁（法国）", "Saint-Martin"},
	Area{277, "圣皮埃尔和密克隆群岛（法国）", "Saint Pierre and Miquelon"},
	Area{278, "圣尤斯特歇斯（荷兰）", "Sint Eustatius"},
	Area{279, "荷属圣马丁（荷兰）", "Sint Maarten"},
	Area{280, "南乔治亚岛和南桑威奇群岛", "South Georgia and The South Sandwich Islands"},
	Area{281, "斯瓦尔巴特群岛（挪威）", "Svalbard"},
	Area{282, "托克劳（新西兰）", "Tokelau"},
	Area{283, "特克斯与凯科斯群岛（英国）", "Turks and Caicos Islands"},
	Area{284, "英属维尔京群岛（英国）", "Virgin Islands, British"},
	Area{285, "美属维尔京群岛（美国）", "Virgin Islands, United States"},
	Area{286, "威克岛（美国）", "Wake Island"},
	Area{287, "瓦利斯和富图纳（法国）", "Wallis and Futuna"},
	// 自治共和国
	Area{288, "阿迪格共和国（俄）", "Adygea"},
	Area{289, "阿尔泰共和国（俄）", "Altay"},
	Area{290, "阿扎尔自治共和国（格）", "Azar"},
	Area{291, "巴什科尔托斯坦共和国（俄）", "Bashkortostan"},
	Area{292, "布里亚特共和国（俄）", "Buryat"},
	Area{293, "车臣共和国（俄）", "Chechnya"},
	Area{294, "楚瓦什共和国（俄）", "Chuvash"},
	Area{295, "克里米亚共和国（乌克兰、俄争议）", "Crimea"},
	Area{296, "达吉斯坦共和国（俄）", "Dagestan"},
	Area{297, "哈卡斯共和国（俄）", "Khakas"},
	Area{298, "印古什共和国（俄）", "Ingushetia"},
	Area{299, "卡巴尔达－巴尔卡尔共和国（俄）", "Kabardino-Balkar"},
	Area{300, "卡拉恰伊－切尔克斯共和国（俄）", "Kalachei - Circassian"},
	Area{301, "卡尔梅克共和国－哈利姆格坦格奇（俄）", "Kalmyk - Halim Gdansk"},
	Area{302, "卡拉卡尔帕克斯坦自治共和国（乌兹）", "Karakalpakstan"},
	Area{303, "卡累利阿共和国（俄）", "Karelia"},
	Area{304, "科米共和国（俄）", "Komi"},
	Area{305, "马里埃尔共和国（俄）", "Mariel"},
	Area{306, "莫尔多瓦共和国（俄）", "Moldova"},
	Area{307, "纳希切万自治共和国（阿）", "Naxçıvan"},
	Area{308, "萨哈（雅库特）共和国（俄）", "Saha (Yakutia)"},
	Area{309, "北奥塞梯共和国（俄）", "North Ossetia"},
	Area{310, "鞑靼斯坦共和国（俄）", "Tatarstan"},
	Area{311, "图瓦共和国（俄）", "Tuva"},
	Area{312, "乌德穆尔特共和国（俄）", "Udmurt"},
})
