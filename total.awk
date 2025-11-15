/^Benchmark/ { total += $3 }
END { printf("%.2f ns, %.3f µs, %.3f ms, %.6f s\n", total, total/1e3, total/1e6, total/1e9) }
