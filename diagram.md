cmd/
  └─ api/                → main.go (bootstrap)
internal/
  ├─ domain/             → entity + interface
  ├─ usecase/            → business logic
  ├─ adapter/
  │    ├─ http/          → fiber handler
  │    └─ persistence/   → gorm repository
  └─ infrastructure/
       └─ database/      → postgres connection