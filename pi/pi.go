package main
import (
    "fmt"
    "math/big"
)
func computePi(digits int) *big.Float {
    prec := uint(float64(digits)*3.32193) + 32
    C := new(big.Float).SetPrec(prec).SetFloat64(426880)
    val10005 := new(big.Float).SetPrec(prec).SetFloat64(10005)
    sqrt10005 := new(big.Float).SetPrec(prec)
    sqrt10005.Sqrt(val10005)
    C.Mul(C, sqrt10005)
    M := new(big.Float).SetPrec(prec).SetFloat64(1)
    L := new(big.Float).SetPrec(prec).SetFloat64(13591409)
    X := new(big.Float).SetPrec(prec).SetFloat64(1)
    K := new(big.Float).SetPrec(prec).SetFloat64(6)
    S := new(big.Float).SetPrec(prec).Set(L)
    terms := digits/14 + 1
    tmp1 := new(big.Float).SetPrec(prec)
    tmp2 := new(big.Float).SetPrec(prec)
    iCubed := new(big.Float).SetPrec(prec)
    for i := 1; i < terms; i++ {
        tmp1.Mul(K, K)
        tmp1.Mul(tmp1, K)
        tmp2.Mul(big.NewFloat(16), K)
        tmp1.Sub(tmp1, tmp2)
        iFloat := new(big.Float).SetPrec(prec).SetFloat64(float64(i))
        iCubed.Mul(iFloat, iFloat)
        iCubed.Mul(iCubed, iFloat)
        M.Mul(M, tmp1)
        M.Quo(M, iCubed)
        L.Add(L, big.NewFloat(545140134))
        X.Mul(X, big.NewFloat(-262537412640768000))
        term := new(big.Float).SetPrec(prec)
        term.Mul(M, L)
        term.Quo(term, X)
        S.Add(S, term)
        K.Add(K, big.NewFloat(12))
    }
    pi := new(big.Float).SetPrec(prec)
    pi.Quo(C, S)
    return pi
}
func main() {
    var digits int
    fmt.Print("Digits(stops being accurate at ~158k+): ")
    fmt.Scan(&digits)
    pi := computePi(digits)
    piStr := pi.Text('f', digits)
    fmt.Printf("π ≈ %s\n", piStr)
}
