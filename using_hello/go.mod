module using_example/using_hello

go 1.23.4

replace example/hello => ../hello

require example/hello v0.0.0-00010101000000-000000000000 // indirect
