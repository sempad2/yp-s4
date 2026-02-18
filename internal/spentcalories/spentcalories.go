package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	//lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	s := strings.Split(data, ",")
	if len(s) != 3 {
		return 0, "", 0, fmt.Errorf("Неверные входные данные, строка должна делиться на три состатвные части с разделителем \",\"")
	}

	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка преобразования в число количества шагов")
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("Количество шагов менее или равно 0")
	}

	typeOfAct := s[1]

	t, err := time.ParseDuration(s[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка преобразования времени")
	}
	if t <= 0 {
		err := errors.New("Продолжительность меньше или равна 0")
		log.Println(err)
		return 0, "", 0, err
	}

	return steps, typeOfAct, t, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	dist := float64(steps) * stepLength / mInKm

	return dist
}

var dur time.Duration

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		err := errors.New("Продолжительность меньше или равна 0")
		log.Println(err)
		return 0
	}

	dist := distance(steps, height)

	averageV := dist / duration.Hours()

	dur = duration

	return averageV
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	var (
		dist     float64
		averageV float64
		cals     float64
	)

	steps, typeOfAct, t, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}

	if typeOfAct != "Бег" && typeOfAct != "Ходьба" {
		//err = fmt.Errorf("неизвестный тип тренировки")
		err := errors.New("неизвестный тип тренировки")
		log.Println(err)
	}

	//if dur <= 0 {
	//err = fmt.Errorf("Продолжительность меньше или равна 0")
	//err := errors.New("Продолжительность меньше или равна 0")
	//log.Println(err)
	//}

	switch typeOfAct {
	case "Ходьба":
		dist = distance(steps, height)
		averageV = meanSpeed(steps, height, t)
		cals, err = WalkingSpentCalories(steps, weight, height, t)
		if err != nil {
			log.Println(err)
		}
	case "Бег":
		dist = distance(steps, height)
		averageV = meanSpeed(steps, height, t)
		cals, err = RunningSpentCalories(steps, weight, height, t)
		if err != nil {
			log.Println(err)
		}
	}

	s := "Тип тренировки: " + typeOfAct + "\nДлительность: " + strconv.FormatFloat(dur.Hours(), 'f', 2, 64) + " ч.\nДистанция: " + strconv.FormatFloat(dist, 'f', 2, 64) + " км.\nСкорость: " + strconv.FormatFloat(averageV, 'f', 2, 64) + " км/ч\nСожгли калорий: " + strconv.FormatFloat(cals, 'f', 2, 64) + "\n"

	return s, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		//return 0, fmt.Errorf("Одно или несколько входящих значений меньше или равно 0")
		err := errors.New("Одно или несколько входящих значений меньше или равно 0")
		log.Println(err)
		return 0, err
	}

	averageV := meanSpeed(steps, height, duration)

	cals := weight * averageV * duration.Minutes() / minInH

	return cals, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	cals, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		//return 0, fmt.Errorf("Одно или несколько входящих значений меньше или равно 0")
		return 0, err
	}

	return cals * walkingCaloriesCoefficient, nil
}
