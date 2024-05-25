1. Add 4 variables to your Github `settings > secrets`

```
FRONTEND_FLY_TOKEN
BACKEND_FLY_TOKEN
DB_FLY_TOKEN
DATABASE_URL
```

```
flyctl tokens create deploy -x 999999h -a imagegen
```

2. Run the script to deploy