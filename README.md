# Виконанено

1. Розроблено http-сервер без використання фреймворків з CRUD операціями над сутностями 

2. Оптимізовано функцію конкатенації та написано бенчмарк тест

# Описання
## HTTP сервер

 Базова авторизація: User : `admin`, Password : `password`

 Сутність продавець(sellers) з полями ім'я та телефон.
 
 Отримання всіх продавців - `GET http://localhost:8080/api/seller`
 
 Отримання одного продавця за id - `GET http://localhost:8080/api/seller/:seller_id`
 
 Додавання нового продавця - `POST http://localhost:8080/api/seller`
```
{
   "name":"Maks2",
   "phone": "+389476543"
}
```
 
 Оновлення продавця за id - `PATCH http://localhost:8080/api/seller/:seller_id`
```
{
   "id":3,
   "name":"Maks3",
   "phone": "+329476543"
}
```
 
 Видалення продавця за id - `DELETE http://localhost:8080/api/seller/:seller_id`

 Запуск програми за допомогою команди : `docker-compose up -d`
 
 Для створення сутностей виконати команду: `make migrate-up`

 Тести: `make test`   

 Лінтери: `make lint`

## Оптимізація функції конкатенації

 Знаходиться: [pkg/concatenations/concatenations.go](https://github.com/LivanaKi/simple-rest/tree/main/pkg/concatenations) 

 Розроблено додатково дві функції (ConcatTwo, ConcatThree) з використанням strings. Для отримання результатів тесту виконати команду: `make test-bench`

# Тестове завдання

1. Розробка http сервера

Уявіть, що ви отримали новий проект інтернет-магазину, і вам потрібно закласти архітектуру для його розробки та підтримки. Як тестове завдання повністю спроектуйте базу даних (sql), а також зробіть CRUD однієї (будь-якої) сутності. 
HTTP сервер повинен бути написаний на Golang, максимально просто, без використання фреймворків. 
Необхідно створити docker compose, в якому серверна частина і БД будуть запускатися. 

Технічне завдання:
- розробити HTTP API з базовою авторизацією, яка дозволятиме виконувати CRUD операції над сутностями. Користувач буде один (адміністратор,
який і створюватиме ці сутності)
- формат відповіді: JSON
- опис сутностей і полів (якщо ви вважаєте, що якогось поля не вистачає, ви можете сміливо його додати):
   продавець (ім'я, телефон)
   товар (назва, опис, ціна, продавець)
   покупець (ім'я, телефон)
   замовлення (покупець, кілька товарів)
- docker для зручності запуску Go програми
- docker compose для запуску програми разом із бд.

2. Оптимізація функції конкатенації.

Оптимізуйте швидкість виконання функції. Кількість значень у вхідному параметрі (len(str)) >= 30.
Напишіть бенчмарк тест на цю функцію та її оптимізовану версію
func concat(str []string) string  {
    result := ""
    for _, v := range str {
        result += v
    }
    return result
}
Виконане тестове завдання (або два завдання) розмістіть на гітхабі та надайте посилання. 
