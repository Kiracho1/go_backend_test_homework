package main

import "fmt"

func buildReport(users map[string][]string, requestedUsers []string) ([]string, []string) {
	var report []string
	var missingUsers []string

	for _, user := range requestedUsers {
		if roles, ok := users[user]; ok {
			line := user + ": "

			for i, role := range roles {
				if i > 0 {
					line += ", "
				}
				line += role
			}
			report = append(report, line)
		} else {
			missingUsers = append(missingUsers, user)
		}
	}
	return report, missingUsers
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
		"petr",
		"anna",
	}

	report, missingUsers := buildReport(users, requestedUsers)

	fmt.Println("Отчет:")
	for _, line := range report {
		fmt.Println(line)
	}
	fmt.Println()
	fmt.Println("Пользователи не найдены:")
	fmt.Println(missingUsers)
}
