package main

import (
	"errors"
	"fmt"
	"lesson5/student/practice"
)

type Example struct {
	title      string
	notifiers  []practice.Notifier
	messages   []string
	dispatcher func([]practice.Notifier, string) error
}

func RunExamples() {
	examples := []Example{
		{
			title: "All OK",
			notifiers: []practice.Notifier{
				practice.NewEmailNotifier("123@x5.tech"),
				practice.NewPushNotifier("x5tech_some_other_id"),
			},
			messages:   []string{"Заказ готов"},
			dispatcher: practice.Dispatch,
		},
		{
			title: "Empty message between normal messages",
			notifiers: []practice.Notifier{
				practice.NewEmailNotifier("123@x5.tech"),
				practice.NewPushNotifier("x5tech_some_other_id"),
			},
			messages:   []string{"testing", "", "again testing"},
			dispatcher: practice.Dispatch,
		},
		{
			title: "No DeviceID",
			notifiers: []practice.Notifier{
				practice.NewPushNotifier(""),
			},
			messages:   []string{"testing"},
			dispatcher: practice.Dispatch,
		},
		{
			title: "Broken notifier (panic recovering)",
			notifiers: []practice.Notifier{
				practice.NewBrokenNotifier(),
				practice.NewPushNotifier("x5tech_some_other_id"),
			},
			messages: []string{"testing"},
			dispatcher: func(notifiers []practice.Notifier, message string) error {
				errs := practice.SafeDispatch(notifiers, message)

				return errs[0]
			},
		},
	}

	for _, example := range examples {
		runner(example)
	}
}

func runner(example Example) {
	fmt.Println("Running:", example.title)

	var errs []error

	for _, message := range example.messages {
		if err := example.dispatcher(example.notifiers, message); err != nil {
			errs = append(errs, err)
		}
	}

	printResult(example.notifiers, errs)

	fmt.Println("Done:", example.title)
}

func printResult(notifiers []practice.Notifier, errs []error) {
	fmt.Println("\tDispatch errors:")
	var typedErr *practice.SendError

	for i, err := range errs {
		fmt.Printf("\t\tError's #%d:\n", i)
		switch {
		case errors.Is(err, practice.ErrEmptyMessage):
			fmt.Printf("\t\t\tType:    ErrEmptyMessage\n")
			fmt.Printf("\t\t\tMessage: %s\n", err.Error())
		case errors.As(err, &typedErr):
			fmt.Printf("\t\t\tType:    %T\n", typedErr)
			fmt.Printf("\t\t\tChannel: %s\n", typedErr.Channel)
			fmt.Printf("\t\t\tReason:  %s\n", typedErr.Reason)
			fmt.Printf("\t\t\tMessage: %s\n", typedErr.Error())
		case errors.Is(err, practice.ErrNotifierPanic):
			fmt.Printf("\t\t\tType:    ErrNotifierPanic\n")
			fmt.Printf("\t\t\tMessage: %s\n", err.Error())
		default:
			fmt.Printf("\t\t\tUndefined error: %s\n", err.Error())
		}
	}
	if len(errs) == 0 {
		fmt.Println("\t\tNo errors")
	}

	fmt.Println("\tNotifiers:")
	for i, notifier := range notifiers {
		fmt.Printf("\t\tNotifier's #%d:\n", i)
		switch t := notifier.(type) {
		case *practice.EmailNotifier:
			fmt.Printf("\t\t\tType:    %T\n", t)
			fmt.Printf("\t\t\tHistory: %q\n", t.Messages)
		case *practice.PushNotifier:
			fmt.Printf("\t\t\tType:    %T\n", t)
			fmt.Printf("\t\t\tHistory: %q\n", t.Messages)
		case *practice.BrokenNotifier:
			fmt.Printf("\t\t\tType:    %T\n", t)
		default:
			fmt.Println("\t\t\tUndefined")
		}
	}
}

func main() {
	RunExamples()
}
