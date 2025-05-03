package game

/text-rpg/
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── cmd/
│   └── rpg/
│       └── main.go          # Точка входа
├── internal/
│   ├── character/           # Логика персонажей
│   │   ├── character.go     # Интерфейсы и базовые структуры
│   │   ├── knight.go
│   │   ├── paladin.go
│   │   ├── mage.go
│   │   └── priest.go
│   ├── inventory/           # Инвентарь и предметы
│   │   ├── item.go
│   │   ├── inventory.go
│   │   └── equipment.go
│   ├── locations/           # Локации игры
│   │   ├── forest.go
│   │   ├── town.go
│   │   ├── tavern.go
│   │   ├── dungeon.go
│   │   └── inn.go
│   ├── game/                # Логика игры
│   │   ├── game.go          # Основная логика
│   │   └── menu.go          # Меню и интерфейс
│   └── utils/               # Вспомогательные функции
│       └── helpers.go
└── pkg/                     # Возможные общие библиотеки
    └── cli/                 # (опционально) CLI утилиты