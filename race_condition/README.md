### Running tests without race condition

```
go test -v
```

```
=== RUN Test_Increment
--- PASS: Test_Increment (0.00s)
=== RUN Test_IncrementAtomic
--- PASS: Test_IncrementAtomic (0.00s)
PASS
ok race_condition 0.259s

```

### Testing with race condition

```

go test -v --race

```

errors out the Test_Increment test

```

```
