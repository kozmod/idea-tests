package context

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//ToDo -> need finish
func Test(t *testing.T) {
	// Создаем контекст background
	ctx := context.Background()
	// Производим контекст с отменой
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)

	// Отложенная отмена высвобождает все ресурсы
	// для этого и производных от него контекстов
	defer func() {
		fmt.Println("Main Defer: canceling context")
		cancelFunction()
	}()
	// Отмена контекста после случайного тайм-аута
	// Если это происходит, все производные от него контексты должны завершиться
	go func() {
		sleepRandom("Main", nil)
		cancelFunction()
		fmt.Println("Main Sleep complete. canceling context")
	}()

	// Выполнение работы
	doWorkContext(ctxWithCancel)
}

// Медленная функция
func sleepRandom(fromFunction string, ch chan int) {
	// Отложенная функция очистки
	defer func() { fmt.Println(fromFunction, "sleepRandom complete") }()

	// Выполним медленную задачу
	// В качестве примера,
	// «заснем» на рандомное время в мс
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleeptime := randomNumber + 100

	fmt.Println(fromFunction, "Starting sleep for", sleeptime, "ms")
	time.Sleep(time.Duration(sleeptime) * time.Millisecond)
	fmt.Println(fromFunction, "Waking up, slept for ", sleeptime, "ms")

	// Напишем в канал, если он был передан
	if ch != nil {
		ch <- sleeptime
	}
}

// Функция, выполняющая медленную работу с использованием контекста
// Заметьте, что контекст - это первый аргумент
func sleepRandomContext(ctx context.Context, ch chan bool) {

	// Выполнение (прим. пер.: отложенное выполнение) действий по очистке
	// Созданных контекстов больше нет
	// Следовательно, отмена не требуется
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()

	// Создаем канал
	sleeptimeChan := make(chan int)

	// Запускаем выполнение медленной задачи в горутине
	// Передаем канал для коммуникаций
	go sleepRandom("sleepRandomContext", sleeptimeChan)

	// Используем select для выхода по истечении времени жизни контекста
	select {
	case <-ctx.Done():
		// Если контекст отменен, выбирается этот случай
		// Это случается, если заканчивается таймаут doWorkContext или
		// doWorkContext или main вызывает cancelFunction
		// Высвобождаем ресурсы, которые больше не нужны из-за прерывания работы
		// Посылаем сигнал всем горутинам, которые должны завершиться (используя каналы)
		// Обычно вы посылаете что-нибудь в канал,
		// ждете выхода из горутины, затем возвращаетесь
		// Или используете группы ожидания вместо каналов для синхронизации
		fmt.Println("sleepRandomContext: Time to return")

	case sleeptime := <-sleeptimeChan:
		// Этот вариант выбирается, когда работа завершается до отмены контекста
		fmt.Println("Slept for ", sleeptime, "ms")
	}
}

// Вспомогательная функция, которая в реальности может использоваться для разных целей
// Здесь она просто вызывает одну функцию
// В данном случае, она могла бы быть в main
func doWorkContext(ctx context.Context) {

	// От контекста с функцией отмены создаём производный контекст с тайм-аутом
	// Таймаут 150 мс
	// Все контексты, производные от этого, завершатся через 150 мс
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)

	// Функция отмены для освобождения ресурсов после завершения функции
	defer func() {
		fmt.Println("doWorkContext complete")
		cancelFunction()
	}()

	// Создаем канал и вызываем функцию контекста
	// Можно также использовать группы ожидания для этого конкретного случая,
	// поскольку мы не используем возвращаемое значение, отправленное в канал
	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)

	// Используем select для выхода при истечении контекста
	select {
	case <-ctx.Done():
		// Этот случай выбирается, когда переданный в качестве аргумента контекст уведомляет о завершении работы
		// В данном примере это произойдёт, когда в main будет вызвана cancelFunction
		fmt.Println("doWorkContext: Time to return")

	case <-ch:
		// Этот вариант выбирается, когда работа завершается до отмены контекста
		fmt.Println("sleepRandomContext returned")
	}
}
