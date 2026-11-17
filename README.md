# API-тесты для проекта https://testboard.avito.com/

## Полезные ссылки

1. [Swagger-спецификация](https://testboard.avito.com/swagger/index.html)
2. [Техническое задание](https://docs.google.com/document/d/1IZ9SxEcePoQT6lyEmukycp1qNok4HZtok73UQnPwKSA/edit?usp=sharing)
3. [Библиотека для написания API-тестов](https://github.com/steinfletcher/apitest)
4. [Мощная библиотека анализа JSON для Go](https://github.com/tidwall/gjson)

## Запуск тестов

1. Установите любое IDE для разработки: [Jetbrains GoLand](https://www.jetbrains.com/ru-ru/go/) (Платное или по студенческой лицензии), [VS Code](https://code.visualstudio.com/) (бесплатное).
2. Установите [Git](https://git-scm.com/install/) на компьютер
3. [Установите](https://go.dev/doc/install) Go версии 1.25.1 или выше
4. Зарегистрируйтесь на [TestBoard](https://testboard.avito.com/)
5. Склонируйте репозиторий и перейдите в него:

```bash
  git clone https://github.com/WannaBread/api-tests-student.git
  cd api-tests-student
```

6. Подтяните необходимые зависимости через команду в корне проекта:

```bash
  go mod tidy
```

7. В файле `.env.override` (создайте его в корне, если нет) выставьте логин и пароль от TestBoard:

```
TEST_LOGIN=your_login
TEST_PASSWORD=your_password
```

8. Запуск всех тестов из корня проекта (`-count=1`, чтобы `go test` не брал результат из кэша):

```bash
  go test -count=1 -v ./...
```

9. Запуск отдельного пакета:

```bash
  go test -count=1 -v ./tests/scenarios/favorites/
```

10. Запуск одного теста:

```bash
  go test -count=1 -v ./tests/scenarios/favorites/ -run 'TestSuiteRun/TestAddFavoriteOverLimit'
```

## Тестовые данные

Заранее готовить данные на стенде не нужно.

- Объявления создаются в самих тестах (название начинается с `autotest`) и удаляются после теста, в том числе если он упал.
- В избранное нельзя добавить своё объявление, поэтому сьют `favorites` на каждый прогон регистрирует второго
  пользователя `autotest_<случайная строка>@example.com`. Удалить пользователя через API нельзя, так что они остаются
  на стенде.
- Для теста на лимит избранного (50) регистрируется ещё один пользователь, его избранное заполняется самыми старыми
  объявлениями из каталога.

## Структура проекта

- `internal/client/http/` – HTTP-запросы к API
- `internal/managers/` – проверка статуса, разбор ответа, создание и удаление тестовых данных
- `internal/assertions/` – общие проверки
- `internal/constants/` – пути, лимиты (`limits`) и тексты ошибок API (`apierrors`)
- `internal/fixtures/` – фотографии для загрузки в объявления
- `tests/scenarios/` – сами тесты:
  - `advertisementCreate` – создание объявления: позитивный сценарий, валидация полей, запросы без авторизации
  - `favorites` – избранное: добавление и удаление, несуществующий/некорректный `ad_id`, 401, 403, 409, лимит
  - `advertisementSearch` – поиск по подстроке и без учёта регистра
  - `myAdvertisement` – список своих объявлений

## Найденные дефекты

- `POST /advertisement` создаёт объявление с `title` из одних пробелов (`"   "`) и отвечает 201, хотя пустой `title`
  отклоняется с ошибкой `title is required`.
- Ограничения на стенде строже, чем в swagger: `title` до 50 символов (в swagger для обновления 255), `description`
  до 500 (в swagger 1000), `quantity` до 100, `price` до 1 000 000.

## Логирование

В консоль выводятся шаги теста с метками `[Precondition]`, `[Step]`, `[Check]` и `[Cleanup]`.

Для включения логирования request/response в консоль: выставьте переменную окружения `DEBUG=true` в
`.env.override`

## Запуск линтера

1. Установить приложение [golangci-lint](https://golangci-lint.run/docs/welcome/install/local/) версии 2.x
2. Запуск линтера (из корня репозитория):

```bash
  golangci-lint run
```
