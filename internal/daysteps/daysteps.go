package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	s := strings.Split(data, ",")
	if len(s) != 2 {
		//err := fmt.Errorf("Неверные входные данные, строка должна делиться на две состатвные части с разделителем \",\"")
		err := errors.New("Неверные входные данные, строка должна делиться на две состатвные части с разделителем \",\"")
		log.Println(err)
		return 0, 0, err
	}

	steps, err := strconv.Atoi(s[0])
	if err != nil {
		//return 0, 0, fmt.Errorf("Ошибка преобразования в число количества шагов")
		err := errors.New("Ошибка преобразования в число количества шагов")
		log.Println(err)
		return 0, 0, err
	}
	if steps <= 0 {
		err := errors.New("Количество шагов менее или равно 0")
		log.Println(err)
		return 0, 0, err
	}

	t, err := time.ParseDuration(s[1])
	if err != nil {
		//return 0, 0, fmt.Errorf("Ошибка преобразования времени")
		err := errors.New("Ошибка преобразования времени")
		log.Println(err)
		return 0, 0, err
	}

	return steps, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, t, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distKm := float64(steps) * stepLength / float64(mInKm)

	cals, err := spentcalories.WalkingSpentCalories(steps, weight, height, t)

	s := "Количество шагов: " + strconv.Itoa(steps) + ".\nДистанция составила " + strconv.FormatFloat(distKm, 'f', 2, 64) + " км.\nВы сожгли " + strconv.FormatFloat(cals, 'f', 2, 64) + " ккал.\n"

	return s
}
