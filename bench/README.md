# Benchmarks

## Quick benchmark

```
go test -bench . -benchmem
```

## Full benchmark

```
go test -timeout 30m -bench . -benchmem | tee base.txt
```

## Switch JSON backends

```
go test -timeout 30ms -tags jwx_goccy -bench . -benchmem | tee goccy.txt
```

## Comparison

```
% benchstat -sort -delta stdlib.txt v0.4.6.txt
name                                                      old time/op    new time/op    delta
JWK/Serialization/RSA/PrivateKey/jwk.Parse-8                70.0µs ± 1%    18.6µs ± 5%  -73.36%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PrivateKey/jwk.ParseString-8          71.7µs ± 4%    19.5µs ±14%  -72.75%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PrivateKey/jwk.ParseReader-8          72.2µs ± 3%    20.2µs ± 5%  -71.96%  (p=0.008 n=5+5)
JWS/Serialization/JSON/jws.ParseString-8                    36.7µs ±27%    11.3µs ± 0%  -69.05%  (p=0.008 n=5+5)
JWS/Serialization/JSON/jws.ParseReader-8                    34.4µs ±10%    11.6µs ± 1%  -66.30%  (p=0.008 n=5+5)
JWK/Serialization/EC/PrivateKey/jwk.Parse-8                 22.5µs ± 2%     7.9µs ± 1%  -64.97%  (p=0.008 n=5+5)
JWK/Serialization/EC/PrivateKey/jwk.ParseString-8           22.8µs ± 3%     8.0µs ± 2%  -64.87%  (p=0.008 n=5+5)
JWK/Serialization/EC/PrivateKey/jwk.ParseReader-8           22.6µs ± 2%     8.0µs ± 2%  -64.56%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PublicKey/jwk.Parse-8                 21.9µs ± 2%     8.0µs ± 4%  -63.29%  (p=0.008 n=5+5)
JWS/Serialization/JSON/jws.Parse-8                          30.0µs ± 1%    11.1µs ± 2%  -63.18%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PublicKey/jwk.ParseString-8           22.2µs ± 2%     8.5µs ±10%  -61.66%  (p=0.008 n=5+5)
JWK/Serialization/EC/PublicKey/jwk.Parse-8                  18.2µs ± 3%     7.0µs ± 1%  -61.43%  (p=0.008 n=5+5)
JWK/Serialization/EC/PublicKey/jwk.ParseString-8            18.4µs ± 1%     7.2µs ± 2%  -61.18%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PublicKey/jwk.ParseReader-8           21.8µs ± 1%     8.6µs ± 9%  -60.50%  (p=0.008 n=5+5)
JWK/Serialization/EC/PublicKey/jwk.ParseReader-8            18.3µs ± 1%     7.3µs ± 2%  -60.10%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PrivateKey/jwk.Parse-8          12.6µs ± 4%     5.8µs ± 4%  -54.15%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PublicKey/jwk.Parse-8           12.5µs ± 2%     5.8µs ± 1%  -53.88%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PublicKey/jwk.ParseString-8     12.7µs ± 2%     5.9µs ± 1%  -53.65%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PrivateKey/jwk.ParseString-8    12.8µs ± 1%     5.9µs ± 6%  -53.60%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PrivateKey/jwk.ParseReader-8    12.9µs ± 1%     6.0µs ± 3%  -53.39%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PublicKey/jwk.ParseReader-8     12.7µs ± 2%     6.0µs ± 3%  -53.01%  (p=0.008 n=5+5)
JWT/Serialization/Sign/jwt.Parse-8                          15.7µs ±24%     8.1µs ± 2%  -48.76%  (p=0.008 n=5+5)
JWT/Serialization/JSON/jwt.Parse-8                          16.5µs ±10%     9.2µs ± 3%  -44.68%  (p=0.008 n=5+5)
JWT/Serialization/Sign/jwt.ParseString-8                    14.2µs ± 8%     8.3µs ± 2%  -41.44%  (p=0.008 n=5+5)
JWT/Serialization/JSON/jwt.ParseReader-8                    15.9µs ± 4%     9.5µs ± 1%  -40.32%  (p=0.008 n=5+5)
JWT/Serialization/Sign/jwt.ParseReader-8                    13.7µs ± 4%     8.3µs ± 2%  -39.50%  (p=0.008 n=5+5)
JWT/Serialization/JSON/jwt.ParseString-8                    15.1µs ± 2%     9.2µs ± 2%  -39.34%  (p=0.008 n=5+5)
JWS/Serialization/Compact/jws.Parse-8                       6.12µs ± 3%    3.76µs ± 2%  -38.61%  (p=0.008 n=5+5)
JWS/Serialization/Compact/jws.ParseString-8                 6.27µs ± 3%    3.85µs ± 2%  -38.54%  (p=0.008 n=5+5)
JWS/Serialization/Compact/jws.ParseReader-8                 6.29µs ± 2%    3.97µs ± 1%  -36.84%  (p=0.008 n=5+5)
JWT/Serialization/JSON/json.Unmarshal-8                     2.32µs ± 5%    2.04µs ± 1%  -11.86%  (p=0.008 n=5+5)
JWS/Serialization/JSON/json.Marshal-8                       18.3µs ± 7%    16.9µs ± 3%   -7.57%  (p=0.008 n=5+5)
JWT/Serialization/JSON/json.Marshal-8                       5.86µs ± 6%    5.46µs ± 5%   -6.76%  (p=0.032 n=5+5)
JWK/Serialization/EC/PublicKey/json.Marshal-8               9.15µs ± 3%    8.71µs ± 1%   -4.79%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PublicKey/json.Marshal-8        6.80µs ± 1%    6.58µs ± 2%   -3.21%  (p=0.008 n=5+5)
JWE/Serialization/JSON/json.Marshal-8                       19.8µs ±11%    19.7µs ±13%     ~     (p=1.000 n=5+5)
JWE/Serialization/JSON/json.Unmarshal-8                     6.76µs ± 5%    6.79µs ± 7%     ~     (p=1.000 n=5+5)
JWK/Serialization/RSA/PublicKey/json.Marshal-8              9.74µs ± 9%    9.73µs ± 6%     ~     (p=0.548 n=5+5)
JWK/Serialization/RSA/PrivateKey/json.Marshal-8             22.2µs ± 3%    23.1µs ± 7%     ~     (p=0.056 n=5+5)
JWK/Serialization/EC/PrivateKey/json.Marshal-8              10.1µs ± 1%     9.8µs ± 4%     ~     (p=0.151 n=5+5)
JWK/Serialization/Symmetric/PrivateKey/json.Marshal-8       6.93µs ± 9%    6.66µs ± 4%     ~     (p=0.310 n=5+5)
JWT/Serialization/Sign/jwt.Sign-8                           2.05ms ± 6%    1.96ms ± 7%     ~     (p=0.222 n=5+5)

name                                                      old alloc/op   new alloc/op   delta
JWK/Serialization/RSA/PrivateKey/jwk.Parse-8                36.5kB ± 0%    22.9kB ± 0%  -37.32%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PrivateKey/jwk.ParseString-8          38.3kB ± 0%    24.6kB ± 0%  -35.57%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PrivateKey/jwk.ParseReader-8          41.1kB ± 0%    27.5kB ± 0%  -33.13%  (p=0.008 n=5+5)
JWS/Serialization/JSON/jws.Parse-8                          16.3kB ± 0%    14.9kB ± 0%   -8.71%  (p=0.008 n=5+5)
JWS/Serialization/JSON/jws.ParseString-8                    17.2kB ± 0%    15.8kB ± 0%   -8.25%  (p=0.008 n=5+5)
JWS/Serialization/JSON/jws.ParseReader-8                    17.9kB ± 0%    16.5kB ± 0%   -7.96%  (p=0.008 n=5+5)
JWE/Serialization/JSON/json.Marshal-8                       3.96kB ± 0%    3.96kB ± 0%   -0.13%  (p=0.016 n=4+5)
JWK/Serialization/RSA/PrivateKey/json.Marshal-8             8.83kB ± 0%    8.82kB ± 0%   -0.12%  (p=0.008 n=5+5)
JWS/Serialization/JSON/json.Marshal-8                       5.17kB ± 0%    5.17kB ± 0%   -0.11%  (p=0.008 n=5+5)
JWT/Serialization/Sign/jwt.Sign-8                           36.3kB ± 0%    36.2kB ± 0%   -0.11%  (p=0.008 n=5+5)
JWK/Serialization/EC/PublicKey/json.Marshal-8               1.79kB ± 0%    1.79kB ± 0%   -0.06%  (p=0.000 n=5+4)
JWK/Serialization/EC/PrivateKey/json.Marshal-8              2.28kB ± 0%    2.28kB ± 0%   -0.04%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PublicKey/json.Marshal-8              2.34kB ± 0%    2.34kB ± 0%   -0.04%  (p=0.008 n=5+5)
JWE/Serialization/JSON/json.Unmarshal-8                     1.58kB ± 0%    1.58kB ± 0%     ~     (all equal)
JWK/Serialization/Symmetric/PublicKey/json.Marshal-8        1.15kB ± 0%    1.15kB ± 0%     ~     (all equal)
JWK/Serialization/Symmetric/PrivateKey/json.Marshal-8       1.15kB ± 0%    1.15kB ± 0%     ~     (all equal)
JWT/Serialization/JSON/json.Unmarshal-8                       592B ± 0%      592B ± 0%     ~     (all equal)
JWT/Serialization/JSON/json.Marshal-8                         720B ± 0%      720B ± 0%     ~     (all equal)
JWS/Serialization/Compact/jws.ParseReader-8                 3.34kB ± 0%    3.47kB ± 0%   +3.83%  (p=0.008 n=5+5)
JWS/Serialization/Compact/jws.ParseString-8                 3.02kB ± 0%    3.15kB ± 0%   +4.23%  (p=0.008 n=5+5)
JWS/Serialization/Compact/jws.Parse-8                       2.83kB ± 0%    2.96kB ± 0%   +4.52%  (p=0.008 n=5+5)
JWK/Serialization/EC/PrivateKey/jwk.ParseReader-8           8.13kB ± 0%    8.54kB ± 0%   +5.02%  (p=0.008 n=5+5)
JWK/Serialization/EC/PrivateKey/jwk.ParseString-8           7.97kB ± 0%    8.38kB ± 0%   +5.12%  (p=0.008 n=5+5)
JWK/Serialization/EC/PrivateKey/jwk.Parse-8                 7.62kB ± 0%    8.02kB ± 0%   +5.36%  (p=0.008 n=5+5)
JWK/Serialization/EC/PublicKey/jwk.ParseReader-8            7.52kB ± 0%    7.99kB ± 0%   +6.28%  (p=0.008 n=5+5)
JWK/Serialization/EC/PublicKey/jwk.ParseString-8            7.25kB ± 0%    7.72kB ± 0%   +6.51%  (p=0.008 n=5+5)
JWK/Serialization/EC/PublicKey/jwk.Parse-8                  7.01kB ± 0%    7.48kB ± 0%   +6.74%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PublicKey/jwk.ParseReader-8     6.75kB ± 0%    7.48kB ± 0%  +10.78%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PrivateKey/jwk.ParseReader-8    6.75kB ± 0%    7.48kB ± 0%  +10.78%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PublicKey/jwk.ParseReader-8           7.70kB ± 0%    8.55kB ± 0%  +11.12%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PublicKey/jwk.ParseString-8           7.60kB ± 0%    8.46kB ± 0%  +11.26%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PublicKey/jwk.ParseString-8     6.37kB ± 0%    7.10kB ± 0%  +11.43%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PrivateKey/jwk.ParseString-8    6.37kB ± 0%    7.10kB ± 0%  +11.43%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PublicKey/jwk.Parse-8           6.24kB ± 0%    6.97kB ± 0%  +11.67%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PrivateKey/jwk.Parse-8          6.24kB ± 0%    6.97kB ± 0%  +11.67%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PublicKey/jwk.Parse-8                 7.18kB ± 0%    8.04kB ± 0%  +11.92%  (p=0.008 n=5+5)
JWT/Serialization/Sign/jwt.ParseReader-8                    11.0kB ± 0%    12.9kB ± 0%  +16.99%  (p=0.008 n=5+5)
JWT/Serialization/Sign/jwt.ParseString-8                    11.0kB ± 0%    12.8kB ± 0%  +17.09%  (p=0.008 n=5+5)
JWT/Serialization/Sign/jwt.Parse-8                          10.5kB ± 0%    12.4kB ± 0%  +17.82%  (p=0.008 n=5+5)
JWT/Serialization/JSON/jwt.ParseReader-8                    10.6kB ± 0%    12.9kB ± 0%  +21.28%  (p=0.008 n=5+5)
JWT/Serialization/JSON/jwt.ParseString-8                    10.1kB ± 0%    12.4kB ± 0%  +22.26%  (p=0.008 n=5+5)
JWT/Serialization/JSON/jwt.Parse-8                          10.1kB ± 0%    12.3kB ± 0%  +22.36%  (p=0.008 n=5+5)

name                                                      old allocs/op  new allocs/op  delta
JWK/Serialization/RSA/PrivateKey/jwk.Parse-8                   205 ± 0%        83 ± 0%  -59.51%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PrivateKey/jwk.ParseString-8             206 ± 0%        84 ± 0%  -59.22%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PrivateKey/jwk.ParseReader-8             209 ± 0%        87 ± 0%  -58.37%  (p=0.008 n=5+5)
JWK/Serialization/EC/PrivateKey/jwk.Parse-8                    129 ± 0%        59 ± 0%  -54.26%  (p=0.008 n=5+5)
JWK/Serialization/EC/PrivateKey/jwk.ParseString-8              130 ± 0%        60 ± 0%  -53.85%  (p=0.008 n=5+5)
JWK/Serialization/EC/PrivateKey/jwk.ParseReader-8              130 ± 0%        60 ± 0%  -53.85%  (p=0.008 n=5+5)
JWK/Serialization/EC/PublicKey/jwk.Parse-8                     112 ± 0%        54 ± 0%  -51.79%  (p=0.008 n=5+5)
JWK/Serialization/EC/PublicKey/jwk.ParseString-8               113 ± 0%        55 ± 0%  -51.33%  (p=0.008 n=5+5)
JWK/Serialization/EC/PublicKey/jwk.ParseReader-8               113 ± 0%        55 ± 0%  -51.33%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PublicKey/jwk.Parse-8                   96.0 ± 0%      50.0 ± 0%  -47.92%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PublicKey/jwk.ParseString-8             97.0 ± 0%      51.0 ± 0%  -47.42%  (p=0.008 n=5+5)
JWK/Serialization/RSA/PublicKey/jwk.ParseReader-8             97.0 ± 0%      51.0 ± 0%  -47.42%  (p=0.008 n=5+5)
JWS/Serialization/Compact/jws.Parse-8                         49.0 ± 0%      27.0 ± 0%  -44.90%  (p=0.008 n=5+5)
JWS/Serialization/Compact/jws.ParseString-8                   50.0 ± 0%      28.0 ± 0%  -44.00%  (p=0.008 n=5+5)
JWS/Serialization/Compact/jws.ParseReader-8                   50.0 ± 0%      28.0 ± 0%  -44.00%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PublicKey/jwk.Parse-8             81.0 ± 0%      47.0 ± 0%  -41.98%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PrivateKey/jwk.Parse-8            81.0 ± 0%      47.0 ± 0%  -41.98%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PublicKey/jwk.ParseString-8       82.0 ± 0%      48.0 ± 0%  -41.46%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PublicKey/jwk.ParseReader-8       82.0 ± 0%      48.0 ± 0%  -41.46%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PrivateKey/jwk.ParseString-8      82.0 ± 0%      48.0 ± 0%  -41.46%  (p=0.008 n=5+5)
JWK/Serialization/Symmetric/PrivateKey/jwk.ParseReader-8      82.0 ± 0%      48.0 ± 0%  -41.46%  (p=0.008 n=5+5)
JWS/Serialization/JSON/jws.Parse-8                             153 ± 0%        96 ± 0%  -37.25%  (p=0.008 n=5+5)
JWS/Serialization/JSON/jws.ParseString-8                       154 ± 0%        97 ± 0%  -37.01%  (p=0.008 n=5+5)
JWS/Serialization/JSON/jws.ParseReader-8                       155 ± 0%        98 ± 0%  -36.77%  (p=0.008 n=5+5)
JWT/Serialization/Sign/jwt.Parse-8                             103 ± 0%        67 ± 0%  -34.95%  (p=0.008 n=5+5)
JWT/Serialization/Sign/jwt.ParseString-8                       104 ± 0%        68 ± 0%  -34.62%  (p=0.008 n=5+5)
JWT/Serialization/Sign/jwt.ParseReader-8                       104 ± 0%        68 ± 0%  -34.62%  (p=0.008 n=5+5)
JWT/Serialization/JSON/jwt.Parse-8                            73.0 ± 0%      60.0 ± 0%  -17.81%  (p=0.008 n=5+5)
JWT/Serialization/JSON/jwt.ParseString-8                      74.0 ± 0%      61.0 ± 0%  -17.57%  (p=0.008 n=5+5)
JWT/Serialization/JSON/jwt.ParseReader-8                      74.0 ± 0%      61.0 ± 0%  -17.57%  (p=0.008 n=5+5)
JWE/Serialization/JSON/json.Marshal-8                         45.0 ± 0%      45.0 ± 0%     ~     (all equal)
JWE/Serialization/JSON/json.Unmarshal-8                       26.0 ± 0%      26.0 ± 0%     ~     (all equal)
JWK/Serialization/RSA/PublicKey/json.Marshal-8                24.0 ± 0%      24.0 ± 0%     ~     (all equal)
JWK/Serialization/RSA/PrivateKey/json.Marshal-8               51.0 ± 0%      51.0 ± 0%     ~     (all equal)
JWK/Serialization/EC/PublicKey/json.Marshal-8                 27.0 ± 0%      27.0 ± 0%     ~     (all equal)
JWK/Serialization/EC/PrivateKey/json.Marshal-8                31.0 ± 0%      31.0 ± 0%     ~     (all equal)
JWK/Serialization/Symmetric/PublicKey/json.Marshal-8          20.0 ± 0%      20.0 ± 0%     ~     (all equal)
JWK/Serialization/Symmetric/PrivateKey/json.Marshal-8         20.0 ± 0%      20.0 ± 0%     ~     (all equal)
JWS/Serialization/JSON/json.Marshal-8                         60.0 ± 0%      60.0 ± 0%     ~     (all equal)
JWT/Serialization/Sign/jwt.Sign-8                              192 ± 0%       192 ± 0%     ~     (all equal)
JWT/Serialization/JSON/json.Unmarshal-8                       10.0 ± 0%      10.0 ± 0%     ~     (all equal)
JWT/Serialization/JSON/json.Marshal-8                         18.0 ± 0%      18.0 ± 0%     ~     (all equal)
```
