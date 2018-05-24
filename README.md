# go-contains-benchmark

# Output
    $>go run main.go
    use slice for size (500): 2.635µs 
    use map for size (500) exist (true): 248ns 
    use map for size (500) exist (false): 199ns 
    use string contains for size (500) exist (true): 1.124µs 
    use string contains for size (500) exist (false): 949ns 

    use slice for size (5000): 22.033µs 
    use map for size (5000) exist (true): 148ns 
    use map for size (5000) exist (false): 158ns 
    use string contains for size (5000) exist (true): 21.595µs 
    use string contains for size (5000) exist (false): 10.003µs 

    use slice for size (50000): 212.945µs 
    use map for size (50000) exist (true): 750ns 
    use map for size (50000) exist (false): 611ns 
    use string contains for size (50000) exist (true): 119.333µs 
    use string contains for size (50000) exist (false): 127.84µs 

    use slice for size (500000): 2.236291ms 
    use map for size (500000) exist (true): 714ns 
    use map for size (500000) exist (false): 587ns 
    use string contains for size (500000) exist (true): 1.892981ms 
    use string contains for size (500000) exist (false): 1.891026ms 

    use slice for size (5000000): 21.048319ms 
    use map for size (5000000) exist (true): 841ns 
    use map for size (5000000) exist (false): 767ns 
    use string contains for size (5000000) exist (true): 15.229476ms 
    use string contains for size (5000000) exist (false): 15.348878m

# Conclusion
  using slice: the execution time is proportional to the size  
  using map: the best performance (regardless of size)  
  using string: the worst performance (regardless of size)   
