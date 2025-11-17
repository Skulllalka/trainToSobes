package mymap

import "sync"

var cache sync.Map

func GetOrCompute (key string, computeFunc() string) string{

}

func compute(userID string) string{
	return fmt.Sprintf("вычисление значения %s", userID)
}

func Task5() {
	usersID := []string{"user1", "user2","user1","user2","user3","user4","user7"}

	for _, userID:= range usersID{
		res := GetOrCompute(userID,func ()string{
			return compute(userID)
		})
		asd
		asd
	}
}asd
