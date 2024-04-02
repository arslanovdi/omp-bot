package fake

import (
	"github.com/arslanovdi/omp-bot/internal/model"
	"github.com/arslanovdi/omp-bot/internal/service"
	"github.com/brianvoe/gofakeit/v7"
	"log/slog"
	"math/rand"
	"time"
)

const (
	create = 30
	delete = 10
	get    = 20
	list   = 20
	update = 20
)

var counter uint64

// Emulate эмуляция работы пользователей телеграм бота
func Emulate(d time.Duration, pkgService *service.LogisticPackageService) {
	for {
		time.Sleep(d)

		rnd := rand.Intn(100) // 100%
		switch {
		case rnd < create: // create % созданий пакета
			pkg := model.Package{}

			if gofakeit.Bool() {
				pkg.Weight = new(uint64)
				*pkg.Weight = gofakeit.Uint64()
			}

			pkg.Title = gofakeit.ProductName()
			pkg.Created = gofakeit.DateRange(time.Now().AddDate(0, -2, 0), time.Now())

			_, err := pkgService.Create(pkg)
			if err != nil {
				slog.Error("FAKE fail to create package", "error", err)
			}
			counter++
		case rnd < create+delete: // delete % удаления пакета

			if counter <= 2 {
				continue
			}

			id := uint64(rand.Int63n(int64(counter))) + 1

			err := pkgService.Delete(id)
			if err != nil {
				slog.Error("FAKE fail to delete package", "error", err)
			}
			counter--
		case rnd < create+delete+get: // get% получения пакета

			if counter <= 2 {
				continue
			}
			id := uint64(rand.Int63n(int64(counter))) + 1

			_, err := pkgService.Get(id)
			if err != nil {
				slog.Error("FAKE fail to get package", "id", id)
			}
		case rnd < create+delete+get+list: // list % получения списка пакетов
			if counter < 10 {
				continue
			}
			offset := rand.Int63n(int64(counter/2) + 1)
			limit := rand.Int63n(int64(counter) - offset)
			_, err := pkgService.List(uint64(offset), uint64(limit))
			if err != nil {
				slog.Error("FAKE fail to list package", "error", err)
			}

		case rnd < create+delete+get+list+update: // update % обновления пакета

			if counter <= 2 {
				continue
			}

			pkg := model.Package{}

			pkg.ID = uint64(rand.Int63n(int64(counter)) + 1)

			if gofakeit.Bool() {
				pkg.Weight = new(uint64)
				*pkg.Weight = gofakeit.Uint64()
			}

			if gofakeit.Bool() {
				pkg.Title = gofakeit.ProductName()
			}

			pkg.Updated = new(time.Time)
			*pkg.Updated = time.Now()

			err := pkgService.Update(pkg)
			if err != nil {
				slog.Error("FAKE fail to update package", "error", err)
			}
		}
	}
}
