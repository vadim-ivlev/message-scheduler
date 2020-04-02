package main

import (
	"flag"
	"fmt"

	"message-scheduler/pkg/fire"
	"message-scheduler/server"
	"strconv"
)

func main() {
	// считать параметры командной строки
	servePort, env := readCommandLineParams()
	// считать конфиг файлы
	fire.ReadConfig("./configs/firebase.yaml", env)

	// если порт > 0, печатаем приветствие и запускаем сервер
	if servePort > 0 {
		printGreetings(servePort, env)
		server.Serve(":" + strconv.Itoa(servePort))
	}
}

// Вспомогательные функции =========================================

// readCommandLineParams читает параметры командной строки
func readCommandLineParams() (serverPort int, env string) {
	flag.IntVar(&serverPort, "port", 8088, "Запустить приложение на порту с номером")
	flag.StringVar(&env, "env", "dev", "Окружение. Возможные значения: dev - разработка, front - в докере для фронтэнд разработчиков. prod - по умолчанию для продакшн.")
	flag.Parse()
	fmt.Println("\nПример запуска: ./message-scheduler -port=8088 -env=dev")
	flag.Usage()
	return
}

// printGreetings печатаем приветственное сообщение
func printGreetings(serverPort int, env string) {
	msg := `
	**********************************************
	MESSAGE_SCHEDULER started. 
	Environment: %v
	GraphQL endpoint -> http://localhost:%v/schema
	GraphQL test     -> https://graphql-test.now.sh/?end_point=http://localhost:%v/schema&tab_name=message-scheduler:%v
	CTRL-C to interrupt.
	**********************************************
`
	fmt.Printf(msg, env, serverPort, serverPort, serverPort)
}
