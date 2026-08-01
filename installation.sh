#!/bin/bash

# === Core tools ===
go install golang.org/x/tools/gopls@latest          # Language server (VS Code / Neovim)
go install github.com/go-delve/delve/cmd/dlv@latest # Debugger
go install github.com/air-verse/air@latest          # Live reload
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest  # Linter

# === Popular web frameworks ===
go install github.com/gin-gonic/gin@latest
go install github.com/labstack/echo/v4@latest
go install github.com/gofiber/fiber/v2@latest
go install github.com/go-chi/chi/v5@latest

# === Database ===
go install gorm.io/gorm@latest
go install gorm.io/driver/sqlite@latest
go install gorm.io/driver/postgres@latest
go install github.com/jackc/pgx/v5@latest
go install github.com/redis/go-redis/v9@latest

# === Auth, Config, CLI, Testing ===
go install github.com/golang-jwt/jwt/v5@latest
go install github.com/spf13/viper@latest
go install github.com/spf13/cobra@latest
go install github.com/stretchr/testify@latest
go install github.com/google/uuid@latest

# === Useful utilities ===
go install golang.org/x/crypto@latest
go install golang.org/x/sync@latest
go install github.com/joho/godotenv@latest

