package task01statuses

import (
	"fmt"
)

func analyzeStatuses(statuses []int) (int, int, int) {
    var successCount, clientErrorCount, serverErrorCount int
    for _, status := range statuses {
        if status >= 200 && status < 300 {
            successCount++
        }
        if status >= 400 && status < 500 {
            clientErrorCount++
        }
        if status >= 500 && status < 600 {
            serverErrorCount++
        }
    }
    return successCount, clientErrorCount, serverErrorCount
}

func main() {
    statuses := []int{
        200, 201, 404, 500, 200,
        403, 204, 502, 301, 400,
    }
    success, clientError, serverError := analyzeStatuses(statuses)
    fmt.Printf("Успешных ответов: %d\n", success)
    fmt.Printf("Ошибок клиента: %d\n", clientError)
    fmt.Printf("Ошибок сервера: %d\n", serverError)
}