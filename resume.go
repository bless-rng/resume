package main

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

func main() {
	now := time.Now().Format("Jan 2006")
	r := Resume{
		Info: Info{
			FullName: "Веснин Максим Викторович",
			Birthday: "02-11-1990",
			Email:    "maks@bless-rng.ru",
			Phone:    "+79313255172",
		},
		Skills: Skills{
			MainLanguage: Language(PHP),
			Technoligies: []string{
				"Symfony 4,5,6",
				"SQL/PostgreSQL/MySQL",
				"Redis",
				"RabbitMQ",
				"Docker",
				"Composer",
				"Linux",
				"Git/Gitlab, CI/CD",
				"Nginx",
			},
			SecondaryLanguages: []Language{Language(Csharp), Language(JAVA), Language(GO)},
		},
		Experiences: []Experience{
			{
				Company:     "Milky-Games (closed)",
				Position:    "Intern C#/Unity developer",
				Period:      Period{From: "Сентябрь 2015", To: "Май 2016"},
				Description: `Разработка игр под мобильные платформы для детей c исопльзованием движка Unity3d`,
			},
			{
				Company:     "Rockstone",
				Position:    "Middle C#/Unity developer",
				Period:      Period{From: "Июль 2016", To: "Декабрь 2017"},
				Description: `Разработка игр под мобильные платформы и WebGL. Написание поведения/логики для юнитов, UI`,
			},
			{
				Company:  "Tektosoft",
				Position: "Middle developer",
				Period:   Period{From: "Январь 2018", To: "Август 2020"},
				Description: `C#/Unity 3D:
	Разработка проекта с использованием дополненной реальности(Vuforia) для демонстрации развлекательного центра "Чудо Остров".
Java:
	Разработка и поддержка плагина для Jira крупной бухгалтерской аутсорс компании.
	(Java, Atlassian SDK, RestAPI, JQuery, Ajax, SOY, Velocity). Релизация взаимодействия тасок в JIRA с 1С, личным кабинетом, почтой клиентов по АПИ.

	Доработка и исправление бэкенда для компании такси(Spring boot, web socket)
PHP:
	Разработка сервиса для размещения площадок для проведения мероприятий. Стек технологий: PHP 7.3, Symfony 4, JS, HTML, CSS(частичный фулстек), Docker, MySQL.`,
			},
			{
				Company:  "Petshop.ru",
				Position: "Middle PHP developer",
				Period:   Period{From: "Август 2020", To: "Март 2021"},
				Description: `Участие в разработки новой системы интернет магазина. Разработка новых модулей системы. Работа в рамках архитектуры DDD. Написание юнит тестов.
Используемые технологии: PHP, PostgreSQL, Symfony, Docker`,
			},
			{
				Company:  "МТС/Stream",
				Position: "Senior PHP Developer",
				Period:   Period{From: "Март 2021", To: "Апрель 2022"},
				Description: `Поддержка и развитие биллинговой платформы в рамках микросервисной архитектуры. 
Реализация нового api, исправление багов, оптимизация запросов в бд, ускорение работы  админки, добавление нового функционала.
Накатка индексов на боевую высоконагруежнную базу.
Технологии: PHP, Yii, PostgreSQL, RabbitMQ, Docker, Kubernetes, Helm`,
			},
			{
				Company:     "ГК Eqvanta",
				Position:    "Senior PHP Developer",
				Period:      Period{From: "Апрель 2022", To: "Август 2022"},
				Description: `Поддержка и развитие админки. На проекте была реализованна мульти-тенантная архитектура. PHP, Yii, PostgreSQL, RabbitMQ, Docker`,
			},
			{
				Company:  "Банк 131",
				Position: "Lead PHP Developer",
				Period:   Period{From: "Август 2023", To: strings.Replace(now, now[0:3], months[now[0:3]], 1)},
				Description: `Интеграция с платежными системами через партнерское апи
Анализ и декомпозиция планируемых интеграций
Участие в бординге партнеров
Code-review`,
			},
		},
	}
	data, _ := json.MarshalIndent(r, "", "\t")
	os.WriteFile("Resume.json", data, 0666)
}

var months = map[string]string{
	"Jan": "Январь",
	"Feb": "Февряль",
	"Mar": "Март",
	"Apr": "Апрель",
	"May": "Май",
	"Jun": "Июнь",
	"Jul": "Июль",
	"Aug": "Август",
	"Sep": "Сентябрь",
	"Oct": "Октябрь",
	"Nov": "Ноябрь",
	"Dec": "Декабрь",
}

type Resume struct {
	Info        Info
	Skills      Skills
	Experiences []Experience
}

type Info struct {
	FullName string
	Birthday string
	Email    string
	Phone    string
}

type Language string

const (
	PHP    string = "PHP"
	GO     string = "GO"
	JAVA   string = "JAVA"
	Csharp string = "C#"
)

type Skills struct {
	MainLanguage       Language
	Technoligies       []string
	SecondaryLanguages []Language
}

type Period struct {
	From string
	To   string
}
type Experience struct {
	Company     string
	Position    string
	Period      Period
	Description string
}
