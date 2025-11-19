package interfaces

type CusstomError struct{
	message string
}

func (c *CusstomError) Error() string{
	return c.message
}

func Handle() error {
	return &CusstomError{"ты большая ошибка жизни!"}
}
