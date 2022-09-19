package constant

import "fmt"

const (
	IE84WindowsXP = "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET CLR 1.1.4322; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)"
)

// 国际音标
func UniCode() {
	fmt.Println("\u0250")
	fmt.Println("\u0251")
	fmt.Println("\u0252")
	fmt.Println("\u0253")
	fmt.Println("\u0254")
	fmt.Println("\u0255")
	fmt.Println("\u0256")
	fmt.Println("\u0257")
	fmt.Println("\u0258")
	fmt.Println("\u0259")
	fmt.Println("\u025A")
	fmt.Println("\u025B")
	fmt.Println("\u025C")
	fmt.Println("\u025D")
	fmt.Println("\u025E")
	fmt.Println("\u025F")
	fmt.Println("\u0260")
	fmt.Println("\u0261")
	fmt.Println("\u0262")
	fmt.Println("\u0263")
	fmt.Println("\u0264")
	fmt.Println("\u0265")
	fmt.Println("\u0266")
	fmt.Println("\u0267")
	fmt.Println("\u0268")
	fmt.Println("\u0269")
	fmt.Println("\u026A")
	fmt.Println("\u026B")
	fmt.Println("\u026C")
	fmt.Println("\u026D")
	fmt.Println("\u026E")
	fmt.Println("\u026F")
	fmt.Println("\u0270")
	fmt.Println("\u0271")
	fmt.Println("\u0272")
	fmt.Println("\u0273")
	fmt.Println("\u0274")
	fmt.Println("\u0275")
	fmt.Println("\u0276")
	fmt.Println("\u0277")
	fmt.Println("\u0278")
	fmt.Println("\u0279")
	fmt.Println("\u027A")
	fmt.Println("\u027B")
	fmt.Println("\u027C")
	fmt.Println("\u027D")
	fmt.Println("\u027E")
	fmt.Println("\u027F")
	fmt.Println("\u0280")
	fmt.Println("\u0281")
	fmt.Println("\u0282")
	fmt.Println("\u0283")
	fmt.Println("\u0284")
	fmt.Println("\u0285")
	fmt.Println("\u0286")
	fmt.Println("\u0287")
	fmt.Println("\u0288")
	fmt.Println("\u0289")
	fmt.Println("\u028A")
	fmt.Println("\u028B")
	fmt.Println("\u028C")
	fmt.Println("\u028D")
	fmt.Println("\u028E")
	fmt.Println("\u028F")
	fmt.Println("\u0290")
	fmt.Println("\u0291")
	fmt.Println("\u0292")
	fmt.Println("\u0293")
	fmt.Println("\u0294")
	fmt.Println("\u0295")
	fmt.Println("\u0296")
	fmt.Println("\u0297")
	fmt.Println("\u0298")
	fmt.Println("\u0299")
	fmt.Println("\u029A")
	fmt.Println("\u029B")
	fmt.Println("\u029C")
	fmt.Println("\u029D")
	fmt.Println("\u029E")
	fmt.Println("\u029F")
	fmt.Println("\u02A0")
	fmt.Println("\u02A1")
	fmt.Println("\u02A2")
	fmt.Println("\u02A3")
	fmt.Println("\u02A4")
	fmt.Println("\u02A5")
	fmt.Println("\u02A6")
	fmt.Println("\u02A7")
	fmt.Println("\u02A8")
	fmt.Println("\u02A9")
	fmt.Println("\u02AA")
	fmt.Println("\u02AB")
	fmt.Println("\u02AC")
	fmt.Println("\u02AD")
	fmt.Println("\u02AE")
	fmt.Println("\u02AF")
}

//汉语拼音
func Pinyin() {
	fmt.Println("\u00E1")
	fmt.Println("\u00EC")
	fmt.Println("\u012B")
	fmt.Println("\u00E9")
	/*
	   组合附加符号	̄ (U+0304)	́ (U+0301)	̌ (U+030C)	̀ (U+0300)
	   常用音节
	   大写	A	Ā (U+0100)	Á (U+00C1)	Ǎ (U+01CD)	À (U+00C0)
	   E	Ē (U+0112)	É (U+00C9)	Ě (U+011A)	È (U+00C8)
	   I	Ī (U+012A)	Í (U+00CD)	Ǐ (U+01CF)	Ì (U+00CC)
	   O	Ō (U+014C)	Ó (U+00D3)	Ǒ (U+01D1)	Ò (U+00D2)
	   U	Ū (U+016A)	Ú (U+00DA)	Ǔ (U+01D3)	Ù (U+00D9)
	   Ü (U+00DC)	Ǖ (U+01D5)	Ǘ (U+01D7)	Ǚ (U+01D9)	Ǜ (U+01DB)
	   小写	a	ā (U+0101)	á (U+00E1)	ǎ (U+01CE)	à (U+00E0)
	   e	ē (U+0113)	é (U+00E9)	ě (U+011B)	è (U+00E8)
	   i	ī (U+012B)	í (U+00ED)	ǐ (U+01D0)	ì (U+00EC)
	   o	ō (U+014D)	ó (U+00F3)	ǒ (U+01D2)	ò (U+00F2)
	   u	ū (U+016B)	ú (U+00FA)	ǔ (U+01D4)	ù (U+00F9)
	   ü (U+00FC)	ǖ (U+01D6)	ǘ (U+01D8)	ǚ (U+01DA)	ǜ (U+01DC)
	   罕用音节
	   大写	Ê (U+00CA)	Ê̄ (U+00CA U+0304)	Ế (U+1EBE)	Ê̌ (U+00CA U+030C)	Ề (U+1EC0)
	   M	M̄ (U+004D U+0304)	Ḿ (U+1E3E)	M̌ (U+004D U+030C)	M̀ (U+004D U+0300)
	   N	N̄ (U+004E U+0304)	Ń (U+0143)	Ň (U+0147)	Ǹ (U+01F8)
	   小写	ê (U+00EA)	ê̄ (U+00EA U+0304)	ế (U+1EBF)	ê̌ (U+00EA U+030C)	ề (U+1EC1)
	   m	m̄ (U+006D U+0304)	ḿ (U+1E3F)	m̌ (U+006D U+030C)	m̀ (U+006D U+0300)
	   n	n̄ (U+006E U+0304)	ń (U+0144)	ň (U+0148)	ǹ (U+01F9)
	*/

}
func emoji() {
	//fmt.Printf("%u", "♂")
	fmt.Println("\u2642")     //男
	fmt.Println("\u2640")     //女
	fmt.Println("\u263F")     //双性人或雌雄同体生物
	fmt.Println("\u26A5")     //双性人或跨性别
	fmt.Println("\u26A7")     //跨性别
	fmt.Println("\u26A2")     //女同性恋
	fmt.Println("\u26A3")     //男同性恋
	fmt.Println("\u26A4")     //异性恋
	fmt.Println("\u26A6")     //跨性别
	fmt.Println("\u26AA")     //无性恋
	fmt.Println("\U0001F600") //笑脸
	fmt.Println("\U0001F446") //上
	fmt.Println("\U0001F447") //下
	fmt.Println("\U0001F448") //左
	fmt.Println("\U0001F449") //右
	fmt.Println("\U0001F595") //Fxxc
	fmt.Println("\u2122")     //商标
	fmt.Println("\u00A9")     //版权符号
	fmt.Println("\u00AE")     //注册商标
	fmt.Println("\U0001F1E8\U0001F1F3")
}
