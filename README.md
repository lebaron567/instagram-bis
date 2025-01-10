
# Instagram-Bis API

Instagram-Bis est une API développée en Go, utilisant les frameworks et bibliothèques suivants :

- **[Gorm](https://gorm.io/)** : ORM pour la gestion des bases de données.
- **[go-chi](https://github.com/go-chi/chi)** : Framework léger pour le routage HTTP.
- **[swaggo](https://github.com/swaggo/swag)** : Génération de documentation Swagger pour l'API.

## Fonctionnalités

- Inscription et authentification des utilisateurs
- Création, lecture, mise à jour et suppression de publications
- Aimer et ne plus aimer des publications
- Commenter des publications
- Suivre et ne plus suivre des utilisateurs
- Messagerie entre utilisateurs
- Discussions

## Prérequis

- **Go** (version recommandée : 1.20 ou supérieure)
- Une base de données (sqlite)

## Installation

1. Clonez le dépôt :

```bash
git clone https://github.com/lebaron567/instagram-bis.git
cd instagram-bis
```

2. Installez les dépendances :

```bash
go mod tidy
```


3. Exécutez l'application :

```bash
go run main.go
```

## Documentation Swagger

1. Générez la documentation Swagger :

```bash
swag init
```

2. La documentation Swagger est disponible à l'adresse suivante une fois le serveur démarré :

```
http://localhost:8080/swagger/index.html
```

## Structure du projet

Voici un aperçu de la structure du projet :

```
.
├── config/
│   └── config.go
├── database/
│   ├── database.go
│   └── dbmodel/
│       ├── comments.go
│       ├── discussions.go
│       ├── followers.go
│       ├── likes.go
│       ├── members.go
│       ├── messages.go
│       ├── posts.go
│       └── users.go
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── pkg/
│   ├── authentication/
│   │   ├── controller.go
│   │   ├── jwt.go
│   │   ├── middleware.go
│   │   └── routes.go
│   ├── comment/
│   │   ├── controller.go
│   │   └── routes.go
│   ├── conversation/
│   │   ├── controller.go
│   │   └── routes.go
│   ├── like/
│   │   ├── controller.go
│   │   └── routes.go
│   ├── messagerie/
│   │   ├── controller.go
│   │   └── routes.go
│   ├── models/
│   │   ├── comments.go
│   │   ├── discussions.go
│   │   ├── followers.go
│   │   ├── like.go
│   │   ├── membres.go
│   │   ├── messages.go
│   │   ├── posts.go
│   │   └── users.go
│   ├── notification/
│   │   └── controller.go
│   ├── post/
│   │   ├── controller.go
│   │   └── routes.go
│   └── user/
│       ├── controller.go
│       └── routes.go
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```



