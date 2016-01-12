var dbError error
maxAttempts := 20
for attempts := 1; attempts <= maxAttempts; attempts++ {
    dbError = db.Ping()
    if dbError == nil {
        break
    }
    log.Println(dbError)
    time.Sleep(time.Duration(attempts) * time.Second)
}
if dbError != nil {
    log.Fatal(dbError)
}
