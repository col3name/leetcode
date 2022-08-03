func isHappy(n int) bool {
    set := make(map[int]struct{})
    return helper(set, n)
}

func helper(set map[int]struct{}, n int) bool {
    if n == 1 {
        return true
    }
    if _, ok := set[n]; ok {
        return false
    } else {
        set[n] = struct{}{}
    }
    
    m := 0
    for {
        if n/10 == 0 {
            m += (n%10)*(n%10)
            break
        } else {
            m += (n%10)*(n%10)
            n /= 10
        }
    }
    return helper(set, m)
}
