package main

import (
    "database/sql"
    "fmt"
    "html/template"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

// Структура для передачи данных в HTML-шаблон
type PageData struct {
    Message string
}

// Глобальная переменная для подключения к БД (можно и по-другому, но для примера норм)
var db *sql.DB

func main() {
    // Подключение к MySQL — замени пароль и имя базы на свои
    connStr := "root:devops366@tcp(172.17.0.1:3308)/proekt"
    var err error
    db, err = sql.Open("mysql", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Проверка подключения
    err = db.Ping()
    if err != nil {
        panic("Не могу подключиться к базе: " + err.Error())
    }
    fmt.Println("Подключено к MySQL")

    // Маршруты
    http.HandleFunc("/", indexHandler)      // Главная страница с формой
    http.HandleFunc("/add", addHandler)     // Обработка POST-запроса

    // Запуск сервера на порту 3333
    fmt.Println("Сервер запущен на http://localhost:3333")
    panic(http.ListenAndServe(":3333", nil))
}

// Главная страница с формой
func indexHandler(w http.ResponseWriter, r *http.Request) {
    data := PageData{}
    // Если передан параметр msg, показываем его (после вставки)
    msg := r.URL.Query().Get("msg")
    if msg != "" {
        data.Message = msg
    }

    tmpl := template.Must(template.New("index").Parse(indexHTML))
    tmpl.Execute(w, data)
}

// Обработчик добавления записи
func addHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    // Читаем значение поля "text" из формы
    text := r.FormValue("text")
    if text == "" {
        http.Redirect(w, r, "/?msg=Пустая+строка,+попробуйте+снова", http.StatusSeeOther)
        return
    }
	if len(text) > 40 {
        http.Redirect(w, r, "/?msg=Ошибка:+максимальная+длина+40+символов", http.StatusSeeOther)
        return
    }
    // Вставка в таблицу proekt в столбец out
    stmt, err := db.Prepare("INSERT INTO proekt(`out`) VALUES(?)")
    if err != nil {
        http.Error(w, "Ошибка подготовки запроса: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(text)
    if err != nil {
        http.Error(w, "Ошибка вставки: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Перенаправляем обратно на главную с сообщением об успехе
    http.Redirect(w, r, "/?msg=Успешно+добавлено:+ "+text, http.StatusSeeOther)
}

// HTML-шаблон (храним прямо в коде для простоты)
const indexHTML = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Вставка в proekt</title>
    <style>
        body { font-family: Arial; margin: 40px; }
        input[type="text"] { width: 300px; padding: 6px; }
        button { padding: 6px 12px; }
        .message { color: green; margin-bottom: 15px; }
    </style>
</head>
<body>
    <h1>Добавить запись в таблицу proekt</h1>
    {{if .Message}}
        <div class="message">{{.Message}}</div>
    {{end}}
    <form action="/add" method="POST">
        <input type="text" name="text" placeholder="Введите что угодно (до 40 символов)">
        <button type="submit">Отправить</button>
    </form>
</body>
</html>
`
