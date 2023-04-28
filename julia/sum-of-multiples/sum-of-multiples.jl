function sum_of_multiples(limit, factors)
    set_of_factors = [Set(factor:factor:limit-1) for factor in factors if factor > 0]
    sum(union(Set(0), set_of_factors...))
end
