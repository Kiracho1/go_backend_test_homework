package main

import "fmt"

func buildReport(users map[string][]string, requestedUsers []string) ([]string, []string) {
	var findedUsers []string
	var notFindedUsers []string
	for _, user := range requestedUsers {
		if _, ok := users[user]; ok {
			findedUsers = append(findedUsers, user)
		} else {
			notFindedUsers = append(notFindedUsers, user)
		}
	}
	return findedUsers, notFindedUsers
}

func main() {
	users := map[string][]string{
	"kirill": {"admin", "developer"},
	"anna":   {"developer"},
	"petr":   {"support", "moderator"},
	"olga":   {"developer", "tester"},
	}

	requestedUsers := []string{
		"kirill",
		"olga",
		"unknown",
		"anna",
	}

	user, request := buildReport(users, requestedUsers)

	fmt.Printf(
		"Отчет:\n%v\n"+
		"\n"+
		"Пользователи не найдены:\n%v\n",
		user[0], request,
	)
}