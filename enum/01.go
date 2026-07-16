package enum

type Season int

// const 声明了一个常量 则必须 赋值 否则报错

const (
	Spring Season = iota + 1
	Summer
	Autumn
	Winter
)

type WeekDay string

const (
	day1 WeekDay = "Monday"
	day2 WeekDay = "Tuesday"
	day3 WeekDay = "Wednesday"
	day4 WeekDay = "Thursday"
	day5 WeekDay = "Friday"
	day6 WeekDay = "Saturday"
	day7 WeekDay = "Sunday"
)

// 为什么能实现枚举效果呢？ ----> 新定义了一个 类型 --> 这个类型 只能取 常量中定义的值 --> 这样就实现了枚举效果
