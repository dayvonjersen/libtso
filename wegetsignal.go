// for great justice
zig := make(chan os.Signal, 1)
signal.Notify(zig, os.Interrupt)
go func() { <-zig; os.Exit(0) }()
