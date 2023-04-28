function transform(input::AbstractDict)
    # Awesome feature. Stolen from community solutions
    Dict(
        lowercase(l) => k
        for (k, v) in input
        for l in v
    )
end

