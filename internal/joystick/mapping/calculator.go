package mapping

func GetCalculatorKeysMapping() JoystickMapping {
	return JoystickMapping{
		Up:     "/",
		Down:   "-",
		Left:   "*",
		Right:  "+",
		Center: "AC",
	}
}
