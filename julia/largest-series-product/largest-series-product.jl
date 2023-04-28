function largest_product(str, span)
    if span < 0 || length(str) < span
        throw(ArgumentError("Incorrect span"))
    end

    span == 0 && return 1

    digits = map(x -> parse(Int64, x), split(str, ""))
    l = length(digits)
    sequences = [digits[i:i+span-1] for i=1:l+1-span]
    maximum(prod, sequences)
end
