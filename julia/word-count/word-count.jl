function wordcount(sentence)
    words = [ strip(match.match, '\'') for match in eachmatch(r"[\w\']+", lowercase(sentence)) ]
    without_empty = filter((w) -> w != "", words)

    dict = Dict()

    for word in without_empty
        haskey(dict, word) ? dict[word] += 1 : dict[word] = 1
    end

    dict
end
