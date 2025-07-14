package user

func GetCourse(c Course) (string, error) {
	return c.Name, nil
}
