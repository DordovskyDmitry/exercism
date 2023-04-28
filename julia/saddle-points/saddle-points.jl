function saddlepoints(M)
    ndims(M) != 2 && return []

    r,c = size(M)
    [(i,j) for i=1:r, j=1:c if M[i, j] == maximum(M[i, :]) && M[i, j] == minimum(M[:, j])]
end
