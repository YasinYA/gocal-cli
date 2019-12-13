package main

import (
    "fmt"
    "log"
    "time"
)

func GetEvents() {
    srv, err := GetCalendarSerivce()
    if err != nil {
        log.Fatalf("Unable to retrieve Calendar client: %v", err)
    }

    t := time.Now().Format(time.RFC3339)
    cl, _ := srv.CalendarList.List().Do()
    for _, item := range cl.Items {
        fmt.Printf("====> %v\n\n", item.Summary)
        fmt.Println("Upcoming events:")
        events, err := srv.Events.List(item.Id).ShowDeleted(false).
            SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
        if err != nil {
            log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
        }
        if len(events.Items) == 0 {
            fmt.Println("No upcoming events found.")
        } else {
            for _, eventItem := range events.Items {
                date := eventItem.Start.DateTime
                if date == "" {
                    date = eventItem.Start.Date
                }
                fmt.Printf("%v on %v \n\n", eventItem.Summary, date)
            }
        }

    }
    // events, err := srv.Events.List("planing").ShowDeleted(false).
    //     SingleEvents(true).TimeMin(t).MaxResults(20).OrderBy("startTime").Do()
    // if err != nil {
    //     log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
    // }
    // fmt.Println("Upcoming events:")
    // if len(events.Items) == 0 {
    //     fmt.Println("No upcoming events found.")
    // } else {
    //     for _, item := range cl.Items {
    //         date := item.Start.DateTime
    //         if date == "" {
    //             date = item.Start.Date
    //         }
    //         fmt.Printf("%v (%v)\n", item.Summary, date)
    //     }
    // }
}
