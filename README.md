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
4. Зарегистрируйте аккаунт на [TestBoard](https://testboard.avito.com/), если его ещё нет.
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

8. Для запуска всех тестов выполните команду в корне проекта (`-count=1` отключает кэш результатов `go test`,
   чтобы тесты действительно ходили в API):

```bash
  go test -count=1 -v ./...
```

9. Отдельный пакет с тестами можно запустить командой:

```bash
  go test -count=1 -v ./tests/scenarios/advertisementCreate/
  go test -count=1 -v ./tests/scenarios/favorites/
```

10. Отдельный тест — через `-run`:

```bash
  go test -count=1 -v ./tests/scenarios/favorites/ -run 'TestSuiteRun/TestAddFavoriteOverLimit'
```

## Тестовые данные

Тесты не требуют заранее подготовленных данных на стенде и идемпотентны — их можно запускать многократно подряд:

- каждое объявление создаётся самим тестом (название начинается с `autotest`) и удаляется после завершения теста,
  даже если тест упал;
- для сценариев избранного нужен второй пользователь: API запрещает добавлять в избранное своё объявление.
  Поэтому сьют `favorites` на каждый прогон регистрирует нового пользователя `autotest_<случайная строка>@example.com`
  — его избранное изначально пустое и не пересекается с другими прогонами. Удалить пользователя через API нельзя,
  поэтому такие пользователи остаются на стенде;
- тест на лимит избранного (50 объявлений) регистрирует ещё одного пользователя и заполняет его избранное
  самыми старыми объявлениями из общего каталога.

## Структура проекта

Тесты построены по слоям **Client → Manager → Scenario**:

| Слой                | Каталог                                     | Назначение                                                                   |
|---------------------|---------------------------------------------|------------------------------------------------------------------------------|
| Client              | `internal/client/http/`                     | Сборка и отправка HTTP-запроса, возвращает сырой `*http.Response`            |
| Manager             | `internal/managers/`                        | Проверка статус-кода, разбор JSON в модели, подготовка и очистка данных      |
| Scenario            | `tests/scenarios/`                          | Тестовые сценарии: шаги, проверки, логирование                               |
| Assertions          | `internal/assertions/`                      | Переиспользуемые проверки с понятными сообщениями об ошибке                  |
| Константы           | `internal/constants/`                       | Пути API, бизнес-лимиты (`limits`), коды и тексты ошибок API (`apierrors`)   |
| Фикстуры            | `internal/fixtures/`                        | Фотографии (png, jpg, gif) для загрузки в объявления                         |
| Хелперы             | `internal/helpers/`                         | Запуск запросов (`api-runner`), проверка статуса и разбор ответа (`response`)|

### Сценарии

| Пакет                                  | Файл                               | Что проверяется                                                                                 |
|----------------------------------------|------------------------------------|-------------------------------------------------------------------------------------------------|
| `tests/scenarios/advertisementCreate`  | `create_positive_test.go`          | Создание объявления со всеми параметрами и 3 фото: GET по id, поиск по названию, GET фотографий |
|                                        | `create_validation_test.go`        | 400 при отсутствии обязательных полей и значениях вне допустимых границ                         |
|                                        | `create_unauthorized_test.go`      | 401 без токена и с невалидным токеном                                                           |
| `tests/scenarios/favorites`            | `favorites_positive_test.go`       | Добавление → есть в списке → удаление → нет в списке; удалённое объявление пропадает из избранного |
|                                        | `favorites_nonexistent_test.go`    | Добавление/удаление несуществующего, пустого и некорректного `ad_id`                            |
|                                        | `favorites_unauthorized_test.go`   | 401 на GET, POST и DELETE избранного без токена и с невалидным токеном                          |
|                                        | `favorites_business_rules_test.go` | 409 при повторном добавлении, 403 для своего объявления, 400 при превышении лимита в 50         |
| `tests/scenarios/advertisementSearch`  | `search_test.go`, `search_substring_test.go` | Поиск по подстроке названия и без учёта регистра                                      |
| `tests/scenarios/myAdvertisement`      | `myAdvertisements_test.go`         | Список собственных объявлений                                                                   |

## Найденные дефекты

- `POST /advertisement` принимает `title`, состоящий только из пробелов (`"   "`), и создаёт объявление (HTTP 201),
  хотя пустой `title` отклоняется с ошибкой `title is required`.
- Фактические ограничения строже, чем указано в swagger: `title` — до 50 символов (в swagger для обновления — 255),
  `description` — до 500 символов (в swagger — 1000), `quantity` — до 100, `price` — до 1 000 000.

## Логирование

Тесты пишут в консоль метки `[Precondition]`, `[Step]`, `[Check]` и `[Cleanup]` — подготовку данных, шаги,
результаты проверок и очистку.

Для включения логирования request/response в консоль: выставьте переменную окружения `DEBUG=true` в
`.env.override`

## Запуск линтера

1. Установить приложение [golangci-lint](https://golangci-lint.run/docs/welcome/install/local/) версии 2.x
   (например, `brew install golangci-lint` на macOS)
2. Запуск линтера (из корня репозитория):

```bash
  golangci-lint run
```
