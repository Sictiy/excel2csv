package entry

import "regexp"

type Column struct {
	Name string
	Comment string
	// 字段类型
	JavaType string
	// 字段是否为list
	IsList bool
	// 字段类型为list时的元素类型
	JavaSubType string
	// csv 字段注解
	CsvBind string
	/**模板用**/
}

/**
根据字段类型 得到csvbind 与子类型
 */
func (c *Column)ProcessJavaType()  {
	c.IsList, _ = regexp.MatchString("^List.*", c.JavaType)
	if c.IsList {
		reg := regexp.MustCompile("[^<>]*")
		allMatch := reg.FindAllString(c.JavaType, -1)
		if len(allMatch) > 1 {
			c.JavaSubType = allMatch[1]
		}
	}
	if c.IsList && c.JavaSubType != "" {
		c.CsvBind = "@CsvBindAndSplitByName(elementType= "+ c.JavaSubType + ".class, splitOn = \"\\\\,\")"
	}else {
		c.CsvBind = "@ScvBindByName"
	}
}
